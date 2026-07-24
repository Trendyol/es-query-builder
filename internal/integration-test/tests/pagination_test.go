package tests_test

import (
	"github.com/Trendyol/es-query-builder/es"
	Order "github.com/Trendyol/es-query-builder/es/enums/sort/order"
	"github.com/stretchr/testify/assert"
)

func (s *testSuite) Test_it_should_paginate_results_with_from_and_size() {
	// Given
	firstPageQuery := es.NewQuery(es.MatchAll()).
		From(0).
		Size(3).
		Sort(es.Sort("id").Order(Order.Asc))

	firstPage, err := s.PokedexElasticsearchRepository.Search(s.TestContext, firstPageQuery)
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), 3, len(firstPage.Hits.Hits))

	// When
	secondPageQuery := es.NewQuery(es.MatchAll()).
		From(3).
		Size(3).
		Sort(es.Sort("id").Order(Order.Asc))

	secondPage, err := s.PokedexElasticsearchRepository.Search(s.TestContext, secondPageQuery)

	// Then
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), 3, len(secondPage.Hits.Hits))

	firstIDs := map[string]struct{}{
		firstPage.Hits.Hits[0].Id: {},
		firstPage.Hits.Hits[1].Id: {},
		firstPage.Hits.Hits[2].Id: {},
	}
	for _, hit := range secondPage.Hits.Hits {
		_, exists := firstIDs[hit.Id]
		assert.False(s.T(), exists, "document %s appeared in both pages", hit.Id)
	}
}

func (s *testSuite) Test_it_should_track_total_hits() {
	// Given
	query := es.NewQuery(es.MatchAll()).
		Size(1).
		TrackTotalHits(true)

	// When
	response, err := s.PokedexElasticsearchRepository.Search(s.TestContext, query)

	// Then
	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), response.Hits)
	assert.NotNil(s.T(), response.Hits.Total)
	assert.Greater(s.T(), response.Hits.Total.Value, int64(1))
	assert.Equal(s.T(), "eq", response.Hits.Total.Relation)
}
