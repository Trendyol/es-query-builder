package es

type dateHistogramAggType Object

// DateHistogramAgg creates a date histogram aggregation for a given field.
//
// A date histogram aggregation is similar to the normal histogram aggregation, but it can
// only be applied on date or date range values. The interval is specified using date/time
// expressions.
//
// Example usage:
//
//	agg := es.DateHistogramAgg("timestamp").CalendarInterval("month")
//	// This creates a date histogram aggregation on "timestamp" with monthly intervals.
//
// Parameters:
//   - field: The date field on which the date histogram aggregation is applied.
//
// Returns:
//
//	An es.dateHistogramAggType object representing the date histogram aggregation.
func DateHistogramAgg(field string) dateHistogramAggType {
	return dateHistogramAggType{
		"date_histogram": Object{
			"field": field,
		},
	}
}

// CalendarInterval sets the "calendar_interval" parameter in the date histogram aggregation.
//
// Calendar-aware intervals understand that daylight savings changes the length of specific
// days, months have different amounts of days, and leap seconds can be tacked on to a
// particular year. Supported values: minute, 1m, hour, 1h, day, 1d, week, 1w, month, 1M,
// quarter, 1q, year, 1y.
//
// Example usage:
//
//	agg := es.DateHistogramAgg("timestamp").CalendarInterval("month")
//
// Parameters:
//   - calendarInterval: A string representing the calendar interval.
//
// Returns:
//
//	The updated es.dateHistogramAggType object with the "calendar_interval" parameter set.
func (d dateHistogramAggType) CalendarInterval(calendarInterval string) dateHistogramAggType {
	return d.putInTheField("calendar_interval", calendarInterval)
}

// FixedInterval sets the "fixed_interval" parameter in the date histogram aggregation.
//
// Fixed intervals are a fixed number of SI units and never deviate, regardless of where
// they fall on the calendar. Supported units: ms (milliseconds), s (seconds), m (minutes),
// h (hours), d (days).
//
// Example usage:
//
//	agg := es.DateHistogramAgg("timestamp").FixedInterval("30d")
//
// Parameters:
//   - fixedInterval: A string representing the fixed interval.
//
// Returns:
//
//	The updated es.dateHistogramAggType object with the "fixed_interval" parameter set.
func (d dateHistogramAggType) FixedInterval(fixedInterval string) dateHistogramAggType {
	return d.putInTheField("fixed_interval", fixedInterval)
}

// Format sets the "format" parameter in the date histogram aggregation.
//
// This method specifies the date format used to format the bucket keys in the response.
//
// Example usage:
//
//	agg := es.DateHistogramAgg("timestamp").CalendarInterval("month").Format("yyyy-MM-dd")
//
// Parameters:
//   - format: A string representing the date format pattern.
//
// Returns:
//
//	The updated es.dateHistogramAggType object with the "format" parameter set.
func (d dateHistogramAggType) Format(format string) dateHistogramAggType {
	return d.putInTheField("format", format)
}

// TimeZone sets the "time_zone" parameter in the date histogram aggregation.
//
// This method specifies the time zone to use for date calculations. Date-times are stored
// in UTC in Elasticsearch. By default, all bucketing and rounding is done in UTC. The
// time_zone parameter can be used to indicate that bucketing should use a different time zone.
//
// Example usage:
//
//	agg := es.DateHistogramAgg("timestamp").CalendarInterval("day").TimeZone("-01:00")
//
// Parameters:
//   - timeZone: A string representing the time zone (e.g., "CET", "UTC", "+01:00").
//
// Returns:
//
//	The updated es.dateHistogramAggType object with the "time_zone" parameter set.
func (d dateHistogramAggType) TimeZone(timeZone string) dateHistogramAggType {
	return d.putInTheField("time_zone", timeZone)
}

// Offset sets the "offset" parameter in the date histogram aggregation.
//
// This method changes the start value of each bucket by the specified positive (+) or
// negative offset (-) duration, such as 1h for an hour, or 1d for a day.
//
// Example usage:
//
//	agg := es.DateHistogramAgg("timestamp").CalendarInterval("day").Offset("+6h")
//
// Parameters:
//   - offset: A string representing the offset duration.
//
// Returns:
//
//	The updated es.dateHistogramAggType object with the "offset" parameter set.
func (d dateHistogramAggType) Offset(offset string) dateHistogramAggType {
	return d.putInTheField("offset", offset)
}

// MinDocCount sets the minimum document count required for a bucket to be included.
//
// Example usage:
//
//	agg := es.DateHistogramAgg("timestamp").CalendarInterval("month").MinDocCount(1)
//
// Parameters:
//   - minDocCount: The minimum number of documents required for a bucket.
//
// Returns:
//
//	The updated es.dateHistogramAggType object with the "min_doc_count" parameter set.
func (d dateHistogramAggType) MinDocCount(minDocCount int) dateHistogramAggType {
	return d.putInTheField("min_doc_count", minDocCount)
}

// ExtendedBounds sets the "extended_bounds" parameter in the date histogram aggregation.
//
// This method forces the date histogram to start building buckets on a specific min date
// and keep building buckets up to a max date, even if there are no documents in some buckets.
//
// Example usage:
//
//	agg := es.DateHistogramAgg("timestamp").CalendarInterval("month").
//		ExtendedBounds("2020-01-01", "2020-12-31")
//
// Parameters:
//   - min: The minimum boundary for the date histogram.
//   - max: The maximum boundary for the date histogram.
//
// Returns:
//
//	The updated es.dateHistogramAggType object with the "extended_bounds" parameter set.
func (d dateHistogramAggType) ExtendedBounds(min, max any) dateHistogramAggType {
	return d.putInTheField("extended_bounds", Object{
		"min": min,
		"max": max,
	})
}

// HardBounds sets the "hard_bounds" parameter in the date histogram aggregation.
//
// This method limits the range of buckets in the date histogram. Buckets outside the
// hard bounds will not be generated.
//
// Example usage:
//
//	agg := es.DateHistogramAgg("timestamp").CalendarInterval("month").
//		HardBounds("2020-01-01", "2020-12-31")
//
// Parameters:
//   - min: The minimum hard boundary.
//   - max: The maximum hard boundary.
//
// Returns:
//
//	The updated es.dateHistogramAggType object with the "hard_bounds" parameter set.
func (d dateHistogramAggType) HardBounds(min, max any) dateHistogramAggType {
	return d.putInTheField("hard_bounds", Object{
		"min": min,
		"max": max,
	})
}

// Keyed sets the "keyed" parameter in the date histogram aggregation.
//
// This method specifies whether the buckets should be returned as a hash instead of an array.
//
// Example usage:
//
//	agg := es.DateHistogramAgg("timestamp").CalendarInterval("month").Keyed(true)
//
// Parameters:
//   - keyed: A boolean indicating whether to return keyed buckets.
//
// Returns:
//
//	The updated es.dateHistogramAggType object with the "keyed" parameter set.
func (d dateHistogramAggType) Keyed(keyed bool) dateHistogramAggType {
	return d.putInTheField("keyed", keyed)
}

// Missing sets a default value to use for documents that do not contain the field.
//
// Example usage:
//
//	agg := es.DateHistogramAgg("timestamp").CalendarInterval("month").Missing("2000-01-01")
//
// Parameters:
//   - missing: The value to use when a document lacks the field.
//
// Returns:
//
//	An es.dateHistogramAggType object with the "missing" field set.
func (d dateHistogramAggType) Missing(missing any) dateHistogramAggType {
	return d.putInTheField("missing", missing)
}

// Order sets the sorting order of the date histogram buckets.
//
// Example usage:
//
//	agg := es.DateHistogramAgg("timestamp").CalendarInterval("month").
//		Order(es.AggOrder("_count", Order.Desc))
//
// Parameters:
//   - orders: A variadic list of sorting rules.
//
// Returns:
//
//	The updated es.dateHistogramAggType object with the "order" parameter set.
func (d dateHistogramAggType) Order(orders ...aggOrder) dateHistogramAggType {
	if len(orders) == 1 && orders[0] == nil {
		return d
	}
	return d.putInTheField("order", orders)
}

// Aggs adds sub-aggregations to the date histogram aggregation.
//
// Example usage:
//
//	agg := es.DateHistogramAgg("timestamp").CalendarInterval("month").
//		Aggs(es.Agg("avg_price", es.AvgAgg("price")))
//
// Parameters:
//   - aggs: A variadic list of sub-aggregations.
//
// Returns:
//
//	An es.dateHistogramAggType object with the specified sub-aggregations added.
func (d dateHistogramAggType) Aggs(aggs ...aggsType) dateHistogramAggType {
	return genericPutAggsInRoot(d, aggs)
}

// Meta adds metadata to the date histogram aggregation.
//
// Example usage:
//
//	agg := es.DateHistogramAgg("timestamp").Meta("description", "Monthly histogram")
//
// Parameters:
//   - key: Metadata key.
//   - value: Metadata value.
//
// Returns:
//
//	A modified es.dateHistogramAggType with the meta field set.
func (d dateHistogramAggType) Meta(key string, value any) dateHistogramAggType {
	meta, ok := d["meta"].(Object)
	if !ok {
		meta = Object{}
	}
	meta[key] = value
	d["meta"] = meta
	return d
}

func (d dateHistogramAggType) putInTheField(key string, value any) dateHistogramAggType {
	return genericPutInTheField(d, "date_histogram", key, value)
}
