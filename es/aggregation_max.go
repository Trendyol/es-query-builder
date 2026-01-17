package es

type maxAggType Object

// MaxAgg creates a max aggregation for the given field.
//
// The max aggregation calculates the maximum value of a specified numeric field.
//
// Example usage:
//
//	agg := es.MaxAgg("price")
//	// agg now contains an es.maxAggType object that calculates the maximum value for the "price" field.
//
// Parameters:
//   - field: The field for which the maximum value should be computed.
//
// Returns:
//
//	An es.maxAggType object representing the max aggregation.
func MaxAgg(field string) maxAggType {
	return maxAggType{
		"max": Object{
			"field": field,
		},
	}
}

// Missing sets the "missing" value for the max aggregation.
//
// This value is used when documents do not have a value for the specified field. It allows missing values
// to be treated as a specific number instead of being ignored.
//
// Example usage:
//
//	agg := es.MaxAgg("price").Missing(0)
//	// agg now contains an es.maxAggType object that treats missing values as 0.
//
// Parameters:
//   - missing: The value to use when a document lacks a field value.
//
// Returns:
//
//	An es.maxAggType object with the "missing" field set.
func (max maxAggType) Missing(missing any) maxAggType {
	return max.putInTheField("missing", missing)
}

// Script sets a script for the max aggregation instead of using a field value.
//
// This allows the aggregation to be computed based on a script, rather than directly referencing
// a field in the document.
//
// Example usage:
//
//	agg := es.MaxAgg("price").Script(es.ScriptSource("doc['price'].value * 1.2", ScriptLanguage.Painless))
//	// agg now contains an es.maxAggType object that applies a script for computing the values.
//
// Parameters:
//   - script: A script to calculate values dynamically.
//
// Returns:
//
//	An es.maxAggType object with the "script" field set.
func (max maxAggType) Script(script scriptType) maxAggType {
	return max.putInTheField("script", script)
}

// Format sets the output format for the max aggregation.
//
// This is used to specify how the results should be formatted when returned.
//
// Example usage:
//
//	agg := es.MaxAgg("price").Format("000.0")
//	// agg now contains an es.maxAggType object with a defined number format.
//
// Parameters:
//   - format: A string specifying the output format.
//
// Returns:
//
//	An es.maxAggType object with the "format" field set.
func (max maxAggType) Format(format string) maxAggType {
	return max.putInTheField("format", format)
}

// Meta sets a custom metadata field for the max aggregation.
//
// This allows additional information to be attached to the aggregation result for tracking or processing.
//
// Example usage:
//
//	agg := es.MaxAgg("price").Meta("source", "sales_data")
//	// agg now contains an es.maxAggType object with metadata indicating the source is "sales_data".
//
// Parameters:
//   - key: The metadata key.
//   - value: The metadata value.
//
// Returns:
//
//	An es.maxAggType object with the "meta" field set.
func (max maxAggType) Meta(key string, value any) maxAggType {
	meta, ok := max["meta"].(Object)
	if !ok {
		meta = Object{}
	}
	meta[key] = value
	max["meta"] = meta
	return max
}

// Aggs adds sub-aggregations to the max aggregation.
//
// This allows performing additional calculations on top of the max aggregation.
//
// Example usage:
//
//	agg := es.MaxAgg("price").Aggs(es.Agg("price_distribution", es.HistogramAgg("price", 10)))
//	// agg now contains an es.maxAggType object with a histogram sub-aggregation.
//
// Parameters:
//   - aggs: A variadic list of aggsType representing the sub-aggregations.
//
// Returns:
//
//	An es.maxAggType object with the specified sub-aggregations added.
func (max maxAggType) Aggs(aggs ...aggsType) maxAggType {
	return genericPutAggsInRoot(max, aggs)
}

func (max maxAggType) putInTheField(key string, value any) maxAggType {
	return genericPutInTheField(max, "max", key, value)
}
