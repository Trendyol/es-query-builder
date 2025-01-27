package es_test

import (
	"testing"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/Trendyol/es-query-builder/test/assert"
)

////   Term   ////

func Test_Term_should_exist_on_es_package(t *testing.T) {
	t.Parallel()
	// Given When Then
	assert.NotNil(t, es.Term[any])
}

func Test_Term_should_create_json_with_term_field_inside_query(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.Term("key", "value"),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"term\":{\"key\":{\"value\":\"value\"}}}}", bodyJSON)
}

func Test_Term_method_should_create_termType(t *testing.T) {
	t.Parallel()
	// Given
	b := es.Term("key", "value")

	// Then
	assert.NotNil(t, b)
	assert.IsTypeString(t, "es.termType", b)
}

func Test_Term_should_have_CaseInsensitive_method(t *testing.T) {
	t.Parallel()
	// Given
	term := es.Term("key", "value")

	// When Then
	assert.NotNil(t, term.CaseInsensitive)
}

func Test_Term_CaseInsensitive_should_create_json_with_case_insensitive_field_inside_term(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.Term("type", "File").
			CaseInsensitive(true),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"term\":{\"type\":{\"case_insensitive\":true,\"value\":\"File\"}}}}", bodyJSON)
}

func Test_Term_should_have_Boost_method(t *testing.T) {
	t.Parallel()
	// Given
	term := es.Term("key", "value")

	// When Then
	assert.NotNil(t, term.Boost)
}

func Test_Term_Boost_should_create_json_with_boost_field_inside_term(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.Term("type", "File").
			Boost(3.14),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"term\":{\"type\":{\"boost\":3.14,\"value\":\"File\"}}}}", bodyJSON)
}

////   TermFunc   ////

func Test_TermFunc_should_exist_on_es_package(t *testing.T) {
	t.Parallel()
	// Given When Then
	assert.NotNil(t, es.TermFunc[any])
}

func Test_TermFunc_should_create_json_with_term_field_inside_query(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.TermFunc("key", "value", func(key string, value string) bool {
			return true
		}),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"term\":{\"key\":{\"value\":\"value\"}}}}", bodyJSON)
}

func Test_TermFunc_should_not_add_term_field_inside_query_when_callback_result_is_false(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.TermFunc("key", "value", func(key string, value string) bool {
			return false
		}),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{}}", bodyJSON)
}

func Test_TermFunc_should_add_only_term_fields_inside_the_query_when_callback_result_is_true(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.Bool().
			Filter(
				es.TermFunc("a", "b", func(key string, value string) bool {
					return true
				}),
				es.TermFunc("c", "d", func(key string, value string) bool {
					return false
				}),
				es.TermFunc("e", "f", func(key string, value string) bool {
					return true
				}),
			),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	// nolint:golint,lll
	assert.Equal(t, "{\"query\":{\"bool\":{\"filter\":[{\"term\":{\"a\":{\"value\":\"b\"}}},{\"term\":{\"e\":{\"value\":\"f\"}}}]}}}", bodyJSON)
}

func Test_TermFunc_method_should_create_termType(t *testing.T) {
	t.Parallel()
	// Given
	b := es.TermFunc("key", "value", func(key string, value string) bool {
		return true
	})

	// Then
	assert.NotNil(t, b)
	assert.IsTypeString(t, "es.termType", b)
}

////   TermIf   ////

func Test_TermIf_should_exist_on_es_package(t *testing.T) {
	t.Parallel()
	// Given When Then
	assert.NotNil(t, es.TermIf[any])
}

func Test_TermIf_should_create_json_with_term_field_inside_query(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.TermIf("key", "value", true),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"term\":{\"key\":{\"value\":\"value\"}}}}", bodyJSON)
}

func Test_TermIf_should_not_add_term_field_inside_query_when_condition_is_false(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.TermIf("key", "value", false),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{}}", bodyJSON)
}

func Test_TermIf_should_add_only_term_fields_inside_the_query_when_condition_is_true(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.Bool().
			Filter(
				es.TermIf("a", "b", true),
				es.TermIf("c", "d", false),
				es.TermIf("e", "f", true),
			),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	// nolint:golint,lll
	assert.Equal(t, "{\"query\":{\"bool\":{\"filter\":[{\"term\":{\"a\":{\"value\":\"b\"}}},{\"term\":{\"e\":{\"value\":\"f\"}}}]}}}", bodyJSON)
}

func Test_TermIf_method_should_create_termType(t *testing.T) {
	t.Parallel()
	// Given
	b := es.TermIf("key", "value", true)

	// Then
	assert.NotNil(t, b)
	assert.IsTypeString(t, "es.termType", b)
}
