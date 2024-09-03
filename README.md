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
- **Go Version**: go1.22.1

![arm64 combined](https://github.com/user-attachments/assets/eade143d-c31c-4caf-96f4-8005fa1b11bc)

<details>
  <summary><b>ARM64 Detailed Benchmark Results</b></summary>

![arm64 simple](https://github.com/user-attachments/assets/818c1381-5a31-47ab-bc94-5133b1713c38)

- **es-query-builder** is 23% less efficient than **vanilla Go**.
- **[aquasecurity/esquery](https://github.com/aquasecurity/esquery)** is 84% less efficient than **es-query-builder**.

Benchmark test file at [simple query benchmark](./benchmarks/simple_example_test.go)

---

![arm64 intermediate](https://github.com/user-attachments/assets/d2c72cc2-27d2-4e0d-908a-b49bf8fd7f9d)

- **es-query-builder** is 24% less efficient than **vanilla Go**.
- **[aquasecurity/esquery](https://github.com/aquasecurity/esquery)** is 75% less efficient than **es-query-builder**.

Benchmark test file at [intermediate query benchmark](./benchmarks/intermediate_example_test.go)

---

![arm64 complex](https://github.com/user-attachments/assets/70dfff75-1e37-4c4f-b102-cc3a9900aa05)

- **es-query-builder** is 29% less efficient than **vanilla Go**.
- **[aquasecurity/esquery](https://github.com/aquasecurity/esquery)** is 71% less efficient than **es-query-builder**.

Benchmark test file at [complex query benchmark](./benchmarks/complex_example_test.go)

---

![arm64 mixed](https://github.com/user-attachments/assets/2b3778ea-500b-421b-96cc-18d2425ef4ac)

- **es-query-builder** is 19% less efficient than **vanilla Go**.
- **[aquasecurity/esquery](https://github.com/aquasecurity/esquery)** is 64% less efficient than **es-query-builder**.

Benchmark test file at [mixed query benchmark](./benchmarks/mixed_example_test.go)

---

![arm64 conditional](https://github.com/user-attachments/assets/d0dd2e69-4169-48a0-9e9b-5cd85e33ebe3)

- **es-query-builder** is 32% less efficient than **vanilla Go**.
- **[aquasecurity/esquery](https://github.com/aquasecurity/esquery)** is 69% less efficient than **es-query-builder**.

Benchmark test file at [conditional query benchmark](./benchmarks/conditional_example_test.go)

---

![arm64 aggs](https://github.com/user-attachments/assets/c102d174-1b50-4b1c-91d2-d50f7ab2aed3)

- **es-query-builder** is 23% less efficient than **vanilla Go**.
- **[aquasecurity/esquery](https://github.com/aquasecurity/esquery)** is 69% less efficient than **es-query-builder**.

Benchmark test file at [aggs query benchmark](./benchmarks/aggs_example_test.go)

---

### MacBook M1 Pro 10 Core Benchmark Result Table 

|Benchmark Name      |vanilla go score    |vanilla go ns/op|aquasecurity/esquery score|aquasecurity/esquery ns/op|es-query-builder score|es-query-builder ns/op|
|--------------------|--------------------|----------------|--------------------------|--------------------------|----------------------|----------------------|
|simple              |16002007            |376,6           |1935308                   |3099                      |12279682              |486,8                 |
|simple              |15991658            |376,7           |1935002                   |3100                      |12303226              |486,5                 |
|simple              |16034337            |373,5           |1935006                   |3086                      |12072054              |485                   |
|simple              |15873903            |374,8           |1942242                   |3091                      |12350944              |484,7                 |
|simple              |15957768            |374,6           |1941484                   |3091                      |12305442              |484,4                 |
|simple              |15948505            |375,2           |1939986                   |3092                      |12294543              |486,5                 |
|simple              |16036609            |374,1           |1941672                   |3094                      |12399751              |483,3                 |
|simple              |15903747            |374,5           |1944676                   |3094                      |12381858              |484,5                 |
|simple              |15937150            |376,6           |1942200                   |3083                      |12333574              |484,4                 |
|simple              |15887662            |376             |1941118                   |3089                      |12324646              |486,2                 |
|simple avg          |15957334,60         |375,26          |1939869,40                |3091,90                   |12304572,00           |485,23                |
|simple median       |15953136,50         |375,00          |1941301,00                |3091,50                   |12315044,00           |484,85                |
|simple stddev       |55280,94            |1,09            |3312,76                   |4,99                      |85565,12              |1,12                  |
|                    |                    |                |                          |                          |                      |                      |
|complex             |2295715             |2612            |468993                    |12791                     |1627998               |3697                  |
|complex             |2293550             |2614            |468337                    |12758                     |1623253               |3699                  |
|complex             |2308629             |2599            |468853                    |12693                     |1627789               |3687                  |
|complex             |2303484             |2604            |471734                    |12701                     |1638758               |3672                  |
|complex             |2301466             |2633            |447708                    |12957                     |1628677               |3706                  |
|complex             |2314368             |2606            |466209                    |12749                     |1638372               |3674                  |
|complex             |2309028             |2604            |471897                    |12732                     |1630087               |3692                  |
|complex             |2304069             |2611            |468198                    |12724                     |1626214               |3686                  |
|complex             |2309613             |2601            |471686                    |12733                     |1630682               |3681                  |
|complex             |2297032             |2623            |468930                    |12731                     |1632646               |3676                  |
|complex avg         |2303695,40          |2610,70         |467254,50                 |12756,90                  |1630447,60            |3687,00               |
|complex median      |2303776,50          |2608,50         |468891,50                 |12732,50                  |1629382,00            |3686,50               |
|complex stddev      |6464,29             |10,02           |6744,35                   |71,71                     |4719,81               |10,87                 |
|                    |                    |                |                          |                          |                      |                      |
|conditional         |4013814             |1506            |833030                    |7156                      |2696304               |2223                  |
|conditional         |3960637             |1512            |833611                    |7170                      |2705523               |2226                  |
|conditional         |3937759             |1516            |832034                    |7166                      |2697906               |2220                  |
|conditional         |3977565             |1511            |838292                    |7145                      |2707563               |2220                  |
|conditional         |3986996             |1504            |824229                    |7140                      |2713401               |2219                  |
|conditional         |3961573             |1503            |828835                    |7132                      |2700866               |2216                  |
|conditional         |4001875             |1499            |823173                    |7135                      |2710687               |2218                  |
|conditional         |3974684             |1514            |829016                    |7139                      |2705721               |2221                  |
|conditional         |3995692             |1503            |828115                    |7135                      |2716176               |2217                  |
|conditional         |3996382             |1505            |828490                    |7174                      |2713070               |2220                  |
|conditional avg     |3980697,70          |1507,30         |829882,50                 |7149,20                   |2706721,70            |2220,00               |
|conditional median  |3982280,50          |1505,50         |828925,50                 |7142,50                   |2706642,00            |2220,00               |
|conditional stddev  |21710,73            |5,29            |4278,41                   |15,10                     |6431,96               |2,76                  |
|                    |                    |                |                          |                          |                      |                      |
|intermediate        |4411344             |1354            |829754                    |6984                      |3343968               |1780                  |
|intermediate        |4406073             |1357            |856807                    |6983                      |3420316               |1770                  |
|intermediate        |4480772             |1345            |850144                    |6959                      |3401730               |1761                  |
|intermediate        |4447161             |1347            |850741                    |6949                      |3417213               |1762                  |
|intermediate        |4464565             |1344            |847190                    |6939                      |3408784               |1765                  |
|intermediate        |4515195             |1329            |848419                    |6967                      |3409892               |1763                  |
|intermediate        |4524844             |1331            |855118                    |6953                      |3418483               |1758                  |
|intermediate        |4454905             |1350            |842991                    |6968                      |3398035               |1765                  |
|intermediate        |4447119             |1347            |847276                    |6951                      |3418030               |1765                  |
|intermediate        |4529850             |1328            |843906                    |6948                      |3420828               |1760                  |
|intermediate avg    |4468182,80          |1343,20         |847234,60                 |6960,10                   |3405727,90            |1764,90               |
|intermediate median |4459735,00          |1346,00         |847847,50                 |6956,00                   |3413552,50            |1764,00               |
|intermediate stddev |41887,03            |9,84            |7145,66                   |14,33                     |21915,51              |5,94                  |
|                    |                    |                |                          |                          |                      |                      |
|mixed               |3397195             |1762            |1000000                   |5208                      |2762512               |2201                  |
|mixed               |3398920             |1754            |1000000                   |5216                      |2755569               |2201                  |
|mixed               |3437743             |1752            |1000000                   |5218                      |2747965               |2177                  |
|mixed               |3435486             |1752            |1000000                   |5212                      |2762694               |2172                  |
|mixed               |3423386             |1745            |1000000                   |5189                      |2767053               |2166                  |
|mixed               |3415612             |1751            |1000000                   |5207                      |2752192               |2179                  |
|mixed               |3413092             |1757            |1000000                   |5212                      |2758905               |2185                  |
|mixed               |3441566             |1764            |1000000                   |5208                      |2772634               |2172                  |
|mixed               |3412839             |1754            |1000000                   |5205                      |2757806               |2176                  |
|mixed               |3418503             |1753            |1000000                   |5202                      |2768888               |2167                  |
|mixed avg           |3419434,20          |1754,40         |1000000,00                |5207,70                   |2760621,80            |2179,60               |
|mixed median        |3417057,50          |1753,50         |1000000,00                |5208,00                   |2760708,50            |2176,50               |
|mixed stddev        |14535,88            |5,20            |0,00                      |7,76                      |7276,77               |11,93                 |
|                    |                    |                |                          |                          |                      |                      |
|aggs                |2469640             |2338            |606574                    |9845                      |1864624               |3215                  |
|aggs                |2586796             |2322            |606202                    |9839                      |1978429               |3038                  |
|aggs                |2587370             |2325            |606057                    |9822                      |1984966               |3037                  |
|aggs                |2593220             |2321            |608720                    |9785                      |1981557               |3015                  |
|aggs                |2600806             |2304            |611236                    |9762                      |1991346               |3003                  |
|aggs                |2594593             |2311            |606961                    |9777                      |1993276               |3027                  |
|aggs                |2587419             |2318            |596414                    |9801                      |1974360               |3019                  |
|aggs                |2590785             |2319            |609549                    |9784                      |1988269               |3021                  |
|aggs                |2596502             |2317            |603372                    |9774                      |1983280               |3024                  |
|aggs                |2579762             |2323            |603915                    |9790                      |2012487               |2996                  |
|aggs avg            |2578689,30          |2319,80         |605900,00                 |9797,90                   |1975259,40            |3039,50               |
|aggs median         |2589102,00          |2320,00         |606388,00                 |9787,50                   |1984123,00            |3022,50               |
|aggs stddev         |36779,07            |8,45            |3895,53                   |26,86                     |38193,95              |59,83                 |

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