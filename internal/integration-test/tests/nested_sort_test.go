package tests_test

import (
	Mode "github.com/Trendyol/es-query-builder/es/enums/sort/mode"
	Order "github.com/Trendyol/es-query-builder/es/enums/sort/order"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/stretchr/testify/assert"
)

func (s *testSuite) Test_it_should_sort_documents_with_nested_sort() {
	// Given
	query := es.NewQuery(es.MatchAll()).
		Size(5).
		Sort(
			es.Sort("abilities.slot").
				Order(Order.Asc).
				Mode(Mode.Min).
				Nested(es.NestedSort("abilities")),
		)

	// When
	response, err := s.PokedexElasticsearchRepository.Search(s.TestContext, query)

	// Then
	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), response.Hits)
	assert.Equal(s.T(), 5, len(response.Hits.Hits))
	assert.NotEmpty(s.T(), response.Hits.Hits[0].Sort)
}
