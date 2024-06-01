package benchmarks_test

import (
	"encoding/json"
	"testing"

	"github.com/GokselKUCUKSAHIN/es-query-builder/es"
	"github.com/GokselKUCUKSAHIN/es-query-builder/es/enums/sort/order"
	"github.com/GokselKUCUKSAHIN/es-query-builder/test/assert"
)

////    Aggs Example    ////

func createAggsQuery() string {
	query := es.NewQuery(
		es.Bool().
			Must(
				es.Term("type", "File"),
			).
			MustNot(
				es.Exists("file.name"),
			),
	)

	query.
		Size(5_000).
		Sort(
			es.Sort("modifiedDate").Order(order.Desc),
			es.Sort("indexedAt").Order(order.Desc),
		)

	query.
		Range("indexedAt").
		GreaterThan("2020-06-01").
		LesserThanOrEqual("now")

	query.Aggs("DocumentIds",
		es.AggTerms().
			Field("document.id").
			Size(250).
			Aggs("OrderCounts",
				es.AggMultiTerms().
					Terms(
						es.AggTerm("document.orders.count"),
						es.AggTerm("files.order.count").
							Missing("book.meta.author"),
					),
			),
	)

	marshal, err := json.Marshal(query)
	if err != nil {
		return ""
	}
	return string(marshal)
}

func createAggsQueryVanillaGo() string {
	query := map[string]interface{}{
		"size": 5000,
		"sort": []map[string]interface{}{
			{
				"modifiedDate": map[string]interface{}{
					"order": "desc",
				},
			},
			{
				"indexedAt": map[string]interface{}{
					"order": "desc",
				},
			},
		},
		"query": map[string]interface{}{
			"range": map[string]interface{}{
				"indexedAt": map[string]interface{}{
					"gt":  "2020-06-01",
					"lte": "now",
				},
			},
			"bool": map[string]interface{}{
				"must": []map[string]interface{}{
					{
						"term": map[string]interface{}{
							"type": "File",
						},
					},
				},
				"must_not": []map[string]interface{}{
					{
						"exists": map[string]interface{}{
							"field": "file.name",
						},
					},
				},
			},
		},
		"aggs": map[string]interface{}{
			"DocumentIds": map[string]interface{}{
				"terms": map[string]interface{}{
					"field": "document.id",
					"size":  250,
				},
				"aggs": map[string]interface{}{
					"OrderCounts": map[string]interface{}{
						"multi_terms": map[string]interface{}{
							"terms": []map[string]interface{}{
								{
									"field": "document.orders.count",
								},
								{
									"field":   "files.order.count",
									"missing": "book.meta.author",
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

func Test_Aggs_Queries_are_equal(t *testing.T) {
	build := createAggsQuery()
	pure := createAggsQueryVanillaGo()
	assert.Equal(t, pure, build)
}

func Benchmark_Aggs_Example_Builder(b *testing.B) {
	createAggsQuery()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		createAggsQuery()
	}
}

func Benchmark_Aggs_Example_VanillaGo(b *testing.B) {
	createAggsQueryVanillaGo()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		createAggsQueryVanillaGo()
	}
}
