package es

type extendedStatsAggType Object

// ExtendedStatsAgg creates an extended stats aggregation for the given field.
//
// The extended stats aggregation calculates statistical metrics such as count, min, max, sum, avg,
// variance, standard deviation, and more for the specified field.
//
// Example usage:
//
//	agg := es.ExtendedStatsAgg("price")
//	// agg now contains an es.extendedStatsAggType object that calculates extended statistics for the "price" field.
//
// Parameters:
//   - field: The field for which extended statistical metrics should be computed.
//
// Returns:
//
//	An es.extendedStatsAggType object representing the extended stats aggregation.
func ExtendedStatsAgg(field string) extendedStatsAggType {
	return extendedStatsAggType{
		"extended_stats": Object{
			"field": field,
		},
	}
}

// Missing sets the "missing" value for the extended stats aggregation.
//
// This value is used when documents do not have a value for the specified field. It allows missing values
// to be treated as a specific number instead of being ignored.
//
// Example usage:
//
//	agg := es.ExtendedStatsAgg("price").Missing(0)
//	// agg now contains an es.extendedStatsAggType object that treats missing values as 0.
//
// Parameters:
//   - missing: The value to use when a document lacks a field value.
//
// Returns:
//
//	An es.extendedStatsAggType object with the "missing" field set.
func (extendedStats extendedStatsAggType) Missing(missing any) extendedStatsAggType {
	return extendedStats.putInTheField("missing", missing)
}

// Script sets a script for the extended stats aggregation instead of using a field value.
//
// This allows the aggregation to be computed based on a script, rather than directly referencing
// a field in the document.
//
// Example usage:
//
//	agg := es.ExtendedStatsAgg("price").Script(es.ScriptSource("doc['price'].value * 1.2", ScriptLanguage.Painless))
//	// agg now contains an es.extendedStatsAggType object that applies a script for computing the values.
//
// Parameters:
//   - script: A script to calculate values dynamically.
//
// Returns:
//
//	An es.extendedStatsAggType object with the "script" field set.
func (extendedStats extendedStatsAggType) Script(script scriptType) extendedStatsAggType {
	return extendedStats.putInTheField("script", script)
}

// Format sets the output format for the extended stats aggregation.
//
// This is used to specify how the results should be formatted when returned.
//
// Example usage:
//
//	agg := es.ExtendedStatsAgg("price").Format("000.0")
//	// agg now contains an es.extendedStatsAggType object with a defined number format.
//
// Parameters:
//   - format: A string specifying the output format.
//
// Returns:
//
//	An es.extendedStatsAggType object with the "format" field set.
func (extendedStats extendedStatsAggType) Format(format string) extendedStatsAggType {
	return extendedStats.putInTheField("format", format)
}

// Meta sets a custom metadata field for the extended stats aggregation.
//
// This allows additional information to be attached to the aggregation result for tracking or processing.
//
// Example usage:
//
//	agg := es.ExtendedStatsAgg("price").Meta("source", "sales_data")
//	// agg now contains an es.extendedStatsAggType object with metadata indicating the source is "sales_data".
//
// Parameters:
//   - key: The metadata key.
//   - value: The metadata value.
//
// Returns:
//
//	An es.extendedStatsAggType object with the "meta" field set.
func (extendedStats extendedStatsAggType) Meta(key string, value any) extendedStatsAggType {
	meta, exists := getFieldFromAggs(extendedStats, "extended_stats", "meta")
	if !exists {
		meta = Object{}
	}
	meta[key] = value
	return extendedStats.putInTheField("meta", meta)
}

// Aggs adds sub-aggregations to the extended stats aggregation.
//
// This allows performing additional calculations on top of the extended stats aggregation.
//
// Example usage:
//
//	agg := es.ExtendedStatsAgg("price").Aggs(es.Agg("price_distribution", es.HistogramAgg("price", 10)))
//	// agg now contains an es.extendedStatsAggType object with a histogram sub-aggregation.
//
// Parameters:
//   - aggs: A variadic list of aggsType representing the sub-aggregations.
//
// Returns:
//
//	An es.extendedStatsAggType object with the specified sub-aggregations added.
func (extendedStats extendedStatsAggType) Aggs(aggs ...aggsType) extendedStatsAggType {
	if len(aggs) == 1 && aggs[0] == nil {
		return extendedStats
	}
	extendedStats["aggs"] = reduceAggs(aggs...)
	return extendedStats
}

func (extendedStats extendedStatsAggType) putInTheField(key string, value any) extendedStatsAggType {
	if extendedStatsAgg, ok := extendedStats["extended_stats"].(Object); ok {
		extendedStatsAgg[key] = value
	}
	return extendedStats
}
