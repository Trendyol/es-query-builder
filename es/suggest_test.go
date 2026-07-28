package es_test

import (
	"testing"

	StringDistance "github.com/Trendyol/es-query-builder/es/enums/string-distance"
	SuggestMode "github.com/Trendyol/es-query-builder/es/enums/suggest-mode"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/Trendyol/es-query-builder/test/assert"
)

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
