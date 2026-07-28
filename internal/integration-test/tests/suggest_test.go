package tests_test

import (
	"encoding/json"
	"integration-tests/model"
	"integration-tests/tests"

	SuggestMode "github.com/Trendyol/es-query-builder/es/enums/suggest-mode"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/stretchr/testify/assert"
)

func (s *testSuite) Test_it_should_return_term_suggestions() {
	// Given
	query := es.NewQuery(es.MatchAll()).
		Size(0).
		Suggest(
			es.Suggest().
				Text("pikacu").
				Suggester("name_suggest",
					es.TermSuggester("name.keyword").
						SuggestMode(SuggestMode.Always).
						Size(5),
				),
		)

	// When
	response, err := s.PokedexElasticsearchRepository.Search(s.TestContext, query)

	// Then
	assert.Nil(s.T(), err)
	raw, exists := response.Suggest["name_suggest"]
	assert.True(s.T(), exists)

	var suggestions []struct {
		Text    string `json:"text"`
		Options []struct {
			Text  string  `json:"text"`
			Score float64 `json:"score"`
		} `json:"options"`
	}
	assert.NoError(s.T(), json.Unmarshal(raw, &suggestions))
	assert.Greater(s.T(), len(suggestions), 0)
	assert.Equal(s.T(), "pikacu", suggestions[0].Text)
	assert.Greater(s.T(), len(suggestions[0].Options), 0)

	found := false
	for _, option := range suggestions[0].Options {
		if option.Text == "pikachu" {
			found = true
			break
		}
	}
	assert.True(s.T(), found, "expected pikachu among term suggestions, got %#v", suggestions[0].Options)
}

func (s *testSuite) Test_it_should_return_completion_suggestions() {
	// Given
	nike := model.FooDocument{
		ID:  "suggest-1",
		Foo: "nike",
		Suggest: &model.CompletionSuggest{
			Input:  []string{"nike", "nike shoes"},
			Weight: 10,
		},
	}
	adidas := model.FooDocument{
		ID:  "suggest-2",
		Foo: "adidas",
		Suggest: &model.CompletionSuggest{
			Input:  []string{"adidas", "adidas shoes"},
			Weight: 5,
		},
	}
	puma := model.FooDocument{
		ID:  "suggest-3",
		Foo: "puma",
		Suggest: &model.CompletionSuggest{
			Input:  []string{"puma", "puma shoes"},
			Weight: 3,
		},
	}

	s.FooElasticsearchRepository.BulkInsert(s.TestContext, nike, adidas, puma)
	tests.WaitExists(s.TestContext, s.FooElasticsearchRepository, nike.ID)
	tests.WaitExists(s.TestContext, s.FooElasticsearchRepository, adidas.ID)
	tests.WaitExists(s.TestContext, s.FooElasticsearchRepository, puma.ID)

	query := es.NewQuery(es.MatchAll()).
		Size(0).
		Suggest(
			es.Suggest().
				Suggester("brand_suggest",
					es.CompletionSuggester("suggest").
						Prefix("ni").
						Size(5).
						SkipDuplicates(true),
				),
		)

	// When
	response, err := s.FooElasticsearchRepository.Search(s.TestContext, query)

	// Then
	assert.Nil(s.T(), err)
	raw, exists := response.Suggest["brand_suggest"]
	assert.True(s.T(), exists)

	var suggestions []struct {
		Text    string `json:"text"`
		Options []struct {
			Text  string  `json:"text"`
			Score float64 `json:"score"`
		} `json:"options"`
	}
	assert.NoError(s.T(), json.Unmarshal(raw, &suggestions))
	assert.Greater(s.T(), len(suggestions), 0)
	assert.Equal(s.T(), "ni", suggestions[0].Text)
	assert.Greater(s.T(), len(suggestions[0].Options), 0)
	assert.Contains(s.T(), suggestions[0].Options[0].Text, "nike")

	s.FooElasticsearchRepository.BulkDelete(s.TestContext, nike.ID, adidas.ID, puma.ID)
}

func (s *testSuite) Test_it_should_return_phrase_suggestions() {
	// Given
	query := es.NewQuery(es.MatchAll()).
		Size(0).
		Suggest(
			es.Suggest().
				Text("pikacu").
				Suggester("phrase_suggest",
					es.PhraseSuggester("name.keyword").
						Size(1).
						Confidence(0.0).
						MaxErrors(2),
				),
		)

	// When
	response, err := s.PokedexElasticsearchRepository.Search(s.TestContext, query)

	// Then
	assert.Nil(s.T(), err)
	raw, exists := response.Suggest["phrase_suggest"]
	assert.True(s.T(), exists)

	var suggestions []struct {
		Text    string `json:"text"`
		Options []struct {
			Text  string  `json:"text"`
			Score float64 `json:"score"`
		} `json:"options"`
	}
	assert.NoError(s.T(), json.Unmarshal(raw, &suggestions))
	assert.Greater(s.T(), len(suggestions), 0)
	assert.Equal(s.T(), "pikacu", suggestions[0].Text)
	assert.Greater(s.T(), len(suggestions[0].Options), 0)
	assert.Equal(s.T(), "pikachu", suggestions[0].Options[0].Text)
}
