package tests_test

import (
	"github.com/Trendyol/es-query-builder/es"
	"github.com/stretchr/testify/assert"
)

func (s *testSuite) Test_it_should_return_documents_with_constant_score_query() {
	// Given
	query := es.NewQuery(
		es.ConstantScore(
			es.Term("name.keyword", "pikachu"),
		).Boost(2.0),
	)

	// When
	result, err := s.PokedexElasticsearchRepository.GetSearchHits(s.TestContext, query)

	// Then
	assert.Nil(s.T(), err)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), 1, len(result))

	pokemon, exists := result["25_35"]
	assert.True(s.T(), exists)
	assert.Equal(s.T(), "pikachu", pokemon.Name)
}

func (s *testSuite) Test_it_should_return_documents_with_constant_score_and_range_filter() {
	// Given
	query := es.NewQuery(
		es.ConstantScore(
			es.Range("baseExperience").GreaterThanOrEqual(250),
		),
	).Size(50)

	// When
	result, err := s.PokedexElasticsearchRepository.GetSearchHits(s.TestContext, query)

	// Then
	assert.Nil(s.T(), err)
	assert.NoError(s.T(), err)
	assert.Greater(s.T(), len(result), 0)

	// Tüm sonuçlar baseExperience >= 250 olmalı
	for _, pokemon := range result {
		assert.GreaterOrEqual(s.T(), pokemon.BaseExperience, uint16(250))
	}
}
