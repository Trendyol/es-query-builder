package es

type sumAggType Object

// SumAgg creates a sum aggregation for a given field.
//
// This aggregation calculates the sum of values for the specified field.
//
// Example usage:
//
//	sumAgg := es.SumAgg("price")
//
// Parameters:
//   - field: The field on which the sum aggregation is applied.
//
// Returns:
//
//	A sumAggType representing the sum aggregation.
func SumAgg(field string) sumAggType {
	return sumAggType{
		"sum": Object{
			"field": field,
		},
	}
}

// Missing sets the "missing" parameter in the sum aggregation.
//
// This is useful when some documents do not contain the field, allowing Elasticsearch
// to treat missing values as a specified default.
//
// Example usage:
//
//	sumAgg := es.SumAgg("price").Missing(0)
//
// Parameters:
//   - missing: The value to be used when the field is missing.
//
// Returns:
//
//	A modified sumAggType with the "missing" value set.
func (sum sumAggType) Missing(missing any) sumAggType {
	return sum.putInTheField("missing", missing)
}

// Script applies a script to the sum aggregation.
//
// Instead of using a field, you can provide a script to compute the sum dynamically.
//
// Example usage:
//
//	sumAgg := es.SumAgg("price").
//	    ScriptSource("scriptName", es.Script("doc['price'].value * 1.2", ScriptLanguage.Painless))
//
// Parameters:
//   - script: The script to execute for the aggregation.
//
// Returns:
//
//	A modified sumAggType with the script applied.
func (sum sumAggType) Script(script scriptType) sumAggType {
	return sum.putInTheField("script", script)
}

// Format sets the output format for the sum aggregation.
//
// Example usage:
//
//	sumAgg := es.SumAgg("price").Format("0.00")
//
// Parameters:
//   - format: A format string, such as `"0.00"` for decimal formatting.
//
// Returns:
//
//	A modified sumAggType with the format set.
func (sum sumAggType) Format(format string) sumAggType {
	return sum.putInTheField("format", format)
}

// Meta adds metadata to the sum aggregation.
//
// This can store additional information that does not affect the aggregation execution.
//
// Example usage:
//
//	sumAgg := es.SumAgg("price").Meta("description", "Total revenue")
//
// Parameters:
//   - key: Metadata key.
//   - value: Metadata value.
//
// Returns:
//
//	A modified sumAggType with the meta field set.
func (sum sumAggType) Meta(key string, value any) sumAggType {
	meta, exists := getObjectFromAggs(sum, "sum", "meta")
	if !exists {
		meta = Object{}
	}
	meta[key] = value
	return sum.putInTheField("meta", meta)
}

// Aggs adds sub-aggregations to the sum aggregation.
//
// Example usage:
//
//	sumAgg := es.SumAgg("price").Aggs(
//	    es.Agg("max_price", es.MaxAgg("price")),
//	)
//
// Parameters:
//   - aggs: A variadic list of sub-aggregations.
//
// Returns:
//
//	A modified sumAggType containing the nested aggregations.
func (sum sumAggType) Aggs(aggs ...aggsType) sumAggType {
	if len(aggs) == 1 && aggs[0] == nil {
		return sum
	}
	sum["aggs"] = reduceAggs(aggs...)
	return sum
}

func (sum sumAggType) putInTheField(key string, value any) sumAggType {
	if sumAgg, ok := sum["sum"].(Object); ok {
		sumAgg[key] = value
	}
	return sum
}
