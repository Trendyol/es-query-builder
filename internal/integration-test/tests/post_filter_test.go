package tests_test

import (
	"github.com/Trendyol/es-query-builder/es"
	"github.com/stretchr/testify/assert"
)

func (s *testSuite) Test_it_should_apply_post_filter_without_affecting_query_matches() {
	// Given — query matches electric-type-ish names via wildcard; post_filter keeps only pikachu
	query := es.NewQuery(
		es.Bool().Should(
			es.Term("name.keyword", "pikachu"),
			es.Term("name.keyword", "pichu"),
			es.Term("name.keyword", "raichu"),
		),
	).PostFilter(es.Term("name.keyword", "pikachu"))

	// When
	result, err := s.PokedexElasticsearchRepository.GetSearchHits(s.TestContext, query)

	// Then
	assert.Nil(s.T(), err)
	assert.Equal(s.T(), 1, len(result))
	assert.Equal(s.T(), "pikachu", result["25_35"].Name)
}
