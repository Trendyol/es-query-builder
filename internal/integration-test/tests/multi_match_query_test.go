package tests_test

import (
	"github.com/Trendyol/es-query-builder/es"
	"github.com/stretchr/testify/assert"
)

func (s *testSuite) Test_it_should_return_documents_that_filtered_by_multi_match_query() {
	// Given
	query := es.NewQuery(
		es.MultiMatch("pikachu").
			Fields("name", "abilities.name"),
	)

	// When
	result, err := s.PokedexElasticsearchRepository.GetSearchHits(s.TestContext, query)

	// Then
	assert.Nil(s.T(), err)
	assert.NoError(s.T(), err)
	assert.GreaterOrEqual(s.T(), len(result), 1)

	// Pikachu olmalı
	pokemon, exists := result["25_35"]
	assert.True(s.T(), exists)
	assert.Equal(s.T(), "pikachu", pokemon.Name)
}
