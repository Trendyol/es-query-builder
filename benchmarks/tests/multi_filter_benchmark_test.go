package tests_test

import (
	"testing"

	Order "github.com/Trendyol/es-query-builder/es/enums/sort/order"

	"github.com/Trendyol/es-query-builder/benchmarks/tests/marshal"
	"github.com/Trendyol/es-query-builder/es"
	"github.com/Trendyol/es-query-builder/test/assert"
)

func createMultiFilterQuery() map[string]any {
	query := es.NewQuery(
		es.Bool().
			Must(
				es.Term("status", "active"),
				es.Range("price").GreaterThanOrEqual(100).LessThanOrEqual(1000),
			).
			Filter(
				es.Terms("category", "Electronics", "Home"),
				es.Exists("stock"),
			),
	).
		Size(20).
		Sort(es.Sort("price").Order(Order.Desc)).
		SourceIncludes("name", "price").
		SourceExcludes("internal_code")

	return query
}

func createMultiFilterQueryVanilla() map[string]any {
	return map[string]interface{}{
		"_source": map[string]interface{}{
			"includes": []string{"name", "price"},
			"excludes": []string{"internal_code"},
		},
		"size": 20,
		"sort": []map[string]interface{}{
			{"price": map[string]interface{}{"order": "desc"}},
		},
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": []map[string]interface{}{
					{"term": map[string]interface{}{"status": "active"}},
					{"range": map[string]interface{}{
						"price": map[string]interface{}{"gte": 100, "lte": 1000},
					}},
				},
				"filter": []map[string]interface{}{
					{"terms": map[string]interface{}{"category": []string{"Electronics", "Home"}}},
					{"exists": map[string]interface{}{"field": "stock"}},
				},
			},
		},
	}
}

func Benchmark_MultiFilter_Builder(b *testing.B) {
	createMultiFilterQuery()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		createMultiFilterQuery()
	}
}

func Benchmark_MultiFilter_Vanilla(b *testing.B) {
	createMultiFilterQueryVanilla()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		createMultiFilterQueryVanilla()
	}
}

func Test_MultiFilter_Queries_are_equal(t *testing.T) {
	build := marshal.String(t, createMultiFilterQuery())
	vanilla := marshal.String(t, createMultiFilterQueryVanilla())
	assert.Equal(t, vanilla, build)
}
