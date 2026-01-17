package es

type avgAggType Object

// AvgAgg creates a new es.avgAggType object to calculate the average value of a specified field.
//
// This function initializes an es.avgAggType object that computes the average of values in a specific field.
// The field parameter represents the name of the field for which the average should be calculated.
//
// Example usage:
//
//	agg := es.AvgAgg("price")
//	// agg now contains an es.avgAggType object that calculates the average of the "price" field.
//
// Parameters:
//   - field: The name of the field for which the average value is calculated. It must be a valid field name.
//
// Returns:
//
//	An es.avgAggType object that calculates the average value for the specified field.
func AvgAgg(field string) avgAggType {
	return avgAggType{
		"avg": Object{
			"field": field,
		},
	}
}

// Missing sets the "missing" field for the average aggregation, which defines the value to use when a document
// does not have a value for the specified field.
//
// This method allows you to specify a missing value for documents that lack a value in the field.
// The missing value is included in the calculation of the average aggregation.
//
// Example usage:
//
//	agg := es.AvgAgg("price").Missing(0)
//	// agg now contains an es.avgAggType object with a missing value of 0 for documents without the "price" field.
//
// Parameters:
//   - missing: The value to use for documents that do not have the field value. It can be any value that matches
//     the field's type.
//
// Returns:
//
//	An es.avgAggType object with the "missing" field set to the provided value.
func (avg avgAggType) Missing(missing any) avgAggType {
	return avg.putInTheField("missing", missing)
}

// Script sets the "script" field for the average aggregation, allowing you to use a script to compute the values
// for the aggregation instead of using the field value.
//
// This method enables the calculation of the average using a script, which can be useful for dynamic field values.
//
// Example usage:
//
//	agg := es.AvgAgg("price").Script(es.ScriptSource("doc['price'].value * 1.1", ScriptLanguage.Painless))
//	// agg now contains an es.avgAggType object that calculates the average using the provided script.
//
// Parameters:
//   - script: The script to be used for the aggregation. It must be of type scriptType.
//
// Returns:
//
//	An es.avgAggType object with the "script" field set to the provided script.
func (avg avgAggType) Script(script scriptType) avgAggType {
	return avg.putInTheField("script", script)
}

// Format sets the "format" field for the average aggregation, which defines the format in which the result
// of the aggregation should be returned.
//
// This method allows you to specify a format for the computed average value (e.g., for date or numeric formatting).
//
// Example usage:
//
//	agg := es.AvgAgg("price").Format("0.00")
//	// agg now contains an es.avgAggType object with the format "0.00" for the average value.
//
// Parameters:
//   - format: The format string to apply to the average result. It must be a valid format string.
//
// Returns:
//
//	An es.avgAggType object with the "format" field set to the provided format.
func (avg avgAggType) Format(format string) avgAggType {
	return avg.putInTheField("format", format)
}

// Meta sets the "meta" field for the average aggregation, which is a custom metadata field that can be used to
// store additional information related to the aggregation.
//
// This method allows you to attach arbitrary metadata to the aggregation, which can be helpful for tracking or
// processing purposes.
//
// Example usage:
//
//	agg := es.AvgAgg("price").Meta("source", "sales_data")
//	// agg now contains an es.avgAggType object with a meta field storing the source as "sales_data".
//
// Parameters:
//   - key: The key for the metadata entry. It must be a string.
//   - value: The value associated with the metadata key. It can be any type.
//
// Returns:
//
//	An es.avgAggType object with the "meta" field set to the provided metadata.
func (avg avgAggType) Meta(key string, value any) avgAggType {
	meta, ok := avg["meta"].(Object)
	if !ok {
		meta = Object{}
	}
	meta[key] = value
	avg["meta"] = meta
	return avg
}

// Aggs adds sub-aggregations to the average aggregation. This method allows you to perform further
// aggregation operations on the result of the average aggregation.
//
// Example usage:
//
//	agg := es.AvgAgg("price").Aggs(es.Agg("category_price", es.TermsAgg("category")))
//	// agg now contains an avgAggType object with a sub-aggregation for the "category" field.
//
// Parameters:
//   - aggs: A variadic list of es.aggsType representing the sub-aggregations to apply.
//
// Returns:
//
//	An es.avgAggType object with the specified sub-aggregations added.
func (avg avgAggType) Aggs(aggs ...aggsType) avgAggType {
	if len(aggs) == 1 && aggs[0] == nil {
		return avg
	}
	avg["aggs"] = reduceAggs(aggs...)
	return avg
}

func (avg avgAggType) putInTheField(key string, value any) avgAggType {
	return genericPutInTheField(avg, "avg", key, value)
}
