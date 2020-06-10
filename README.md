# gcselect

select query to GCS

## How to work

It uses GCS external data source with temporary table on BigQuery. So it requires BigQuery.
https://cloud.google.com/bigquery/external-data-cloud-storage?hl=en#temporary-tables

## How to install

```go
GO111MODULE=off go get github.com/syucream/gcselect
```

## How to use

```sh
$ cat user.avsc
{
  "name": "User",
  "type": "record",
  "fields": [
    {
      "name": "id",
      "type": "long"
    },
    {
      "name": "name",
      "type": "string"
    },
    {
      "name": "age",
      "type": "long"
    }
  ]
}
```

```sh
$ gcselect gs://syucream-dev/test/user.avro 'SELECT id, name, age FROM __gcselect'
[{"age":7532107857140623916,"id":6006445776448206505,"name":"gktla"},{"age":-3205607056053824364,"id":851631935230924000,"name":"hlwt"},{"age":8540484859928717878,"id":-1574784563590716287,"name":"ekxi"},{"age":1202055264913828667,"id":3601873433973733209,"name":"gc"},{"age":5505536880286505386,"id":2044201735586106587,"name":"kvdeltmytm"},{"age":439990441965119387,"id":2270904139985944258,"name":"shtkbwyt"},{"age":-3430001363178382567,"id":-8437730792950743681,"name":"dudiktu"},{"age":-4911467699644869312,"id":5877633693208886821,"name":"swbnjdfdtt"},{"age":-3014264803582721781,"id":-3929798637128994569,"name":"ngvdckclpkcbe"},{"age":865908852910145979,"id":-1588204925166013109,"name":"nwiw"}]
```
