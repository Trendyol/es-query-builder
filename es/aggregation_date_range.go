package es

type dateRangeAggType Object

type dateRangeAggEntry Object

// DateRangeAgg creates a date range aggregation for a given field.
//
// A date range aggregation is dedicated for date values. The main difference between
// this aggregation and the normal range aggregation is that the from and to values can
// be expressed in date math expressions.
//
// Example usage:
//
//	agg := es.DateRangeAgg("date").
//		Format("MM-yyyy").
//		Range(es.DateRangeEntry().To("now-10M/M")).
//		Range(es.DateRangeEntry().From("now-10M/M"))
//
// Parameters:
//   - field: The date field on which the date range aggregation is applied.
//
// Returns:
//
//	An es.dateRangeAggType object representing the date range aggregation.
func DateRangeAgg(field string) dateRangeAggType {
	return dateRangeAggType{
		"date_range": Object{
			"field": field,
		},
	}
}

// Range adds a date range entry to the date range aggregation.
//
// This method appends a date range entry to the "ranges" array of the date range aggregation.
// Each range entry defines a bucket with optional "from" and "to" date boundaries.
//
// Example usage:
//
//	agg := es.DateRangeAgg("date").
//		Range(es.DateRangeEntry().From("2020-01-01").To("2021-01-01")).
//		Range(es.DateRangeEntry().From("2021-01-01"))
//
// Parameters:
//   - entry: An es.dateRangeAggEntry object defining the date range boundaries.
//
// Returns:
//
//	The updated es.dateRangeAggType object with the date range entry added.
func (d dateRangeAggType) Range(entry dateRangeAggEntry) dateRangeAggType {
	if dateRangeObj, ok := d["date_range"].(Object); ok {
		ranges, rOk := dateRangeObj["ranges"].([]dateRangeAggEntry)
		if !rOk {
			ranges = make([]dateRangeAggEntry, 0, 1)
		}
		dateRangeObj["ranges"] = append(ranges, entry)
	}
	return d
}

// Format sets the "format" parameter in the date range aggregation.
//
// This method specifies the date format used to parse the from and to values and
// to format the bucket keys in the response.
//
// Example usage:
//
//	agg := es.DateRangeAgg("date").Format("MM-yyyy")
//
// Parameters:
//   - format: A string representing the date format pattern.
//
// Returns:
//
//	The updated es.dateRangeAggType object with the "format" parameter set.
func (d dateRangeAggType) Format(format string) dateRangeAggType {
	return d.putInTheField("format", format)
}

// Keyed sets the "keyed" parameter in the date range aggregation.
//
// This method specifies whether the buckets should be returned as a hash instead of an array.
//
// Example usage:
//
//	agg := es.DateRangeAgg("date").Keyed(true)
//
// Parameters:
//   - keyed: A boolean indicating whether to return keyed buckets.
//
// Returns:
//
//	The updated es.dateRangeAggType object with the "keyed" parameter set.
func (d dateRangeAggType) Keyed(keyed bool) dateRangeAggType {
	return d.putInTheField("keyed", keyed)
}

// Missing sets a default value to use for documents that do not contain the field.
//
// Example usage:
//
//	agg := es.DateRangeAgg("date").Missing("1970-01-01")
//
// Parameters:
//   - missing: The value to use when a document lacks the field.
//
// Returns:
//
//	An es.dateRangeAggType object with the "missing" field set.
func (d dateRangeAggType) Missing(missing any) dateRangeAggType {
	return d.putInTheField("missing", missing)
}

// TimeZone sets the "time_zone" parameter in the date range aggregation.
//
// This method specifies the time zone to use for date calculations.
//
// Example usage:
//
//	agg := es.DateRangeAgg("date").TimeZone("CET")
//
// Parameters:
//   - timeZone: A string representing the time zone (e.g., "CET", "UTC", "+01:00").
//
// Returns:
//
//	The updated es.dateRangeAggType object with the "time_zone" parameter set.
func (d dateRangeAggType) TimeZone(timeZone string) dateRangeAggType {
	return d.putInTheField("time_zone", timeZone)
}

// Aggs adds sub-aggregations to the date range aggregation.
//
// Example usage:
//
//	agg := es.DateRangeAgg("date").
//		Range(es.DateRangeEntry().To("now-10M/M")).
//		Aggs(es.Agg("avg_price", es.AvgAgg("price")))
//
// Parameters:
//   - aggs: A variadic list of sub-aggregations.
//
// Returns:
//
//	An es.dateRangeAggType object with the specified sub-aggregations added.
func (d dateRangeAggType) Aggs(aggs ...aggsType) dateRangeAggType {
	return genericPutAggsInRoot(d, aggs)
}

// Meta adds metadata to the date range aggregation.
//
// Example usage:
//
//	agg := es.DateRangeAgg("date").Meta("description", "Date ranges")
//
// Parameters:
//   - key: Metadata key.
//   - value: Metadata value.
//
// Returns:
//
//	A modified es.dateRangeAggType with the meta field set.
func (d dateRangeAggType) Meta(key string, value any) dateRangeAggType {
	meta, ok := d["meta"].(Object)
	if !ok {
		meta = Object{}
	}
	meta[key] = value
	d["meta"] = meta
	return d
}

func (d dateRangeAggType) putInTheField(key string, value any) dateRangeAggType {
	return genericPutInTheField(d, "date_range", key, value)
}

// DateRangeEntry creates a new date range entry for use in a date range aggregation.
//
// Example usage:
//
//	entry := es.DateRangeEntry().From("2020-01-01").To("2021-01-01")
//
// Returns:
//
//	An es.dateRangeAggEntry object ready for configuration.
func DateRangeEntry() dateRangeAggEntry {
	return dateRangeAggEntry{}
}

// From sets the "from" boundary of the date range entry (inclusive).
//
// Example usage:
//
//	entry := es.DateRangeEntry().From("now-10M/M")
//
// Parameters:
//   - from: The lower date boundary of the range.
//
// Returns:
//
//	The updated es.dateRangeAggEntry object with the "from" boundary set.
func (e dateRangeAggEntry) From(from any) dateRangeAggEntry {
	e["from"] = from
	return e
}

// To sets the "to" boundary of the date range entry (exclusive).
//
// Example usage:
//
//	entry := es.DateRangeEntry().To("now-10M/M")
//
// Parameters:
//   - to: The upper date boundary of the range.
//
// Returns:
//
//	The updated es.dateRangeAggEntry object with the "to" boundary set.
func (e dateRangeAggEntry) To(to any) dateRangeAggEntry {
	e["to"] = to
	return e
}

// Key sets a custom key for the date range entry bucket.
//
// Example usage:
//
//	entry := es.DateRangeEntry().Key("last_year").From("now-1y/y").To("now/y")
//
// Parameters:
//   - key: A string representing the custom bucket key.
//
// Returns:
//
//	The updated es.dateRangeAggEntry object with the "key" set.
func (e dateRangeAggEntry) Key(key string) dateRangeAggEntry {
	e["key"] = key
	return e
}
