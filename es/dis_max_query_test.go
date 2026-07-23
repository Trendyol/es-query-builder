package es_test

import (
	"testing"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/Trendyol/es-query-builder/test/assert"
)

////   DisMax   ////

func Test_DisMax_should_exist_on_es_package(t *testing.T) {
	t.Parallel()
	// Given When Then
	assert.NotNil(t, es.DisMax)
}

func Test_DisMax_should_create_json_with_dis_max_field_inside_query(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.DisMax(
			es.Term("title", "quick"),
			es.Term("body", "brown"),
		),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	// nolint:golint,lll
	assert.Equal(t, "{\"query\":{\"dis_max\":{\"queries\":[{\"term\":{\"title\":{\"value\":\"quick\"}}},{\"term\":{\"body\":{\"value\":\"brown\"}}}]}}}", bodyJSON)
}

func Test_DisMax_method_should_create_disMaxType(t *testing.T) {
	t.Parallel()
	// Given
	d := es.DisMax(es.Term("title", "quick"))

	// Then
	assert.NotNil(t, d)
	assert.IsTypeString(t, "es.disMaxType", d)
}

func Test_DisMax_should_skip_nil_queries(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.DisMax(
			es.Term("title", "quick"),
			nil,
			es.TermIf("body", "brown", false),
		),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"dis_max\":{\"queries\":[{\"term\":{\"title\":{\"value\":\"quick\"}}}]}}}", bodyJSON)
}

func Test_DisMax_should_have_TieBreaker_method(t *testing.T) {
	t.Parallel()
	// Given
	d := es.DisMax(es.Term("title", "quick"))

	// When Then
	assert.NotNil(t, d.TieBreaker)
}

func Test_DisMax_TieBreaker_should_create_json_with_tie_breaker_field(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.DisMax(es.Term("title", "quick")).
			TieBreaker(0.7),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"dis_max\":{\"queries\":[{\"term\":{\"title\":{\"value\":\"quick\"}}}],\"tie_breaker\":0.7}}}", bodyJSON)
}

func Test_DisMax_should_have_Boost_method(t *testing.T) {
	t.Parallel()
	// Given
	d := es.DisMax(es.Term("title", "quick"))

	// When Then
	assert.NotNil(t, d.Boost)
}

func Test_DisMax_Boost_should_create_json_with_boost_field(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.DisMax(es.Term("title", "quick")).
			Boost(1.2),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"dis_max\":{\"boost\":1.2,\"queries\":[{\"term\":{\"title\":{\"value\":\"quick\"}}}]}}}", bodyJSON)
}

func Test_DisMax_should_have_Name_method(t *testing.T) {
	t.Parallel()
	// Given
	d := es.DisMax(es.Term("title", "quick"))

	// When Then
	assert.NotNil(t, d.Name)
}

func Test_DisMax_Name_should_create_json_with_name_field(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.DisMax(es.Term("title", "quick")).
			Name("best_fields"),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	// nolint:golint,lll
	assert.Equal(t, "{\"query\":{\"dis_max\":{\"_name\":\"best_fields\",\"queries\":[{\"term\":{\"title\":{\"value\":\"quick\"}}}]}}}", bodyJSON)
}

func Test_DisMax_should_create_json_with_all_fields(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.DisMax(
			es.Match("title", "quick brown"),
			es.Match("body", "quick brown"),
		).
			TieBreaker(0.3).
			Boost(1.5).
			Name("dismax_search"),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	// nolint:golint,lll
	assert.Equal(t, "{\"query\":{\"dis_max\":{\"_name\":\"dismax_search\",\"boost\":1.5,\"queries\":[{\"match\":{\"title\":{\"query\":\"quick brown\"}}},{\"match\":{\"body\":{\"query\":\"quick brown\"}}}],\"tie_breaker\":0.3}}}", bodyJSON)
}

func Test_DisMax_with_Bool_query_should_wrap_correctly(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.DisMax(
			es.Bool().Must(es.Term("status", "active")),
		),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	// nolint:golint,lll
	assert.Equal(t, "{\"query\":{\"dis_max\":{\"queries\":[{\"bool\":{\"must\":[{\"term\":{\"status\":{\"value\":\"active\"}}}]}}]}}}", bodyJSON)
}
