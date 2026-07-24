package tests_test

import (
	"encoding/json"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/stretchr/testify/assert"
)

func (s *testSuite) Test_it_should_include_only_requested_source_fields() {
	// Given
	query := es.NewQuery(es.Term("name.keyword", "pikachu")).
		SourceIncludes("name", "id")

	// When
	response, err := s.PokedexElasticsearchRepository.Search(s.TestContext, query)

	// Then
	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), response.Hits)
	assert.Equal(s.T(), 1, len(response.Hits.Hits))

	var source map[string]any
	assert.NoError(s.T(), json.Unmarshal(response.Hits.Hits[0].Source, &source))
	assert.Contains(s.T(), source, "name")
	assert.Contains(s.T(), source, "id")
	assert.NotContains(s.T(), source, "weight")
	assert.NotContains(s.T(), source, "abilities")
}

func (s *testSuite) Test_it_should_exclude_requested_source_fields() {
	// Given
	query := es.NewQuery(es.Term("name.keyword", "pikachu")).
		SourceExcludes("abilities", "moves", "stats")

	// When
	response, err := s.PokedexElasticsearchRepository.Search(s.TestContext, query)

	// Then
	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), response.Hits)
	assert.Equal(s.T(), 1, len(response.Hits.Hits))

	var source map[string]any
	assert.NoError(s.T(), json.Unmarshal(response.Hits.Hits[0].Source, &source))
	assert.Contains(s.T(), source, "name")
	assert.NotContains(s.T(), source, "abilities")
	assert.NotContains(s.T(), source, "moves")
	assert.NotContains(s.T(), source, "stats")
}

func (s *testSuite) Test_it_should_omit_source_when_source_false() {
	// Given
	query := es.NewQuery(es.Term("name.keyword", "pikachu")).
		SourceFalse()

	// When
	response, err := s.PokedexElasticsearchRepository.Search(s.TestContext, query)

	// Then
	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), response.Hits)
	assert.Equal(s.T(), 1, len(response.Hits.Hits))
	assert.True(s.T(), len(response.Hits.Hits[0].Source) == 0 || string(response.Hits.Hits[0].Source) == "null")
}
