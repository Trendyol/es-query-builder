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
- **es-query-builder Version**: v0.4.2
- **Benchmark Date**: 02/01/2025

![arm64 combined](https://github.com/user-attachments/assets/61a38526-7ec7-47b2-85c8-037f5394acd8)

<details>
  <summary><b>ARM64 Detailed Benchmark Results</b></summary>

![arm64 simple](https://github.com/user-attachments/assets/6f8244fc-702c-4570-bef4-0e44a2514131)

- **es-query-builder** is 32% less efficient than **vanilla Go**.
- **[aquasecurity/esquery](https://github.com/aquasecurity/esquery)** is 82% less efficient than **es-query-builder**.
- **[defensestation/osquery](https://github.com/defensestation/osquery)** is 84% less efficient than **es-query-builder**.

Benchmark test file at [simple query benchmark](./benchmarks/tests/simple_benchmark_test.go)

---

![arm64 intermediate](https://github.com/user-attachments/assets/a31b01dc-8c49-4621-8993-303c0df0cf22)

- **es-query-builder** is 25% less efficient than **vanilla Go**.
- **[aquasecurity/esquery](https://github.com/aquasecurity/esquery)** is 75% less efficient than **es-query-builder**.
- **[defensestation/osquery](https://github.com/defensestation/osquery)** is 77% less efficient than **es-query-builder**.

Benchmark test file at [intermediate query benchmark](./benchmarks/tests/intermediate_benchmark_test.go)

---

![arm64 complex](https://github.com/user-attachments/assets/c3115f6b-35b5-422a-81e8-6af203c42ebb)

- **es-query-builder** is 31% less efficient than **vanilla Go**.
- **[aquasecurity/esquery](https://github.com/aquasecurity/esquery)** is 71% less efficient than **es-query-builder**.
- **[defensestation/osquery](https://github.com/defensestation/osquery)** is 71% less efficient than **es-query-builder**.

Benchmark test file at [complex query benchmark](./benchmarks/tests/complex_benchmark_test.go)

---

![arm64 mixed](https://github.com/user-attachments/assets/6b6a7a98-ff88-4be5-bc4f-2d6970553d8b)

- **es-query-builder** is 16% less efficient than **vanilla Go**.
- **[aquasecurity/esquery](https://github.com/aquasecurity/esquery)** is 66% less efficient than **es-query-builder**.
- **[defensestation/osquery](https://github.com/defensestation/osquery)** is 66% less efficient than **es-query-builder**.

Benchmark test file at [mixed query benchmark](./benchmarks/tests/mixed_benchmark_test.go)

---

![arm64 conditional](https://github.com/user-attachments/assets/be44f4f2-1b33-4f05-b049-08bfc78ae283)

- **es-query-builder** is 28% less efficient than **vanilla Go**.
- **[aquasecurity/esquery](https://github.com/aquasecurity/esquery)** is 71% less efficient than **es-query-builder**.
- **[defensestation/osquery](https://github.com/defensestation/osquery)** is 72% less efficient than **es-query-builder**.

Benchmark test file at [conditional query benchmark](./benchmarks/tests/conditional_benchmark_test.go)

---

![arm64 multi filter](https://github.com/user-attachments/assets/5956c4f5-4f85-436f-ae83-ec0dfda9c170)

- **es-query-builder** is 24% less efficient than **vanilla Go**.
- **[aquasecurity/esquery](https://github.com/aquasecurity/esquery)** is 73% less efficient than **es-query-builder**.
- **[defensestation/osquery](https://github.com/defensestation/osquery)** is 74% less efficient than **es-query-builder**.

Benchmark test file at [multi filter query benchmark](./benchmarks/tests/multi_filter_benchmark_test.go)

---

![arm64 aggs](https://github.com/user-attachments/assets/7c6434ab-fa1e-4fd0-99e1-67429168e8ac)

- **es-query-builder** is 29% less efficient than **vanilla Go**.
- **[aquasecurity/esquery](https://github.com/aquasecurity/esquery)** is 67% less efficient than **es-query-builder**.
- **[defensestation/osquery](https://github.com/defensestation/osquery)** is 68% less efficient than **es-query-builder**.

Benchmark test file at [aggs query benchmark](./benchmarks/tests/aggs_benchmark_test.go)

---

### MacBook M1 Pro 10 Core Benchmark Result Table 

|Benchmark Name     |vanilla go score|vanilla go ns/op|aquasecurity/esquery score|aquasecurity/esquery ns/op|defensestation/osquery score|defensestation/osquery ns/op|es-query-builder score|es-query-builder ns/op|
|-------------------|----------------|----------------|--------------------------|--------------------------|----------------------------|----------------------------|----------------------|----------------------|
|simple             |16289468        |368             |1949335                   |3079                      |1747318                     |3425                        |11026227              |533                   |
|simple             |16274949        |368             |1948286                   |3077                      |1748682                     |3431                        |11158036              |536                   |
|simple             |16157809        |370             |1944097                   |3082                      |1747022                     |3446                        |10935156              |536                   |
|simple             |16272403        |370             |1945136                   |3082                      |1742745                     |3439                        |11048505              |538                   |
|simple             |16122024        |370             |1948617                   |3075                      |1746759                     |3433                        |11024264              |533                   |
|simple avg         |16223330,60     |369,20          |1947094,20                |3079,00                   |1746505,20                  |3434,80                     |11038437,60           |535,20                |
|simple median      |16272403,00     |370,00          |1948286,00                |3079,00                   |1747022,00                  |3433,00                     |11026227,00           |536,00                |
|simple stddev      |69286,11        |0,98            |2077,41                   |2,76                      |1993,47                     |7,17                        |71305,27              |1,94                  |
|                   |                |                |                          |                          |                            |                            |                      |                      |
|complex            |2328243         |2575            |471954                    |12725                     |426900                      |14162                       |1609507               |3705                  |
|complex            |2336179         |2565            |470151                    |12656                     |428108                      |14103                       |1627707               |3705                  |
|complex            |2343412         |2573            |477055                    |12630                     |426849                      |14037                       |1622174               |3698                  |
|complex            |2340090         |2563            |474279                    |12627                     |427857                      |14038                       |1619374               |3695                  |
|complex            |2333827         |2571            |474705                    |12645                     |425506                      |14062                       |1625364               |3702                  |
|complex avg        |2336350,20      |2569,40         |473628,80                 |12656,60                  |427044,00                   |14080,40                    |1620825,20            |3701,00               |
|complex median     |2336179,00      |2571,00         |474279,00                 |12645,00                  |426900,00                   |14062,00                    |1622174,00            |3702,00               |
|complex stddev     |5214,93         |4,63            |2375,70                   |35,77                     |918,40                      |47,31                       |6324,35               |3,95                  |
|                   |                |                |                          |                          |                            |                            |                      |                      |
|conditional        |3983532         |1491            |836204                    |7168                      |793478                      |7522                        |2894722               |2087                  |
|conditional        |3918195         |1534            |845427                    |7106                      |805912                      |7486                        |2884924               |2121                  |
|conditional        |4044111         |1487            |840529                    |7157                      |794827                      |7545                        |2898068               |2081                  |
|conditional        |4048710         |1488            |846610                    |7125                      |804211                      |7468                        |2881280               |2082                  |
|conditional        |4039155         |1484            |840256                    |7105                      |802660                      |7475                        |2891826               |2072                  |
|conditional avg    |4006740,60      |1496,80         |841805,20                 |7132,20                   |800217,60                   |7499,20                     |2890164,00            |2088,60               |
|conditional median |4039155,00      |1488,00         |840529,00                 |7125,00                   |802660,00                   |7486,00                     |2891826,00            |2082,00               |
|conditional stddev |50174,96        |18,73           |3784,33                   |25,98                     |5075,82                     |29,50                       |6203,13               |16,91                 |
|                   |                |                |                          |                          |                            |                            |                      |                      |
|intermediate       |4573310         |1312            |867562                    |6954                      |779059                      |7674                        |3395592               |1774                  |
|intermediate       |4567904         |1316            |861070                    |6973                      |771276                      |7714                        |3413900               |1769                  |
|intermediate       |4563678         |1319            |862593                    |6924                      |782931                      |7651                        |3408219               |1777                  |
|intermediate       |4556863         |1317            |867105                    |6904                      |783907                      |7647                        |3398814               |1770                  |
|intermediate       |4559426         |1320            |863877                    |6945                      |780218                      |7773                        |3416360               |1768                  |
|intermediate avg   |4564236,20      |1316,80         |864441,40                 |6940,00                   |779478,20                   |7691,80                     |3406577,00            |1771,60               |
|intermediate median|4563678,00      |1317,00         |863877,00                 |6945,00                   |780218,00                   |7674,00                     |3408219,00            |1770,00               |
|intermediate stddev|5892,37         |2,79            |2527,23                   |23,92                     |4461,73                     |47,06                       |8160,44               |3,38                  |
|                   |                |                |                          |                          |                            |                            |                      |                      |
|mixed              |3452348         |1739            |1000000                   |5576                      |1000000                     |5519                        |2926890               |2057                  |
|mixed              |3467690         |1734            |1000000                   |5143                      |1000000                     |5525                        |2930038               |2052                  |
|mixed              |3489582         |1726            |1000000                   |5224                      |1000000                     |5520                        |2920724               |2053                  |
|mixed              |3465878         |1728            |1000000                   |5154                      |1000000                     |5541                        |2935546               |2053                  |
|mixed              |3474674         |1734            |1000000                   |5141                      |1000000                     |5513                        |2932101               |2056                  |
|mixed avg          |3470034,40      |1732,20         |1000000,00                |5247,60                   |1000000,00                  |5523,60                     |2929059,80            |2054,20               |
|mixed median       |3467690,00      |1734,00         |1000000,00                |5154,00                   |1000000,00                  |5520,00                     |2930038,00            |2053,00               |
|mixed stddev       |12159,07        |4,66            |0,00                      |167,01                    |0,00                        |9,50                        |5029,32               |1,94                  |
|                   |                |                |                          |                          |                            |                            |                      |                      |
|multi filter       |4035531         |1481            |841471                    |7113                      |797647                      |7482                        |3064244               |1964                  |
|multi filter       |4089892         |1474            |836565                    |7091                      |800364                      |7492                        |3077422               |1961                  |
|multi filter       |4062204         |1477            |840290                    |7105                      |801840                      |7493                        |3064178               |1957                  |
|multi filter       |4034770         |1481            |838080                    |7136                      |795669                      |7491                        |3077875               |1967                  |
|multi filter       |4026196         |1474            |848299                    |7074                      |806515                      |7483                        |3078358               |1962                  |
|multi filter avg   |4049718,60      |1477,40         |840941,00                 |7103,80                   |800407,00                   |7488,20                     |3072415,40            |1962,20               |
|multi filter median|4035531,00      |1477,00         |840290,00                 |7105,00                   |800364,00                   |7491,00                     |3077422,00            |1962,00               |
|multi filter stdev |23442,90        |3,14            |4054,12                   |20,86                     |3726,37                     |4,71                        |6705,44               |3,31                  |
|                   |                |                |                          |                          |                            |                            |                      |                      |
|aggs               |2591437         |2314            |619117                    |9737                      |604124                      |10140                       |1859446               |3232                  |
|aggs               |2625234         |2291            |611305                    |9711                      |586738                      |10074                       |1872366               |3220                  |
|aggs               |2622649         |2292            |613981                    |9707                      |596095                      |10684                       |1871360               |3216                  |
|aggs               |2626522         |2284            |605070                    |9692                      |598178                      |10066                       |1857014               |3209                  |
|aggs               |2623935         |2295            |613622                    |9687                      |593932                      |10063                       |1858818               |3224                  |
|aggs avg           |2617955,40      |2295,20         |612619,00                 |9706,80                   |595813,40                   |10205,40                    |1863800,80            |3220,20               |
|aggs median        |2623935,00      |2292,00         |613622,00                 |9707,00                   |596095,00                   |10074,00                    |1859446,00            |3220,00               |
|aggs stddev        |13321,98        |10,07           |4556,29                   |17,55                     |5668,14                     |240,96                      |6638,64               |7,70                  |

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
