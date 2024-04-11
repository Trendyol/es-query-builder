# es-query-builder [![GoDoc][doc-img]][doc] [![Go Report Card][go-report-img]][go-report]
Simple Query builder for Elasticsearch

```bash
go get github.com/GokselKUCUKSAHIN/es-query-builder
```

# TODOs

## Project
- [ ] improve README
  - [ ] add examples
  - [ ] add benchmark results
  - [x] add `go get` command
- [x] add benchmark
- [x] add missing tests
- [x] add makefile
- [x] setup linter

## Builder fields
- [x] range
- [x] sort
- [ ] nested
- [ ] aggs
- [ ] match
- [ ] match_all
- [ ] match_none
- [x] minimum_should_match
- [x] boost


### Examples 

# 🚧 Still under construction 🚧

```json
{
  "query": {
    "bool": {
      "must": [
        {
          "bool": {
            "should": [
              {
                "term": {
                  "doc.id": "293"
                }
              },
              {
                "term": {
                  "file.fileId": "293"
                }
              }
            ]
          }
        }
      ],
      "filter": [
        {
          "terms": {
            "type": [
              "DOC",
              "FILE"
            ]
          }
        }
      ]
    }
  }
}
```

With pure Go
```go
query := map[string]interface{}{
  "query": map[string]interface{}{
    "bool": map[string]interface{}{
      "must": []interface{}{
        map[string]interface{}{
          "bool": map[string]interface{}{
            "should": []interface{}{
              map[string]interface{}{
                "term": map[string]interface{}{
                  "doc.id": id,
                },
              },
              map[string]interface{}{
                "term": map[string]interface{}{
                  "file.fileId": id,
                },
              },
            },
          },
        },
      },
      "filter": []interface{}{
        map[string]interface{}{
          "terms": map[string]interface{}{
            "type": []string{
              "DOC", "FILE",
            },
          },
        },
      },
    },
  },
}
```

```go
body := es.NewQuery(
    es.Bool().
        Must(
            es.Bool().
                Should(
                    es.Term("doc.id", id),
                    es.Term("file.fileId", id),
                ), 
        ).
        Filter(
            es.Terms("type", "DOC", "FILE"),
        ),
)
```

[//]: # ([![Coverage Status][cov-img]][cov])

[doc-img]: https://godoc.org/github.com/GokselKUCUKSAHIN/es-query-builder?status.svg
[doc]: https://godoc.org/github.com/GokselKUCUKSAHIN/es-query-builder
[go-report-img]: https://goreportcard.com/badge/github.com/GokselKUCUKSAHIN/es-query-builder
[go-report]: https://goreportcard.com/report/github.com/GokselKUCUKSAHIN/es-query-builder

[cov-img]: https://codecov.io/gh/GokselKUCUKSAHIN/es-query-builder/branch/main/graph/badge.svg
[cov]: https://codecov.io/gh/GokselKUCUKSAHIN/es-query-builder