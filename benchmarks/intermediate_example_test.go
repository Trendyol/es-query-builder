package benchmarks_test

import (
	"testing"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/Trendyol/es-query-builder/es/enums/sort/order"
	"github.com/Trendyol/es-query-builder/test/assert"
)

func createIntermediateQuery(id int) map[string]any {
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
			)).
		Size(45).
		Sort(es.Sort("name").Order(order.Asc)).
		SourceIncludes("id", "type", "indexedAt", "chapters")

	return query
}

func createIntermediateQueryVanilla(id int) map[string]any {
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
										"doc.id": map[string]interface{}{
											"value": id,
										},
									},
								},
								map[string]interface{}{
									"term": map[string]interface{}{
										"file.fileId": map[string]interface{}{
											"value": id,
										},
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
	return query
}

func Benchmark_Intermediate_Builder(b *testing.B) {
	id := 42
	createIntermediateQuery(id)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		createIntermediateQuery(id)
	}
}

func Benchmark_Intermediate_Vanilla(b *testing.B) {
	id := 42
	createIntermediateQueryVanilla(id)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		createIntermediateQueryVanilla(id)
	}
}

func Test_Intermediate_Queries_are_equal(t *testing.T) {
	id := 42
	build := marshalString(t, createIntermediateQuery(id))
	vanilla := marshalString(t, createIntermediateQueryVanilla(id))
	assert.Equal(t, vanilla, build)
}
