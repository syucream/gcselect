package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"cloud.google.com/go/bigquery"
	"google.golang.org/api/iterator"
)

const (
	tmpTableName = "__gcselect"
)

var (
	toSourceFormat = map[string]bigquery.DataFormat{
		string(bigquery.Avro):    bigquery.Avro,
		string(bigquery.CSV):     bigquery.CSV,
		string(bigquery.JSON):    bigquery.JSON,
		string(bigquery.Parquet): bigquery.Parquet,
		string(bigquery.ORC):     bigquery.ORC,
	}
)

func query(ctx context.Context, projectId, query string, exdata bigquery.ExternalData) ([]map[string]bigquery.Value, error) {
	client, err := bigquery.NewClient(ctx, projectId)
	if err != nil {
		return nil, err
	}

	q := client.Query(query)

	q.TableDefinitions = map[string]bigquery.ExternalData{
		tmpTableName: exdata,
	}

	job, err := q.Run(ctx)
	if err != nil {
		return nil, err
	}

	ri, err := job.Read(ctx)
	if err != nil {
		return nil, err
	}

	rows := make([]map[string]bigquery.Value, 0, ri.TotalRows)
	for {
		var r map[string]bigquery.Value
		err = ri.Next(&r)

		if err == iterator.Done {
			break
		} else if err != nil {
			return nil, err
		}

		rows = append(rows, r)
	}

	return rows, nil
}

func main() {
	projectId := flag.String("projectId", "'", "GCP project id")
	sourceFormat := flag.String("sourceFormat", string(bigquery.Avro), "source format of gcs objects")

	flag.Parse()

	args := flag.Args()
	numArgs := len(args)

	if numArgs < 2 {
		log.Fatalf("no required arguments")
	}

	sf, sfOk := toSourceFormat[*sourceFormat]
	if !sfOk {
		log.Fatalf("unsupported source format: %s", *sourceFormat)
	}

	conf := &bigquery.ExternalDataConfig{
		SourceFormat: sf,
		SourceURIs:   []string{args[0]},
	}

	ctx := context.Background()

	rows, err := query(ctx, *projectId, args[numArgs-1], conf)
	if err != nil {
		log.Fatal(err)
	}

	jsonStr, err := json.Marshal(rows)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(jsonStr))
}
