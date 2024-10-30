package tests_test

import (
	"testing"

	"github.com/Trendyol/es-query-builder/benchmarks/tests/marshal"
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
											"driver.vehicle.make": map[string]interface{}{
												"value": "Powell Motors",
											},
										},
									},
									{
										"term": map[string]interface{}{
											"driver.vehicle.model": map[string]interface{}{
												"value": "Canyonero",
											},
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

func Benchmark_Nested_Builder(b *testing.B) {
	createNestedQuery()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		createNestedQuery()
	}
}

func Benchmark_Nested_Vanilla(b *testing.B) {
	createNestedQueryVanilla()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		createNestedQueryVanilla()
	}
}

func Test_Nested_Queries_are_equal(t *testing.T) {
	build := marshal.String(t, createNestedQuery())
	vanilla := marshal.String(t, createNestedQueryVanilla())
	assert.Equal(t, vanilla, build)
}
