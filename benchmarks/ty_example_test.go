package benchmarks_test

import (
	"encoding/json"
	"testing"

	"github.com/GokselKUCUKSAHIN/es-query-builder/es"
	"github.com/GokselKUCUKSAHIN/es-query-builder/test/assert"
)

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

func createTyExampleQueryVanillaGo(brandIds []int64, storefrontIds []string) string {
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
	pure := createTyExampleQueryVanillaGo(brandIds, storefrontIds)
	assert.Equal(t, pure, build)
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

func Benchmark_Ty_Example_VanillaGo(b *testing.B) {
	brandIds := []int64{11, 22, 33, 44}
	storefrontIds := []string{"35", "36", "43", "48", "49", "50"}
	createTyExampleQueryVanillaGo(brandIds, storefrontIds)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		createTyExampleQueryVanillaGo(brandIds, storefrontIds)
	}
}
