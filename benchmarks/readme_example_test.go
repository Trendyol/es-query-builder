package benchmarks_test

import (
	"testing"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/Trendyol/es-query-builder/test/assert"
)

func createReadmeQuery() map[string]any {
	query := es.NewQuery(
		es.Bool().
			Must(
				es.Term("author", "George Orwell"),
			).
			MustNot(
				es.Terms("genre", "Fantasy", "Science Fiction"),
				es.Exists("out_of_print"),
			).
			Should(
				es.Terms("title", "1984", "Animal Farm"),
			),
	).Aggs("genres_count",
		es.AggTerms().
			Field("genre"),
	).Aggs("authors_and_genres",
		es.AggTerms().
			Field("author").
			Aggs("genres",
				es.AggTerms().
					Field("genre"),
			),
	)

	return query
}

func createReadmeQueryVanilla() map[string]any {
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": []map[string]interface{}{
					{
						"term": map[string]interface{}{
							"author": "George Orwell",
						},
					},
				},
				"must_not": []map[string]interface{}{
					{
						"terms": map[string]interface{}{
							"genre": []string{
								"Fantasy",
								"Science Fiction",
							},
						},
					},
					{
						"exists": map[string]interface{}{
							"field": "out_of_print",
						},
					},
				},
				"should": []map[string]interface{}{
					{
						"terms": map[string]interface{}{
							"title": []string{
								"1984",
								"Animal Farm",
							},
						},
					},
				},
			},
		},
		"aggs": map[string]interface{}{
			"genres_count": map[string]interface{}{
				"terms": map[string]interface{}{
					"field": "genre",
				},
			},
			"authors_and_genres": map[string]interface{}{
				"terms": map[string]interface{}{
					"field": "author",
				},
				"aggs": map[string]interface{}{
					"genres": map[string]interface{}{
						"terms": map[string]interface{}{
							"field": "genre",
						},
					},
				},
			},
		},
	}
	return query
}

func Benchmark_Readme_Example_Builder(b *testing.B) {
	createReadmeQuery()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		createReadmeQuery()
	}
}

func Benchmark_Readme_Example_Vanilla(b *testing.B) {
	createReadmeQueryVanilla()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		createReadmeQueryVanilla()
	}
}

func Test_Readme_Queries_are_equal(t *testing.T) {
	build := marshalString(t, createReadmeQuery())
	vanilla := marshalString(t, createReadmeQueryVanilla())
	assert.Equal(t, vanilla, build)
}
