package es_test

import (
	"testing"

	SuggestMode "github.com/Trendyol/es-query-builder/es/enums/suggest-mode"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/Trendyol/es-query-builder/test/assert"
)

////   PhraseSuggester   ////

func Test_PhraseSuggester_should_exist_on_es_package(t *testing.T) {
	t.Parallel()
	// Given When Then
	assert.NotNil(t, es.PhraseSuggester)
}

func Test_PhraseSuggester_should_return_type_of_phraseSuggesterType(t *testing.T) {
	t.Parallel()
	// Given
	s := es.PhraseSuggester("message")

	// When Then
	assert.NotNil(t, s)
	assert.IsTypeString(t, "es.phraseSuggesterType", s)
	assert.MarshalWithoutError(t, s)
}

func Test_PhraseSuggester_should_create_json_with_phrase_field(t *testing.T) {
	t.Parallel()
	// Given
	s := es.PhraseSuggester("message").Text("noble prize").Size(1)

	// When Then
	assert.NotNil(t, s)
	bodyJSON := assert.MarshalWithoutError(t, s)
	assert.Equal(t, "{\"phrase\":{\"field\":\"message\",\"size\":1},\"text\":\"noble prize\"}", bodyJSON)
}

func Test_PhraseSuggester_all_options_should_create_correct_json(t *testing.T) {
	t.Parallel()
	// Given
	s := es.PhraseSuggester("message").
		Text("noble prize").
		Size(1).
		GramSize(3).
		Confidence(1.0).
		MaxErrors(2).
		RealWordErrorLikelihood(0.95).
		Separator(" ").
		DirectGenerator(
			es.DirectGenerator("message.trigram").
				SuggestMode(SuggestMode.Always).
				MinWordLength(3).
				PrefixLength(1).
				MaxEdits(2),
		)

	// When Then
	assert.NotNil(t, s)
	bodyJSON := assert.MarshalWithoutError(t, s)
	// nolint:golint,lll
	assert.Equal(t, "{\"phrase\":{\"confidence\":1,\"direct_generator\":[{\"field\":\"message.trigram\",\"max_edits\":2,\"min_word_length\":3,\"prefix_length\":1,\"suggest_mode\":\"always\"}],\"field\":\"message\",\"gram_size\":3,\"max_errors\":2,\"real_word_error_likelihood\":0.95,\"separator\":\" \",\"size\":1},\"text\":\"noble prize\"}", bodyJSON)
}

func Test_DirectGenerator_should_exist_on_es_package(t *testing.T) {
	t.Parallel()
	// Given When Then
	assert.NotNil(t, es.DirectGenerator)
}

func Test_DirectGenerator_should_create_json(t *testing.T) {
	t.Parallel()
	// Given
	d := es.DirectGenerator("message.trigram")

	// When Then
	assert.NotNil(t, d)
	assert.IsTypeString(t, "es.directGeneratorType", d)
	bodyJSON := assert.MarshalWithoutError(t, d)
	assert.Equal(t, "{\"field\":\"message.trigram\"}", bodyJSON)
}
