package benchmarks

import (
	"encoding/json"
	"testing"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/Trendyol/es-query-builder/test/assert"
)

////    Readme Example    ////

func createReadmeQuery() string {
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
	marshal, err := json.Marshal(query)
	if err != nil {
		return ""
	}
	return string(marshal)
}

func createReadmeQueryVanillaGo() string {
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
	marshal, err := json.Marshal(query)
	if err != nil {
		return ""
	}
	return string(marshal)
}

func Test_Readme_Queries_are_equal(t *testing.T) {
	build := createReadmeQuery()
	vanilla := createReadmeQueryVanillaGo()
	assert.Equal(t, vanilla, build)
}

func Benchmark_Readme_Example_Builder(b *testing.B) {
	createReadmeQuery()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		createReadmeQuery()
	}
}

func Benchmark_Readme_Example_VanillaGo(b *testing.B) {
	createReadmeQueryVanillaGo()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		createReadmeQueryVanillaGo()
	}
}
