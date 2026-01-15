# es-query-builder [![GoDoc][doc-img]][doc] [![Release][release-img]][release] [![Build Status][ci-img]][ci] [![Go Report Card][go-report-img]][go-report] [![Coverage Status][cov-img]][cov] [![OpenSSF Scorecard][scorecard-img]][scorecard]

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
        )).
	Aggs(
            es.Agg("genres_count", es.TermsAgg("genre")), 
            es.Agg("authors_and_genres", es.TermsAgg("author").
                Aggs(es.Agg("genres", es.TermsAgg("genre"))),
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
- **OS**: macOS Tahoe 26.2 
- **CPU**: Apple Silicon M1 Pro 10 Core
- **Arch**: ARM64
- **Memory**: 32GB LPDDR5
- **Go Version**: go1.25.1
- **es-query-builder Version**: v1.0.3
- **Benchmark Date**: 16/01/2026

![arm64 combined](./benchmarks/results/combined.png)

<details>
  <summary><b>ARM64 Detailed Benchmark Results</b></summary>

![arm64 simple](./benchmarks/results/simple.png)

- **es-query-builder** is **32%** less efficient than **vanilla Go**.
- **[aquasecurity/esquery](https://github.com/aquasecurity/esquery)** is **82%** less efficient than **es-query-builder**.
- **[defensestation/osquery](https://github.com/defensestation/osquery)** is **83%** less efficient than **es-query-builder**.

Benchmark test file at [simple query benchmark](./benchmarks/tests/simple_benchmark_test.go)

---

![arm64 intermediate](./benchmarks/results/intermediate.png)

- **es-query-builder** is **20%** less efficient than **vanilla Go**.
- **[aquasecurity/esquery](https://github.com/aquasecurity/esquery)** is **74%** less efficient than **es-query-builder**.
- **[defensestation/osquery](https://github.com/defensestation/osquery)** is **76%** less efficient than **es-query-builder**.

Benchmark test file at [intermediate query benchmark](./benchmarks/tests/intermediate_benchmark_test.go)

---

![arm64 complex](./benchmarks/results/complex.png)

- **es-query-builder** is **26%** less efficient than **vanilla Go**.
- **[aquasecurity/esquery](https://github.com/aquasecurity/esquery)** is **70%** less efficient than **es-query-builder**.
- **[defensestation/osquery](https://github.com/defensestation/osquery)** is **73%** less efficient than **es-query-builder**.

Benchmark test file at [complex query benchmark](./benchmarks/tests/complex_benchmark_test.go)

---

![arm64 mixed](./benchmarks/results/mixed.png)

- **es-query-builder** is **12%** less efficient than **vanilla Go**.
- **[aquasecurity/esquery](https://github.com/aquasecurity/esquery)** is **58%** less efficient than **es-query-builder**.
- **[defensestation/osquery](https://github.com/defensestation/osquery)** is **61%** less efficient than **es-query-builder**.

Benchmark test file at [mixed query benchmark](./benchmarks/tests/mixed_benchmark_test.go)

---

![arm64 conditional](./benchmarks/results/conditional.png)

- **es-query-builder** is **25%** less efficient than **vanilla Go**.
- **[aquasecurity/esquery](https://github.com/aquasecurity/esquery)** is **70%** less efficient than **es-query-builder**.
- **[defensestation/osquery](https://github.com/defensestation/osquery)** is **71%** less efficient than **es-query-builder**.

Benchmark test file at [conditional query benchmark](./benchmarks/tests/conditional_benchmark_test.go)

---

![arm64 multi filter](https://github.com/user-attachments/assets/5956c4f5-4f85-436f-ae83-ec0dfda9c170)

- **es-query-builder** is **23%** less efficient than **vanilla Go**.
- **[aquasecurity/esquery](https://github.com/aquasecurity/esquery)** is **71%** less efficient than **es-query-builder**.
- **[defensestation/osquery](https://github.com/defensestation/osquery)** is **72%** less efficient than **es-query-builder**.

Benchmark test file at [multi filter query benchmark](./benchmarks/tests/multi_filter_benchmark_test.go)

---

![arm64 aggs](https://github.com/user-attachments/assets/7c6434ab-fa1e-4fd0-99e1-67429168e8ac)

- **es-query-builder** is **18%** less efficient than **vanilla Go**.
- **[aquasecurity/esquery](https://github.com/aquasecurity/esquery)** is **64%** less efficient than **es-query-builder**.
- **[defensestation/osquery](https://github.com/defensestation/osquery)** is **66%** less efficient than **es-query-builder**.

Benchmark test file at [aggs query benchmark](./benchmarks/tests/aggs_benchmark_test.go)

---

### MacBook M1 Pro 10 Core Benchmark Result Table 

| Benchmark Name      | vanilla go score | vanilla go ns/op | aquasecurity/esquery score | aquasecurity/esquery ns/op | defensestation/osquery score | defensestation/osquery ns/op | es-query-builder score | es-query-builder ns/op |
| ------------------- | ---------------- | ---------------- | -------------------------- | -------------------------- | ---------------------------- | ---------------------------- | ---------------------- | ---------------------- |
| simple              | 7136586          | 335,7            | 910437                     | 2627                       | 810501                       | 2934                         | 4922604                | 474,6                  |
| simple              | 7047357          | 338,7            | 909422                     | 2635                       | 816651                       | 2949                         | 4899026                | 480,4                  |
| simple              | 7112192          | 333,1            | 915338                     | 2626                       | 792463                       | 2927                         | 4735447                | 475,4                  |
| simple              | 7095946          | 339,1            | 903325                     | 2621                       | 816793                       | 2984                         | 4868560                | 484,3                  |
| simple              | 7126017          | 338,3            | 902650                     | 2630                       | 819847                       | 2929                         | 4911927                | 483,6                  |
| simple              | 7145269          | 339,9            | 894154                     | 2635                       | 823344                       | 2933                         | 4817287                | 483,4                  |
| simple              | 7138833          | 336,1            | 912168                     | 2629                       | 800422                       | 2931                         | 4846564                | 476,6                  |
| simple              | 7110584          | 337,3            | 905252                     | 2626                       | 809718                       | 2920                         | 4804375                | 475,4                  |
| simple              | 7126405          | 337,8            | 914166                     | 2627                       | 813373                       | 2933                         | 5006870                | 478,7                  |
| simple              | 6878625          | 352              | 888853                     | 2678                       | 811065                       | 2991                         | 4775053                | 501,4                  |
| simple              | 7199766          | 335,3            | 917606                     | 2621                       | 821077                       | 2928                         | 4980991                | 475,8                  |
| simple              | 7179116          | 343,5            | 822660                     | 2780                       | 815586                       | 2927                         | 4947879                | 477,1                  |
| simple              | 7119448          | 338              | 898772                     | 2628                       | 817172                       | 2925                         | 4888240                | 477,2                  |
| simple              | 7096699          | 364,6            | 756648                     | 2803                       | 811969                       | 2989                         | 4635244                | 508,6                  |
| simple              | 6877939          | 348,4            | 897672                     | 2676                       | 808742                       | 2969                         | 4422966                | 498,4                  |
| simple              | 7191384          | 336,9            | 900682                     | 2626                       | 818016                       | 2930                         | 4926795                | 476,3                  |
| simple              | 7130270          | 336,6            | 916179                     | 2773                       | 812930                       | 2993                         | 4911207                | 475,2                  |
| simple              | 7231214          | 336,5            | 912652                     | 2629                       | 806228                       | 2923                         | 4871220                | 475                    |
| simple              | 7133586          | 339,6            | 839488                     | 2803                       | 811442                       | 2929                         | 4778890                | 476,2                  |
| simple              | 7199564          | 336,6            | 912165                     | 2630                       | 811531                       | 2926                         | 4908123                | 474,6                  |
| simple              | 7063579          | 338,8            | 887763                     | 2630                       | 808515                       | 2932                         | 4947600                | 478,6                  |
| simple              | 7154626          | 336,7            | 897507                     | 2620                       | 821241                       | 2926                         | 4962855                | 475,3                  |
| simple              | 7184446          | 336,2            | 909108                     | 2621                       | 819040                       | 2920                         | 4812297                | 475,4                  |
| simple              | 6938464          | 340,3            | 913638                     | 2661                       | 810810                       | 2963                         | 4818706                | 490,2                  |
| simple              | 6945331          | 343,7            | 911082                     | 2668                       | 808646                       | 2952                         | 4773291                | 481                    |
| simple avg          | 7102529,84       | 339,99           | 893975,48                  | 2660,12                    | 812684,88                    | 2942,52                      | 4846960,68             | 481,15                 |
| simple median       | 7126405          | 338              | 905252                     | 2630                       | 811969                       | 2931                         | 4871220                | 477,1                  |
| simple stddev       | 96406,57         | 6,59             | 36346,88                   | 60,35                      | 6808,04                      | 24,15                        | 123134,43              | 9,13                   |
|                     |                  |                  |                            |                            |                              |                              |                        |                        |
| complex             | 1000000          | 2345             | 220923                     | 10774                      | 200908                       | 11993                        | 750070                 | 3190                   |
| complex             | 1000000          | 2365             | 221221                     | 10763                      | 201042                       | 11989                        | 750168                 | 3410                   |
| complex             | 972150           | 2364             | 221136                     | 10775                      | 198964                       | 12014                        | 749394                 | 3174                   |
| complex             | 983293           | 2354             | 219159                     | 10735                      | 199822                       | 11956                        | 685695                 | 3436                   |
| complex             | 998304           | 2373             | 223593                     | 10815                      | 200030                       | 12122                        | 697144                 | 3203                   |
| complex             | 1000000          | 2347             | 221875                     | 10760                      | 199878                       | 12012                        | 748706                 | 3197                   |
| complex             | 973760           | 2543             | 222940                     | 10775                      | 199466                       | 11956                        | 633584                 | 3169                   |
| complex             | 1000000          | 2348             | 222703                     | 10728                      | 200450                       | 11948                        | 752290                 | 3193                   |
| complex             | 1000000          | 2444             | 191391                     | 10984                      | 201798                       | 12051                        | 741266                 | 3245                   |
| complex             | 1000000          | 2397             | 222880                     | 10750                      | 201038                       | 12054                        | 747745                 | 3192                   |
| complex             | 1000000          | 3502             | 181063                     | 11646                      | 199801                       | 12075                        | 745404                 | 3176                   |
| complex             | 1000000          | 2345             | 221311                     | 10733                      | 200516                       | 11957                        | 751840                 | 3185                   |
| complex             | 1000000          | 2355             | 223248                     | 10765                      | 202176                       | 11954                        | 754708                 | 3184                   |
| complex             | 1000000          | 2368             | 223058                     | 10739                      | 200386                       | 11950                        | 748453                 | 3187                   |
| complex             | 945394           | 2413             | 216142                     | 10996                      | 200929                       | 12329                        | 739760                 | 3251                   |
| complex             | 1000000          | 2345             | 225157                     | 10730                      | 200350                       | 11954                        | 753710                 | 3168                   |
| complex             | 1000000          | 2357             | 222384                     | 10761                      | 201603                       | 12001                        | 748792                 | 3207                   |
| complex             | 1000000          | 2374             | 222255                     | 10775                      | 199756                       | 11972                        | 751460                 | 3191                   |
| complex             | 1000000          | 2360             | 222795                     | 10759                      | 198070                       | 11984                        | 760131                 | 3184                   |
| complex             | 1000000          | 2362             | 222974                     | 10744                      | 200458                       | 11967                        | 655891                 | 3179                   |
| complex             | 1000000          | 2360             | 223491                     | 10735                      | 198692                       | 12018                        | 741771                 | 3192                   |
| complex             | 1000000          | 2356             | 221150                     | 10725                      | 201123                       | 11969                        | 752702                 | 3187                   |
| complex             | 1000000          | 2358             | 221970                     | 10725                      | 200338                       | 12006                        | 748818                 | 3190                   |
| complex             | 1000000          | 2458             | 221637                     | 10883                      | 201889                       | 12095                        | 754812                 | 3242                   |
| complex             | 973845           | 2387             | 219902                     | 10843                      | 198200                       | 12197                        | 754618                 | 3187                   |
| complex avg         | 993869,84        | 2423,2           | 219054,32                  | 10816,72                   | 200307,32                    | 12020,92                     | 736757,28              | 3212,76                |
| complex median      | 1000000          | 2362             | 221970                     | 10761                      | 200386                       | 11993                        | 748818                 | 3190                   |
| complex stddev      | 13651,29         | 229,21           | 10140,31                   | 187,48                     | 1082,08                      | 89,04                        | 32555,94               | 66,95                  |
|                     |                  |                  |                            |                            |                              |                              |                        |                        |
| conditional         | 1761163          | 1357             | 401946                     | 5939                       | 383113                       | 6261                         | 1326986                | 1796                   |
| conditional         | 1771482          | 1354             | 401154                     | 5946                       | 381466                       | 6256                         | 1333767                | 1798                   |
| conditional         | 1787192          | 1355             | 401908                     | 5946                       | 385698                       | 6255                         | 1342350                | 1784                   |
| conditional         | 1707636          | 1359             | 400123                     | 5955                       | 381170                       | 6274                         | 1337961                | 1795                   |
| conditional         | 1754228          | 1358             | 406419                     | 5934                       | 378148                       | 6260                         | 1348107                | 1781                   |
| conditional         | 1762357          | 1356             | 406225                     | 5960                       | 379110                       | 6289                         | 1337082                | 1810                   |
| conditional         | 1770903          | 1354             | 401331                     | 5953                       | 379926                       | 6268                         | 1335477                | 1810                   |
| conditional         | 1789513          | 1355             | 402715                     | 5925                       | 373424                       | 6247                         | 1340494                | 1802                   |
| conditional         | 1685074          | 1380             | 405048                     | 5995                       | 382274                       | 6288                         | 1341414                | 1807                   |
| conditional         | 1776829          | 1425             | 374980                     | 6093                       | 384249                       | 6245                         | 1345015                | 1787                   |
| conditional         | 1766887          | 1353             | 406766                     | 5927                       | 386780                       | 6252                         | 1344804                | 1793                   |
| conditional         | 1773517          | 1366             | 404851                     | 5941                       | 382365                       | 6254                         | 1302368                | 1797                   |
| conditional         | 1772199          | 1353             | 403869                     | 5941                       | 381678                       | 6246                         | 1336604                | 1790                   |
| conditional         | 1766043          | 1359             | 403119                     | 6183                       | 385885                       | 6383                         | 1282988                | 1785                   |
| conditional         | 1745331          | 1369             | 392877                     | 6016                       | 379905                       | 6431                         | 1315065                | 1807                   |
| conditional         | 1775966          | 1353             | 404760                     | 5936                       | 379540                       | 6252                         | 1319792                | 1803                   |
| conditional         | 1673844          | 1382             | 406208                     | 5936                       | 376844                       | 6251                         | 1193227                | 1794                   |
| conditional         | 1758682          | 1361             | 400184                     | 6003                       | 380378                       | 6270                         | 1326451                | 1797                   |
| conditional         | 1778167          | 1358             | 400514                     | 5943                       | 380977                       | 6283                         | 1321360                | 1797                   |
| conditional         | 1765039          | 1377             | 399458                     | 5948                       | 384547                       | 6258                         | 1339388                | 1803                   |
| conditional         | 1759484          | 1361             | 399762                     | 5925                       | 385018                       | 6248                         | 1330646                | 1798                   |
| conditional         | 1760367          | 1357             | 399334                     | 5933                       | 384848                       | 6235                         | 1342532                | 1787                   |
| conditional         | 1763613          | 1354             | 405427                     | 5922                       | 382321                       | 6255                         | 1329721                | 1797                   |
| conditional         | 1749752          | 1364             | 392264                     | 6072                       | 361263                       | 6319                         | 1332295                | 1800                   |
| conditional         | 1750826          | 1372             | 402261                     | 6000                       | 379243                       | 6326                         | 1332682                | 1807                   |
| conditional avg     | 1757043,76       | 1363,68          | 400940,12                  | 5970,88                    | 380806,8                     | 6276,24                      | 1325543,04             | 1797                   |
| conditional median  | 1763613          | 1358             | 401946                     | 5946                       | 381466                       | 6258                         | 1333767                | 1797                   |
| conditional stddev  | 28277,06         | 15,39            | 6546,95                    | 62,72                      | 5110,37                      | 45,55                        | 31112,94               | 8,15                   |
|                     |                  |                  |                            |                            |                              |                              |                        |                        |
| intermediate        | 1967211          | 1257             | 364634                     | 6193                       | 370027                       | 6439                         | 1539102                | 1534                   |
| intermediate        | 1964376          | 1236             | 409650                     | 5841                       | 370682                       | 6449                         | 1572568                | 1528                   |
| intermediate        | 1949232          | 1352             | 410047                     | 5844                       | 374745                       | 6436                         | 1565110                | 1524                   |
| intermediate        | 1939333          | 1251             | 407359                     | 5853                       | 369106                       | 6465                         | 1580144                | 1532                   |
| intermediate        | 1914984          | 1248             | 380869                     | 6215                       | 371192                       | 6452                         | 1564984                | 1541                   |
| intermediate        | 1947264          | 1350             | 356775                     | 5979                       | 369568                       | 6448                         | 1564012                | 1533                   |
| intermediate        | 1945606          | 1225             | 410191                     | 5841                       | 372290                       | 6468                         | 1560674                | 1523                   |
| intermediate        | 1963743          | 1249             | 407761                     | 5839                       | 371035                       | 6909                         | 1566199                | 1529                   |
| intermediate        | 1937840          | 1242             | 405748                     | 5911                       | 366315                       | 6492                         | 1550152                | 1521                   |
| intermediate        | 1939778          | 1239             | 405986                     | 6264                       | 367609                       | 6460                         | 1551794                | 1535                   |
| intermediate        | 1960746          | 1224             | 409324                     | 5841                       | 375855                       | 6454                         | 1574971                | 1530                   |
| intermediate        | 1951579          | 1231             | 408160                     | 5845                       | 367803                       | 6452                         | 1551181                | 1539                   |
| intermediate        | 1962092          | 1226             | 409024                     | 5834                       | 333604                       | 6597                         | 1571673                | 1519                   |
| intermediate        | 1877978          | 1274             | 401834                     | 6000                       | 368730                       | 6610                         | 1530961                | 1559                   |
| intermediate        | 1899608          | 1268             | 407286                     | 5928                       | 364371                       | 6699                         | 1545860                | 1551                   |
| intermediate        | 1960666          | 1232             | 407598                     | 5829                       | 372921                       | 6459                         | 1565779                | 1526                   |
| intermediate        | 1937190          | 1232             | 413080                     | 5843                       | 372584                       | 6451                         | 1566754                | 1534                   |
| intermediate        | 1934011          | 1244             | 411706                     | 5854                       | 366188                       | 6469                         | 1566883                | 1529                   |
| intermediate        | 1959883          | 1232             | 408008                     | 5841                       | 372967                       | 6472                         | 1571113                | 1524                   |
| intermediate        | 1949336          | 1224             | 407044                     | 5845                       | 373170                       | 6439                         | 1563824                | 1532                   |
| intermediate        | 1943166          | 1233             | 413164                     | 5847                       | 372246                       | 6457                         | 1381070                | 1537                   |
| intermediate        | 1951704          | 1233             | 407367                     | 5842                       | 370960                       | 6450                         | 1564404                | 1526                   |
| intermediate        | 1944936          | 1248             | 410502                     | 5853                       | 366796                       | 6462                         | 1531525                | 1639                   |
| intermediate        | 1907518          | 1257             | 412392                     | 5910                       | 355419                       | 6558                         | 1562833                | 1540                   |
| intermediate        | 1764538          | 1282             | 407710                     | 5913                       | 364524                       | 6520                         | 1546740                | 1649                   |
| intermediate avg    | 1934972,72       | 1251,56          | 403728,76                  | 5912,2                     | 368028,28                    | 6502,68                      | 1552412,4              | 1541,36                |
| intermediate median | 1945606          | 1242             | 407761                     | 5847                       | 370027                       | 6460                         | 1564012                | 1532                   |
| intermediate stddev | 41615,49         | 33,7             | 14347,76                   | 126,4                      | 8312,39                      | 106,28                       | 37973,3                | 32,2                   |
|                     |                  |                  |                            |                            |                              |                              |                        |                        |
| mixed               | 1491925          | 1610             | 530078                     | 4457                       | 502545                       | 4772                         | 1285450                | 1855                   |
| mixed               | 1490300          | 1605             | 531625                     | 4460                       | 504578                       | 5002                         | 1283505                | 1862                   |
| mixed               | 1470877          | 1615             | 539704                     | 4472                       | 500138                       | 4770                         | 1279606                | 1869                   |
| mixed               | 1370437          | 1627             | 534024                     | 4447                       | 505645                       | 4791                         | 1292632                | 1862                   |
| mixed               | 1488138          | 1620             | 534123                     | 4463                       | 494505                       | 4805                         | 1287218                | 1861                   |
| mixed               | 1481264          | 1615             | 533299                     | 4466                       | 497670                       | 4784                         | 1274191                | 1863                   |
| mixed               | 1492581          | 1619             | 543978                     | 4448                       | 507468                       | 4758                         | 1284162                | 1864                   |
| mixed               | 1489264          | 1612             | 536594                     | 4449                       | 505935                       | 4769                         | 1275082                | 1871                   |
| mixed               | 1482362          | 1616             | 533193                     | 4559                       | 465102                       | 4834                         | 1280503                | 1877                   |
| mixed               | 1478761          | 1631             | 533559                     | 4470                       | 503614                       | 4894                         | 1281306                | 1856                   |
| mixed               | 1475074          | 1621             | 543597                     | 4448                       | 496706                       | 4774                         | 1286600                | 2052                   |
| mixed               | 1470812          | 1620             | 532965                     | 4447                       | 503172                       | 4756                         | 1276438                | 1878                   |
| mixed               | 1481493          | 1612             | 545342                     | 4450                       | 502586                       | 4763                         | 1278556                | 2058                   |
| mixed               | 1416348          | 1673             | 520384                     | 4556                       | 466088                       | 4909                         | 1232283                | 1924                   |
| mixed               | 1446158          | 1666             | 529260                     | 4540                       | 500601                       | 4912                         | 1233720                | 2093                   |
| mixed               | 1367568          | 1623             | 531343                     | 4460                       | 502404                       | 4793                         | 1278543                | 1944                   |
| mixed               | 1488856          | 1610             | 524438                     | 4462                       | 505186                       | 4915                         | 1275585                | 1876                   |
| mixed               | 1484791          | 1622             | 531183                     | 4455                       | 497931                       | 4776                         | 1286622                | 2058                   |
| mixed               | 1482290          | 1616             | 534446                     | 4447                       | 503613                       | 4758                         | 1291989                | 1861                   |
| mixed               | 1279209          | 1619             | 532794                     | 4455                       | 501841                       | 4775                         | 1294134                | 1869                   |
| mixed               | 1495063          | 1692             | 533358                     | 4446                       | 497517                       | 4776                         | 1288670                | 2011                   |
| mixed               | 1495306          | 1605             | 537343                     | 4441                       | 498868                       | 4767                         | 1299070                | 1870                   |
| mixed               | 1481564          | 1627             | 538863                     | 4463                       | 490764                       | 4783                         | 1288177                | 1876                   |
| mixed               | 1468518          | 1631             | 542824                     | 4509                       | 492010                       | 4820                         | 1281916                | 1881                   |
| mixed               | 1459990          | 1646             | 538268                     | 4498                       | 502053                       | 4818                         | 1269405                | 1909                   |
| mixed avg           | 1461157,96       | 1626,12          | 534663,4                   | 4470,72                    | 497941,6                     | 4810,96                      | 1279414,52             | 1912                   |
| mixed median        | 1481493          | 1620             | 533559                     | 4460                       | 501841                       | 4783                         | 1281916                | 1876                   |
| mixed stddev        | 51294,61         | 21,47            | 5842,63                    | 34,35                      | 10612,75                     | 64,51                        | 15585,96               | 76,47                  |
|                     |                  |                  |                            |                            |                              |                              |                        |                        |
| multifilter         | 1785645          | 1349             | 405847                     | 5929                       | 385944                       | 6249                         | 1367917                | 1746                   |
| multifilter         | 1775078          | 1355             | 407239                     | 6304                       | 369632                       | 6254                         | 1366550                | 1750                   |
| multifilter         | 1765948          | 1347             | 397543                     | 5956                       | 376857                       | 6255                         | 1361536                | 1747                   |
| multifilter         | 1770699          | 1373             | 372030                     | 6026                       | 385699                       | 6630                         | 1366354                | 1749                   |
| multifilter         | 1765532          | 1365             | 404883                     | 5932                       | 382356                       | 6250                         | 1367367                | 1749                   |
| multifilter         | 1772985          | 1352             | 405282                     | 5928                       | 373760                       | 6250                         | 1366009                | 1744                   |
| multifilter         | 1779619          | 1348             | 404727                     | 5911                       | 381897                       | 6245                         | 1364995                | 1764                   |
| multifilter         | 1779584          | 1356             | 402487                     | 5931                       | 382945                       | 6235                         | 1361892                | 1749                   |
| multifilter         | 1797445          | 1347             | 403706                     | 5946                       | 375079                       | 6304                         | 1364300                | 1744                   |
| multifilter         | 1741188          | 1375             | 395217                     | 6093                       | 357421                       | 6425                         | 1341567                | 1784                   |
| multifilter         | 1759147          | 1348             | 399774                     | 5935                       | 382672                       | 6235                         | 1370491                | 1752                   |
| multifilter         | 1783940          | 1350             | 401028                     | 5931                       | 382639                       | 6259                         | 1369666                | 1750                   |
| multifilter         | 1772955          | 1340             | 403993                     | 5920                       | 381873                       | 6249                         | 1376023                | 1747                   |
| multifilter         | 1726716          | 1381             | 398599                     | 6078                       | 355581                       | 6406                         | 1341860                | 1784                   |
| multifilter         | 1759033          | 1368             | 395697                     | 6034                       | 380331                       | 6364                         | 1357093                | 1766                   |
| multifilter         | 1768196          | 1345             | 403572                     | 5926                       | 382366                       | 6250                         | 1357777                | 1749                   |
| multifilter         | 1786311          | 1345             | 401508                     | 5954                       | 379269                       | 6240                         | 1351516                | 1758                   |
| multifilter         | 1776760          | 1345             | 404756                     | 5918                       | 382744                       | 6250                         | 1356824                | 1754                   |
| multifilter         | 1780686          | 1346             | 402231                     | 5934                       | 381162                       | 6246                         | 1366863                | 1749                   |
| multifilter         | 1783516          | 1337             | 401127                     | 5930                       | 385303                       | 6241                         | 1374058                | 1745                   |
| multifilter         | 1790011          | 1352             | 402717                     | 5923                       | 383881                       | 6263                         | 1379800                | 1747                   |
| multifilter         | 1574276          | 1359             | 399538                     | 5930                       | 385738                       | 6250                         | 1367830                | 1752                   |
| multifilter         | 1782702          | 1339             | 404448                     | 5912                       | 377984                       | 6238                         | 1370560                | 1746                   |
| multifilter         | 1609984          | 1372             | 400779                     | 5977                       | 375669                       | 6315                         | 1316362                | 1757                   |
| multifilter         | 1773374          | 1354             | 402198                     | 6005                       | 383467                       | 6343                         | 1371200                | 1775                   |
| multifilter avg     | 1758453,2        | 1353,92          | 400837,04                  | 5970,52                    | 378890,76                    | 6289,84                      | 1362256,4              | 1754,28                |
| multifilter median  | 1773374          | 1350             | 402231                     | 5932                       | 381897                       | 6250                         | 1366354                | 1749                   |
| multifilter stddev  | 52505,37         | 11,97            | 6734,75                    | 86,05                      | 7866,55                      | 88,88                        | 13263,76               | 11,6                   |
|                     |                  |                  |                            |                            |                              |                              |                        |                        |
| aggs                | 1107342          | 2152             | 294852                     | 8214                       | 280825                       | 8528                         | 816226                 | 2908                   |
| aggs                | 1107857          | 2151             | 290728                     | 8213                       | 281715                       | 8578                         | 819967                 | 2921                   |
| aggs                | 1100693          | 2165             | 291958                     | 8215                       | 281131                       | 8542                         | 825214                 | 2907                   |
| aggs                | 1111474          | 2144             | 289671                     | 8190                       | 282400                       | 8522                         | 828920                 | 2898                   |
| aggs                | 1102730          | 2161             | 289561                     | 8246                       | 279996                       | 8555                         | 811369                 | 2890                   |
| aggs                | 1109921          | 2147             | 291009                     | 8177                       | 283255                       | 8512                         | 829220                 | 2905                   |
| aggs                | 1107857          | 2151             | 291703                     | 8182                       | 282865                       | 8523                         | 823189                 | 2897                   |
| aggs                | 1109921          | 2147             | 290378                     | 8198                       | 279508                       | 8508                         | 819733                 | 2903                   |
| aggs                | 1109921          | 2147             | 293406                     | 8189                       | 280531                       | 8528                         | 823610                 | 2895                   |
| aggs                | 1108372          | 2150             | 290304                     | 8223                       | 282766                       | 8556                         | 816241                 | 2893                   |
| aggs                | 1098663          | 2169             | 292372                     | 8203                       | 282356                       | 8540                         | 824234                 | 2902                   |
| aggs                | 1112512          | 2142             | 290540                     | 8198                       | 280502                       | 8559                         | 823286                 | 2900                   |
| aggs                | 1103752          | 2159             | 290900                     | 8279                       | 257852                       | 8565                         | 827701                 | 2910                   |
| aggs                | 1109404          | 2148             | 291739                     | 8179                       | 280419                       | 8526                         | 821042                 | 2901                   |
| aggs                | 1067174          | 2233             | 291428                     | 8364                       | 279945                       | 8767                         | 805053                 | 2961                   |
| aggs                | 1105800          | 2155             | 291122                     | 8181                       | 284082                       | 8496                         | 804098                 | 2902                   |
| aggs                | 1098663          | 2169             | 290558                     | 8222                       | 282330                       | 8532                         | 833383                 | 2898                   |
| aggs                | 1107342          | 2152             | 289677                     | 8195                       | 280248                       | 8507                         | 826546                 | 2886                   |
| aggs                | 1108888          | 2149             | 289767                     | 8197                       | 278280                       | 8522                         | 829855                 | 2889                   |
| aggs                | 1110956          | 2145             | 291128                     | 8235                       | 279548                       | 8529                         | 819082                 | 2898                   |
| aggs                | 1088128          | 2190             | 292560                     | 8213                       | 283500                       | 8514                         | 759892                 | 3109                   |
| aggs                | 1114593          | 2138             | 292746                     | 8196                       | 281385                       | 9107                         | 825134                 | 2901                   |
| aggs                | 1106828          | 2153             | 294103                     | 8208                       | 281836                       | 8518                         | 817861                 | 2889                   |
| aggs                | 1107342          | 2152             | 286256                     | 8247                       | 280830                       | 9103                         | 825268                 | 2893                   |
| aggs                | 1072940          | 2221             | 290528                     | 8271                       | 281358                       | 8607                         | 821478                 | 2903                   |
| aggs avg            | 1103563          | 2159,6           | 291159,76                  | 8217,4                     | 280378,52                    | 8589,76                      | 819104,08              | 2910,36                |
| aggs median         | 1107342          | 2152             | 291009                     | 8208                       | 281131                       | 8529                         | 823189                 | 2901                   |
| aggs stddev         | 11507,89175      | 23,03            | 1712,12                    | 40,79                      | 4898,81                      | 163,69                       | 14225,85               | 43,83                  |
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

[scorecard]: https://scorecard.dev/viewer/?uri=github.com/Trendyol/es-query-builder

[scorecard-img]: https://api.scorecard.dev/projects/github.com/Trendyol/es-query-builder/badge
