package es

import (
	Order "github.com/Trendyol/es-query-builder/es/enums/sort/order"
)

type compositeSourceType Object

// CompositeTerms creates a terms source for a composite aggregation.
//
// Example usage:
//
//	source := es.CompositeTerms("brand", "brand.keyword")
//
// Parameters:
//   - name: The name of the source (used as the key in after_key).
//   - field: The field to bucket on.
//
// Returns:
//
//	An es.compositeSourceType representing a terms composite source.
func CompositeTerms(name, field string) compositeSourceType {
	return compositeSourceType{
		name: Object{
			"terms": Object{
				"field": field,
			},
		},
	}
}

// CompositeHistogram creates a histogram source for a composite aggregation.
//
// Example usage:
//
//	source := es.CompositeHistogram("price", "price", 50)
//
// Parameters:
//   - name: The name of the source (used as the key in after_key).
//   - field: The numeric field to bucket on.
//   - interval: The interval size for each bucket.
//
// Returns:
//
//	An es.compositeSourceType representing a histogram composite source.
func CompositeHistogram(name, field string, interval float64) compositeSourceType {
	return compositeSourceType{
		name: Object{
			"histogram": Object{
				"field":    field,
				"interval": interval,
			},
		},
	}
}

// CompositeDateHistogram creates a date_histogram source for a composite aggregation.
//
// Example usage:
//
//	source := es.CompositeDateHistogram("by_day", "timestamp").CalendarInterval("day")
//
// Parameters:
//   - name: The name of the source (used as the key in after_key).
//   - field: The date field to bucket on.
//
// Returns:
//
//	An es.compositeSourceType representing a date_histogram composite source.
func CompositeDateHistogram(name, field string) compositeSourceType {
	return compositeSourceType{
		name: Object{
			"date_histogram": Object{
				"field": field,
			},
		},
	}
}

// MissingBucket includes documents that are missing the source field in a dedicated bucket.
//
// Example usage:
//
//	source := es.CompositeTerms("brand", "brand.keyword").MissingBucket(true)
//
// Parameters:
//   - missingBucket: Whether to include a bucket for missing values.
//
// Returns:
//
//	The updated es.compositeSourceType with the "missing_bucket" field set.
func (cs compositeSourceType) MissingBucket(missingBucket bool) compositeSourceType {
	return cs.putInSource("missing_bucket", missingBucket)
}

// Order sets the sort order of the composite source buckets.
//
// Example usage:
//
//	source := es.CompositeTerms("brand", "brand.keyword").Order(Order.Desc)
//
// Parameters:
//   - order: An Order.Order value (asc or desc).
//
// Returns:
//
//	The updated es.compositeSourceType with the "order" field set.
func (cs compositeSourceType) Order(order Order.Order) compositeSourceType {
	return cs.putInSource("order", order)
}

// Offset sets the "offset" parameter for a histogram composite source.
//
// Example usage:
//
//	source := es.CompositeHistogram("price", "price", 10).Offset(5)
//
// Parameters:
//   - offset: The offset value to shift bucket boundaries.
//
// Returns:
//
//	The updated es.compositeSourceType with the "offset" field set.
func (cs compositeSourceType) Offset(offset float64) compositeSourceType {
	return cs.putInSource("offset", offset)
}

// CalendarInterval sets the "calendar_interval" for a date_histogram composite source.
//
// Example usage:
//
//	source := es.CompositeDateHistogram("by_day", "timestamp").CalendarInterval("day")
//
// Parameters:
//   - calendarInterval: A string representing the calendar interval.
//
// Returns:
//
//	The updated es.compositeSourceType with the "calendar_interval" field set.
func (cs compositeSourceType) CalendarInterval(calendarInterval string) compositeSourceType {
	return cs.putInSource("calendar_interval", calendarInterval)
}

// FixedInterval sets the "fixed_interval" for a date_histogram composite source.
//
// Example usage:
//
//	source := es.CompositeDateHistogram("by_day", "timestamp").FixedInterval("1d")
//
// Parameters:
//   - fixedInterval: A string representing the fixed interval.
//
// Returns:
//
//	The updated es.compositeSourceType with the "fixed_interval" field set.
func (cs compositeSourceType) FixedInterval(fixedInterval string) compositeSourceType {
	return cs.putInSource("fixed_interval", fixedInterval)
}

// Format sets the "format" for a date_histogram composite source.
//
// Example usage:
//
//	source := es.CompositeDateHistogram("by_day", "timestamp").
//		CalendarInterval("day").Format("yyyy-MM-dd")
//
// Parameters:
//   - format: A string representing the date format pattern.
//
// Returns:
//
//	The updated es.compositeSourceType with the "format" field set.
func (cs compositeSourceType) Format(format string) compositeSourceType {
	return cs.putInSource("format", format)
}

// TimeZone sets the "time_zone" for a date_histogram composite source.
//
// Example usage:
//
//	source := es.CompositeDateHistogram("by_day", "timestamp").
//		CalendarInterval("day").TimeZone("+03:00")
//
// Parameters:
//   - timeZone: A string representing the time zone.
//
// Returns:
//
//	The updated es.compositeSourceType with the "time_zone" field set.
func (cs compositeSourceType) TimeZone(timeZone string) compositeSourceType {
	return cs.putInSource("time_zone", timeZone)
}

func (cs compositeSourceType) putInSource(key string, value any) compositeSourceType {
	for _, v := range cs {
		if sourceObj, ok := v.(Object); ok {
			for _, nested := range sourceObj {
				if nestedObj, ok := nested.(Object); ok {
					nestedObj[key] = value
					return cs
				}
			}
		}
	}
	return cs
}
