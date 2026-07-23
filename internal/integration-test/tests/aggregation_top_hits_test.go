package tests_test

import (
	"encoding/json"

	Order "github.com/Trendyol/es-query-builder/es/enums/sort/order"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/stretchr/testify/assert"
)

func (s *testSuite) Test_it_should_return_top_hits_inside_terms_aggregation() {
	// Given
	query := es.NewQuery(es.MatchAll()).
		Size(0).
		Aggs(
			es.Agg("by_is_default",
				es.TermsAgg("isDefault").Size(2).Aggs(
					es.Agg("sample",
						es.TopHitsAgg().
							Size(1).
							Sort(es.Sort("id").Order(Order.Asc)).
							SourceIncludes("name", "id"),
					),
				),
			),
		)

	// When
	response, err := s.PokedexElasticsearchRepository.Search(s.TestContext, query)

	// Then
	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), response.Aggregations)

	raw, exists := response.Aggregations["by_is_default"]
	assert.True(s.T(), exists)

	var termsAgg struct {
		Buckets []struct {
			KeyAsString string `json:"key_as_string"`
			DocCount    int    `json:"doc_count"`
			Sample      struct {
				Hits struct {
					Hits []struct {
						Source struct {
							Name string `json:"name"`
							ID   int    `json:"id"`
						} `json:"_source"`
					} `json:"hits"`
				} `json:"hits"`
			} `json:"sample"`
		} `json:"buckets"`
	}
	assert.NoError(s.T(), json.Unmarshal(raw, &termsAgg))
	assert.Greater(s.T(), len(termsAgg.Buckets), 0)

	for _, bucket := range termsAgg.Buckets {
		assert.Greater(s.T(), bucket.DocCount, 0)
		assert.Equal(s.T(), 1, len(bucket.Sample.Hits.Hits))
		assert.NotEmpty(s.T(), bucket.Sample.Hits.Hits[0].Source.Name)
	}
}
