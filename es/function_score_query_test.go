package es_test

import (
	"testing"

	BoostMode "github.com/Trendyol/es-query-builder/es/enums/boost-mode"
	ScoreMode "github.com/Trendyol/es-query-builder/es/enums/score-mode"
	ScriptLanguage "github.com/Trendyol/es-query-builder/es/enums/script-language"

	"github.com/Trendyol/es-query-builder/es"
	"github.com/Trendyol/es-query-builder/test/assert"
)

////   FunctionScore   ////

func Test_FunctionScore_should_exist_on_es_package(t *testing.T) {
	t.Parallel()
	// Given When Then
	assert.NotNil(t, es.FunctionScore)
}

func Test_FunctionScore_method_should_create_functionScoreType(t *testing.T) {
	t.Parallel()
	// Given
	fs := es.FunctionScore(es.MatchAll())

	// Then
	assert.NotNil(t, fs)
	assert.IsTypeString(t, "es.functionScoreType", fs)
}

func Test_FunctionScore_should_create_json_with_function_score_field_inside_query(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.FunctionScore(es.MatchAll()),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"function_score\":{\"query\":{\"match_all\":{}}}}}", bodyJSON)
}

func Test_FunctionScore_should_create_json_with_nil_query(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.FunctionScore(nil),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"function_score\":{}}}", bodyJSON)
}

func Test_FunctionScore_should_create_json_with_bool_query(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.FunctionScore(
			es.Bool().Must(es.Term("status", "active")),
		),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	// nolint:golint,lll
	assert.Equal(t, "{\"query\":{\"function_score\":{\"query\":{\"bool\":{\"must\":[{\"term\":{\"status\":{\"value\":\"active\"}}}]}}}}}", bodyJSON)
}

////   Boost   ////

func Test_FunctionScore_should_have_Boost_method(t *testing.T) {
	t.Parallel()
	// Given
	fs := es.FunctionScore(es.MatchAll())

	// When Then
	assert.NotNil(t, fs.Boost)
}

func Test_FunctionScore_Boost_should_create_json_with_boost_field_inside_function_score(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.FunctionScore(es.MatchAll()).Boost(5),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"function_score\":{\"boost\":5,\"query\":{\"match_all\":{}}}}}", bodyJSON)
}

////   MaxBoost   ////

func Test_FunctionScore_should_have_MaxBoost_method(t *testing.T) {
	t.Parallel()
	// Given
	fs := es.FunctionScore(es.MatchAll())

	// When Then
	assert.NotNil(t, fs.MaxBoost)
}

func Test_FunctionScore_MaxBoost_should_create_json_with_max_boost_field_inside_function_score(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.FunctionScore(es.MatchAll()).MaxBoost(42),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"function_score\":{\"max_boost\":42,\"query\":{\"match_all\":{}}}}}", bodyJSON)
}

////   ScoreMode   ////

func Test_FunctionScore_should_have_ScoreMode_method(t *testing.T) {
	t.Parallel()
	// Given
	fs := es.FunctionScore(es.MatchAll())

	// When Then
	assert.NotNil(t, fs.ScoreMode)
}

func Test_FunctionScore_ScoreMode_should_create_json_with_score_mode_field_inside_function_score(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.FunctionScore(es.MatchAll()).ScoreMode(ScoreMode.Sum),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"function_score\":{\"query\":{\"match_all\":{}},\"score_mode\":\"sum\"}}}", bodyJSON)
}

////   BoostMode   ////

func Test_FunctionScore_should_have_BoostMode_method(t *testing.T) {
	t.Parallel()
	// Given
	fs := es.FunctionScore(es.MatchAll())

	// When Then
	assert.NotNil(t, fs.BoostMode)
}

func Test_FunctionScore_BoostMode_should_create_json_with_boost_mode_field_inside_function_score(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.FunctionScore(es.MatchAll()).BoostMode(BoostMode.Replace),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"function_score\":{\"boost_mode\":\"replace\",\"query\":{\"match_all\":{}}}}}", bodyJSON)
}

////   MinScore   ////

func Test_FunctionScore_should_have_MinScore_method(t *testing.T) {
	t.Parallel()
	// Given
	fs := es.FunctionScore(es.MatchAll())

	// When Then
	assert.NotNil(t, fs.MinScore)
}

func Test_FunctionScore_MinScore_should_create_json_with_min_score_field_inside_function_score(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.FunctionScore(es.MatchAll()).MinScore(5),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"function_score\":{\"min_score\":5,\"query\":{\"match_all\":{}}}}}", bodyJSON)
}

////   Weight   ////

func Test_FunctionScore_should_have_Weight_method(t *testing.T) {
	t.Parallel()
	// Given
	fs := es.FunctionScore(es.MatchAll())

	// When Then
	assert.NotNil(t, fs.Weight)
}

func Test_FunctionScore_Weight_should_create_json_with_weight_field_inside_function_score(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.FunctionScore(es.MatchAll()).Weight(2),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"function_score\":{\"query\":{\"match_all\":{}},\"weight\":2}}}", bodyJSON)
}

////   ScriptScore   ////

func Test_FunctionScore_should_have_ScriptScore_method(t *testing.T) {
	t.Parallel()
	// Given
	fs := es.FunctionScore(es.MatchAll())

	// When Then
	assert.NotNil(t, fs.ScriptScore)
}

func Test_FunctionScore_ScriptScore_should_create_json_with_script_score_field_inside_function_score(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.FunctionScore(es.MatchAll()).
			ScriptScore(es.ScriptSource("_score * doc['likes'].value", ScriptLanguage.Painless)),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	// nolint:golint,lll
	assert.Equal(t, "{\"query\":{\"function_score\":{\"query\":{\"match_all\":{}},\"script_score\":{\"script\":{\"lang\":\"painless\",\"source\":\"_score * doc['likes'].value\"}}}}}", bodyJSON)
}

////   RandomScore   ////

func Test_FunctionScore_should_have_RandomScore_method(t *testing.T) {
	t.Parallel()
	// Given
	fs := es.FunctionScore(es.MatchAll())

	// When Then
	assert.NotNil(t, fs.RandomScore)
}

func Test_FunctionScore_RandomScore_should_create_json_with_random_score_field_inside_function_score(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.FunctionScore(es.MatchAll()).RandomScore(42, "_seq_no"),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	// nolint:golint,lll
	assert.Equal(t, "{\"query\":{\"function_score\":{\"query\":{\"match_all\":{}},\"random_score\":{\"field\":\"_seq_no\",\"seed\":42}}}}", bodyJSON)
}

////   FieldValueFactor   ////

func Test_FunctionScore_should_have_FieldValueFactor_method(t *testing.T) {
	t.Parallel()
	// Given
	fs := es.FunctionScore(es.MatchAll())

	// When Then
	assert.NotNil(t, fs.FieldValueFactor)
}

func Test_FunctionScore_FieldValueFactor_should_create_json_with_field_value_factor_field_inside_function_score(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.FunctionScore(es.MatchAll()).
			FieldValueFactor(es.FieldValueFactor("likes")),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	// nolint:golint,lll
	assert.Equal(t, "{\"query\":{\"function_score\":{\"field_value_factor\":{\"field\":\"likes\"},\"query\":{\"match_all\":{}}}}}", bodyJSON)
}

////   Functions   ////

func Test_FunctionScore_should_have_Functions_method(t *testing.T) {
	t.Parallel()
	// Given
	fs := es.FunctionScore(es.MatchAll())

	// When Then
	assert.NotNil(t, fs.Functions)
}

func Test_FunctionScore_Functions_should_create_json_with_functions_array_inside_function_score(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.FunctionScore(es.MatchAll()).
			Functions(
				es.WeightFunction(2),
			),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"function_score\":{\"functions\":[{\"weight\":2}],\"query\":{\"match_all\":{}}}}}", bodyJSON)
}

func Test_FunctionScore_Functions_should_skip_nil_functions(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.FunctionScore(es.MatchAll()).
			Functions(nil),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	assert.Equal(t, "{\"query\":{\"function_score\":{\"query\":{\"match_all\":{}}}}}", bodyJSON)
}

func Test_FunctionScore_Functions_should_create_json_with_multiple_functions(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.FunctionScore(es.MatchAll()).
			Functions(
				es.WeightFunction(2).Filter(es.Term("status", "published")),
				es.RandomScoreFunction().Seed(42).Field("_seq_no"),
			),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	// nolint:golint,lll
	assert.Equal(t, "{\"query\":{\"function_score\":{\"functions\":[{\"filter\":{\"term\":{\"status\":{\"value\":\"published\"}}},\"weight\":2},{\"random_score\":{\"field\":\"_seq_no\",\"seed\":42}}],\"query\":{\"match_all\":{}}}}}", bodyJSON)
}

////   ScriptScoreFunction   ////

func Test_ScriptScoreFunction_should_exist_on_es_package(t *testing.T) {
	t.Parallel()
	// Given When Then
	assert.NotNil(t, es.ScriptScoreFunction)
}

func Test_ScriptScoreFunction_should_create_functionScoreFunction(t *testing.T) {
	t.Parallel()
	// Given
	fn := es.ScriptScoreFunction(es.ScriptSource("_score * 2", ScriptLanguage.Painless))

	// Then
	assert.NotNil(t, fn)
	assert.IsTypeString(t, "es.functionScoreFunction", fn)
}

func Test_ScriptScoreFunction_should_create_json_with_script_score(t *testing.T) {
	t.Parallel()
	// Given
	fn := es.ScriptScoreFunction(es.ScriptSource("_score * 2", ScriptLanguage.Painless))

	// When Then
	assert.NotNil(t, fn)
	bodyJSON := assert.MarshalWithoutError(t, fn)
	assert.Equal(t, "{\"script_score\":{\"script\":{\"lang\":\"painless\",\"source\":\"_score * 2\"}}}", bodyJSON)
}

////   RandomScoreFunction   ////

func Test_RandomScoreFunction_should_exist_on_es_package(t *testing.T) {
	t.Parallel()
	// Given When Then
	assert.NotNil(t, es.RandomScoreFunction)
}

func Test_RandomScoreFunction_should_create_json_with_seed_and_field(t *testing.T) {
	t.Parallel()
	// Given
	fn := es.RandomScoreFunction().Seed(42).Field("_seq_no")

	// When Then
	assert.NotNil(t, fn)
	bodyJSON := assert.MarshalWithoutError(t, fn)
	assert.Equal(t, "{\"random_score\":{\"field\":\"_seq_no\",\"seed\":42}}", bodyJSON)
}

////   WeightFunction   ////

func Test_WeightFunction_should_exist_on_es_package(t *testing.T) {
	t.Parallel()
	// Given When Then
	assert.NotNil(t, es.WeightFunction)
}

func Test_WeightFunction_should_create_json_with_weight(t *testing.T) {
	t.Parallel()
	// Given
	fn := es.WeightFunction(23)

	// When Then
	assert.NotNil(t, fn)
	bodyJSON := assert.MarshalWithoutError(t, fn)
	assert.Equal(t, "{\"weight\":23}", bodyJSON)
}

////   FieldValueFactorFunction   ////

func Test_FieldValueFactorFunction_should_exist_on_es_package(t *testing.T) {
	t.Parallel()
	// Given When Then
	assert.NotNil(t, es.FieldValueFactorFunction)
}

func Test_FieldValueFactorFunction_should_create_json_with_field_value_factor(t *testing.T) {
	t.Parallel()
	// Given
	fn := es.FieldValueFactorFunction(es.FieldValueFactor("likes"))

	// When Then
	assert.NotNil(t, fn)
	bodyJSON := assert.MarshalWithoutError(t, fn)
	assert.Equal(t, "{\"field_value_factor\":{\"field\":\"likes\"}}", bodyJSON)
}

////   DecayFunction   ////

func Test_DecayFunction_should_exist_on_es_package(t *testing.T) {
	t.Parallel()
	// Given When Then
	assert.NotNil(t, es.DecayFunction)
}

func Test_DecayFunction_should_create_json_with_gauss_decay(t *testing.T) {
	t.Parallel()
	// Given
	fn := es.DecayFunction("gauss", es.Decay("date").Origin("now").Scale("10d"))

	// When Then
	assert.NotNil(t, fn)
	bodyJSON := assert.MarshalWithoutError(t, fn)
	assert.Equal(t, "{\"gauss\":{\"date\":{\"origin\":\"now\",\"scale\":\"10d\"}}}", bodyJSON)
}

func Test_DecayFunction_should_create_json_with_linear_decay(t *testing.T) {
	t.Parallel()
	// Given
	fn := es.DecayFunction("linear", es.Decay("price").Origin(0).Scale(20))

	// When Then
	assert.NotNil(t, fn)
	bodyJSON := assert.MarshalWithoutError(t, fn)
	assert.Equal(t, "{\"linear\":{\"price\":{\"origin\":0,\"scale\":20}}}", bodyJSON)
}

func Test_DecayFunction_should_create_json_with_exp_decay(t *testing.T) {
	t.Parallel()
	// Given
	fn := es.DecayFunction("exp", es.Decay("date").Origin("now").Scale("10d").Offset("5d").DecayValue(0.5))

	// When Then
	assert.NotNil(t, fn)
	bodyJSON := assert.MarshalWithoutError(t, fn)
	assert.Equal(t, "{\"exp\":{\"date\":{\"decay\":0.5,\"offset\":\"5d\",\"origin\":\"now\",\"scale\":\"10d\"}}}", bodyJSON)
}

////   FunctionScoreFunction Filter   ////

func Test_FunctionScoreFunction_Filter_should_create_json_with_filter(t *testing.T) {
	t.Parallel()
	// Given
	fn := es.WeightFunction(2).Filter(es.Term("status", "published"))

	// When Then
	assert.NotNil(t, fn)
	bodyJSON := assert.MarshalWithoutError(t, fn)
	assert.Equal(t, "{\"filter\":{\"term\":{\"status\":{\"value\":\"published\"}}},\"weight\":2}", bodyJSON)
}

func Test_FunctionScoreFunction_Filter_should_create_json_with_bool_filter(t *testing.T) {
	t.Parallel()
	// Given
	fn := es.WeightFunction(2).Filter(es.Bool().Must(es.Term("status", "active")))

	// When Then
	assert.NotNil(t, fn)
	bodyJSON := assert.MarshalWithoutError(t, fn)
	// nolint:golint,lll
	assert.Equal(t, "{\"filter\":{\"bool\":{\"must\":[{\"term\":{\"status\":{\"value\":\"active\"}}}]}},\"weight\":2}", bodyJSON)
}

////   FunctionScoreFunction Weight   ////

func Test_FunctionScoreFunction_Weight_should_add_weight_to_function(t *testing.T) {
	t.Parallel()
	// Given
	fn := es.ScriptScoreFunction(es.ScriptSource("_score * 2", ScriptLanguage.Painless)).Weight(5)

	// When Then
	assert.NotNil(t, fn)
	bodyJSON := assert.MarshalWithoutError(t, fn)
	assert.Equal(t, "{\"script_score\":{\"script\":{\"lang\":\"painless\",\"source\":\"_score * 2\"}},\"weight\":5}", bodyJSON)
}

////   Complex Function Score Query   ////

func Test_FunctionScore_should_create_complex_json_with_all_parameters(t *testing.T) {
	t.Parallel()
	// Given
	query := es.NewQuery(
		es.FunctionScore(es.MatchAll()).
			Boost(5).
			MaxBoost(42).
			ScoreMode(ScoreMode.Max).
			BoostMode(BoostMode.Multiply).
			MinScore(5).
			Functions(
				es.WeightFunction(23).Filter(es.Term("status", "published")),
				es.ScriptScoreFunction(es.ScriptSource("_score * doc['likes'].value", ScriptLanguage.Painless)),
				es.RandomScoreFunction().Seed(42).Field("_seq_no"),
			),
	)

	// When Then
	assert.NotNil(t, query)
	bodyJSON := assert.MarshalWithoutError(t, query)
	// nolint:golint,lll
	assert.Equal(t, "{\"query\":{\"function_score\":{\"boost\":5,\"boost_mode\":\"multiply\",\"functions\":[{\"filter\":{\"term\":{\"status\":{\"value\":\"published\"}}},\"weight\":23},{\"script_score\":{\"script\":{\"lang\":\"painless\",\"source\":\"_score * doc['likes'].value\"}}},{\"random_score\":{\"field\":\"_seq_no\",\"seed\":42}}],\"max_boost\":42,\"min_score\":5,\"query\":{\"match_all\":{}},\"score_mode\":\"max\"}}}", bodyJSON)
}
