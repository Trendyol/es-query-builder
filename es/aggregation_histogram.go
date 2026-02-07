package es

type histogramAggType Object

// HistogramAgg creates a histogram aggregation for a given field.
//
// A histogram aggregation groups documents into fixed-size intervals (buckets) based on
// the values of a numeric field. Each bucket covers an interval of the specified size.
//
// Example usage:
//
//	agg := es.HistogramAgg("price", 50)
//	// This creates a histogram aggregation on "price" with an interval of 50.
//
// Parameters:
//   - field: The numeric field on which the histogram aggregation is applied.
//   - interval: The interval size for each bucket.
//
// Returns:
//
//	An es.histogramAggType object representing the histogram aggregation.
func HistogramAgg(field string, interval float64) histogramAggType {
	return histogramAggType{
		"histogram": Object{
			"field":    field,
			"interval": interval,
		},
	}
}

// MinDocCount sets the minimum document count required for a bucket to be included.
//
// By default, the response will fill gaps in the histogram with empty buckets. Setting
// min_doc_count to a value greater than 0 will only return buckets that meet the threshold.
//
// Example usage:
//
//	agg := es.HistogramAgg("price", 50).MinDocCount(1)
//
// Parameters:
//   - minDocCount: The minimum number of documents required for a bucket.
//
// Returns:
//
//	The updated es.histogramAggType object with the "min_doc_count" parameter set.
func (h histogramAggType) MinDocCount(minDocCount int) histogramAggType {
	return h.putInTheField("min_doc_count", minDocCount)
}

// ExtendedBounds sets the "extended_bounds" parameter in the histogram aggregation.
//
// This method forces the histogram to start building buckets on a specific min value and
// keep building buckets up to a max value, even if there are no documents in some buckets.
//
// Example usage:
//
//	agg := es.HistogramAgg("price", 50).ExtendedBounds(0, 500)
//
// Parameters:
//   - min: The minimum boundary for the histogram.
//   - max: The maximum boundary for the histogram.
//
// Returns:
//
//	The updated es.histogramAggType object with the "extended_bounds" parameter set.
func (h histogramAggType) ExtendedBounds(min, max float64) histogramAggType {
	return h.putInTheField("extended_bounds", Object{
		"min": min,
		"max": max,
	})
}

// HardBounds sets the "hard_bounds" parameter in the histogram aggregation.
//
// This method limits the range of buckets in the histogram. Buckets outside the hard
// bounds will not be generated.
//
// Example usage:
//
//	agg := es.HistogramAgg("price", 50).HardBounds(0, 1000)
//
// Parameters:
//   - min: The minimum hard boundary.
//   - max: The maximum hard boundary.
//
// Returns:
//
//	The updated es.histogramAggType object with the "hard_bounds" parameter set.
func (h histogramAggType) HardBounds(min, max float64) histogramAggType {
	return h.putInTheField("hard_bounds", Object{
		"min": min,
		"max": max,
	})
}

// Offset sets the "offset" parameter in the histogram aggregation.
//
// This method shifts the bucket boundaries by the specified offset. For example, with
// an interval of 10 and an offset of 5, the buckets would be [5, 15), [15, 25), etc.
//
// Example usage:
//
//	agg := es.HistogramAgg("price", 10).Offset(5)
//
// Parameters:
//   - offset: The offset value to shift bucket boundaries.
//
// Returns:
//
//	The updated es.histogramAggType object with the "offset" parameter set.
func (h histogramAggType) Offset(offset float64) histogramAggType {
	return h.putInTheField("offset", offset)
}

// Keyed sets the "keyed" parameter in the histogram aggregation.
//
// This method specifies whether the buckets should be returned as a hash instead of an array.
//
// Example usage:
//
//	agg := es.HistogramAgg("price", 50).Keyed(true)
//
// Parameters:
//   - keyed: A boolean indicating whether to return keyed buckets.
//
// Returns:
//
//	The updated es.histogramAggType object with the "keyed" parameter set.
func (h histogramAggType) Keyed(keyed bool) histogramAggType {
	return h.putInTheField("keyed", keyed)
}

// Missing sets a default value to use for documents that do not contain the field.
//
// Example usage:
//
//	agg := es.HistogramAgg("price", 50).Missing(0)
//
// Parameters:
//   - missing: The value to use when a document lacks the field.
//
// Returns:
//
//	An es.histogramAggType object with the "missing" field set.
func (h histogramAggType) Missing(missing any) histogramAggType {
	return h.putInTheField("missing", missing)
}

// Order sets the sorting order of the histogram buckets.
//
// Example usage:
//
//	agg := es.HistogramAgg("price", 50).Order(es.AggOrder("_count", Order.Desc))
//
// Parameters:
//   - orders: A variadic list of sorting rules.
//
// Returns:
//
//	The updated es.histogramAggType object with the "order" parameter set.
func (h histogramAggType) Order(orders ...aggOrder) histogramAggType {
	if len(orders) == 1 && orders[0] == nil {
		return h
	}
	return h.putInTheField("order", orders)
}

// Aggs adds sub-aggregations to the histogram aggregation.
//
// Example usage:
//
//	agg := es.HistogramAgg("price", 50).
//		Aggs(es.Agg("avg_score", es.AvgAgg("score")))
//
// Parameters:
//   - aggs: A variadic list of sub-aggregations.
//
// Returns:
//
//	An es.histogramAggType object with the specified sub-aggregations added.
func (h histogramAggType) Aggs(aggs ...aggsType) histogramAggType {
	return genericPutAggsInRoot(h, aggs)
}

// Meta adds metadata to the histogram aggregation.
//
// Example usage:
//
//	agg := es.HistogramAgg("price", 50).Meta("description", "Price histogram")
//
// Parameters:
//   - key: Metadata key.
//   - value: Metadata value.
//
// Returns:
//
//	A modified es.histogramAggType with the meta field set.
func (h histogramAggType) Meta(key string, value any) histogramAggType {
	meta, ok := h["meta"].(Object)
	if !ok {
		meta = Object{}
	}
	meta[key] = value
	h["meta"] = meta
	return h
}

func (h histogramAggType) putInTheField(key string, value any) histogramAggType {
	return genericPutInTheField(h, "histogram", key, value)
}
