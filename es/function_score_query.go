package es

import (
	BoostMode "github.com/Trendyol/es-query-builder/es/enums/boost-mode"
	ScoreMode "github.com/Trendyol/es-query-builder/es/enums/score-mode"
)

type functionScoreType Object

type functionScoreFunction Object

// FunctionScore creates a new es.functionScoreType object with the specified query.
//
// This function initializes an es.functionScoreType object that wraps a given query,
// allowing custom scoring functions to be applied to the matching documents. The query
// parameter represents the base query used to determine which documents match.
//
// Example usage:
//
//	fs := es.FunctionScore(es.Bool().Must(es.Term("status", "active")))
//	// fs now contains an es.functionScoreType object with the specified query.
//
// Parameters:
//   - query: An object representing the base query. It can be of any type.
//
// Returns:
//
//	An es.functionScoreType object containing the specified query.
func FunctionScore(query any) functionScoreType {
	o := Object{}
	if field, ok := correctType(query); ok {
		o["query"] = field
	}
	return functionScoreType{
		"function_score": o,
	}
}

// Boost sets the "boost" parameter in an es.functionScoreType query.
//
// This method allows you to specify a boost factor for the function_score query,
// which influences the relevance score of matching documents.
//
// Example usage:
//
//	fs := es.FunctionScore(es.MatchAll()).Boost(5)
//	// fs now includes a "boost" parameter set to 5.
//
// Parameters:
//   - boost: A float64 value representing the boost factor.
//
// Returns:
//
//	The updated es.functionScoreType object with the "boost" parameter set.
func (fs functionScoreType) Boost(boost float64) functionScoreType {
	return fs.putInTheField("boost", boost)
}

// MaxBoost sets the "max_boost" parameter in an es.functionScoreType query.
//
// This method restricts the maximum boost that can be applied by the function_score query.
// Regardless of the computed function score, the final boost will not exceed this value.
//
// Example usage:
//
//	fs := es.FunctionScore(es.MatchAll()).MaxBoost(42)
//	// fs now includes a "max_boost" parameter set to 42.
//
// Parameters:
//   - maxBoost: A float64 value representing the maximum boost allowed.
//
// Returns:
//
//	The updated es.functionScoreType object with the "max_boost" parameter set.
func (fs functionScoreType) MaxBoost(maxBoost float64) functionScoreType {
	return fs.putInTheField("max_boost", maxBoost)
}

// ScoreMode sets the "score_mode" parameter in an es.functionScoreType query.
//
// This method specifies how the scores computed by the individual scoring functions
// should be combined. The available modes include multiply, sum, avg, first, max, and min.
//
// Example usage:
//
//	fs := es.FunctionScore(es.MatchAll()).ScoreMode(ScoreMode.Sum)
//	// fs now includes a "score_mode" parameter set to "sum".
//
// Parameters:
//   - scoreMode: A ScoreMode.ScoreMode value representing the scoring combination mode.
//
// Returns:
//
//	The updated es.functionScoreType object with the "score_mode" parameter set.
func (fs functionScoreType) ScoreMode(scoreMode ScoreMode.ScoreMode) functionScoreType {
	return fs.putInTheField("score_mode", scoreMode)
}

// BoostMode sets the "boost_mode" parameter in an es.functionScoreType query.
//
// This method specifies how the computed function score is combined with the query score.
// The available modes include multiply, replace, sum, avg, max, and min.
//
// Example usage:
//
//	fs := es.FunctionScore(es.MatchAll()).BoostMode(BoostMode.Replace)
//	// fs now includes a "boost_mode" parameter set to "replace".
//
// Parameters:
//   - boostMode: A BoostMode.BoostMode value representing the boost combination mode.
//
// Returns:
//
//	The updated es.functionScoreType object with the "boost_mode" parameter set.
func (fs functionScoreType) BoostMode(boostMode BoostMode.BoostMode) functionScoreType {
	return fs.putInTheField("boost_mode", boostMode)
}

// MinScore sets the "min_score" parameter in an es.functionScoreType query.
//
// This method excludes documents whose combined score is below the specified threshold.
// Documents with a score lower than min_score will not be included in the results.
//
// Example usage:
//
//	fs := es.FunctionScore(es.MatchAll()).MinScore(5)
//	// fs now includes a "min_score" parameter set to 5.
//
// Parameters:
//   - minScore: A float64 value representing the minimum score threshold.
//
// Returns:
//
//	The updated es.functionScoreType object with the "min_score" parameter set.
func (fs functionScoreType) MinScore(minScore float64) functionScoreType {
	return fs.putInTheField("min_score", minScore)
}

// Functions adds one or more scoring functions to the es.functionScoreType query.
//
// This method sets the "functions" array in the function_score query, which defines
// the individual scoring functions to be applied to matching documents. Each function
// can optionally include a filter to restrict which documents it applies to.
//
// Example usage:
//
//	fs := es.FunctionScore(es.MatchAll()).
//		Functions(
//			es.ScriptScoreFunction(es.ScriptSource("_score * doc['likes'].value", ScriptLanguage.Painless)),
//			es.RandomScoreFunction().Seed(42).Field("_seq_no"),
//		)
//	// fs now includes a "functions" array with the specified scoring functions.
//
// Parameters:
//   - functions: A variadic list of es.functionScoreFunction objects.
//
// Returns:
//
//	The updated es.functionScoreType object with the "functions" parameter set.
func (fs functionScoreType) Functions(functions ...functionScoreFunction) functionScoreType {
	if len(functions) == 1 && functions[0] == nil {
		return fs
	}
	funcs := make(Array, 0, len(functions))
	for i := 0; i < len(functions); i++ {
		if functions[i] != nil {
			funcs = append(funcs, functions[i])
		}
	}
	return fs.putInTheField("functions", funcs)
}

// ScriptScore sets the "script_score" parameter directly in the es.functionScoreType query.
//
// This method sets a script_score function at the top level of the function_score query,
// which is used when only a single scoring function is needed without the "functions" array.
//
// Example usage:
//
//	fs := es.FunctionScore(es.MatchAll()).
//		ScriptScore(es.ScriptSource("_score * doc['likes'].value", ScriptLanguage.Painless))
//	// fs now includes a "script_score" parameter with the specified script.
//
// Parameters:
//   - script: A scriptType object representing the script to compute the score.
//
// Returns:
//
//	The updated es.functionScoreType object with the "script_score" parameter set.
func (fs functionScoreType) ScriptScore(script scriptType) functionScoreType {
	return fs.putInTheField("script_score", Object{
		"script": script,
	})
}

// RandomScore sets the "random_score" parameter directly in the es.functionScoreType query.
//
// This method sets a random_score function at the top level of the function_score query,
// which is used when only a single scoring function is needed without the "functions" array.
//
// Example usage:
//
//	fs := es.FunctionScore(es.MatchAll()).RandomScore(42, "_seq_no")
//	// fs now includes a "random_score" parameter with seed 42 and field "_seq_no".
//
// Parameters:
//   - seed: An integer value used as the random seed.
//   - field: A string representing the field to use for generating the random score.
//
// Returns:
//
//	The updated es.functionScoreType object with the "random_score" parameter set.
func (fs functionScoreType) RandomScore(seed int, field string) functionScoreType {
	return fs.putInTheField("random_score", Object{
		"seed":  seed,
		"field": field,
	})
}

// FieldValueFactor sets the "field_value_factor" parameter directly in the es.functionScoreType query.
//
// This method sets a field_value_factor function at the top level of the function_score query,
// which is used when only a single scoring function is needed without the "functions" array.
//
// Example usage:
//
//	fs := es.FunctionScore(es.MatchAll()).
//		FieldValueFactor(es.FieldValueFactor("likes"))
//	// fs now includes a "field_value_factor" parameter for the "likes" field.
//
// Parameters:
//   - fieldValueFactor: A fieldValueFactorType object representing the field value factor configuration.
//
// Returns:
//
//	The updated es.functionScoreType object with the "field_value_factor" parameter set.
func (fs functionScoreType) FieldValueFactor(fieldValueFactor fieldValueFactorType) functionScoreType {
	return fs.putInTheField("field_value_factor", fieldValueFactor)
}

// Weight sets the "weight" parameter directly in the es.functionScoreType query.
//
// This method sets a weight at the top level of the function_score query, which is used
// when only a single scoring function is needed without the "functions" array.
//
// Example usage:
//
//	fs := es.FunctionScore(es.MatchAll()).Weight(2)
//	// fs now includes a "weight" parameter set to 2.
//
// Parameters:
//   - weight: A float64 value representing the weight to apply.
//
// Returns:
//
//	The updated es.functionScoreType object with the "weight" parameter set.
func (fs functionScoreType) Weight(weight float64) functionScoreType {
	return fs.putInTheField("weight", weight)
}

func (fs functionScoreType) putInTheField(key string, value any) functionScoreType {
	return genericPutInTheField(fs, "function_score", key, value)
}

// ScriptScoreFunction creates a new es.functionScoreFunction with a script_score configuration.
//
// This function initializes a function entry for use in the "functions" array of a function_score query.
// The script is used to compute a custom score for each matching document.
//
// Example usage:
//
//	fn := es.ScriptScoreFunction(es.ScriptSource("_score * doc['likes'].value", ScriptLanguage.Painless))
//	// fn now contains a function entry with a script_score configuration.
//
// Parameters:
//   - script: A scriptType object representing the script to compute the score.
//
// Returns:
//
//	An es.functionScoreFunction object containing the script_score configuration.
func ScriptScoreFunction(script scriptType) functionScoreFunction {
	return functionScoreFunction{
		"script_score": Object{
			"script": script,
		},
	}
}

// RandomScoreFunction creates a new es.functionScoreFunction with a random_score configuration.
//
// This function initializes a function entry for use in the "functions" array of a function_score query.
// The random_score generates a uniformly distributed random score.
//
// Example usage:
//
//	fn := es.RandomScoreFunction().Seed(42).Field("_seq_no")
//	// fn now contains a function entry with a random_score configuration.
//
// Returns:
//
//	An es.functionScoreFunction object containing an empty random_score configuration.
func RandomScoreFunction() functionScoreFunction {
	return functionScoreFunction{
		"random_score": Object{},
	}
}

// WeightFunction creates a new es.functionScoreFunction with a weight configuration.
//
// This function initializes a function entry for use in the "functions" array of a function_score query.
// The weight is a number that the score is multiplied by.
//
// Example usage:
//
//	fn := es.WeightFunction(2)
//	// fn now contains a function entry with a weight of 2.
//
// Parameters:
//   - weight: A float64 value representing the weight to apply.
//
// Returns:
//
//	An es.functionScoreFunction object containing the weight configuration.
func WeightFunction(weight float64) functionScoreFunction {
	return functionScoreFunction{
		"weight": weight,
	}
}

// FieldValueFactorFunction creates a new es.functionScoreFunction with a field_value_factor configuration.
//
// This function initializes a function entry for use in the "functions" array of a function_score query.
// The field_value_factor uses the value of a document field to influence the score.
//
// Example usage:
//
//	fn := es.FieldValueFactorFunction(es.FieldValueFactor("likes"))
//	// fn now contains a function entry with a field_value_factor configuration.
//
// Parameters:
//   - fieldValueFactor: A fieldValueFactorType object representing the field value factor configuration.
//
// Returns:
//
//	An es.functionScoreFunction object containing the field_value_factor configuration.
func FieldValueFactorFunction(fieldValueFactor fieldValueFactorType) functionScoreFunction {
	return functionScoreFunction{
		"field_value_factor": fieldValueFactor,
	}
}

// DecayFunction creates a new es.functionScoreFunction with a decay function configuration.
//
// This function initializes a function entry for use in the "functions" array of a function_score query.
// Decay functions score documents based on the distance of a numeric, date, or geo-point field value
// from a given origin. The decayType parameter specifies the decay curve (gauss, linear, or exp).
//
// Example usage:
//
//	fn := es.DecayFunction("gauss", es.Decay("date").Origin("now").Scale("10d").Offset("5d").DecayValue(0.5))
//	// fn now contains a function entry with a gauss decay configuration.
//
// Parameters:
//   - decayType: A string representing the type of decay function ("gauss", "linear", or "exp").
//   - decay: A decayType object representing the decay configuration.
//
// Returns:
//
//	An es.functionScoreFunction object containing the decay function configuration.
func DecayFunction(decayType string, decay decayFunctionType) functionScoreFunction {
	return functionScoreFunction{
		decayType: decay,
	}
}

// Filter sets the "filter" parameter in an es.functionScoreFunction.
//
// This method restricts the scoring function to only apply to documents matching the filter.
// Documents that do not match the filter will not have this function applied to their score.
//
// Example usage:
//
//	fn := es.WeightFunction(2).Filter(es.Term("status", "published"))
//	// fn now includes a "filter" parameter with the specified term query.
//
// Parameters:
//   - filter: An object representing the filter query. It can be of any type.
//
// Returns:
//
//	The updated es.functionScoreFunction object with the "filter" parameter set.
func (f functionScoreFunction) Filter(filter any) functionScoreFunction {
	if field, ok := correctType(filter); ok {
		f["filter"] = field
	}
	return f
}

// Weight sets the "weight" parameter in an es.functionScoreFunction.
//
// This method applies a weight multiplier to the score computed by the function.
// The weight is multiplied with the function score before combining with other functions.
//
// Example usage:
//
//	fn := es.RandomScoreFunction().Seed(42).Field("_seq_no").Weight(2)
//	// fn now includes a "weight" parameter set to 2.
//
// Parameters:
//   - weight: A float64 value representing the weight to apply.
//
// Returns:
//
//	The updated es.functionScoreFunction object with the "weight" parameter set.
func (f functionScoreFunction) Weight(weight float64) functionScoreFunction {
	f["weight"] = weight
	return f
}

// Seed sets the "seed" parameter in the random_score of an es.functionScoreFunction.
//
// This method sets the random seed for the random_score function. Documents with the same
// seed and field value will receive the same score, making the random scoring reproducible.
//
// Example usage:
//
//	fn := es.RandomScoreFunction().Seed(42)
//	// fn now includes a "seed" parameter set to 42 in the random_score.
//
// Parameters:
//   - seed: An integer value used as the random seed.
//
// Returns:
//
//	The updated es.functionScoreFunction object with the "seed" parameter set.
func (f functionScoreFunction) Seed(seed int) functionScoreFunction {
	if randomScore, ok := f["random_score"].(Object); ok {
		randomScore["seed"] = seed
	}
	return f
}

// Field sets the "field" parameter in the random_score of an es.functionScoreFunction.
//
// This method sets the field used for generating the random score. The field value is used
// together with the seed to produce a deterministic random score.
//
// Example usage:
//
//	fn := es.RandomScoreFunction().Seed(42).Field("_seq_no")
//	// fn now includes a "field" parameter set to "_seq_no" in the random_score.
//
// Parameters:
//   - field: A string representing the field to use for random score generation.
//
// Returns:
//
//	The updated es.functionScoreFunction object with the "field" parameter set.
func (f functionScoreFunction) Field(field string) functionScoreFunction {
	if randomScore, ok := f["random_score"].(Object); ok {
		randomScore["field"] = field
	}
	return f
}
