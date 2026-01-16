package es

type statsAggType Object

// StatsAgg creates a statistical aggregation for a given field.
//
// This aggregation computes statistical metrics such as min, max, sum, count, and avg.
//
// Example usage:
//
//	statsAgg := es.StatsAgg("price")
//
// Parameters:
//   - field: The field on which the statistical aggregation is applied.
//
// Returns:
//
//	A statsAggType representing the statistical aggregation.
func StatsAgg(field string) statsAggType {
	return statsAggType{
		"stats": Object{
			"field": field,
		},
	}
}

// Missing sets the "missing" parameter in the stats' aggregation.
//
// This is useful when some documents do not contain the field, allowing Elasticsearch
// to treat missing values as a specified default.
//
// Example usage:
//
//	statsAgg := es.StatsAgg("price").Missing(0)
//
// Parameters:
//   - missing: The value to be used when the field is missing.
//
// Returns:
//
//	An emodified es.statsAggType with the "missing" value set.
func (stats statsAggType) Missing(missing any) statsAggType {
	return stats.putInTheField("missing", missing)
}

// Script applies a script to the stats' aggregation.
//
// Instead of using a field, you can provide a script to compute custom metrics.
//
// Example usage:
//
//	statsAgg := es.StatsAgg("price").Script(es.ScriptSource("doc['price'].value * 2", ScriptLanguage.Painrless))
//
// Parameters:
//   - script: The script to execute for the aggregation.
//
// Returns:
//
//	An emodified es.statsAggType with the script applied.
func (stats statsAggType) Script(script scriptType) statsAggType {
	return stats.putInTheField("script", script)
}

// Format sets the output format for the stats' aggregation.
//
// Example usage:
//
//	statsAgg := es.StatsAgg("price").Format("0.00")
//
// Parameters:
//   - format: A format string, such as `"0.00"` for decimal formatting.
//
// Returns:
//
//	An emodified es.statsAggType with the format set.
func (stats statsAggType) Format(format string) statsAggType {
	return stats.putInTheField("format", format)
}

// Meta adds metadata to the stats' aggregation.
//
// This can store additional information that does not affect the aggregation execution.
//
// Example usage:
//
//	statsAgg := es.StatsAgg("price").Meta("description", "Price statistics")
//
// Parameters:
//   - key: Metadata key.
//   - value: Metadata value.
//
// Returns:
//
//	A modified es.statsAggType with the meta field set.
func (stats statsAggType) Meta(key string, value any) statsAggType {
	meta, ok := stats["meta"].(Object)
	if !ok {
		meta = Object{}
	}
	meta[key] = value
	stats["meta"] = meta
	return stats
}

// Aggs adds sub-aggregations to the stats' aggregation.
//
// Example usage:
//
//	statsAgg := es.StatsAgg("price").Aggs(
//	    es.Agg("max_price", es.MaxAgg("price")),
//	)
//
// Parameters:
//   - aggs: A variadic list of sub-aggregations.
//
// Returns:
//
//	An emodified es.statsAggType containing the nested aggregations.
func (stats statsAggType) Aggs(aggs ...aggsType) statsAggType {
	if len(aggs) == 1 && aggs[0] == nil {
		return stats
	}
	stats["aggs"] = reduceAggs(aggs...)
	return stats
}

func (stats statsAggType) putInTheField(key string, value any) statsAggType {
	if statsAgg, ok := stats["stats"].(Object); ok {
		statsAgg[key] = value
	}
	return stats
}
