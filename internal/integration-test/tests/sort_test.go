package tests_test

import (
	"encoding/json"

	Order "github.com/Trendyol/es-query-builder/es/enums/sort/order"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/stretchr/testify/assert"
)

func (s *testSuite) Test_it_should_sort_documents_by_id_descending() {
	// Given
	query := es.NewQuery(es.MatchAll()).
		Size(5).
		Sort(es.Sort("id").Order(Order.Desc))

	// When
	response, err := s.PokedexElasticsearchRepository.Search(s.TestContext, query)

	// Then
	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), response.Hits)
	assert.Equal(s.T(), 5, len(response.Hits.Hits))

	var previousID float64
	for i, hit := range response.Hits.Hits {
		assert.NotEmpty(s.T(), hit.Sort)
		id, ok := hit.Sort[0].(float64)
		assert.True(s.T(), ok)
		if i > 0 {
			assert.GreaterOrEqual(s.T(), previousID, id)
		}
		previousID = id
	}
}

func (s *testSuite) Test_it_should_sort_documents_by_multiple_fields() {
	// Given
	query := es.NewQuery(es.MatchAll()).
		Size(10).
		Sort(
			es.Sort("isDefault").Order(Order.Desc),
			es.Sort("id").Order(Order.Asc),
		)

	// When
	response, err := s.PokedexElasticsearchRepository.Search(s.TestContext, query)

	// Then
	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), response.Hits)
	assert.Equal(s.T(), 10, len(response.Hits.Hits))

	for _, hit := range response.Hits.Hits {
		assert.Equal(s.T(), 2, len(hit.Sort))
		assert.NotEmpty(s.T(), hit.Source)
		var raw json.RawMessage
		assert.NoError(s.T(), json.Unmarshal(hit.Source, &raw))
	}
}
