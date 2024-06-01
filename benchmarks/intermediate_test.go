package benchmarks_test

import (
	"encoding/json"
	"github.com/GokselKUCUKSAHIN/es-query-builder/es"
	"github.com/GokselKUCUKSAHIN/es-query-builder/es/enums/sort/order"
	"github.com/GokselKUCUKSAHIN/es-query-builder/test/assert"
	"testing"
)

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
	query.Sort(
		es.Sort("name").Order(order.Asc),
	)
	query.Source().
		Includes("id", "type", "indexedAt", "chapters")

	marshal, err := json.Marshal(query)
	if err != nil {
		return ""
	}
	return string(marshal)
}

func createIntermediateQueryVanillaGo(id int) string {
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
	pure := createIntermediateQueryVanillaGo(id)
	assert.Equal(t, pure, build)
}

func Benchmark_Intermediate_Builder(b *testing.B) {
	id := 42
	createIntermediateQuery(id)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		createIntermediateQuery(id)
	}
}

func Benchmark_Intermediate_VanillaGo(b *testing.B) {
	id := 42
	createIntermediateQueryVanillaGo(id)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		createIntermediateQueryVanillaGo(id)
	}
}
