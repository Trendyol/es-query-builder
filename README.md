# es-query-builder [![GoDoc][doc-img]][doc] [![Build Status][ci-img]][ci] [![Go Report Card][go-report-img]][go-report] [![Coverage Status][cov-img]][cov]
A simple, user-friendly, and streamlined library for programmatically building Elasticsearch DSL queries in Go, designed for low overhead and minimal memory usage.

## How to Get
To install the es-query-builder library, run the following command:
```bash
go get github.com/GokselKUCUKSAHIN/es-query-builder
```

### Examples 

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
query := es.NewQuery(
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

# Benchmarks
You can check and run [benchmarks](./benchmarks) on your machine.

### ARMv6l
- **Device**: Raspberry Pi Zero W
- **CPU**: Broadcom BCM2835 1GHz 1 Core
- **Arch**: ARM v6 32 bit
- **Memory**: 512MB LPDDR2
- **Go Version**: go1.22.3

![armv6l](https://github.com/GokselKUCUKSAHIN/es-query-builder/assets/33639948/8972003d-9b00-4021-9f69-347723ac59de)

### ARM64
- **Device**: MacBook Pro 16" 2021
- **CPU**: Apple Silicon M1 Pro 10 Core
- **Arch**: ARM64
- **Memory**: 32GB LPDDR5
- **Go Version**: go1.22.1

![arm64](https://github.com/GokselKUCUKSAHIN/es-query-builder/assets/33639948/ca9e2603-ebcd-4dec-92f4-e501ddcc4abe)

# License
MIT - Please check the [LICENSE](./LICENSE) file for full text.

[doc-img]: https://godoc.org/github.com/GokselKUCUKSAHIN/es-query-builder?status.svg
[doc]: https://godoc.org/github.com/GokselKUCUKSAHIN/es-query-builder
[go-report-img]: https://goreportcard.com/badge/github.com/GokselKUCUKSAHIN/es-query-builder
[go-report]: https://goreportcard.com/report/github.com/GokselKUCUKSAHIN/es-query-builder
[cov-img]: https://codecov.io/gh/GokselKUCUKSAHIN/es-query-builder/branch/main/graph/badge.svg
[cov]: https://codecov.io/gh/GokselKUCUKSAHIN/es-query-builder
[ci-img]: https://github.com/GokselKUCUKSAHIN/es-query-builder/actions/workflows/build-test.yml/badge.svg
[ci]: https://github.com/GokselKUCUKSAHIN/es-query-builder/actions/workflows/build-test.yml