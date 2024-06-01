package benchmarks_test

import (
	"encoding/json"
	"github.com/GokselKUCUKSAHIN/es-query-builder/es"
	"github.com/GokselKUCUKSAHIN/es-query-builder/test/assert"
	"testing"
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

func createSimpleQueryVanillaGo() string {
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
	pure := createSimpleQueryVanillaGo()
	assert.Equal(t, pure, build)
}

func Benchmark_Simple_Builder(b *testing.B) {
	createSimpleQuery()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		createSimpleQuery()
	}
}

func Benchmark_Simple_VanillaGo(b *testing.B) {
	createSimpleQueryVanillaGo()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		createSimpleQueryVanillaGo()
	}
}
