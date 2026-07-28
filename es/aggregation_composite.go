package es

type compositeAggType Object

// CompositeAgg creates a composite aggregation with the given sources.
//
// A composite aggregation allows paginating through all buckets using an after_key
// cursor. Sources define the dimensions used to build composite buckets.
//
// Example usage:
//
//	agg := es.CompositeAgg(
//		es.CompositeTerms("brand", "brand.keyword"),
//		es.CompositeHistogram("price", "price", 50),
//	).Size(100)
//
// Parameters:
//   - sources: A variadic list of composite source definitions.
//
// Returns:
//
//	An es.compositeAggType object representing the composite aggregation.
func CompositeAgg(sources ...compositeSourceType) compositeAggType {
	return compositeAggType{
		"composite": Object{
			"sources": sources,
		},
	}
}

// Size sets the number of composite buckets to return per page.
//
// Example usage:
//
//	agg := es.CompositeAgg(es.CompositeTerms("brand", "brand.keyword")).Size(100)
//
// Parameters:
//   - size: The maximum number of buckets to return.
//
// Returns:
//
//	A modified es.compositeAggType with the "size" field set.
func (c compositeAggType) Size(size int) compositeAggType {
	return c.putInTheField("size", size)
}

// After sets the after_key cursor used to fetch the next page of buckets.
//
// Example usage:
//
//	agg := es.CompositeAgg(es.CompositeTerms("brand", "brand.keyword")).
//		After(es.Object{"brand": "nike"})
//
// Parameters:
//   - after: An es.Object containing the after_key values from the previous response.
//
// Returns:
//
//	A modified es.compositeAggType with the "after" field set.
func (c compositeAggType) After(after Object) compositeAggType {
	return c.putInTheField("after", after)
}

// Aggs adds sub-aggregations to the composite aggregation.
//
// Example usage:
//
//	agg := es.CompositeAgg(es.CompositeTerms("brand", "brand.keyword")).
//		Aggs(es.Agg("avg_score", es.AvgAgg("score")))
//
// Parameters:
//   - aggs: A variadic list of sub-aggregations.
//
// Returns:
//
//	An es.compositeAggType object with the specified sub-aggregations added.
func (c compositeAggType) Aggs(aggs ...aggsType) compositeAggType {
	return genericPutAggsInRoot(c, aggs)
}

// Meta adds metadata to the composite aggregation.
//
// Example usage:
//
//	agg := es.CompositeAgg(es.CompositeTerms("brand", "brand.keyword")).
//		Meta("description", "Brand pages")
//
// Parameters:
//   - key: Metadata key.
//   - value: Metadata value.
//
// Returns:
//
//	A modified es.compositeAggType with the meta field set.
func (c compositeAggType) Meta(key string, value any) compositeAggType {
	meta, ok := c["meta"].(Object)
	if !ok {
		meta = Object{}
	}
	meta[key] = value
	c["meta"] = meta
	return c
}

func (c compositeAggType) putInTheField(key string, value any) compositeAggType {
	return genericPutInTheField(c, "composite", key, value)
}
