package tests_test

import (
	"github.com/Trendyol/es-query-builder/es"
	"github.com/stretchr/testify/assert"
)

func (s *testSuite) Test_it_should_return_all_documents_with_match_all_query() {
	// Given
	query := es.NewQuery(
		es.MatchAll(),
	).Size(300)

	// When
	result, err := s.PokedexElasticsearchRepository.GetSearchHits(s.TestContext, query)

	// Then
	assert.Nil(s.T(), err)
	assert.NoError(s.T(), err)
	assert.Greater(s.T(), len(result), 200) // Pokemon datasında 200+ pokemon olmalı
}
