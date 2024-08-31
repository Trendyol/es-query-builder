package benchmarks_test

import (
	"testing"

	"github.com/GokselKUCUKSAHIN/es-query-builder/es"
	"github.com/GokselKUCUKSAHIN/es-query-builder/test/assert"
)

func createSimpleQuery() map[string]any {
	query := es.NewQuery(
		es.Bool().
			Filter(
				es.Term("id", 123456),
			),
	)
	return query
}

func createSimpleQueryVanilla() map[string]any {
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
	return query
}

func Benchmark_Simple_Builder(b *testing.B) {
	createSimpleQuery()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		createSimpleQuery()
	}
}

func Benchmark_Simple_Vanilla(b *testing.B) {
	createSimpleQueryVanilla()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		createSimpleQueryVanilla()
	}
}

func Test_Simple_Queries_are_equal(t *testing.T) {
	build := marshalString(t, createSimpleQuery())
	vanilla := marshalString(t, createSimpleQueryVanilla())
	assert.Equal(t, vanilla, build)
}
