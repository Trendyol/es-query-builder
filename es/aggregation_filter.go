package es

type filterAggType Object

// FilterAgg creates a filter aggregation that narrows the set of documents to those matching a query.
//
// A filter aggregation defines a single bucket of all the documents in the current document set
// context that match a specified filter. Often this will be used to narrow down the current
// aggregation context to a specific set of documents.
//
// Example usage:
//
//	agg := es.FilterAgg(es.Term("status", "active"))
//	// This creates a filter aggregation that only includes documents where status is "active".
//
// Parameters:
//   - filter: The query clause to filter documents. It can be of any type.
//
// Returns:
//
//	An es.filterAggType object representing the filter aggregation.
func FilterAgg(filter any) filterAggType {
	if field, ok := correctType(filter); ok {
		return filterAggType{
			"filter": field,
		}
	}
	return filterAggType{
		"filter": Object{},
	}
}

// Aggs adds sub-aggregations to the filter aggregation.
//
// This method allows performing additional aggregations on the filtered set of documents.
//
// Example usage:
//
//	agg := es.FilterAgg(es.Term("status", "active")).
//		Aggs(es.Agg("avg_price", es.AvgAgg("price")))
//
// Parameters:
//   - aggs: A variadic list of sub-aggregations.
//
// Returns:
//
//	An es.filterAggType object with the specified sub-aggregations added.
func (f filterAggType) Aggs(aggs ...aggsType) filterAggType {
	return genericPutAggsInRoot(f, aggs)
}

// Meta adds metadata to the filter aggregation.
//
// This can store additional information that does not affect the aggregation execution.
//
// Example usage:
//
//	agg := es.FilterAgg(es.Term("status", "active")).Meta("description", "Active items")
//
// Parameters:
//   - key: Metadata key.
//   - value: Metadata value.
//
// Returns:
//
//	A modified es.filterAggType with the meta field set.
func (f filterAggType) Meta(key string, value any) filterAggType {
	meta, ok := f["meta"].(Object)
	if !ok {
		meta = Object{}
	}
	meta[key] = value
	f["meta"] = meta
	return f
}
