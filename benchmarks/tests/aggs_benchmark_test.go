package tests_test

import (
	"testing"

	"github.com/Trendyol/es-query-builder/benchmarks/tests/marshal"
	"github.com/Trendyol/es-query-builder/es"
	"github.com/Trendyol/es-query-builder/es/enums/sort/order"
	"github.com/Trendyol/es-query-builder/test/assert"
)

func createAggsQuery() map[string]any {
	query := es.NewQuery(
		es.Bool().
			Must(
				es.Term("type", "File"),
				es.Range("indexedAt").
					GreaterThan("2020-06-01").
					LessThanOrEqual("now"),
			).
			MustNot(
				es.Exists("file.name"),
			)).
		Size(5_000).
		Sort(
			es.Sort("modifiedDate").Order(order.Desc),
			es.Sort("indexedAt").Order(order.Asc),
		).
		Aggs(
			es.Agg("by_category", es.TermsAgg("category.keyword").
				Size(250).
				Aggs(
					es.Agg("nested_reviews", es.NestedAgg("reviews").
						Aggs(
							es.Agg("average_rating", es.AvgAgg("reviews.rating")),
							es.Agg("by_reviewer", es.TermsAgg("reviews.reviewer.keyword").
								Aggs(es.Agg("max_reviewer_rating", es.MaxAgg("reviews.rating"))),
							),
						),
					),
				),
			),
		)

	return query
}

func createAggsQueryVanilla() map[string]any {
	query := map[string]interface{}{
		"size": 5000,
		"sort": []map[string]interface{}{
			{
				"modifiedDate": map[string]interface{}{
					"order": "desc",
				},
			},
			{
				"indexedAt": map[string]interface{}{
					"order": "asc",
				},
			},
		},
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": []map[string]interface{}{
					{
						"term": map[string]interface{}{
							"type": map[string]interface{}{
								"value": "File",
							},
						},
					},
					{
						"range": map[string]interface{}{
							"indexedAt": map[string]interface{}{
								"gt":  "2020-06-01",
								"lte": "now",
							},
						},
					},
				},
				"must_not": []map[string]interface{}{
					{
						"exists": map[string]interface{}{
							"field": "file.name",
						},
					},
				},
			},
		},
		"aggs": map[string]interface{}{
			"by_category": map[string]interface{}{
				"terms": map[string]interface{}{
					"field": "category.keyword",
					"size":  250,
				},
				"aggs": map[string]interface{}{
					"nested_reviews": map[string]interface{}{
						"nested": map[string]interface{}{
							"path": "reviews",
						},
						"aggs": map[string]interface{}{
							"average_rating": map[string]interface{}{
								"avg": map[string]interface{}{
									"field": "reviews.rating",
								},
							},
							"by_reviewer": map[string]interface{}{
								"terms": map[string]interface{}{
									"field": "reviews.reviewer.keyword",
								},
								"aggs": map[string]interface{}{
									"max_reviewer_rating": map[string]interface{}{
										"max": map[string]interface{}{
											"field": "reviews.rating",
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

func Benchmark_Aggs_Builder(b *testing.B) {
	createAggsQuery()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		createAggsQuery()
	}
}

func Benchmark_Aggs_Vanilla(b *testing.B) {
	createAggsQueryVanilla()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		createAggsQueryVanilla()
	}
}

func Test_Aggs_Queries_are_equal(t *testing.T) {
	builder := marshal.String(t, createAggsQuery())
	vanilla := marshal.String(t, createAggsQueryVanilla())
	assert.Equal(t, vanilla, builder)
}
