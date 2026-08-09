package es_test

import (
	"testing"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/Trendyol/es-query-builder/test/assert"
)

////   CompletionSuggester   ////

func Test_CompletionSuggester_should_exist_on_es_package(t *testing.T) {
	t.Parallel()
	// Given When Then
	assert.NotNil(t, es.CompletionSuggester)
}

func Test_CompletionSuggester_should_return_type_of_completionSuggesterType(t *testing.T) {
	t.Parallel()
	// Given
	s := es.CompletionSuggester("name.suggest")

	// When Then
	assert.NotNil(t, s)
	assert.IsTypeString(t, "es.completionSuggesterType", s)
	assert.MarshalWithoutError(t, s)
}

func Test_CompletionSuggester_should_create_json_with_completion_field(t *testing.T) {
	t.Parallel()
	// Given
	s := es.CompletionSuggester("name.suggest")

	// When Then
	assert.NotNil(t, s)
	bodyJSON := assert.MarshalWithoutError(t, s)
	assert.Equal(t, "{\"completion\":{\"field\":\"name.suggest\"}}", bodyJSON)
}

func Test_CompletionSuggester_Prefix_should_add_prefix_field(t *testing.T) {
	t.Parallel()
	// Given
	s := es.CompletionSuggester("name.suggest").Prefix("ni").Size(5)

	// When Then
	assert.NotNil(t, s)
	bodyJSON := assert.MarshalWithoutError(t, s)
	assert.Equal(t, "{\"completion\":{\"field\":\"name.suggest\",\"size\":5},\"prefix\":\"ni\"}", bodyJSON)
}

func Test_CompletionSuggester_Regex_should_add_regex_field(t *testing.T) {
	t.Parallel()
	// Given
	s := es.CompletionSuggester("name.suggest").Regex("n[iI]ke.*")

	// When Then
	assert.NotNil(t, s)
	bodyJSON := assert.MarshalWithoutError(t, s)
	assert.Equal(t, "{\"completion\":{\"field\":\"name.suggest\"},\"regex\":\"n[iI]ke.*\"}", bodyJSON)
}

func Test_CompletionSuggester_Fuzzy_should_add_fuzzy_field(t *testing.T) {
	t.Parallel()
	// Given
	s := es.CompletionSuggester("name.suggest").
		Prefix("ni").
		Fuzzy(
			es.CompletionFuzzy().
				Fuzziness(1).
				Transpositions(true).
				MinLength(3).
				PrefixLength(1).
				UnicodeAware(true),
		)

	// When Then
	assert.NotNil(t, s)
	bodyJSON := assert.MarshalWithoutError(t, s)
	// nolint:golint,lll
	assert.Equal(t, "{\"completion\":{\"field\":\"name.suggest\",\"fuzzy\":{\"fuzziness\":1,\"min_length\":3,\"prefix_length\":1,\"transpositions\":true,\"unicode_aware\":true}},\"prefix\":\"ni\"}", bodyJSON)
}

func Test_CompletionFuzzy_should_exist_on_es_package(t *testing.T) {
	t.Parallel()
	// Given When Then
	assert.NotNil(t, es.CompletionFuzzy)
}

func Test_CompletionFuzzy_should_create_json(t *testing.T) {
	t.Parallel()
	// Given
	f := es.CompletionFuzzy().Fuzziness("AUTO")

	// When Then
	assert.NotNil(t, f)
	assert.IsTypeString(t, "es.completionFuzzyType", f)
	bodyJSON := assert.MarshalWithoutError(t, f)
	assert.Equal(t, "{\"fuzziness\":\"AUTO\"}", bodyJSON)
}

func Test_CompletionSuggester_SkipDuplicates_and_Text_should_create_correct_json(t *testing.T) {
	t.Parallel()
	// Given
	s := es.CompletionSuggester("name.suggest").
		Text("nike").
		Size(5).
		SkipDuplicates(true)

	// When Then
	assert.NotNil(t, s)
	bodyJSON := assert.MarshalWithoutError(t, s)
	// nolint:golint,lll
	assert.Equal(t, "{\"completion\":{\"field\":\"name.suggest\",\"size\":5,\"skip_duplicates\":true},\"text\":\"nike\"}", bodyJSON)
}
