package es_test

import (
	"testing"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/Trendyol/es-query-builder/test/assert"
)

////   Boosting   ////

func Test_Boosting_should_exist_on_es_package(t *testing.T) {
	t.Parallel()
	// Given When Then
	assert.NotNil(t, es.Boosting)
}

func Test_Boosting_method_should_create_boostingType(t *testing.T) {
	t.Parallel()
	// Given
	b := es.Boosting(es.MatchAll(), es.Term("status", "archived"), 0.2)

	// Then
	assert.NotNil(t, b)
	assert.IsTypeString(t, "es.boostingType", b)
}

func Test_Boosting_should_create_json_with_required_fields(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.Boosting(
			es.Match("title", "apple"),
			es.Term("status", "archived"),
			0.2,
		),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	// nolint:golint,lll
	assert.Equal(t, "{\"query\":{\"boosting\":{\"negative\":{\"term\":{\"status\":{\"value\":\"archived\"}}},\"negative_boost\":0.2,\"positive\":{\"match\":{\"title\":{\"query\":\"apple\"}}}}}}", bodyJSON)
}

func Test_Boosting_should_wrap_positive_Bool_query_correctly(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.Boosting(
			es.Bool().Must(es.Term("status", "active")),
			es.Term("category", "outdated"),
			0.2,
		),
	)

	// When Then
	bodyJSON := assert.MarshalWithoutError(t, query)
	// nolint:golint,lll
	assert.Equal(t, "{\"query\":{\"boosting\":{\"negative\":{\"term\":{\"category\":{\"value\":\"outdated\"}}},\"negative_boost\":0.2,\"positive\":{\"bool\":{\"must\":[{\"term\":{\"status\":{\"value\":\"active\"}}}]}}}}}", bodyJSON)
}

func Test_Boosting_should_wrap_negative_Bool_query_correctly(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.Boosting(
			es.Term("status", "active"),
			es.Bool().Must(es.Term("category", "outdated")),
			0.2,
		),
	)

	// When Then
	bodyJSON := assert.MarshalWithoutError(t, query)
	// nolint:golint,lll
	assert.Equal(t, "{\"query\":{\"boosting\":{\"negative\":{\"bool\":{\"must\":[{\"term\":{\"category\":{\"value\":\"outdated\"}}}]}},\"negative_boost\":0.2,\"positive\":{\"term\":{\"status\":{\"value\":\"active\"}}}}}}", bodyJSON)
}

func Test_Boosting_should_return_nil_when_positive_query_is_nil(t *testing.T) {
	t.Parallel()
	// Given When
	b := es.Boosting(nil, es.MatchAll(), 0.2)

	// Then
	assert.True(t, b == nil)
}

func Test_Boosting_should_return_nil_when_negative_query_is_nil(t *testing.T) {
	t.Parallel()
	// Given When
	b := es.Boosting(es.MatchAll(), nil, 0.2)

	// Then
	assert.True(t, b == nil)
}

func Test_Boosting_should_return_nil_when_query_is_typed_nil(t *testing.T) {
	t.Parallel()
	// Given
	var typedNil *es.Object

	// When
	positive := es.Boosting(typedNil, es.MatchAll(), 0.2)
	negative := es.Boosting(es.MatchAll(), typedNil, 0.2)

	// Then
	assert.True(t, positive == nil)
	assert.True(t, negative == nil)
}

func Test_Boosting_should_serialize_negative_boost_boundaries(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name          string
		negativeBoost float64
		expected      string
	}{
		{name: "zero", negativeBoost: 0, expected: "{\"query\":{\"boosting\":{\"negative\":{\"match_none\":{}},\"negative_boost\":0,\"positive\":{\"match_all\":{}}}}}"},
		{name: "fraction", negativeBoost: 0.5, expected: "{\"query\":{\"boosting\":{\"negative\":{\"match_none\":{}},\"negative_boost\":0.5,\"positive\":{\"match_all\":{}}}}}"},
		{name: "one", negativeBoost: 1, expected: "{\"query\":{\"boosting\":{\"negative\":{\"match_none\":{}},\"negative_boost\":1,\"positive\":{\"match_all\":{}}}}}"},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			// Given
			query := es.NewQuery(es.Boosting(es.MatchAll(), es.MatchNone(), test.negativeBoost))

			// When Then
			bodyJSON := assert.MarshalWithoutError(t, query)
			assert.Equal(t, test.expected, bodyJSON)
		})
	}
}

func Test_Boosting_should_have_Boost_method(t *testing.T) {
	t.Parallel()
	// Given
	b := es.Boosting(es.MatchAll(), es.MatchNone(), 0.2)

	// When Then
	assert.NotNil(t, b.Boost)
}

func Test_Boosting_Boost_should_create_json_with_boost_field(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.Boosting(es.MatchAll(), es.MatchNone(), 0.2).
			Boost(1.5),
	)

	// When Then
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"boosting\":{\"boost\":1.5,\"negative\":{\"match_none\":{}},\"negative_boost\":0.2,\"positive\":{\"match_all\":{}}}}}", bodyJSON)
}

func Test_Boosting_should_have_Name_method(t *testing.T) {
	t.Parallel()
	// Given
	b := es.Boosting(es.MatchAll(), es.MatchNone(), 0.2)

	// When Then
	assert.NotNil(t, b.Name)
}

func Test_Boosting_Name_should_create_json_with_name_field(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.Boosting(es.MatchAll(), es.MatchNone(), 0.2).
			Name("demote_archived"),
	)

	// When Then
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"boosting\":{\"_name\":\"demote_archived\",\"negative\":{\"match_none\":{}},\"negative_boost\":0.2,\"positive\":{\"match_all\":{}}}}}", bodyJSON)
}

func Test_Boosting_should_create_json_with_all_fields(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.Boosting(
			es.Match("title", "apple"),
			es.Term("status", "archived"),
			0.2,
		).
			Boost(1.5).
			Name("demote_archived"),
	)

	// When Then
	bodyJSON := assert.MarshalWithoutError(t, query)
	// nolint:golint,lll
	assert.Equal(t, "{\"query\":{\"boosting\":{\"_name\":\"demote_archived\",\"boost\":1.5,\"negative\":{\"term\":{\"status\":{\"value\":\"archived\"}}},\"negative_boost\":0.2,\"positive\":{\"match\":{\"title\":{\"query\":\"apple\"}}}}}}", bodyJSON)
}
