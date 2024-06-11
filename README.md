# es-query-builder [![GoDoc][doc-img]][doc] [![Build Status][ci-img]][ci] [![Go Report Card][go-report-img]][go-report] [![Coverage Status][cov-img]][cov]

A simple, user-friendly, and streamlined library for programmatically building Elasticsearch DSL queries in Go, designed
for low overhead and minimal memory usage.

## How to Get

To install the es-query-builder library, run the following command:

```bash
go get github.com/GokselKUCUKSAHIN/es-query-builder
```

### Example
```json
{
  "query": {
    "bool": {
      "must": [
        {
          "term": {
            "author": "George Orwell"
          }
        }
      ],
      "must_not": [
        {
          "terms": {
            "genre": [
              "Fantasy",
              "Science Fiction"
            ]
          }
        },
        {
          "exists": {
            "field": "out_of_print"
          }
        }
      ],
      "should": [
        {
          "terms": {
            "title": [
              "1984",
              "Animal Farm"
            ]
          }
        }
      ]
    }
  },
  "aggs": {
    "genres_count": {
      "terms": {
        "field": "genre"
      }
    },
    "authors_and_genres": {
      "terms": {
        "field": "author"
      },
      "aggs": {
        "genres": {
          "terms": {
            "field": "genre"
          }
        }
      }
    }
  }
}
```

### With es-query-builder

```go
query := es.NewQuery(
    es.Bool().
        Must(
            es.Term("author", "George Orwell"),
        ).
        MustNot(
            es.Terms("genre", "Fantasy", "Science Fiction"),
            es.Exists("out_of_print"),
        ).
        Should(
            es.Terms("title", "1984", "Animal Farm"),
        ),
).Aggs("genres_count",
    es.AggTerms().
        Field("genre"),
).Aggs("authors_and_genres",
    es.AggTerms().
        Field("author").
        Aggs("genres",
            es.AggTerms().
                Field("genre"),
        ),
)
```

### With vanilla Go

```go
query := map[string]interface{}{
  "query": map[string]interface{}{
    "bool": map[string]interface{}{
      "must": []map[string]interface{}{
        {
          "term": map[string]interface{}{
            "author": "George Orwell",
          },
        },
      },
      "must_not": []map[string]interface{}{
        {
          "terms": map[string]interface{}{
            "genre": []string{
              "Fantasy",
              "Science Fiction",
            },
          },
        },
        {
          "exists": map[string]interface{}{
            "field": "out_of_print",
          },
        },
      },
      "should": []map[string]interface{}{
        {
          "terms": map[string]interface{}{
            "title": []string{
              "1984",
              "Animal Farm",
            },
          },
        },
      },
    },
  },
  "aggs": map[string]interface{}{
    "genres_count": map[string]interface{}{
      "terms": map[string]interface{}{
        "field": "genre",
      },
    },
    "authors_and_genres": map[string]interface{}{
      "terms": map[string]interface{}{
        "field": "author",
      },
      "aggs": map[string]interface{}{
        "genres": map[string]interface{}{
          "terms": map[string]interface{}{
            "field": "genre",
          },
        },
      },
    },
  },
}
```



# Benchmarks

You can check and run [benchmarks](./benchmarks) on your machine.

### ARMv6l

- **Device**: Raspberry Pi Zero W
- **CPU**: Broadcom BCM2835 1GHz Single Core
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

# Want to Contribute?

<details>
  <summary><b>Join Us</b></summary>
  <img src="https://github.com/GokselKUCUKSAHIN/es-query-builder/assets/33639948/bc696d14-a55d-4ec4-9cb4-021cc4128760" width="400px" alt="join us"/>
</details>

###  Contribute to Our Project

Want to help out? Awesome! Here’s how you can contribute:

1. **Report Issues:** Got a suggestion, recommendation, or found a bug? Head over to the [Issues](https://github.com/GokselKUCUKSAHIN/es-query-builder/issues) section and let us know.

2. **Make Changes:** Want to improve the code?
   - Fork the repo
   - Create a new branch
   - Make your changes
   - Open a Pull Request (PR)

We’re excited to see your contributions. Thanks for helping make this project better!

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