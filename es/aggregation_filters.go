package es

type filtersAggType Object

// FiltersAgg creates a filters aggregation that defines a multi-bucket aggregation where each
// bucket is associated with a named filter. Each bucket will collect all documents that match
// its associated filter.
//
// Example usage:
//
//	agg := es.FiltersAgg().
//		Filter("errors", es.Term("status", "error")).
//		Filter("warnings", es.Term("status", "warning"))
//	// This creates a filters aggregation with two named buckets.
//
// Returns:
//
//	An es.filtersAggType object representing the filters aggregation.
func FiltersAgg() filtersAggType {
	return filtersAggType{
		"filters": Object{
			"filters": Object{},
		},
	}
}

// Filter adds a named filter to the filters aggregation.
//
// This method adds a named filter bucket to the filters aggregation. Each named filter
// defines a bucket that will collect all documents matching the filter query.
//
// Example usage:
//
//	agg := es.FiltersAgg().
//		Filter("active", es.Term("status", "active")).
//		Filter("inactive", es.Term("status", "inactive"))
//
// Parameters:
//   - name: The name of the filter bucket.
//   - filter: The query clause for this filter bucket.
//
// Returns:
//
//	The updated es.filtersAggType object with the named filter added.
func (f filtersAggType) Filter(name string, filter any) filtersAggType {
	if filtersObj, ok := f["filters"].(Object); ok {
		if filters, fOk := filtersObj["filters"].(Object); fOk {
			if field, cOk := correctType(filter); cOk {
				filters[name] = field
			}
		}
	}
	return f
}

// OtherBucket sets the "other_bucket" parameter in the filters aggregation.
//
// This method adds a bucket to the response which will contain all documents that do
// not match any of the given filters. When set to true, an additional bucket is created
// for documents not matching any filter.
//
// Example usage:
//
//	agg := es.FiltersAgg().
//		Filter("errors", es.Term("status", "error")).
//		OtherBucket(true)
//
// Parameters:
//   - otherBucket: A boolean indicating whether to include an "other" bucket.
//
// Returns:
//
//	The updated es.filtersAggType object with the "other_bucket" parameter set.
func (f filtersAggType) OtherBucket(otherBucket bool) filtersAggType {
	return f.putInTheField("other_bucket", otherBucket)
}

// OtherBucketKey sets the "other_bucket_key" parameter in the filters aggregation.
//
// This method sets the key for the "other" bucket. By default, the other bucket is
// named "_other_". This parameter allows customizing that name.
//
// Example usage:
//
//	agg := es.FiltersAgg().
//		Filter("errors", es.Term("status", "error")).
//		OtherBucket(true).
//		OtherBucketKey("remaining")
//
// Parameters:
//   - otherBucketKey: A string representing the key for the other bucket.
//
// Returns:
//
//	The updated es.filtersAggType object with the "other_bucket_key" parameter set.
func (f filtersAggType) OtherBucketKey(otherBucketKey string) filtersAggType {
	return f.putInTheField("other_bucket_key", otherBucketKey)
}

// Aggs adds sub-aggregations to the filters aggregation.
//
// Example usage:
//
//	agg := es.FiltersAgg().
//		Filter("errors", es.Term("status", "error")).
//		Aggs(es.Agg("avg_price", es.AvgAgg("price")))
//
// Parameters:
//   - aggs: A variadic list of sub-aggregations.
//
// Returns:
//
//	An es.filtersAggType object with the specified sub-aggregations added.
func (f filtersAggType) Aggs(aggs ...aggsType) filtersAggType {
	return genericPutAggsInRoot(f, aggs)
}

// Meta adds metadata to the filters aggregation.
//
// Example usage:
//
//	agg := es.FiltersAgg().Meta("description", "Status filters")
//
// Parameters:
//   - key: Metadata key.
//   - value: Metadata value.
//
// Returns:
//
//	A modified es.filtersAggType with the meta field set.
func (f filtersAggType) Meta(key string, value any) filtersAggType {
	meta, ok := f["meta"].(Object)
	if !ok {
		meta = Object{}
	}
	meta[key] = value
	f["meta"] = meta
	return f
}

func (f filtersAggType) putInTheField(key string, value any) filtersAggType {
	return genericPutInTheField(f, "filters", key, value)
}
