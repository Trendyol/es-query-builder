package benchmarks_test

import (
	"testing"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/Trendyol/es-query-builder/es/enums/sort/order"
	"github.com/Trendyol/es-query-builder/test/assert"
)

func createConditionalQuery(items []int) map[string]any {
	query := es.NewQuery(
		es.Bool().
			Filter(
				es.Range("indexedAt").
					GreaterThan("2021-01-01").
					LesserThanOrEqual("now"),
				es.Term("type", "File"),
				es.Terms("sector", 1, 2, 3),
				es.TermsFunc("id", items, func(key string, values []int) bool {
					for _, value := range values {
						if value == 21 {
							return false
						}
					}
					return true
				}),
			).
			MustNot(
				es.Exists("blocks.reason.id"),
			),
	).
		Size(100).
		Sort(
			es.Sort("modifiedDate").Order(order.Desc),
		)

	query.Source().
		Includes("id", "type", "indexedAt", "chapters").
		Excludes("private.key")

	return query
}

func createConditionalQueryVanilla(items []int) map[string]any {
	var flag bool
	for _, item := range items {
		if item == 21 {
			flag = false
			break
		}
		flag = true
	}

	filter := []map[string]interface{}{
		{
			"range": map[string]interface{}{
				"indexedAt": map[string]interface{}{
					"gt":  "2021-01-01",
					"lte": "now",
				},
			},
		},
		{
			"term": map[string]interface{}{
				"type": "File",
			},
		},
		{
			"terms": map[string]interface{}{
				"sector": []interface{}{1, 2, 3},
			},
		},
	}
	if flag {
		filter = append(filter, map[string]interface{}{
			"terms": map[string]interface{}{
				"id": items,
			},
		})
	}

	query := map[string]interface{}{
		"_source": map[string]interface{}{
			"includes": []interface{}{"id", "type", "indexedAt", "chapters"},
			"excludes": []interface{}{"private.key"},
		},
		"size": 100,
		"sort": []map[string]interface{}{
			{
				"modifiedDate": map[string]interface{}{
					"order": "desc",
				},
			},
		},
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"filter": filter,
				"must_not": []map[string]interface{}{
					{
						"exists": map[string]interface{}{
							"field": "blocks.reason.id",
						},
					},
				},
			},
		},
	}
	return query
}

func Benchmark_Conditional_Builder(b *testing.B) {
	items := []int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55}
	createConditionalQuery(items)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		createConditionalQuery(items)
	}
}

func Benchmark_Conditional_Vanilla(b *testing.B) {
	items := []int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55}
	createConditionalQueryVanilla(items)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		createConditionalQueryVanilla(items)
	}
}

func Test_Conditional_Queries_are_equal(t *testing.T) {
	items := []int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55}
	build := marshalString(t, createConditionalQuery(items))
	vanilla := marshalString(t, createConditionalQueryVanilla(items))
	assert.Equal(t, vanilla, build)
}
