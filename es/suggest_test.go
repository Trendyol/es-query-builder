package es_test

import (
	"testing"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/Trendyol/es-query-builder/test/assert"
)

////   Suggest   ////

func Test_Suggest_should_exist_on_es_package(t *testing.T) {
	t.Parallel()
	// Given When Then
	assert.NotNil(t, es.Suggest)
}

func Test_Suggest_method_should_create_suggestType(t *testing.T) {
	t.Parallel()
	// Given
	s := es.Suggest()

	// Then
	assert.NotNil(t, s)
	assert.IsTypeString(t, "es.suggestType", s)
}

func Test_Object_should_have_Suggest_method(t *testing.T) {
	t.Parallel()
	// Given
	b := es.NewQuery(nil)

	// When Then
	assert.NotNil(t, b.Suggest)
}

func Test_Suggest_should_add_suggest_field_into_Object(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(es.MatchAll()).
		Suggest(es.Suggest())

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"match_all\":{}},\"suggest\":{}}", bodyJSON)
}

func Test_Suggest_Text_should_add_text_field(t *testing.T) {
	t.Parallel()
	// Given
	s := es.Suggest().Text("nike shos")

	// When Then
	assert.NotNil(t, s)
	bodyJSON := assert.MarshalWithoutError(t, s)
	assert.Equal(t, "{\"text\":\"nike shos\"}", bodyJSON)
}

func Test_Suggest_Suggester_should_add_named_suggester(t *testing.T) {
	t.Parallel()
	// Given
	s := es.Suggest().
		Text("nike shos").
		Suggester("product_suggest",
			es.CompletionSuggester("name.suggest").Size(5).SkipDuplicates(true),
		)

	// When Then
	assert.NotNil(t, s)
	bodyJSON := assert.MarshalWithoutError(t, s)
	// nolint:golint,lll
	assert.Equal(t, "{\"product_suggest\":{\"completion\":{\"field\":\"name.suggest\",\"size\":5,\"skip_duplicates\":true}},\"text\":\"nike shos\"}", bodyJSON)
}

func Test_Suggest_with_NewQuery_should_create_correct_json(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(es.MatchAll()).
		Suggest(
			es.Suggest().
				Text("nike shos").
				Suggester("product_suggest",
					es.CompletionSuggester("name.suggest").Size(5).SkipDuplicates(true),
				),
		)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	// nolint:golint,lll
	assert.Equal(t, "{\"query\":{\"match_all\":{}},\"suggest\":{\"product_suggest\":{\"completion\":{\"field\":\"name.suggest\",\"size\":5,\"skip_duplicates\":true}},\"text\":\"nike shos\"}}", bodyJSON)
}
