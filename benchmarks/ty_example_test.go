package benchmarks_test

import (
	"testing"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/Trendyol/es-query-builder/test/assert"
)

func createTyExampleQuery(brandIds []int64, storefrontIds []string) map[string]any {
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

	return query
}

func createTyExampleQueryVanilla(brandIds []int64, storefrontIds []string) map[string]any {
	query := map[string]interface{}{
		"size": 1,
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"filter": []interface{}{
					map[string]interface{}{
						"term": map[string]interface{}{
							"type": map[string]interface{}{
								"value": "LegalRule",
							},
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
	return query
}

func Benchmark_Ty_Example_Builder(b *testing.B) {
	brandIds := []int64{11, 22, 33, 44}
	storefrontIds := []string{"35", "36", "43", "48", "49", "50"}
	createTyExampleQuery(brandIds, storefrontIds)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		createTyExampleQuery(brandIds, storefrontIds)
	}
}

func Benchmark_Ty_Example_Vanilla(b *testing.B) {
	brandIds := []int64{11, 22, 33, 44}
	storefrontIds := []string{"35", "36", "43", "48", "49", "50"}
	createTyExampleQueryVanilla(brandIds, storefrontIds)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		createTyExampleQueryVanilla(brandIds, storefrontIds)
	}
}

func Test_Ty_Example_Queries_are_equal(t *testing.T) {
	brandIds := []int64{11, 22, 33, 44}
	storefrontIds := []string{"35", "36", "43", "48", "49", "50"}
	build := marshalString(t, createTyExampleQuery(brandIds, storefrontIds))
	vanilla := marshalString(t, createTyExampleQueryVanilla(brandIds, storefrontIds))
	assert.Equal(t, vanilla, build)
}
