package tests_test

import (
	"github.com/Trendyol/es-query-builder/es"
	"github.com/stretchr/testify/assert"
)

func (s *testSuite) Test_it_should_return_documents_that_filtered_by_match_phrase_query() {
	// Given
	query := es.NewQuery(
		es.MatchPhrase("name", "pikachu"),
	)

	// When
	result, err := s.PokedexElasticsearchRepository.GetSearchHits(s.TestContext, query)

	// Then
	assert.Nil(s.T(), err)
	assert.NoError(s.T(), err)
	assert.Equal(s.T(), 1, len(result))
	for pokeID, pokemon := range result {
		switch pokeID {
		case "25_35":
			assert.Equal(s.T(), "pikachu", pokemon.Name)
		default:
			s.T().FailNow()
		}
	}
}
