package es

type rangeAggType Object

type rangeAggEntry Object

// RangeAgg creates a range aggregation for a given field.
//
// A range aggregation allows you to define a set of ranges, each representing a bucket.
// During the aggregation process, the values extracted from each document will be checked
// against each bucket range and the relevant document will be "bucketed".
//
// Example usage:
//
//	agg := es.RangeAgg("price").
//		Range(es.RangeEntry().To(50)).
//		Range(es.RangeEntry().From(50).To(100)).
//		Range(es.RangeEntry().From(100))
//
// Parameters:
//   - field: The field on which the range aggregation is applied.
//
// Returns:
//
//	An es.rangeAggType object representing the range aggregation.
func RangeAgg(field string) rangeAggType {
	return rangeAggType{
		"range": Object{
			"field": field,
		},
	}
}

// Range adds a range entry to the range aggregation.
//
// This method appends a range entry to the "ranges" array of the range aggregation.
// Each range entry defines a bucket with optional "from" and "to" boundaries.
//
// Example usage:
//
//	agg := es.RangeAgg("price").
//		Range(es.RangeEntry().From(10).To(50)).
//		Range(es.RangeEntry().From(50).To(100))
//
// Parameters:
//   - entry: An es.rangeAggEntry object defining the range boundaries.
//
// Returns:
//
//	The updated es.rangeAggType object with the range entry added.
func (r rangeAggType) Range(entry rangeAggEntry) rangeAggType {
	if rangeObj, ok := r["range"].(Object); ok {
		ranges, rOk := rangeObj["ranges"].([]rangeAggEntry)
		if !rOk {
			ranges = make([]rangeAggEntry, 0, 1)
		}
		rangeObj["ranges"] = append(ranges, entry)
	}
	return r
}

// Keyed sets the "keyed" parameter in the range aggregation.
//
// This method specifies whether the buckets should be returned as a hash instead of an array.
// When set to true, each bucket is associated with a unique string key.
//
// Example usage:
//
//	agg := es.RangeAgg("price").Keyed(true)
//
// Parameters:
//   - keyed: A boolean indicating whether to return keyed buckets.
//
// Returns:
//
//	The updated es.rangeAggType object with the "keyed" parameter set.
func (r rangeAggType) Keyed(keyed bool) rangeAggType {
	return r.putInTheField("keyed", keyed)
}

// Missing sets a default value to use for documents that do not contain the field.
//
// Example usage:
//
//	agg := es.RangeAgg("price").Missing(0)
//
// Parameters:
//   - missing: The value to use when a document lacks the field.
//
// Returns:
//
//	An es.rangeAggType object with the "missing" field set.
func (r rangeAggType) Missing(missing any) rangeAggType {
	return r.putInTheField("missing", missing)
}

// Aggs adds sub-aggregations to the range aggregation.
//
// Example usage:
//
//	agg := es.RangeAgg("price").
//		Range(es.RangeEntry().To(50)).
//		Aggs(es.Agg("avg_score", es.AvgAgg("score")))
//
// Parameters:
//   - aggs: A variadic list of sub-aggregations.
//
// Returns:
//
//	An es.rangeAggType object with the specified sub-aggregations added.
func (r rangeAggType) Aggs(aggs ...aggsType) rangeAggType {
	return genericPutAggsInRoot(r, aggs)
}

// Meta adds metadata to the range aggregation.
//
// Example usage:
//
//	agg := es.RangeAgg("price").Meta("description", "Price ranges")
//
// Parameters:
//   - key: Metadata key.
//   - value: Metadata value.
//
// Returns:
//
//	A modified es.rangeAggType with the meta field set.
func (r rangeAggType) Meta(key string, value any) rangeAggType {
	meta, ok := r["meta"].(Object)
	if !ok {
		meta = Object{}
	}
	meta[key] = value
	r["meta"] = meta
	return r
}

func (r rangeAggType) putInTheField(key string, value any) rangeAggType {
	return genericPutInTheField(r, "range", key, value)
}

// RangeEntry creates a new range entry for use in a range aggregation.
//
// Example usage:
//
//	entry := es.RangeEntry().From(10).To(50)
//
// Returns:
//
//	An es.rangeAggEntry object ready for configuration.
func RangeEntry() rangeAggEntry {
	return rangeAggEntry{}
}

// From sets the "from" boundary of the range entry (inclusive).
//
// Example usage:
//
//	entry := es.RangeEntry().From(50)
//
// Parameters:
//   - from: The lower boundary of the range.
//
// Returns:
//
//	The updated es.rangeAggEntry object with the "from" boundary set.
func (e rangeAggEntry) From(from any) rangeAggEntry {
	e["from"] = from
	return e
}

// To sets the "to" boundary of the range entry (exclusive).
//
// Example usage:
//
//	entry := es.RangeEntry().To(100)
//
// Parameters:
//   - to: The upper boundary of the range.
//
// Returns:
//
//	The updated es.rangeAggEntry object with the "to" boundary set.
func (e rangeAggEntry) To(to any) rangeAggEntry {
	e["to"] = to
	return e
}

// Key sets a custom key for the range entry bucket.
//
// Example usage:
//
//	entry := es.RangeEntry().Key("cheap").To(50)
//
// Parameters:
//   - key: A string representing the custom bucket key.
//
// Returns:
//
//	The updated es.rangeAggEntry object with the "key" set.
func (e rangeAggEntry) Key(key string) rangeAggEntry {
	e["key"] = key
	return e
}
