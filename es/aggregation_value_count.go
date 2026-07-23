package es

type valueCountAggType Object

// ValueCountAgg creates a new es.valueCountAggType object to count the number of
// values that are extracted from the aggregated documents.
//
// Example usage:
//
//	agg := es.ValueCountAgg("price")
//	// agg now contains an es.valueCountAggType that counts values in the "price" field.
//
// Parameters:
//   - field: The name of the field for which values should be counted.
//
// Returns:
//
//	An es.valueCountAggType object that counts values for the specified field.
func ValueCountAgg(field string) valueCountAggType {
	return valueCountAggType{
		"value_count": Object{
			"field": field,
		},
	}
}

// Missing sets the "missing" field for the value_count aggregation.
//
// Example usage:
//
//	agg := es.ValueCountAgg("price").Missing(0)
//
// Parameters:
//   - missing: The value to use for documents that do not have the field value.
//
// Returns:
//
//	An es.valueCountAggType object with the "missing" field set.
func (vc valueCountAggType) Missing(missing any) valueCountAggType {
	return vc.putInTheField("missing", missing)
}

// Script sets the "script" field for the value_count aggregation.
//
// Example usage:
//
//	agg := es.ValueCountAgg("price").Script(es.ScriptSource("doc['price'].value", ScriptLanguage.Painless))
//
// Parameters:
//   - script: The script to be used for the aggregation.
//
// Returns:
//
//	An es.valueCountAggType object with the "script" field set.
func (vc valueCountAggType) Script(script scriptType) valueCountAggType {
	return vc.putInTheField("script", script)
}

// Meta sets the "meta" field for the value_count aggregation.
//
// Example usage:
//
//	agg := es.ValueCountAgg("price").Meta("source", "sales_data")
//
// Parameters:
//   - key: The key for the metadata entry.
//   - value: The value associated with the metadata key.
//
// Returns:
//
//	An es.valueCountAggType object with the "meta" field set.
func (vc valueCountAggType) Meta(key string, value any) valueCountAggType {
	meta, ok := vc["meta"].(Object)
	if !ok {
		meta = Object{}
	}
	meta[key] = value
	vc["meta"] = meta
	return vc
}

// Aggs adds sub-aggregations to the value_count aggregation.
//
// Example usage:
//
//	agg := es.ValueCountAgg("price").Aggs(es.Agg("by_category", es.TermsAgg("category")))
//
// Parameters:
//   - aggs: A variadic list of es.aggsType representing the sub-aggregations.
//
// Returns:
//
//	An es.valueCountAggType object with the specified sub-aggregations added.
func (vc valueCountAggType) Aggs(aggs ...aggsType) valueCountAggType {
	return genericPutAggsInRoot(vc, aggs)
}

func (vc valueCountAggType) putInTheField(key string, value any) valueCountAggType {
	return genericPutInTheField(vc, "value_count", key, value)
}
