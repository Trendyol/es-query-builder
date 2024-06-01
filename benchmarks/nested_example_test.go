package benchmarks_test

import (
	"encoding/json"
	"testing"

	"github.com/GokselKUCUKSAHIN/es-query-builder/es"
	"github.com/GokselKUCUKSAHIN/es-query-builder/test/assert"
)

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

func createNestedQueryVanillaGo() string {
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
	pure := createNestedQueryVanillaGo()
	assert.Equal(t, pure, build)
}

func Benchmark_Nested_Example_Builder(b *testing.B) {
	createNestedQuery()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		createNestedQuery()
	}
}

func Benchmark_Nested_Example_VanillaGo(b *testing.B) {
	createNestedQueryVanillaGo()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		createNestedQueryVanillaGo()
	}
}
