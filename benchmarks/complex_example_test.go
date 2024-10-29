package benchmarks_test

import (
	"testing"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/Trendyol/es-query-builder/es/enums/sort/order"
	"github.com/Trendyol/es-query-builder/test/assert"
)

func createComplexQuery(id int) map[string]any {
	query := es.NewQuery(
		es.Bool().
			Must(
				es.Range("partition").
					GreaterThan(25).
					LesserThanOrEqual(30),
				es.Bool().
					Should(
						es.Term("doc.id", id),
						es.Term("file.fileId", id),
						es.Term("page.number", id),
					),
			).
			Filter(
				es.Term("type", "File"),
				es.Terms("sector", 1, 2, 3),
			).
			MustNot(
				es.Exists("blocks.reason.id"),
			).
			MinimumShouldMatch(1).
			Boost(3.14)).
		Size(100).
		From(5000).
		Sort(
			es.Sort("modifiedDate").Order(order.Desc),
			es.Sort("name").Order(order.Asc),
			es.Sort("indexedAt").Order(order.Asc),
		).
		SourceIncludes("id", "type", "indexedAt", "chapters").
		SourceExcludes("private.key", "cipher")

	return query
}

func createComplexQueryVanilla(id int) map[string]any {
	query := map[string]interface{}{
		"_source": map[string]interface{}{
			"includes": []interface{}{"id", "type", "indexedAt", "chapters"},
			"excludes": []interface{}{"private.key", "cipher"},
		},
		"size": 100,
		"from": 5000,
		"sort": []map[string]interface{}{
			{
				"modifiedDate": map[string]interface{}{
					"order": "desc",
				},
			},
			{
				"name": map[string]interface{}{
					"order": "asc",
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
				"minimum_should_match": 1,
				"boost":                3.14,
				"must": []map[string]interface{}{
					{
						"range": map[string]interface{}{
							"partition": map[string]interface{}{
								"gt":  25,
								"lte": 30,
							},
						},
					},
					{
						"bool": map[string]interface{}{
							"should": []map[string]interface{}{
								{
									"term": map[string]interface{}{
										"doc.id": map[string]interface{}{
											"value": id,
										},
									},
								},
								{
									"term": map[string]interface{}{
										"file.fileId": map[string]interface{}{
											"value": id,
										},
									},
								},
								{
									"term": map[string]interface{}{
										"page.number": map[string]interface{}{
											"value": id,
										},
									},
								},
							},
						},
					},
				},
				"filter": []map[string]interface{}{
					{
						"term": map[string]interface{}{
							"type": map[string]interface{}{
								"value": "File",
							},
						},
					},
					{
						"terms": map[string]interface{}{
							"sector": []interface{}{
								1, 2, 3,
							},
						},
					},
				},
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

func Benchmark_Complex_Builder(b *testing.B) {
	id := 76
	createComplexQuery(id)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		createComplexQuery(id)
	}
}

func Benchmark_Complex_Vanilla(b *testing.B) {
	id := 76
	createComplexQueryVanilla(id)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		createComplexQueryVanilla(id)
	}
}

func Test_Complex_Queries_are_equal(t *testing.T) {
	id := 76
	build := marshalString(t, createComplexQuery(id))
	vanilla := marshalString(t, createComplexQueryVanilla(id))
	assert.Equal(t, vanilla, build)
}
