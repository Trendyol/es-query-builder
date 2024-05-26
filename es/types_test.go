package es_test

import (
	"reflect"
	"testing"

	"github.com/GokselKUCUKSAHIN/es-query-builder/es"
	"github.com/GokselKUCUKSAHIN/es-query-builder/test/assert"

	Operator "github.com/GokselKUCUKSAHIN/es-query-builder/es/enums/match/operator"
	ScoreMode "github.com/GokselKUCUKSAHIN/es-query-builder/es/enums/nested/score-mode"
	Mode "github.com/GokselKUCUKSAHIN/es-query-builder/es/enums/sort/mode"
	Order "github.com/GokselKUCUKSAHIN/es-query-builder/es/enums/sort/order"
)

////   NewQuery   ////

func Test_NewQuery_should_exist_on_es_package(t *testing.T) {
	// Given When Then
	assert.NotNil(t, es.NewQuery)
}

func Test_NewQuery_should_create_a_new_Object(t *testing.T) {
	// Given When
	bodyA := es.NewQuery(nil)
	bodyB := es.NewQuery(nil)

	// Then
	assert.NotNil(t, bodyA)
	assert.NotNil(t, bodyB)
	assert.Equal(t, bodyA, bodyB)
	assert.NotEqualReference(t, bodyA, bodyB)
	assert.MarshalWithoutError(t, bodyA)
	assert.MarshalWithoutError(t, bodyB)
}

func Test_NewQuery_should_return_type_of_Object(t *testing.T) {
	// Given
	query := es.NewQuery(nil)

	// When
	bodyType := reflect.TypeOf(query).String()

	// Then
	assert.NotNil(t, query)
	assert.Equal(t, "es.Object", bodyType)
	assert.MarshalWithoutError(t, query)
}

func Test_NewQuery_should_add_query_field_into_Object(t *testing.T) {
	// Given
	query := es.NewQuery(nil)

	// When
	q, exists := query["query"]

	// Then
	assert.True(t, exists)
	assert.NotNil(t, q)
}

func Test_NewQuery_should_create_json_with_query_field(t *testing.T) {
	// Given
	query := es.NewQuery(nil)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{}}", bodyJSON)
}

func Test_NewQuery_Bool_should_create_json_with_bool_field_inside_query(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.Bool(),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"bool\":{}}}", bodyJSON)
}

////   Bool   ////

func Test_Bool_should_exist_on_es_package(t *testing.T) {
	// Given When Then
	assert.NotNil(t, es.Bool)
}

func Test_Bool_method_should_create_boolType(t *testing.T) {
	// Given
	b := es.Bool()

	// Then
	assert.NotNil(t, b)
	assert.IsTypeString(t, "es.boolType", b)
}

func Test_Bool_should_have_MinimumShouldMatch_method(t *testing.T) {
	// Given
	b := es.Bool()

	// When Then
	assert.NotNil(t, b.MinimumShouldMatch)
}

func Test_Bool_MinimumShouldMatch_should_create_json_with_minimum_should_match_field_inside_bool(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.Bool().
			MinimumShouldMatch(7),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"bool\":{\"minimum_should_match\":7}}}", bodyJSON)
}

func Test_Bool_should_have_Boost_method(t *testing.T) {
	// Given
	b := es.Bool()

	// When Then
	assert.NotNil(t, b.Boost)
}

func Test_Bool_Boost_should_create_json_with_minimum_should_match_field_inside_bool(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.Bool().
			Boost(3.1415),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"bool\":{\"boost\":3.1415}}}", bodyJSON)
}

func Test_Bool_should_have_Filter_method(t *testing.T) {
	// Given
	b := es.Bool()

	// When Then
	assert.NotNil(t, b.Filter)
}

func Test_Bool_should_have_Must_method(t *testing.T) {
	// Given
	b := es.Bool()

	// When Then
	assert.NotNil(t, b.Must)
}

func Test_Bool_should_have_MustNot_method(t *testing.T) {
	// Given
	b := es.Bool()

	// When Then
	assert.NotNil(t, b.MustNot)
}

func Test_Bool_should_have_Should_method(t *testing.T) {
	// Given
	b := es.Bool()

	// When Then
	assert.NotNil(t, b.Should)
}

////   Object   ////

func Test_Object_should_have_TrackTotalHits_method(t *testing.T) {
	// Given
	b := es.NewQuery(nil)

	// When Then
	assert.NotNil(t, b.TrackTotalHits)
}

func Test_TrackTotalHits_should_add_track_total_hits_field_into_Object(t *testing.T) {
	// Given
	query := es.NewQuery(nil)

	// When
	_, beforeExists := query["track_total_hits"]
	object := query.TrackTotalHits(true)
	trackTotalHits, afterExists := query["track_total_hits"]

	// Then
	assert.NotNil(t, object)
	assert.False(t, beforeExists)
	assert.True(t, afterExists)
	assert.True(t, trackTotalHits.(bool))
}

func Test_Object_should_have_Size_method(t *testing.T) {
	// Given
	b := es.NewQuery(nil)

	// When Then
	assert.NotNil(t, b.Size)
}

func Test_Size_should_add_size_field_into_Object(t *testing.T) {
	// Given
	query := es.NewQuery(nil)

	// When
	_, beforeExists := query["size"]
	object := query.Size(123)
	size, afterExists := query["size"]

	// Then
	assert.NotNil(t, object)
	assert.False(t, beforeExists)
	assert.True(t, afterExists)
	assert.Equal(t, 123, size.(int))
}

func Test_Object_should_have_From_method(t *testing.T) {
	// Given
	b := es.NewQuery(nil)

	// When Then
	assert.NotNil(t, b.From)
}

func Test_From_should_add_from_field_into_Object(t *testing.T) {
	// Given
	query := es.NewQuery(nil)

	// When
	_, beforeExists := query["from"]
	object := query.From(1500)
	from, afterExists := query["from"]

	// Then
	assert.NotNil(t, object)
	assert.False(t, beforeExists)
	assert.True(t, afterExists)
	assert.Equal(t, 1500, from.(int))
}

func Test_Object_should_have_Sort_method(t *testing.T) {
	// Given
	b := es.NewQuery(nil)

	// When Then
	assert.NotNil(t, b.Sort)
}

func Test_Sort_should_exist_on_es_package(t *testing.T) {
	// Given When Then
	assert.NotNil(t, es.Sort)
}

func Test_Sort_should_return_sortType_with_order(t *testing.T) {
	// Given
	sort := es.Sort("name").Order(Order.Asc)

	// When
	bodyType := reflect.TypeOf(sort).String()

	// Then
	assert.NotNil(t, sort)
	assert.Equal(t, "es.sortType", bodyType)
	bodyJSON := assert.MarshalWithoutError(t, sort)
	assert.Equal(t, "{\"name\":{\"order\":\"asc\"}}", bodyJSON)
}

func Test_Sort_should_return_sortType_with_mode(t *testing.T) {
	// Given
	sort := es.Sort("age").Mode(Mode.Median)

	// When
	bodyType := reflect.TypeOf(sort).String()

	// Then
	assert.NotNil(t, sort)
	assert.Equal(t, "es.sortType", bodyType)
	bodyJSON := assert.MarshalWithoutError(t, sort)
	assert.Equal(t, "{\"age\":{\"mode\":\"median\"}}", bodyJSON)
}

func Test_Sort_should_return_sortType_with_order_and_mode(t *testing.T) {
	// Given
	sort := es.Sort("salary").Order(Order.Desc).Mode(Mode.Sum)

	// When
	bodyType := reflect.TypeOf(sort).String()

	// Then
	assert.NotNil(t, sort)
	assert.Equal(t, "es.sortType", bodyType)
	bodyJSON := assert.MarshalWithoutError(t, sort)
	assert.Equal(t, "{\"salary\":{\"mode\":\"sum\",\"order\":\"desc\"}}", bodyJSON)
}

func Test_Sort_should_add_sort_field_into_Object(t *testing.T) {
	// Given
	query := es.NewQuery(nil)

	// When
	_, beforeExists := query["sort"]
	query.Sort(es.Sort("name").Order(Order.Desc))
	sort, afterExists := query["sort"]

	// Then
	assert.NotNil(t, sort)
	assert.False(t, beforeExists)
	assert.True(t, afterExists)
	assert.IsTypeString(t, "[]es.sortType", sort)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{},\"sort\":[{\"name\":{\"order\":\"desc\"}}]}", bodyJSON)
}

func Test_Object_should_have_Source_method(t *testing.T) {
	// Given
	b := es.NewQuery(nil)

	// When Then
	assert.NotNil(t, b.Source)
}

func Test_Source_should_add_source_field_into_Object(t *testing.T) {
	// Given
	query := es.NewQuery(nil)

	// When
	_, beforeExists := query["_source"]
	query.Source()
	source, afterExists := query["_source"]

	// Then
	assert.NotNil(t, source)
	assert.False(t, beforeExists)
	assert.True(t, afterExists)
	assert.IsTypeString(t, "es.sourceType", source)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"_source\":{},\"query\":{}}", bodyJSON)
}

func Test_Source_should_have_Includes_method(t *testing.T) {
	// Given
	query := es.NewQuery(nil)

	// When
	source := query.Source()

	// Then
	assert.NotNil(t, source)
	assert.IsTypeString(t, "es.sourceType", source)
	assert.NotNil(t, source.Includes)
}

func Test_Source_should_have_Excludes_method(t *testing.T) {
	// Given
	query := es.NewQuery(nil)

	// When
	source := query.Source()

	// Then
	assert.NotNil(t, source)
	assert.IsTypeString(t, "es.sourceType", source)
	assert.NotNil(t, source.Excludes)
}

func Test_Source_should_create_json_with_source_field(t *testing.T) {
	// Given
	query := es.NewQuery(nil)

	// When
	query.Source().
		Includes("hello", "world").
		Excludes("Lorem", "Ipsum")

	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"_source\":{\"excludes\":[\"Lorem\",\"Ipsum\"],\"includes\":[\"hello\",\"world\"]},\"query\":{}}", bodyJSON)
}

func Test_Source_should_append_existing_fields(t *testing.T) {
	// Given
	query := es.NewQuery(nil)

	// When
	query.Source().
		Includes("hello", "world").
		Excludes("Lorem", "Ipsum").
		Includes("golang", "gopher").
		Excludes("Metallica", "Iron Maiden")

	bodyJSON := assert.MarshalWithoutError(t, query)
	// nolint:golint,lll
	assert.Equal(t, "{\"_source\":{\"excludes\":[\"Lorem\",\"Ipsum\",\"Metallica\",\"Iron Maiden\"],\"includes\":[\"hello\",\"world\",\"golang\",\"gopher\"]},\"query\":{}}", bodyJSON)
}

func Test_Object_should_have_SourceFalse_method(t *testing.T) {
	// Given
	b := es.NewQuery(nil)

	// When Then
	assert.NotNil(t, b.SourceFalse)
}

func Test_SourceFalse_should_set_source_field_as_false(t *testing.T) {
	// Given
	query := es.NewQuery(nil)

	// When
	_, beforeExists := query["_source"]
	query.SourceFalse()
	source, afterExists := query["_source"]

	// Then
	assert.NotNil(t, source)
	assert.False(t, beforeExists)
	assert.True(t, afterExists)
	assert.False(t, source.(bool))
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"_source\":false,\"query\":{}}", bodyJSON)
}

func Test_Object_should_have_Range_method(t *testing.T) {
	// Given
	b := es.NewQuery(nil)

	// When Then
	assert.NotNil(t, b.Range)
}

func Test_Object_should_have_Aggs_method(t *testing.T) {
	// Given
	b := es.NewQuery(nil)

	// When Then
	assert.NotNil(t, b.Aggs)
}

////   Term   ////

func Test_Term_should_exist_on_es_package(t *testing.T) {
	// Given When Then
	assert.NotNil(t, es.Term[any])
}

func Test_Term_should_create_json_with_term_field_inside_query(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.Term("key", "value"),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"term\":{\"key\":\"value\"}}}", bodyJSON)
}

func Test_Term_method_should_create_termType(t *testing.T) {
	// Given
	b := es.Term("key", "value")

	// Then
	assert.NotNil(t, b)
	assert.IsTypeString(t, "es.termType", b)
}

////   TermFunc   ////

func Test_TermFunc_should_exist_on_es_package(t *testing.T) {
	// Given When Then
	assert.NotNil(t, es.TermFunc[any])
}

func Test_TermFunc_should_create_json_with_term_field_inside_query(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.TermFunc("key", "value", func(key string, value string) bool {
			return true
		}),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"term\":{\"key\":\"value\"}}}", bodyJSON)
}

func Test_TermFunc_should_not_add_term_field_inside_query_when_callback_result_is_false(t *testing.T) {
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
	assert.Equal(t, "{\"query\":{\"bool\":{\"filter\":[{\"term\":{\"a\":\"b\"}},{\"term\":{\"e\":\"f\"}}]}}}", bodyJSON)
}

func Test_TermFunc_method_should_create_termType(t *testing.T) {
	// Given
	b := es.TermFunc("key", "value", func(key string, value string) bool {
		return true
	})

	// Then
	assert.NotNil(t, b)
	assert.IsTypeString(t, "es.termType", b)
}

////   Terms   ////

func Test_Terms_should_exist_on_es_package(t *testing.T) {
	// Given When Then
	assert.NotNil(t, es.Terms)
}

func Test_Terms_should_create_json_with_terms_field_inside_query(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.Terms("key", "value1", "value2", "value3"),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"terms\":{\"key\":[\"value1\",\"value2\",\"value3\"]}}}", bodyJSON)
}

func Test_Terms_method_should_create_termsType(t *testing.T) {
	// Given
	b := es.Terms("key", "value1", "value2", "value3")

	// Then
	assert.NotNil(t, b)
	assert.IsTypeString(t, "es.termsType", b)
}

////   TermsArray   ////

func Test_TermsArray_should_exist_on_es_package(t *testing.T) {
	// Given When Then
	assert.NotNil(t, es.TermsArray[string])
}

func Test_TermsArray_should_create_json_with_terms_field_inside_query(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.TermsArray("key", []any{"value1", "value2", "value3"}),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"terms\":{\"key\":[\"value1\",\"value2\",\"value3\"]}}}", bodyJSON)
}

func Test_TermsArray_method_should_create_termsType(t *testing.T) {
	// Given
	b := es.TermsArray("key", []any{"value1", "value2", "value3"})

	// Then
	assert.NotNil(t, b)
	assert.IsTypeString(t, "es.termsType", b)
}

////   TermsArrayFunc   ////

func Test_TermsArrayFunc_should_exist_on_es_package(t *testing.T) {
	// Given When Then
	assert.NotNil(t, es.TermsArrayFunc[string])
}

func Test_TermsArrayFunc_should_create_json_with_terms_field_inside_query(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.TermsArrayFunc("key", []string{"a", "b", "c"}, func(key string, values []string) bool {
			return true
		}),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"terms\":{\"key\":[\"a\",\"b\",\"c\"]}}}", bodyJSON)
}

func Test_TermsArrayFunc_should_not_add_terms_field_inside_query_when_callback_result_is_false(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.TermsArrayFunc("key", []string{"a", "b", "c"}, func(key string, value []string) bool {
			return false
		}),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{}}", bodyJSON)
}

func Test_TermsArrayFunc_should_add_only_terms_fields_inside_the_query_when_callback_result_is_true(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.Bool().
			Filter(
				es.TermsArrayFunc("a", []string{"10", "11", "12"}, func(key string, value []string) bool {
					return false
				}),
				es.TermsArrayFunc("c", []string{"20", "21", "22"}, func(key string, value []string) bool {
					return false
				}),
				es.TermsArrayFunc("e", []string{"30", "31", "32"}, func(key string, value []string) bool {
					return true
				}),
			),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"bool\":{\"filter\":[{\"terms\":{\"e\":[\"30\",\"31\",\"32\"]}}]}}}", bodyJSON)
}

func Test_TermsArrayFunc_method_should_create_termType(t *testing.T) {
	// Given
	b := es.TermsArrayFunc("key", []string{"a", "b", "c"}, func(key string, value []string) bool {
		return true
	})

	// Then
	assert.NotNil(t, b)
	assert.IsTypeString(t, "es.termsType", b)
}

////   Exists   ////

func Test_Exists_should_exist_on_es_package(t *testing.T) {
	// Given When Then
	assert.NotNil(t, es.Exists)
}

func Test_Exists_should_create_json_with_exists_field_inside_query(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.Exists("key"),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"exists\":{\"field\":\"key\"}}}", bodyJSON)
}

func Test_Exists_method_should_create_existsType(t *testing.T) {
	// Given
	b := es.Exists("key")

	// Then
	assert.NotNil(t, b)
	assert.IsTypeString(t, "es.existsType", b)
}

////   ExistsFunc   ////

func Test_ExistsFunc_should_exist_on_es_package(t *testing.T) {
	// Given When Then
	assert.NotNil(t, es.ExistsFunc)
}

func Test_ExistsFunc_should_create_json_with_exists_field_inside_query(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.ExistsFunc("key", func(key string) bool {
			return true
		}),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"exists\":{\"field\":\"key\"}}}", bodyJSON)
}

func Test_ExistsFunc_should_not_add_exists_field_inside_query_when_callback_result_is_false(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.ExistsFunc("key", func(key string) bool {
			return false
		}),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{}}", bodyJSON)
}

func Test_ExistsFunc_should_add_only_exists_fields_inside_the_query_when_callback_result_is_true(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.Bool().
			Filter(
				es.ExistsFunc("a", func(key string) bool {
					return false
				}),
				es.ExistsFunc("c", func(key string) bool {
					return true
				}),
				es.ExistsFunc("e", func(key string) bool {
					return true
				}),
			),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"bool\":{\"filter\":[{\"exists\":{\"field\":\"c\"}},{\"exists\":{\"field\":\"e\"}}]}}}", bodyJSON)
}

func Test_ExistsFunc_method_should_create_existsType(t *testing.T) {
	// Given
	b := es.ExistsFunc("key", func(key string) bool {
		return true
	})

	// Then
	assert.NotNil(t, b)
	assert.IsTypeString(t, "es.existsType", b)
}

////   Match   ////

func Test_Match_should_exist_on_es_package(t *testing.T) {
	// Given When Then
	assert.NotNil(t, es.Match[any])
}

func Test_Match_method_should_create_matchType(t *testing.T) {
	// Given
	b := es.Match("key", "value")

	// Then
	assert.NotNil(t, b)
	assert.IsTypeString(t, "es.matchType", b)
}

func Test_Match_should_create_json_with_match_field_inside_query(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.Match("message", "this is a test").
			Boost(2.14).
			Operator(Operator.Or),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t,
		"{\"query\":{\"match\":{\"message\":{\"boost\":2.14,\"operator\":\"or\",\"query\":\"this is a test\"}}}}", bodyJSON)
}

////   MatchNone   ////

func Test_MatchNone_should_exist_on_es_package(t *testing.T) {
	// Given When Then
	assert.NotNil(t, es.MatchNone[any])
}

func Test_MatchNone_method_should_create_matchNoneType(t *testing.T) {
	// Given
	b := es.MatchNone("key", "value")

	// Then
	assert.NotNil(t, b)
	assert.IsTypeString(t, "es.matchNoneType", b)
}

func Test_MatchNone_should_create_json_with_match_none_field_inside_query(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.MatchNone("fooBar", "lorem ipsum").
			Boost(6.19).
			Operator(Operator.And),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t,
		"{\"query\":{\"match_none\":{\"fooBar\":{\"boost\":6.19,\"operator\":\"and\",\"query\":\"lorem ipsum\"}}}}", bodyJSON)
}

////    MatchAll    ////

func Test_MatchAll_should_exist_on_es_package(t *testing.T) {
	// Given When Then
	assert.NotNil(t, es.MatchAll)
}

func Test_MatchAll_method_should_create_matchAllType(t *testing.T) {
	// Given
	b := es.MatchAll()

	// Then
	assert.NotNil(t, b)
	assert.IsTypeString(t, "es.matchAllType", b)
}

func Test_MatchAll_should_create_json_with_match_all_field_inside_query(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.MatchAll().
			Boost(2.14),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t,
		"{\"query\":{\"match_all\":{\"boost\":2.14}}}", bodyJSON)
}

////   Range   ////

func Test_Range_should_exist_on_es_package(t *testing.T) {
	// Given When Then
	assert.NotNil(t, es.Range)
}

func Test_Range_should_add_range_field_when_inside_query(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.Bool().
			Must(
				es.Range("age").
					GreaterThanOrEqual(10).
					LesserThanOrEqual(20),
				es.Term("language", "tr"),
			),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t,
		"{\"query\":{\"bool\":{\"must\":[{\"range\":{\"age\":{\"gte\":10,\"lte\":20}}},{\"term\":{\"language\":\"tr\"}}]}}}", bodyJSON)
}

func Test_Range_method_should_create_rangeType(t *testing.T) {
	// Given
	query := es.NewQuery(nil)
	b := query.Range("age")

	// Then
	assert.NotNil(t, b)
	assert.IsTypeString(t, "es.rangeType", b)
}

func Test_Range_should_create_json_with_range_field_inside_query(t *testing.T) {
	// Given
	query := es.NewQuery(nil)
	query.Range("age").
		GreaterThanOrEqual(10).
		LesserThanOrEqual(20)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"range\":{\"age\":{\"gte\":10,\"lte\":20}}}}", bodyJSON)
}

func Test_Range_should_have_LesserThan_method(t *testing.T) {
	// Given
	r := es.Range("age")

	// When Then
	assert.NotNil(t, r)
	assert.NotNil(t, r.LesserThan)
}

func Test_Range_should_have_LesserThanOrEqual_method(t *testing.T) {
	// Given
	r := es.Range("age")

	// When Then
	assert.NotNil(t, r)
	assert.NotNil(t, r.LesserThanOrEqual)
}

func Test_Range_should_have_GreaterThan_method(t *testing.T) {
	// Given
	r := es.Range("age")

	// When Then
	assert.NotNil(t, r)
	assert.NotNil(t, r.GreaterThan)
}

func Test_Range_should_have_GreaterThanOrEqual_method(t *testing.T) {
	// Given
	r := es.Range("age")

	// When Then
	assert.NotNil(t, r)
	assert.NotNil(t, r.GreaterThanOrEqual)
}

func Test_Range_gte_should_override_gt_and_vise_versa(t *testing.T) {
	// Given
	query := es.NewQuery(nil)
	query.Range("age").
		GreaterThanOrEqual(10).
		GreaterThan(20)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"range\":{\"age\":{\"gt\":20}}}}", bodyJSON)
}

func Test_Range_lte_should_override_lt_and_vise_versa(t *testing.T) {
	// Given
	query := es.NewQuery(nil)
	query.Range("age").
		LesserThan(11).
		LesserThanOrEqual(23)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"range\":{\"age\":{\"lte\":23}}}}", bodyJSON)
}

func Test_Range_should_have_Format_method(t *testing.T) {
	// Given
	r := es.Range("age")

	// When Then
	assert.NotNil(t, r)
	assert.NotNil(t, r.Format)
}

func Test_Range_Format_should_create_json_with_range_field_inside_query(t *testing.T) {
	// Given
	query := es.NewQuery(nil)
	query.Range("birth-date").
		GreaterThanOrEqual("1990-01-01").
		LesserThanOrEqual("2024-04-12").
		Format("yyyy-MM-dd")

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t,
		"{\"query\":{\"range\":{\"birth-date\":{\"format\":\"yyyy-MM-dd\",\"gte\":\"1990-01-01\",\"lte\":\"2024-04-12\"}}}}", bodyJSON)
}

func Test_Range_should_have_Boost_method(t *testing.T) {
	// Given
	r := es.Range("age")

	// When Then
	assert.NotNil(t, r)
	assert.NotNil(t, r.Boost)
}

func Test_Range_Boost_should_create_json_with_range_field_inside_query(t *testing.T) {
	// Given
	query := es.NewQuery(nil)
	query.Range("partition").
		GreaterThan(112).
		LesserThan(765).
		Boost(3.2)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t,
		"{\"query\":{\"range\":{\"partition\":{\"boost\":3.2,\"gt\":112,\"lt\":765}}}}", bodyJSON)
}

func Test_Range_should_not_range_field_when_no_query_field_in_Object(t *testing.T) {
	// Given
	object := es.Object{}
	object.Range("age").
		GreaterThanOrEqual(10).
		LesserThanOrEqual(20)

	// When Then
	assert.NotNil(t, object)
	bodyJSON := assert.MarshalWithoutError(t, object)
	assert.Equal(t, "{}", bodyJSON)
}

////   Bool.Filter   ////

func Test_Filter_method_should_return_boolType(t *testing.T) {
	// Given
	b := es.Bool()

	// When
	filter := b.Filter()

	// Then
	assert.NotNil(t, filter)
	assert.IsTypeString(t, "es.boolType", filter)
}

func Test_Filter_method_should_add_filter_if_doesnt_exist_before(t *testing.T) {
	// Given
	b := es.Bool()

	// When
	_, beforeExists := b["filter"]
	filter := b.Filter()
	_, afterExists := b["filter"]

	// Then
	assert.NotNil(t, filter)
	assert.False(t, beforeExists)
	assert.True(t, afterExists)
}

func Test_Filter_method_should_hold_items(t *testing.T) {
	// Given
	b := es.Bool().
		Filter(
			es.Term("id", 12345),
		)

	// When
	filter, exists := b["filter"]

	// Then
	assert.True(t, exists)
	assert.IsTypeString(t, "es.filterType", filter)

	bodyJSON := assert.MarshalWithoutError(t, b)
	assert.Equal(t, "{\"filter\":[{\"term\":{\"id\":12345}}]}", bodyJSON)
}

////   Bool.Must   ////

func Test_Must_method_should_return_boolType(t *testing.T) {
	// Given
	b := es.Bool()

	// When
	must := b.Must()

	// Then
	assert.NotNil(t, must)
	assert.IsTypeString(t, "es.boolType", must)
}

func Test_Must_method_should_add_must_if_doesnt_exist_before(t *testing.T) {
	// Given
	b := es.Bool()

	// When
	_, beforeExists := b["must"]
	filter := b.Must()
	_, afterExists := b["must"]

	// Then
	assert.NotNil(t, filter)
	assert.False(t, beforeExists)
	assert.True(t, afterExists)
}

func Test_Must_method_should_hold_items(t *testing.T) {
	// Given
	b := es.Bool().
		Must(
			es.Term("id", 12345),
		)

	// When
	must, exists := b["must"]

	// Then
	assert.True(t, exists)
	assert.IsTypeString(t, "es.mustType", must)

	bodyJSON := assert.MarshalWithoutError(t, b)
	assert.Equal(t, "{\"must\":[{\"term\":{\"id\":12345}}]}", bodyJSON)
}

////   Bool.MustNot   ////

func Test_MustNot_method_should_return_boolType(t *testing.T) {
	// Given
	b := es.Bool()

	// When
	mustNot := b.MustNot()

	// Then
	assert.NotNil(t, mustNot)
	assert.IsTypeString(t, "es.boolType", mustNot)
}

func Test_MustNot_method_should_add_must_not_if_doesnt_exist_before(t *testing.T) {
	// Given
	b := es.Bool()

	// When
	_, beforeExists := b["must_not"]
	filter := b.MustNot()
	_, afterExists := b["must_not"]

	// Then
	assert.NotNil(t, filter)
	assert.False(t, beforeExists)
	assert.True(t, afterExists)
}

func Test_MustNot_method_should_hold_items(t *testing.T) {
	// Given
	b := es.Bool().
		MustNot(
			es.Term("id", 12345),
		)

	// When
	mustNot, exists := b["must_not"]

	// Then
	assert.True(t, exists)
	assert.IsTypeString(t, "es.mustNotType", mustNot)

	bodyJSON := assert.MarshalWithoutError(t, b)
	assert.Equal(t, "{\"must_not\":[{\"term\":{\"id\":12345}}]}", bodyJSON)
}

////   Bool.Should   ////

func Test_Should_method_should_return_boolType(t *testing.T) {
	// Given
	b := es.Bool()

	// When
	should := b.Should()

	// Then
	assert.NotNil(t, should)
	assert.IsTypeString(t, "es.boolType", should)
}

func Test_Should_method_should_add_should_if_doesnt_exist_before(t *testing.T) {
	// Given
	b := es.Bool()

	// When
	_, beforeExists := b["should"]
	filter := b.Should()
	_, afterExists := b["should"]

	// Then
	assert.NotNil(t, filter)
	assert.False(t, beforeExists)
	assert.True(t, afterExists)
}

func Test_Should_method_should_hold_items(t *testing.T) {
	// Given
	b := es.Bool().
		Should(
			es.Term("id", 12345),
		)

	// When
	should, exists := b["should"]

	// Then
	assert.True(t, exists)
	assert.IsTypeString(t, "es.shouldType", should)

	bodyJSON := assert.MarshalWithoutError(t, b)
	assert.Equal(t, "{\"should\":[{\"term\":{\"id\":12345}}]}", bodyJSON)
}

////    Nested    ////

func Test_Nested_should_exist_on_es_package(t *testing.T) {
	// Given When Then
	assert.NotNil(t, es.Nested[any])
}

func Test_Nested_method_should_create_nestedType(t *testing.T) {
	// Given
	n := es.Nested("path", es.Object{})

	// Then
	assert.NotNil(t, n)
	assert.IsTypeString(t, "es.nestedType", n)
}

func Test_Nested_should_create_query_json_with_nested_field_inside(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.Nested("nested.path",
			es.Object{},
		),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"nested\":{\"path\":\"nested.path\",\"query\":{}}}}", bodyJSON)
}

func Test_Nested_should_have_InnerHits_method(t *testing.T) {
	// Given
	n := es.Nested("path", es.Object{})

	// When Then
	assert.NotNil(t, n.InnerHits)
}

func Test_InnerHits_should_add_inner_hits_field_into_Nested(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.Nested("nested.path", es.Object{}).InnerHits(es.Object{"inner": "hits"}),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"nested\":{\"inner_hits\":{\"inner\":\"hits\"},\"path\":\"nested.path\",\"query\":{}}}}", bodyJSON)
}

func Test_Nested_should_have_ScoreMode_method(t *testing.T) {
	// Given
	n := es.Nested("path", es.Object{})

	// When Then
	assert.NotNil(t, n.ScoreMode)
}

func Test_ScoreMod_should_add_score_mode_field_into_Nested(t *testing.T) {
	// Given
	query := es.NewQuery(
		es.Nested("nested.path", es.Object{}).ScoreMode(ScoreMode.Sum),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"nested\":{\"path\":\"nested.path\",\"query\":{},\"score_mode\":\"sum\"}}}", bodyJSON)
}

////    AGGS    ////

func Test_AggTerm_should_exist_on_es_package(t *testing.T) {
	// Given When Then
	assert.NotNil(t, es.AggTerm)
}

func Test_AggTerm_method_should_create_aggTermType(t *testing.T) {
	// Given
	a := es.AggTerm("path")

	// Then
	assert.NotNil(t, a)
	assert.IsTypeString(t, "es.aggTermType", a)
}

func Test_AggTerm_should_create_json_with_field_field_inside(t *testing.T) {
	// Given
	a := es.AggTerm("path")

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"field\":\"path\"}", bodyJSON)
}

func Test_AggTerm_should_have_Missing_method(t *testing.T) {
	// Given
	a := es.AggTerm("path")

	// When Then
	assert.NotNil(t, a.Missing)
}

func Test_Missing_should_add_missing_field_into_AggTerm(t *testing.T) {
	// Given
	a := es.AggTerm("path").Missing("missing_name")

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"field\":\"path\",\"missing\":\"missing_name\"}", bodyJSON)
}

func Test_AggTerms_should_exist_on_es_package(t *testing.T) {
	// Given When Then
	assert.NotNil(t, es.AggTerms)
}

func Test_AggTerms_method_should_create_aggType(t *testing.T) {
	// Given
	a := es.AggTerms()

	// Then
	assert.NotNil(t, a)
	assert.IsTypeString(t, "es.aggsType", a)
}

func Test_AggTerms_should_create_json_with_terms_field_inside(t *testing.T) {
	// Given
	a := es.AggTerms()

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"terms\":{}}", bodyJSON)
}

func Test_AggMultiTerms_should_exist_on_es_package(t *testing.T) {
	// Given When Then
	assert.NotNil(t, es.AggMultiTerms)
}

func Test_AggMultiTerms_method_should_create_aggType(t *testing.T) {
	// Given
	a := es.AggMultiTerms()

	// Then
	assert.NotNil(t, a)
	assert.IsTypeString(t, "es.aggsType", a)
}

func Test_AggMultiTerms_should_create_json_with_multi_terms_field_inside(t *testing.T) {
	// Given
	a := es.AggMultiTerms()

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"multi_terms\":{}}", bodyJSON)
}

func Test_AggCustom_should_exist_on_es_package(t *testing.T) {
	// Given When Then
	assert.NotNil(t, es.AggCustom)
}

func Test_AggCustom_method_should_create_aggType(t *testing.T) {
	// Given
	a := es.AggCustom(nil)

	// Then
	assert.NotNil(t, a)
	assert.IsTypeString(t, "es.aggsType", a)
}

func Test_AggCustom_should_create_json(t *testing.T) {
	// Given
	a := es.AggCustom(es.Object{
		"custom": es.Object{
			"my_field": es.Array{1, 2, 3},
		},
	})

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"custom\":{\"my_field\":[1,2,3]}}", bodyJSON)
}

func Test_AggType_should_have_Field_method(t *testing.T) {
	// Given
	a := es.AggTerms()

	// When Then
	assert.NotNil(t, a.Field)
}

func Test_Field_should_add_field_field_into_AggTerm(t *testing.T) {
	// Given
	a := es.AggTerms().Field("path")

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"terms\":{\"field\":\"path\"}}", bodyJSON)
}

func Test_AggType_should_have_Size_method(t *testing.T) {
	// Given
	a := es.AggTerms()

	// When Then
	assert.NotNil(t, a.Field)
}

func Test_Size_should_add_size_field_into_AggType(t *testing.T) {
	// Given
	a := es.AggTerms().Size(333)

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"terms\":{\"size\":333}}", bodyJSON)
}

func Test_AggType_should_have_Order_method(t *testing.T) {
	// Given
	a := es.AggTerms()

	// When Then
	assert.NotNil(t, a.Field)
}

func Test_Order_should_add_order_field_into_AggType(t *testing.T) {
	// Given
	a := es.AggTerms().Order("path", Order.Desc)

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"terms\":{\"order\":{\"path\":\"desc\"}}}", bodyJSON)
}

func Test_AggType_should_have_Include_method(t *testing.T) {
	// Given
	a := es.AggTerms()

	// When Then
	assert.NotNil(t, a.Include)
}

func Test_Include_should_add_include_field_into_AggType(t *testing.T) {
	// Given
	a := es.AggTerms().Include("*.2024")

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"terms\":{\"include\":\"*.2024\"}}", bodyJSON)
}

func Test_AggType_should_have_Exclude_method(t *testing.T) {
	// Given
	a := es.AggTerms()

	// When Then
	assert.NotNil(t, a.Exclude)
}

func Test_Exclude_should_add_exclude_field_into_AggType(t *testing.T) {
	// Given
	a := es.AggTerms().Exclude("*.2021")

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	assert.Equal(t, "{\"terms\":{\"exclude\":\"*.2021\"}}", bodyJSON)
}

func Test_AggType_should_have_Terms_method(t *testing.T) {
	// Given
	a := es.AggTerms()

	// When Then
	assert.NotNil(t, a.Terms)
}

func Test_Terms_should_add_terms_field_into_AggType(t *testing.T) {
	// Given
	a := es.AggMultiTerms().
		Terms(
			es.AggTerm("A1"),
			es.AggTerm("B2").
				Missing("Hell Divers"),
			es.AggTerm("C3"),
			es.AggTerm("D4"),
		)

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	// nolint:golint,lll
	assert.Equal(t, "{\"multi_terms\":{\"terms\":[{\"field\":\"A1\"},{\"field\":\"B2\",\"missing\":\"Hell Divers\"},{\"field\":\"C3\"},{\"field\":\"D4\"}]}}", bodyJSON)
}

func Test_AggType_should_have_Aggs_method(t *testing.T) {
	// Given
	a := es.AggTerms()

	// When Then
	assert.NotNil(t, a.Aggs)
}

func Test_Aggs_should_add_aggs_field_into_AggType(t *testing.T) {
	// Given
	a := es.AggTerms().
		Field("path").
		Size(1_000).
		Order("_key", Order.Asc).
		Include("reduces").
		Aggs("test",
			es.AggMultiTerms().
				Terms(
					es.AggTerm("A1").
						Missing("a1"),
					es.AggTerm("B2"),
				),
		)

	// When Then
	assert.NotNil(t, a)
	bodyJSON := assert.MarshalWithoutError(t, a)
	// nolint:golint,lll
	assert.Equal(t, "{\"aggs\":{\"test\":{\"multi_terms\":{\"terms\":[{\"field\":\"A1\",\"missing\":\"a1\"},{\"field\":\"B2\"}]}}},\"terms\":{\"field\":\"path\",\"include\":\"reduces\",\"order\":{\"_key\":\"asc\"},\"size\":1000}}", bodyJSON)
}

func Test_Aggs_should_create_json_with_aggs_field_inside_query(t *testing.T) {
	// Given
	query := es.NewQuery(nil)
	query.Aggs("types",
		es.AggTerms().
			Field("type").
			Size(100),
	)

	// When Then
	assert.NotNil(t, query)
	assert.NotNil(t, query.Aggs)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"aggs\":{\"types\":{\"terms\":{\"field\":\"type\",\"size\":100}}},\"query\":{}}", bodyJSON)
}
