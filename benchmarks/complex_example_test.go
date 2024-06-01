package benchmarks_test

import (
	"encoding/json"
	"testing"

	"github.com/GokselKUCUKSAHIN/es-query-builder/es"
	"github.com/GokselKUCUKSAHIN/es-query-builder/es/enums/sort/mode"
	"github.com/GokselKUCUKSAHIN/es-query-builder/es/enums/sort/order"
	"github.com/GokselKUCUKSAHIN/es-query-builder/test/assert"
)

////    Complex Example   ////

func createComplexQuery(id int) string {
	query := es.NewQuery(
		es.Bool().
			Must(
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
			Boost(3.14),
	)
	query.TrackTotalHits(true)
	query.Size(100)
	query.From(5000)
	query.Sort(
		es.Sort("modifiedDate").Order(order.Desc),
		es.Sort("name").Order(order.Asc).Mode(mode.Median),
		es.Sort("indexedAt").Order(order.Asc),
	)
	query.Source().
		Includes("id", "type", "indexedAt", "chapters").
		Excludes("private.key", "cipher")
	query.Range("partition").
		GreaterThan(25).
		LesserThanOrEqual(30)

	marshal, err := json.Marshal(query)
	if err != nil {
		return ""
	}
	return string(marshal)
}

func createComplexQueryVanillaGo(id int) string {
	query := map[string]interface{}{
		"_source": map[string]interface{}{
			"includes": []interface{}{"id", "type", "indexedAt", "chapters"},
			"excludes": []interface{}{"private.key", "cipher"},
		},
		"size":             100,
		"from":             5000,
		"track_total_hits": true,
		"sort": []map[string]interface{}{
			{
				"modifiedDate": map[string]interface{}{
					"order": "desc",
				},
			},
			{
				"name": map[string]interface{}{
					"order": "asc",
					"mode":  "median",
				},
			},
			{
				"indexedAt": map[string]interface{}{
					"order": "asc",
				},
			},
		},
		"query": map[string]interface{}{
			"range": map[string]interface{}{
				"partition": map[string]interface{}{
					"gt":  25,
					"lte": 30,
				},
			},
			"bool": map[string]interface{}{
				"minimum_should_match": 1,
				"boost":                3.14,
				"must": []map[string]interface{}{
					{
						"bool": map[string]interface{}{
							"should": []map[string]interface{}{
								{
									"term": map[string]interface{}{
										"doc.id": id,
									},
								},
								{
									"term": map[string]interface{}{
										"file.fileId": id,
									},
								},
								{
									"term": map[string]interface{}{
										"page.number": id,
									},
								},
							},
						},
					},
				},
				"filter": []map[string]interface{}{
					{
						"term": map[string]interface{}{
							"type": "File",
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

	marshal, err := json.Marshal(query)
	if err != nil {
		return ""
	}
	return string(marshal)
}

func Test_Complex_Queries_are_equal(t *testing.T) {
	id := 76
	build := createComplexQuery(id)
	pure := createComplexQueryVanillaGo(id)
	assert.Equal(t, pure, build)
}

func Benchmark_Complex_Builder(b *testing.B) {
	id := 76
	createComplexQuery(id)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		createComplexQuery(id)
	}
}

func Benchmark_Complex_VanillaGo(b *testing.B) {
	id := 76
	createComplexQueryVanillaGo(id)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		createComplexQueryVanillaGo(id)
	}
}
