package tests_test

import (
	"encoding/json"

	Order "github.com/Trendyol/es-query-builder/es/enums/sort/order"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/stretchr/testify/assert"
)

func (s *testSuite) Test_it_should_paginate_results_with_search_after() {
	// Given — first page
	firstPageQuery := es.NewQuery(es.MatchAll()).
		Size(2).
		Sort(es.Sort("id").Order(Order.Asc))

	firstPage, err := s.PokedexElasticsearchRepository.Search(s.TestContext, firstPageQuery)
	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), firstPage.Hits)
	assert.Equal(s.T(), 2, len(firstPage.Hits.Hits))

	lastHit := firstPage.Hits.Hits[len(firstPage.Hits.Hits)-1]
	assert.NotEmpty(s.T(), lastHit.Sort)

	// When — second page via search_after
	secondPageQuery := es.NewQuery(es.MatchAll()).
		Size(2).
		Sort(es.Sort("id").Order(Order.Asc)).
		SearchAfter(lastHit.Sort...)

	secondPage, err := s.PokedexElasticsearchRepository.Search(s.TestContext, secondPageQuery)
	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), secondPage.Hits)
	assert.Equal(s.T(), 2, len(secondPage.Hits.Hits))

	// Then — pages must not overlap
	firstIDs := map[string]struct{}{
		firstPage.Hits.Hits[0].Id: {},
		firstPage.Hits.Hits[1].Id: {},
	}
	for _, hit := range secondPage.Hits.Hits {
		_, exists := firstIDs[hit.Id]
		assert.False(s.T(), exists, "document %s appeared in both pages", hit.Id)
	}

	// sanity: sources decode
	var pokemon json.RawMessage
	assert.NotEmpty(s.T(), secondPage.Hits.Hits[0].Source)
	assert.NoError(s.T(), json.Unmarshal(secondPage.Hits.Hits[0].Source, &pokemon))
}
