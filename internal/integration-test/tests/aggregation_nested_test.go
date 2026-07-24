package tests_test

import (
	"encoding/json"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/stretchr/testify/assert"
)

func (s *testSuite) Test_it_should_return_nested_terms_aggregation_on_abilities() {
	// Given
	query := es.NewQuery(es.MatchAll()).
		Size(0).
		Aggs(
			es.Agg("abilities_nested",
				es.NestedAgg("abilities").Aggs(
					es.Agg("by_ability_name", es.TermsAgg("abilities.name").Size(10)),
				),
			),
		)

	// When
	response, err := s.PokedexElasticsearchRepository.Search(s.TestContext, query)

	// Then
	assert.Nil(s.T(), err)
	raw, exists := response.Aggregations["abilities_nested"]
	assert.True(s.T(), exists)

	var nestedAgg struct {
		DocCount      int `json:"doc_count"`
		ByAbilityName struct {
			Buckets []struct {
				Key      string `json:"key"`
				DocCount int    `json:"doc_count"`
			} `json:"buckets"`
		} `json:"by_ability_name"`
	}
	assert.NoError(s.T(), json.Unmarshal(raw, &nestedAgg))
	assert.Greater(s.T(), nestedAgg.DocCount, 0)
	assert.Greater(s.T(), len(nestedAgg.ByAbilityName.Buckets), 0)
	assert.NotEmpty(s.T(), nestedAgg.ByAbilityName.Buckets[0].Key)
}

func (s *testSuite) Test_it_should_return_reverse_nested_aggregation() {
	// Given
	query := es.NewQuery(es.MatchAll()).
		Size(0).
		Aggs(
			es.Agg("abilities_nested",
				es.NestedAgg("abilities").Aggs(
					es.Agg("by_ability_name",
						es.TermsAgg("abilities.name").Size(5).Aggs(
							es.Agg("to_pokemon",
								es.ReverseNestedAgg().Aggs(
									es.Agg("unique_pokemon", es.CardinalityAgg("name.keyword")),
								),
							),
						),
					),
				),
			),
		)

	// When
	response, err := s.PokedexElasticsearchRepository.Search(s.TestContext, query)

	// Then
	assert.Nil(s.T(), err)
	raw, exists := response.Aggregations["abilities_nested"]
	assert.True(s.T(), exists)

	var nestedAgg struct {
		ByAbilityName struct {
			Buckets []struct {
				Key       string `json:"key"`
				DocCount  int    `json:"doc_count"`
				ToPokemon struct {
					DocCount      int `json:"doc_count"`
					UniquePokemon struct {
						Value float64 `json:"value"`
					} `json:"unique_pokemon"`
				} `json:"to_pokemon"`
			} `json:"buckets"`
		} `json:"by_ability_name"`
	}
	assert.NoError(s.T(), json.Unmarshal(raw, &nestedAgg))
	assert.Greater(s.T(), len(nestedAgg.ByAbilityName.Buckets), 0)

	for _, bucket := range nestedAgg.ByAbilityName.Buckets {
		assert.NotEmpty(s.T(), bucket.Key)
		assert.Greater(s.T(), bucket.ToPokemon.DocCount, 0)
		assert.Greater(s.T(), bucket.ToPokemon.UniquePokemon.Value, float64(0))
	}
}
