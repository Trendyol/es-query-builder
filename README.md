# es-query-builder [![GoDoc][doc-img]][doc] [![Release][release-img]][release] [![Build Status][ci-img]][ci] [![Go Report Card][go-report-img]][go-report] [![Coverage Status][cov-img]][cov]

A simple, user-friendly, and streamlined library for programmatically building Elasticsearch DSL queries in Go, designed
for low overhead and minimal memory usage.

## Install
With [Go's module support](https://go.dev/wiki/Modules#how-to-use-modules), `go [build|run|test]` automatically fetches the necessary dependencies when you add the import in your code:

```sh
import "github.com/Trendyol/es-query-builder"
```

Alternatively, use `go get`:

```sh
go get -u github.com/Trendyol/es-query-builder
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

### ARM64

- **Device**: MacBook Pro 16" 2021
- **CPU**: Apple Silicon M1 Pro 10 Core
- **Arch**: ARM64
- **Memory**: 32GB LPDDR5
- **Go Version**: go1.23.5
- **es-query-builder Version**: v0.4.0
- **Benchmark Date**: 01/22/2025

![arm64 combined](https://github.com/user-attachments/assets/3d462d23-b9be-4e6b-82c8-ba8bc40de241)

<details>
  <summary><b>ARM64 Detailed Benchmark Results</b></summary>

![arm64 simple](https://github.com/user-attachments/assets/dcb1303b-d384-424a-9f79-41369d3c2b82)

- **es-query-builder** is 23% less efficient than **vanilla Go**.
- **[aquasecurity/esquery](https://github.com/aquasecurity/esquery)** is 82% less efficient than **es-query-builder**.

Benchmark test file at [simple query benchmark](./benchmarks/simple_example_test.go)

---

![arm64 intermediate](https://github.com/user-attachments/assets/63cc99de-7590-4266-be3d-0b9fc9dce66e)

- **es-query-builder** is 25% less efficient than **vanilla Go**.
- **[aquasecurity/esquery](https://github.com/aquasecurity/esquery)** is 74% less efficient than **es-query-builder**.

Benchmark test file at [intermediate query benchmark](./benchmarks/intermediate_example_test.go)

---

![arm64 complex](https://github.com/user-attachments/assets/94c364c7-d0f8-4fba-ab7c-ecce9f790c4f)

- **es-query-builder** is 30% less efficient than **vanilla Go**.
- **[aquasecurity/esquery](https://github.com/aquasecurity/esquery)** is 70% less efficient than **es-query-builder**.

Benchmark test file at [complex query benchmark](./benchmarks/complex_example_test.go)

---

![arm64 mixed](https://github.com/user-attachments/assets/507d3d9c-dbb3-44c3-b052-ff46a4b11b5e)

- **es-query-builder** is 16% less efficient than **vanilla Go**.
- **[aquasecurity/esquery](https://github.com/aquasecurity/esquery)** is 82% less efficient than **es-query-builder**.

Benchmark test file at [mixed query benchmark](./benchmarks/mixed_example_test.go)

---

![arm64 conditional](https://github.com/user-attachments/assets/4b0bd815-eb01-4bd6-8c3d-4840f5150291)

- **es-query-builder** is 32% less efficient than **vanilla Go**.
- **[aquasecurity/esquery](https://github.com/aquasecurity/esquery)** is 69% less efficient than **es-query-builder**.

Benchmark test file at [conditional query benchmark](./benchmarks/conditional_example_test.go)

---

![arm64 aggs](https://github.com/user-attachments/assets/1e115a3d-6a38-4796-8d3c-e3b1bd67b3bf)

- **es-query-builder** is 29% less efficient than **vanilla Go**.
- **[aquasecurity/esquery](https://github.com/aquasecurity/esquery)** is 81% less efficient than **es-query-builder**.

Benchmark test file at [aggs query benchmark](./benchmarks/aggs_example_test.go)

---

### MacBook M1 Pro 10 Core Benchmark Result Table 

|Benchmark Name     |vanilla go score|vanilla go ns/op|aquasecurity/esquery score|aquasecurity/esquery ns/op|es-query-builder score|es-query-builder ns/op|
|-------------------|----------------|----------------|--------------------------|--------------------------|----------------------|----------------------|
|simple             |15997112        |372,4           |1953010                   |3075                      |10861778              |541,4                 |
|simple             |15932286        |376,7           |1942562                   |3079                      |10861311              |537,7                 |
|simple             |15981565        |376,8           |1952407                   |3074                      |10896735              |539,2                 |
|simple             |16174654        |373,9           |1945106                   |3253                      |10965974              |538,5                 |
|simple             |16186130        |372,6           |1943928                   |3069                      |11010657              |538,7                 |
|simple avg         |16054349,40     |374,48          |1947402,60                |3110,00                   |10919291,00           |539,10                |
|simple median      |15997112,00     |373,90          |1945106,00                |3075,00                   |10896735,00           |538,70                |
|simple stddev      |105178,41       |1,92            |4410,57                   |71,57                     |59506,72              |1,25                  |
|                   |                |                |                          |                          |                      |                      |
|complex            |2324410         |2572            |476419                    |12588                     |1612232               |3708                  |
|complex            |2315541         |2576            |472420                    |12610                     |1612942               |3714                  |
|complex            |2321216         |2582            |473491                    |12608                     |1603651               |3705                  |
|complex            |2325516         |2585            |478119                    |12642                     |1616520               |3711                  |
|complex            |2327028         |2586            |471618                    |12594                     |1615902               |3715                  |
|complex avg        |2322742,20      |2580,20         |474413,40                 |12608,40                  |1612249,40            |3710,60               |
|complex median     |2324410,00      |2582,00         |473491,00                 |12608,00                  |1612942,00            |3711,00               |
|complex stddev     |4075,05         |5,38            |2465,47                   |18,74                     |4604,17               |3,72                  |
|                   |                |                |                          |                          |                      |                      |
|conditional        |4023123         |1491            |845757                    |7077                      |2885863               |2076                  |
|conditional        |4041091         |1488            |853201                    |7091                      |2853171               |2078                  |
|conditional        |4012322         |1503            |841545                    |7081                      |2866297               |2076                  |
|conditional        |4013662         |1488            |847909                    |7104                      |2890788               |2082                  |
|conditional        |4026201         |1496            |840610                    |7090                      |2883222               |2092                  |
|conditional avg    |4023279,80      |1493,20         |845804,40                 |7088,60                   |2875868,20            |2080,80               |
|conditional median |4023123,00      |1491,00         |845757,00                 |7090,00                   |2883222,00            |2078,00               |
|conditional stddev |10376,63        |5,71            |4566,52                   |9,35                      |14023,88              |6,01                  |
|                   |                |                |                          |                          |                      |                      |
|intermediate       |4519802         |1323            |867354                    |6914                      |3373606               |1773                  |
|intermediate       |4488582         |1333            |864675                    |6904                      |3372969               |1771                  |
|intermediate       |4461734         |1335            |864476                    |6903                      |3391429               |1770                  |
|intermediate       |4476802         |1335            |865116                    |6901                      |3345444               |1766                  |
|intermediate       |4487874         |1335            |857962                    |6910                      |3358254               |1777                  |
|intermediate avg   |4486958,80      |1332,20         |863916,60                 |6906,40                   |3368340,40            |1771,40               |
|intermediate median|4487874,00      |1335,00         |864675,00                 |6904,00                   |3372969,00            |1771,00               |
|intermediate stddev|19087,16        |4,66            |3149,59                   |4,84                      |15544,87              |3,61                  |
|                   |                |                |                          |                          |                      |                      |
|mixed              |3430264         |1733            |1000000                   |5158                      |2880654               |2074                  |
|mixed              |3468865         |1737            |1000000                   |5167                      |2927622               |2057                  |
|mixed              |3445854         |1741            |1000000                   |5182                      |2892280               |2074                  |
|mixed              |3453601         |1730            |1000000                   |5158                      |2887956               |2064                  |
|mixed              |3445731         |1733            |1000000                   |5157                      |2901133               |2056                  |
|mixed avg          |3448863,00      |1734,80         |1000000,00                |5164,40                   |2897929,00            |2065,00               |
|mixed median       |3445854,00      |1733,00         |1000000,00                |5158,00                   |2892280,00            |2064,00               |
|mixed stddev       |12548,43        |3,82            |0,00                      |9,52                      |16258,77              |7,85                  |
|                   |                |                |                          |                          |                      |                      |
|aggs               |2604656         |2296            |618597                    |9704                      |1856065               |3279                  |
|aggs               |2616678         |2289            |620818                    |9719                      |1861394               |3228                  |
|aggs               |2616700         |2300            |618390                    |9711                      |1866285               |3235                  |
|aggs               |2613950         |2295            |615925                    |9705                      |1855579               |3211                  |
|aggs               |2611999         |2300            |617900                    |9697                      |1857831               |3216                  |
|aggs avg           |2612796,60      |2296,00         |618326,00                 |9707,20                   |1859430,80            |3233,80               |
|aggs median        |2613950,00      |2296,00         |618390,00                 |9705,00                   |1857831,00            |3228,00               |
|aggs stddev        |4439,32         |4,05            |1564,22                   |7,39                      |3988,95               |24,14                 |


</details>


# Want to Contribute?

<details>
  <summary><b>Join Us</b></summary>
  <img src="https://github.com/user-attachments/assets/34bb6fc2-237b-49df-bae9-8ce2b14096ca" width="400px" alt="join us"/>
</details>

###  Contribute to Our Project

Want to help out? Awesome! Here’s how you can contribute:

1. **Report Issues:** Got a suggestion, recommendation, or found a bug? Head over to the [Issues](https://github.com/Trendyol/es-query-builder/issues) section and let us know.

2. **Make Changes:** Want to improve the code?
   - Fork the repo
   - Create a new branch
   - Make your changes
   - Open a Pull Request (PR)

We’re excited to see your contributions. Thanks for helping make this project better!

# License

MIT - Please check the [LICENSE](./LICENSE) file for full text.

[doc-img]: https://godoc.org/github.com/Trendyol/es-query-builder?status.svg

[doc]: https://godoc.org/github.com/Trendyol/es-query-builder

[release]: https://github.com/Trendyol/es-query-builder/releases

[release-img]: https://img.shields.io/github/v/release/Trendyol/es-query-builder.svg

[go-report-img]: https://goreportcard.com/badge/github.com/Trendyol/es-query-builder

[go-report]: https://goreportcard.com/report/github.com/Trendyol/es-query-builder

[cov-img]: https://codecov.io/gh/Trendyol/es-query-builder/branch/main/graph/badge.svg

[cov]: https://codecov.io/gh/Trendyol/es-query-builder

[ci-img]: https://github.com/Trendyol/es-query-builder/actions/workflows/build-test.yml/badge.svg

[ci]: https://github.com/Trendyol/es-query-builder/actions/workflows/build-test.yml
