package es_test

import (
	"testing"

	StringDistance "github.com/Trendyol/es-query-builder/es/enums/string-distance"
	SuggestMode "github.com/Trendyol/es-query-builder/es/enums/suggest-mode"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/Trendyol/es-query-builder/test/assert"
)

////   TermSuggester   ////

func Test_TermSuggester_should_exist_on_es_package(t *testing.T) {
	t.Parallel()
	// Given When Then
	assert.NotNil(t, es.TermSuggester)
}

func Test_TermSuggester_should_return_type_of_termSuggesterType(t *testing.T) {
	t.Parallel()
	// Given
	s := es.TermSuggester("message")

	// When Then
	assert.NotNil(t, s)
	assert.IsTypeString(t, "es.termSuggesterType", s)
	assert.MarshalWithoutError(t, s)
}

func Test_TermSuggester_should_create_json_with_term_field(t *testing.T) {
	t.Parallel()
	// Given
	s := es.TermSuggester("message")

	// When Then
	assert.NotNil(t, s)
	bodyJSON := assert.MarshalWithoutError(t, s)
	assert.Equal(t, "{\"term\":{\"field\":\"message\"}}", bodyJSON)
}

func Test_TermSuggester_Text_should_add_text_field(t *testing.T) {
	t.Parallel()
	// Given
	s := es.TermSuggester("message").Text("triyng")

	// When Then
	assert.NotNil(t, s)
	bodyJSON := assert.MarshalWithoutError(t, s)
	assert.Equal(t, "{\"term\":{\"field\":\"message\"},\"text\":\"triyng\"}", bodyJSON)
}

func Test_TermSuggester_Size_should_add_size_field(t *testing.T) {
	t.Parallel()
	// Given
	s := es.TermSuggester("message").Size(5)

	// When Then
	assert.NotNil(t, s)
	bodyJSON := assert.MarshalWithoutError(t, s)
	assert.Equal(t, "{\"term\":{\"field\":\"message\",\"size\":5}}", bodyJSON)
}

func Test_TermSuggester_SuggestMode_should_add_suggest_mode_field(t *testing.T) {
	t.Parallel()
	// Given
	s := es.TermSuggester("message").SuggestMode(SuggestMode.Popular)

	// When Then
	assert.NotNil(t, s)
	bodyJSON := assert.MarshalWithoutError(t, s)
	assert.Equal(t, "{\"term\":{\"field\":\"message\",\"suggest_mode\":\"popular\"}}", bodyJSON)
}

func Test_TermSuggester_all_options_should_create_correct_json(t *testing.T) {
	t.Parallel()
	// Given
	s := es.TermSuggester("message").
		Text("triyng").
		Size(5).
		SuggestMode(SuggestMode.Always).
		MinWordLength(4).
		MaxEdits(2).
		PrefixLength(1).
		MinDocFreq(0.01).
		MaxTermFreq(0.01).
		StringDistance(StringDistance.Levenshtein).
		ShardSize(5)

	// When Then
	assert.NotNil(t, s)
	bodyJSON := assert.MarshalWithoutError(t, s)
	// nolint:golint,lll
	assert.Equal(t, "{\"term\":{\"field\":\"message\",\"max_edits\":2,\"max_term_freq\":0.01,\"min_doc_freq\":0.01,\"min_word_length\":4,\"prefix_length\":1,\"shard_size\":5,\"size\":5,\"string_distance\":\"levenshtein\",\"suggest_mode\":\"always\"},\"text\":\"triyng\"}", bodyJSON)
}
