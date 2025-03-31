package es

type nestedAggType Object

// NestedAgg creates a nested aggregation for querying nested fields in Elasticsearch.
//
// A nested aggregation allows searching within objects in a `nested` field type.
// It enables running sub-aggregations on the nested objects rather than on the parent documents.
//
// Example usage:
//
//	agg := es.NestedAgg("products")
//	// This creates a nested aggregation on the "products" field.
//
// Parameters:
//   - path: The nested field path to aggregate on.
//
// Returns:
//
//	An es.nestedAggType object representing the nested aggregation.
func NestedAgg(path string) nestedAggType {
	return nestedAggType{
		"nested": Object{
			"path": path,
		},
	}
}

// Aggs adds sub-aggregations to the nested aggregation.
//
// This method allows performing additional aggregations inside the nested context.
// It is commonly used to perform further breakdowns on nested fields.
//
// Example usage:
//
//	agg := es.NestedAgg("products").
//		Aggs(es.Agg("max_price", es.MaxAgg("products.price")))
//
// Parameters:
//   - aggs: A variadic list of sub-aggregations.
//
// Returns:
//
//	An es.nestedAggType object with the specified sub-aggregations added.
func (nested nestedAggType) Aggs(aggs ...aggsType) nestedAggType {
	if len(aggs) == 1 && aggs[0] == nil {
		return nested
	}
	nested["aggs"] = reduceAggs(aggs...)
	return nested
}
