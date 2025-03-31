package es

type minAggType Object

// MinAgg creates a min aggregation for the given field.
//
// The min aggregation calculates the minimum value of a specified numeric field.
//
// Example usage:
//
//	agg := es.MinAgg("price")
//	// agg now contains an es.minAggType object that calculates the minimum value for the "price" field.
//
// Parameters:
//   - field: The field for which the minimum value should be computed.
//
// Returns:
//
//	An es.minAggType object representing the min aggregation.
func MinAgg(field string) minAggType {
	return minAggType{
		"min": Object{
			"field": field,
		},
	}
}

// Missing sets the "missing" value for the min aggregation.
//
// This value is used when documents do not have a value for the specified field. It allows missing values
// to be treated as a specific number instead of being ignored.
//
// Example usage:
//
//	agg := es.MinAgg("price").Missing(0)
//	// agg now contains an es.minAggType object that treats missing values as 0.
//
// Parameters:
//   - missing: The value to use when a document lacks a field value.
//
// Returns:
//
//	An es.minAggType object with the "missing" field set.
func (min minAggType) Missing(missing any) minAggType {
	return min.putInTheField("missing", missing)
}

// Script sets a script for the min aggregation instead of using a field value.
//
// This allows the aggregation to be computed based on a script, rather than directly referencing
// a field in the document.
//
// Example usage:
//
//	agg := es.MinAgg("price").Script(es.ScriptSource("doc['price'].value * 1.2", ScriptLanguage.Painless))
//	// agg now contains an es.minAggType object that applies a script for computing the values.
//
// Parameters:
//   - script: A script to calculate values dynamically.
//
// Returns:
//
//	An es.minAggType object with the "script" field set.
func (min minAggType) Script(script scriptType) minAggType {
	return min.putInTheField("script", script)
}

// Format sets the output format for the min aggregation.
//
// This is used to specify how the results should be formatted when returned.
//
// Example usage:
//
//	agg := es.MinAgg("price").Format("000.0")
//	// agg now contains an es.minAggType object with a defined number format.
//
// Parameters:
//   - format: A string specifying the output format.
//
// Returns:
//
//	An es.minAggType object with the "format" field set.
func (min minAggType) Format(format string) minAggType {
	return min.putInTheField("format", format)
}

// Meta sets a custom metadata field for the min aggregation.
//
// This allows additional information to be attached to the aggregation result for tracking or processing.
//
// Example usage:
//
//	agg := es.MinAgg("price").Meta("source", "sales_data")
//	// agg now contains an es.minAggType object with metadata indicating the source is "sales_data".
//
// Parameters:
//   - key: The metadata key.
//   - value: The metadata value.
//
// Returns:
//
//	An es.minAggType object with the "meta" field set.
func (min minAggType) Meta(key string, value any) minAggType {
	meta, exists := getObjectFromAggs(min, "min", "meta")
	if !exists {
		meta = Object{}
	}
	meta[key] = value
	return min.putInTheField("meta", meta)
}

// Aggs adds sub-aggregations to the min aggregation.
//
// This allows performing additional calculations on top of the min aggregation.
//
// Example usage:
//
//	agg := es.MinAgg("price").Aggs(es.Agg("price_distribution", es.HistogramAgg("price", 10)))
//	// agg now contains an es.minAggType object with a histogram sub-aggregation.
//
// Parameters:
//   - aggs: A variadic list of aggsType representing the sub-aggregations.
//
// Returns:
//
//	An es.minAggType object with the specified sub-aggregations added.
func (min minAggType) Aggs(aggs ...aggsType) minAggType {
	if len(aggs) == 1 && aggs[0] == nil {
		return min
	}
	min["aggs"] = reduceAggs(aggs...)
	return min
}

func (min minAggType) putInTheField(key string, value any) minAggType {
	if minAgg, ok := min["min"].(Object); ok {
		minAgg[key] = value
	}
	return min
}
