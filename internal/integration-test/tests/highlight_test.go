package tests_test

import (
	"github.com/Trendyol/es-query-builder/es"
	"github.com/stretchr/testify/assert"
)

func (s *testSuite) Test_it_should_return_highlight_snippets_for_matched_fields() {
	// Given
	query := es.NewQuery(es.Match("name", "pikachu")).
		Highlight(
			es.Highlight().
				PreTags("<em>").
				PostTags("</em>").
				Field(es.HighlightField("name")),
		)

	// When
	response, err := s.PokedexElasticsearchRepository.Search(s.TestContext, query)

	// Then
	assert.Nil(s.T(), err)
	assert.NotNil(s.T(), response.Hits)
	assert.Greater(s.T(), len(response.Hits.Hits), 0)

	hit := response.Hits.Hits[0]
	assert.NotEmpty(s.T(), hit.Highlight)
	snippets, ok := hit.Highlight["name"]
	assert.True(s.T(), ok)
	assert.Greater(s.T(), len(snippets), 0)
	assert.Contains(s.T(), snippets[0], "<em>")
}
