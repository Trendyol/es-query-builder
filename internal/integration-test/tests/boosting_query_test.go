package tests_test

import (
	"github.com/Trendyol/es-query-builder/es"
	"github.com/stretchr/testify/assert"
)

func (s *testSuite) Test_it_should_demote_documents_matching_negative_boosting_query() {
	// Given
	query := es.NewQuery(
		es.Boosting(
			es.Terms("name.keyword", "pikachu", "bulbasaur"),
			es.Term("name.keyword", "pikachu"),
			0.2,
		),
	)

	// When
	response, err := s.PokedexElasticsearchRepository.Search(s.TestContext, query)

	// Then
	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), response.Hits)
	assert.Len(s.T(), response.Hits.Hits, 2)
	assert.Equal(s.T(), "1_1", response.Hits.Hits[0].Id)
	assert.Equal(s.T(), "25_35", response.Hits.Hits[1].Id)
	assert.Greater(s.T(), response.Hits.Hits[0].Score, response.Hits.Hits[1].Score)
}
