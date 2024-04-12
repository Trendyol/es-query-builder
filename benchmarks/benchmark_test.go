package benchmarks

import (
	"encoding/json"
	"testing"

	"github.com/GokselKUCUKSAHIN/es-query-builder/es"
	"github.com/GokselKUCUKSAHIN/es-query-builder/es/enums/sort/mode"
	"github.com/GokselKUCUKSAHIN/es-query-builder/es/enums/sort/order"
	"github.com/GokselKUCUKSAHIN/es-query-builder/test/assert"
)

////    Simple    ////

func createSimpleQuery() string {
	query := es.NewQuery(
		es.Bool().
			Filter(
				es.Term("id", 123456),
			),
	)

	marshal, err := json.Marshal(query)
	if err != nil {
		return ""
	}
	return string(marshal)
}

func createSimpleQueryPureGo() string {
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"filter": []map[string]interface{}{
					{
						"term": map[string]interface{}{
							"id": 123456,
						},
					},
				},
			},
		},
	}

	marshal, err := json.Marshal(query)
	if err != nil {
		return ""
	}
	return string(marshal)
}

func Test_Simple_Queries_are_equal(t *testing.T) {
	build := createSimpleQuery()
	pure := createSimpleQueryPureGo()
	assert.Equal(t, pure, build)
}

func Benchmark_Simple_Builder(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		createSimpleQuery()
	}
}

func Benchmark_Simple_PureGo(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		createSimpleQueryPureGo()
	}
}

////    Intermediate    ////

func createIntermediateQuery(id int) string {
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
	query.Size(45)
	query.Sort(es.Sort("name", order.Asc))
	query.Source().
		Includes("id", "type", "indexedAt", "chapters")

	marshal, err := json.Marshal(query)
	if err != nil {
		return ""
	}
	return string(marshal)
}

func createIntermediateQueryPureGo(id int) string {
	query := map[string]interface{}{
		"_source": map[string]interface{}{
			"includes": []interface{}{"id", "type", "indexedAt", "chapters"},
		},
		"size": 45,
		"sort": []map[string]interface{}{
			{
				"name": map[string]interface{}{
					"order": "asc",
				},
			},
		},
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

	marshal, err := json.Marshal(query)
	if err != nil {
		return ""
	}
	return string(marshal)
}

func Test_Intermediate_Queries_are_equal(t *testing.T) {
	id := 42
	build := createIntermediateQuery(id)
	pure := createIntermediateQueryPureGo(id)
	assert.Equal(t, pure, build)
}

func Benchmark_Intermediate_Builder(b *testing.B) {
	id := 42
	for i := 0; i <= b.N; i++ {
		createIntermediateQuery(id)
	}
}

func Benchmark_Intermediate_PureGo(b *testing.B) {
	id := 42
	for i := 0; i <= b.N; i++ {
		createIntermediateQueryPureGo(id)
	}
}

////    Conditional Complex   ////

func createConditionalQuery(items []int) string {
	query := es.NewQuery(
		es.Bool().
			Filter(
				es.Term("type", "File"),
				es.Terms("sector", 1, 2, 3),
				es.TermsArrayFunc("id", items, func(key string, values []int) bool {
					for _, value := range values {
						if value == 21 {
							return false
						}
					}
					return true
				}),
			).
			MustNot(
				es.Exists("blocks.reason.id"),
			),
	)
	query.Size(100)
	query.Sort(es.Sort("modifiedDate", order.Desc))
	query.Source().
		Includes("id", "type", "indexedAt", "chapters").
		Excludes("private.key")
	query.TrackTotalHits(true)
	query.Range("indexedAt").
		GreaterThan("2021-01-01").
		LesserThanOrEqual("now")

	marshal, err := json.Marshal(query)
	if err != nil {
		return ""
	}
	return string(marshal)
}

func createConditionalQueryPureGo(items []int) string {
	var flag bool
	for _, item := range items {
		if item == 21 {
			flag = false
			break
		}
		flag = true
	}

	filter := []map[string]interface{}{
		{
			"term": map[string]interface{}{
				"type": "File",
			},
		},
		{
			"terms": map[string]interface{}{
				"sector": []interface{}{1, 2, 3},
			},
		},
	}
	if flag {
		filter = append(filter, map[string]interface{}{
			"terms": map[string]interface{}{
				"id": items,
			},
		})
	}

	query := map[string]interface{}{
		"_source": map[string]interface{}{
			"includes": []interface{}{"id", "type", "indexedAt", "chapters"},
			"excludes": []interface{}{"private.key"},
		},
		"size":             100,
		"track_total_hits": true,
		"sort": []map[string]interface{}{
			{
				"modifiedDate": map[string]interface{}{
					"order": "desc",
				},
			},
		},
		"query": map[string]interface{}{
			"range": map[string]interface{}{
				"indexedAt": map[string]interface{}{
					"gt":  "2021-01-01",
					"lte": "now",
				},
			},
			"bool": map[string]interface{}{
				"filter": filter,
				"must_not": []map[string]interface{}{
					{
						"exists": map[string]interface{}{
							"field": "blocks.reason.id",
						},
					},
				},
			},
		},
	}

	marshal, err := json.Marshal(query)
	if err != nil {
		return ""
	}
	return string(marshal)
}

func Test_Conditional_Queries_are_equal(t *testing.T) {
	items := []int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55}

	build := createConditionalQuery(items)
	pure := createConditionalQueryPureGo(items)
	assert.Equal(t, pure, build)
}

func Benchmark_Conditional_Builder(b *testing.B) {
	items := []int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55}
	for i := 0; i <= b.N; i++ {
		createConditionalQuery(items)
	}
}

func Benchmark_Conditional_PureGo(b *testing.B) {
	items := []int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55}
	for i := 0; i <= b.N; i++ {
		createConditionalQueryPureGo(items)
	}
}

////    Complex    ////

func createComplexQuery(id int) string {
	query := es.NewQuery(
		es.Bool().
			Must(
				es.Bool().
					Should(
						es.Term("doc.id", id),
						es.Term("file.fileId", id),
						es.Term("page.number", id),
					),
			).
			Filter(
				es.Term("type", "File"),
				es.Terms("sector", 1, 2, 3),
			).
			MustNot(
				es.Exists("blocks.reason.id"),
			).
			MinimumShouldMatch(1).
			Boost(3.14),
	)
	query.TrackTotalHits(true)
	query.Size(100)
	query.From(5000)
	query.Sort(
		es.Sort("modifiedDate", order.Desc),
		es.SortWithMode("name", order.Asc, mode.Median),
		es.Sort("indexedAt", order.Asc),
	)
	query.Source().
		Includes("id", "type", "indexedAt", "chapters").
		Excludes("private.key", "cipher")
	query.Range("partition").
		GreaterThan(25).
		LesserThanOrEqual(30)

	marshal, err := json.Marshal(query)
	if err != nil {
		return ""
	}
	return string(marshal)
}

func createComplexQueryPureGo(id int) string {
	query := map[string]interface{}{
		"_source": map[string]interface{}{
			"includes": []interface{}{"id", "type", "indexedAt", "chapters"},
			"excludes": []interface{}{"private.key", "cipher"},
		},
		"size":             100,
		"from":             5000,
		"track_total_hits": true,
		"sort": []map[string]interface{}{
			{
				"modifiedDate": map[string]interface{}{
					"order": "desc",
				},
			},
			{
				"name": map[string]interface{}{
					"order": "asc",
					"mode":  "median",
				},
			},
			{
				"indexedAt": map[string]interface{}{
					"order": "asc",
				},
			},
		},
		"query": map[string]interface{}{
			"range": map[string]interface{}{
				"partition": map[string]interface{}{
					"gt":  25,
					"lte": 30,
				},
			},
			"bool": map[string]interface{}{
				"minimum_should_match": 1,
				"boost":                3.14,
				"must": []map[string]interface{}{
					{
						"bool": map[string]interface{}{
							"should": []map[string]interface{}{
								{
									"term": map[string]interface{}{
										"doc.id": id,
									},
								},
								{
									"term": map[string]interface{}{
										"file.fileId": id,
									},
								},
								{
									"term": map[string]interface{}{
										"page.number": id,
									},
								},
							},
						},
					},
				},
				"filter": []map[string]interface{}{
					{
						"term": map[string]interface{}{
							"type": "File",
						},
					},
					{
						"terms": map[string]interface{}{
							"sector": []interface{}{
								1, 2, 3,
							},
						},
					},
				},
				"must_not": []map[string]interface{}{
					{
						"exists": map[string]interface{}{
							"field": "blocks.reason.id",
						},
					},
				},
			},
		},
	}

	marshal, err := json.Marshal(query)
	if err != nil {
		return ""
	}
	return string(marshal)
}

func Test_Complex_Queries_are_equal(t *testing.T) {
	id := 76
	build := createComplexQuery(id)
	pure := createComplexQueryPureGo(id)
	assert.Equal(t, pure, build)
}

func Benchmark_Complex_Builder(b *testing.B) {
	id := 76
	for i := 0; i <= b.N; i++ {
		createComplexQuery(id)
	}
}

func Benchmark_Complex_PureGo(b *testing.B) {
	id := 76
	for i := 0; i <= b.N; i++ {
		createComplexQueryPureGo(id)
	}
}

////    TY Example    ////

func createTyExampleQuery(brandIds []int64, storefrontIds []string) string {
	query := es.NewQuery(
		es.Bool().
			Filter(
				es.Term("type", "LegalRule"),
				es.TermsArray("brandId", brandIds),
				es.TermsArray("allowedStorefronts.storefrontId", storefrontIds),
			),
	)
	query.Size(1)
	query.SourceFalse()

	marshal, err := json.Marshal(query)
	if err != nil {
		return ""
	}
	return string(marshal)
}

func createTyExampleQueryPureGo(brandIds []int64, storefrontIds []string) string {
	query := map[string]interface{}{
		"size": 1,
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"filter": []interface{}{
					map[string]interface{}{
						"term": map[string]interface{}{
							"type": "LegalRule",
						},
					},
					map[string]interface{}{
						"terms": map[string]interface{}{
							"brandId": brandIds,
						},
					},
					map[string]interface{}{
						"terms": map[string]interface{}{
							"allowedStorefronts.storefrontId": storefrontIds,
						},
					},
				},
			},
		},
		"_source": false,
	}

	marshal, err := json.Marshal(query)
	if err != nil {
		return ""
	}
	return string(marshal)
}

func Test_TyExample_Queries_are_equal(t *testing.T) {
	brandIds := []int64{11, 22, 33, 44}
	storefrontIds := []string{"35", "36", "43", "48", "49", "50"}
	build := createTyExampleQuery(brandIds, storefrontIds)
	pure := createTyExampleQueryPureGo(brandIds, storefrontIds)
	assert.Equal(t, pure, build)
}

func Benchmark_Ty_Example_Builder(b *testing.B) {
	brandIds := []int64{11, 22, 33, 44}
	storefrontIds := []string{"35", "36", "43", "48", "49", "50"}
	for i := 0; i <= b.N; i++ {
		createTyExampleQuery(brandIds, storefrontIds)
	}
}

func Benchmark_Ty_Example_PureGo(b *testing.B) {
	brandIds := []int64{11, 22, 33, 44}
	storefrontIds := []string{"35", "36", "43", "48", "49", "50"}
	for i := 0; i <= b.N; i++ {
		createTyExampleQueryPureGo(brandIds, storefrontIds)
	}
}

////    Nested Example    ////

func createNestedQuery() string {
	query := es.NewQuery(
		es.Nested("driver",
			es.Nested("driver.vehicle",
				es.Bool().
					Must(
						es.Term("driver.vehicle.make", "Powell Motors"),
						es.Term("driver.vehicle.model", "Canyonero"),
					),
			),
		),
	)

	marshal, err := json.Marshal(query)
	if err != nil {
		return ""
	}
	return string(marshal)
}

func createNestedQueryPureGo() string {
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"nested": map[string]interface{}{
				"path": "driver",
				"query": map[string]interface{}{
					"nested": map[string]interface{}{
						"path": "driver.vehicle",
						"query": map[string]interface{}{
							"bool": map[string]interface{}{
								"must": []map[string]interface{}{
									{
										"term": map[string]interface{}{
											"driver.vehicle.make": "Powell Motors",
										},
									},
									{
										"term": map[string]interface{}{
											"driver.vehicle.model": "Canyonero",
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	marshal, err := json.Marshal(query)
	if err != nil {
		return ""
	}
	return string(marshal)
}

func Test_Nested_Queries_are_equal(t *testing.T) {
	build := createNestedQuery()
	pure := createNestedQueryPureGo()
	assert.Equal(t, pure, build)
}

func Benchmark_Nested_Example_Builder(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		createNestedQuery()
	}
}

func Benchmark_Nested_Example_PureGo(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		createNestedQueryPureGo()
	}
}
