package es

type cardinalityAggType Object

// CardinalityAgg creates a new es.cardinalityAggType object to calculate the cardinality (distinct count)
// of a specified field.
//
// This function initializes an es.cardinalityAggType object that computes the number of unique values
// for a specified field.
//
// Example usage:
//
//	agg := es.CardinalityAgg("user_id")
//	// agg now contains an es.cardinalityAggType object that calculates the distinct count of the "user_id" field.
//
// Parameters:
//   - field: The name of the field for which the cardinality should be calculated. It must be a valid field name.
//
// Returns:
//
//	An es.cardinalityAggType object that calculates the cardinality of the specified field.
func CardinalityAgg(field string) cardinalityAggType {
	return cardinalityAggType{
		"cardinality": Object{
			"field": field,
		},
	}
}

// Missing sets the "missing" field for the cardinality aggregation, which defines the value to use when a document
// does not have a value for the specified field.
//
// This method allows you to specify a missing value for documents that lack a value in the field. The missing value
// is treated as a unique value in the cardinality calculation.
//
// Example usage:
//
//	agg := es.CardinalityAgg("user_id").Missing(0)
//	// agg now contains an es.cardinalityAggType object with a missing value of 0 for documents without the "user_id" field.
//
// Parameters:
//   - missing: The value to use for documents that do not have the field value. It can be any value that matches
//     the field's type.
//
// Returns:
//
//	An es.cardinalityAggType object with the "missing" field set to the provided value.
func (cardinality cardinalityAggType) Missing(missing any) cardinalityAggType {
	return cardinality.putInTheField("missing", missing)
}

// PrecisionThreshold sets the "precision_threshold" field for the cardinality aggregation, which controls the
// accuracy of the cardinality computation. Lower precision values will result in faster but less accurate results.
//
// Example usage:
//
//	agg := es.CardinalityAgg("user_id").PrecisionThreshold(1000)
//	// agg now contains an es.cardinalityAggType object with a precision threshold of 1000.
//
// Parameters:
//   - precisionThreshold: The precision threshold for cardinality. Higher values increase accuracy but may affect
//     performance.
//
// Returns:
//
//	An es.cardinalityAggType object with the "precision_threshold" field set to the provided value.
func (cardinality cardinalityAggType) PrecisionThreshold(precisionThreshold int) cardinalityAggType {
	return cardinality.putInTheField("precision_threshold", precisionThreshold)
}

// Script sets the "script" field for the cardinality aggregation, allowing you to use a script to compute the values
// for the aggregation instead of using the field value.
//
// This method enables the calculation of cardinality using a script, which can be useful for dynamic field values.
//
// Example usage:
//
//	agg := es.CardinalityAgg("user_id").Script(es.ScriptSource("doc['user_id'].value * 2", ScriptLanguage.Painless))
//	// agg now contains an es.cardinalityAggType object that calculates the cardinality using the provided script.
//
// Parameters:
//   - script: The script to be used for the aggregation. It must be of type scriptType.
//
// Returns:
//
//	An es.cardinalityAggType object with the "script" field set to the provided script.
func (cardinality cardinalityAggType) Script(script scriptType) cardinalityAggType {
	return cardinality.putInTheField("script", script)
}

// Meta sets the "meta" field for the cardinality aggregation, which is a custom metadata field that can be used to
// store additional information related to the aggregation.
//
// This method allows you to attach arbitrary metadata to the aggregation, which can be helpful for tracking or
// processing purposes.
//
// Example usage:
//
//	agg := es.CardinalityAgg("user_id").Meta("source", "user_data")
//	// agg now contains an es.cardinalityAggType object with a meta field storing the source as "user_data".
//
// Parameters:
//   - key: The key for the metadata entry. It must be a string.
//   - value: The value associated with the metadata key. It can be any type.
//
// Returns:
//
//	An es.cardinalityAggType object with the "meta" field set to the provided metadata.
func (cardinality cardinalityAggType) Meta(key string, value any) cardinalityAggType {
	meta, exists := getObjectFromAggs(cardinality, "cardinality", "meta")
	if !exists {
		meta = Object{}
	}
	meta[key] = value
	return cardinality.putInTheField("meta", meta)
}

// Aggs adds sub-aggregations to the cardinality aggregation. This method allows you to perform further
// aggregation operations on the result of the cardinality aggregation.
//
// Example usage:
//
//	agg := es.CardinalityAgg("user_id").Aggs(es.Agg("category_count", es.TermsAgg("category")))
//	// agg now contains an es.cardinalityAggType object with a sub-aggregation for the "category" field.
//
// Parameters:
//   - aggs: A variadic list of aggsType representing the sub-aggregations to apply.
//
// Returns:
//
//	An es.cardinalityAggType object with the specified sub-aggregations added.
func (cardinality cardinalityAggType) Aggs(aggs ...aggsType) cardinalityAggType {
	if len(aggs) == 1 && aggs[0] == nil {
		return cardinality
	}
	cardinality["aggs"] = reduceAggs(aggs...)
	return cardinality
}

func (cardinality cardinalityAggType) putInTheField(key string, value any) cardinalityAggType {
	if cardinalityAgg, ok := cardinality["cardinality"].(Object); ok {
		cardinalityAgg[key] = value
	}
	return cardinality
}
