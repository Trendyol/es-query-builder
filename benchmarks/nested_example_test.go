package benchmarks_test

import (
	"testing"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/Trendyol/es-query-builder/test/assert"
)

func createNestedQuery() map[string]any {
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

	return query
}

func createNestedQueryVanilla() map[string]any {
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
	return query
}

func Benchmark_Nested_Example_Builder(b *testing.B) {
	createNestedQuery()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		createNestedQuery()
	}
}

func Benchmark_Nested_Example_Vanilla(b *testing.B) {
	createNestedQueryVanilla()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		createNestedQueryVanilla()
	}
}

func Test_Nested_Queries_are_equal(t *testing.T) {
	build := marshalString(t, createNestedQuery())
	vanilla := marshalString(t, createNestedQueryVanilla())
	assert.Equal(t, vanilla, build)
}
