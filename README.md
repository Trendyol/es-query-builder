# es-query-builder
Simple Query builder for Elasticsearch

# TODOs

## Project
- [ ] improve README
  - [ ] add examples
  - [ ] add benchmark results
  - [ ] add `go get` command
- [ ] add benchmark
- [ ] add missing tests
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
body := es.New()
body.
  Query().
  Bool().
  Must(
    es.Bool().
      Should(
        es.Term("doc.id", id),
        es.Term("file.fileId", id),
      ).Build(),
  ).
  Filter(
    es.Terms("type", "DOC", "FILE"),
  )
```