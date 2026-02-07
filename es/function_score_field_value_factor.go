package es

import Modifier "github.com/Trendyol/es-query-builder/es/enums/modifier"

type fieldValueFactorType Object

// FieldValueFactor creates a new es.fieldValueFactorType object with the specified field.
//
// This function initializes a field_value_factor configuration that uses the value of a
// document field to influence the score. The field parameter represents the name of the
// field whose value will be used in the score calculation.
//
// Example usage:
//
//	fvf := es.FieldValueFactor("likes")
//	// fvf now contains a field_value_factor configuration for the "likes" field.
//
// Parameters:
//   - field: A string representing the field name whose value will be used for scoring.
//
// Returns:
//
//	An es.fieldValueFactorType object with the specified field.
func FieldValueFactor(field string) fieldValueFactorType {
	return fieldValueFactorType{
		"field": field,
	}
}

// Factor sets the "factor" parameter in an es.fieldValueFactorType.
//
// This method specifies an optional factor to multiply the field value with before
// applying the modifier. The default value is 1.
//
// Example usage:
//
//	fvf := es.FieldValueFactor("likes").Factor(1.2)
//	// fvf now includes a "factor" parameter set to 1.2.
//
// Parameters:
//   - factor: A float64 value representing the factor to multiply the field value with.
//
// Returns:
//
//	The updated es.fieldValueFactorType object with the "factor" parameter set.
func (fvf fieldValueFactorType) Factor(factor float64) fieldValueFactorType {
	fvf["factor"] = factor
	return fvf
}

// Modifier sets the "modifier" parameter in an es.fieldValueFactorType.
//
// This method specifies the mathematical function to apply to the field value before
// using it in the score calculation. Available modifiers include none, log, log1p,
// log2p, ln, ln1p, ln2p, square, sqrt, and reciprocal.
//
// Example usage:
//
//	fvf := es.FieldValueFactor("likes").Modifier(Modifier.Log1p)
//	// fvf now includes a "modifier" parameter set to "log1p".
//
// Parameters:
//   - modifier: A Modifier.Modifier value representing the modifier function.
//
// Returns:
//
//	The updated es.fieldValueFactorType object with the "modifier" parameter set.
func (fvf fieldValueFactorType) Modifier(modifier Modifier.Modifier) fieldValueFactorType {
	fvf["modifier"] = modifier
	return fvf
}

// Missing sets the "missing" parameter in an es.fieldValueFactorType.
//
// This method specifies the value to use if the document does not have the specified field.
// The modifier and factor are still applied to the missing value as though it were coming
// from the document.
//
// Example usage:
//
//	fvf := es.FieldValueFactor("likes").Missing(1)
//	// fvf now includes a "missing" parameter set to 1.
//
// Parameters:
//   - missing: A value to use when the document lacks the field.
//
// Returns:
//
//	The updated es.fieldValueFactorType object with the "missing" parameter set.
func (fvf fieldValueFactorType) Missing(missing any) fieldValueFactorType {
	fvf["missing"] = missing
	return fvf
}
