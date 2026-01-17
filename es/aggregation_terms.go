package es

import (
	CollectMode "github.com/Trendyol/es-query-builder/es/enums/collect-mode"
	ExecutionHint "github.com/Trendyol/es-query-builder/es/enums/execution-hint"
)

type termsAggType Object

// TermsAgg creates a terms aggregation for a given field.
//
// This aggregation groups documents by unique terms in the specified field.
//
// Example usage:
//
//	termsAgg := es.TermsAgg("category")
//
// Parameters:
//   - field: The field on which the terms aggregation is applied.
//
// Returns:
//
//	An es.termsAggType representing the terms' aggregation.
func TermsAgg(field string) termsAggType {
	return termsAggType{
		"terms": Object{
			"field": field,
		},
	}
}

// Missing sets a default value to use for documents that do not contain the field.
//
// Example usage:
//
//	agg := es.TermsAgg("category").Missing("unknown")
//	// Documents without "category" will be assigned "unknown".
//
// Parameters:
//   - missing: The value to use when a document lacks the field.
//
// Returns:
//
//	An es.termsAggType object with the "missing" field set.
func (terms termsAggType) Missing(missing any) termsAggType {
	return terms.putInTheField("missing", missing)
}

// Script sets a script to compute dynamic bucket values instead of using a field.
//
// Example usage:
//
//	agg := es.TermsAgg("category").Script(es.ScriptSource("doc['category'].value + '_modified'", ScriptLanguage.Painless))
//
// Parameters:
//   - script: A script for computing dynamic term values.
//
// Returns:
//
//	An es.termsAggType object with the "script" field set.
func (terms termsAggType) Script(script scriptType) termsAggType {
	return terms.putInTheField("script", script)
}

// Size sets the number of term buckets to return.
//
// Example usage:
//
//	termsAgg := es.TermsAgg("category").Size(10)
//
// Parameters:
//   - size: The maximum number of term buckets to return.
//
// Returns:
//
//	A modified es.termsAggType with the "size" field set.
func (terms termsAggType) Size(size int) termsAggType {
	return terms.putInTheField("size", size)
}

// ShardSize sets the number of candidate term buckets per shard before final reduction.
//
// Example usage:
//
//	termsAgg := es.TermsAgg("category").ShardSize(500)
//
// Parameters:
//   - shardSize: The number of term buckets per shard.
//
// Returns:
//
//	A modified es.termsAggType with the "shard_size" field set.
func (terms termsAggType) ShardSize(shardSize float64) termsAggType {
	return terms.putInTheField("shard_size", shardSize)
}

// ShowTermDocCountError indicates whether to return doc count error estimates.
//
// Example usage:
//
//	termsAgg := es.TermsAgg("category").ShowTermDocCountError(true)
//
// Parameters:
//   - showTermDocCountError: Boolean flag to enable error estimation.
//
// Returns:
//
//	A modified es.termsAggType with the "show_term_doc_count_error" field set.
func (terms termsAggType) ShowTermDocCountError(showTermDocCountError bool) termsAggType {
	return terms.putInTheField("show_term_doc_count_error", showTermDocCountError)
}

// Include filters the terms aggregation to include only specified values.
//
// Example usage:
//
//	termsAgg := es.TermsAgg("category").Include("electronics", "furniture")
//
// Parameters:
//   - include: Variadic list of terms to include.
//
// Returns:
//
//	A modified es.termsAggType with the "include" field set.
func (terms termsAggType) Include(include ...string) termsAggType {
	return terms.putInTheField("include", include)
}

// Exclude filters the terms aggregation to exclude specified values.
//
// Example usage:
//
//	termsAgg := es.TermsAgg("category").Exclude("uncategorized")
//
// Parameters:
//   - exclude: Variadic list of terms to exclude.
//
// Returns:
//
//	A modified es.termsAggType with the "exclude" field set.
func (terms termsAggType) Exclude(exclude ...string) termsAggType {
	return terms.putInTheField("exclude", exclude)
}

// MinDocCount sets the minimum document count required for a term to be included.
//
// Example usage:
//
//	termsAgg := es.TermsAgg("category").MinDocCount(5)
//
// Parameters:
//   - minDocCount: The minimum number of documents required for a term bucket.
//
// Returns:
//
//	A modified es.termsAggType with the "min_doc_count" field set.
func (terms termsAggType) MinDocCount(minDocCount int) termsAggType {
	return terms.putInTheField("min_doc_count", minDocCount)
}

// ExecutionHint sets how terms aggregation is executed.
//
// Example usage:
//
//	termsAgg := es.TermsAgg("category").ExecutionHint(ExecutionHint.Map)
//
// Parameters:
//   - executionHint: The execution strategy (e.g., "map" or "global_ordinals").
//
// Returns:
//
//	A modified es.termsAggType with the "execution_hint" field set.
func (terms termsAggType) ExecutionHint(executionHint ExecutionHint.ExecutionHint) termsAggType {
	return terms.putInTheField("execution_hint", executionHint)
}

// CollectMode sets the collection mode for the aggregation.
//
// Example usage:
//
//	termsAgg := es.TermsAgg("category").CollecMode(CollectMode.DepthFirst)
//
// Parameters:
//   - collectMode: The collection mode (e.g., "depth_first" or "breadth_first").
//
// Returns:
//
//	A modified es.termsAggType with the "collect_mode" field set.
func (terms termsAggType) CollectMode(collectMode CollectMode.CollectMode) termsAggType {
	return terms.putInTheField("collect_mode", collectMode)
}

// Order sets the sorting order of the term buckets.
//
// Example usage:
//
//	termsAgg := es.TermsAgg("category").Order(es.AggOrder("_count", Order.Desc))
//
// Parameters:
//   - orders: A variadic list of sorting rules.
//
// Returns:
//
//	A modified es.termsAggType with the "order" field set.
func (terms termsAggType) Order(orders ...aggOrder) termsAggType {
	if len(orders) == 1 && orders[0] == nil {
		return terms
	}
	return terms.putInTheField("order", orders)
}

// Aggs adds sub-aggregations to the terms' aggregation.
//
// Example usage:
//
//	termsAgg := es.TermsAgg("category").Aggs(
//	    es.Agg("avg_price", es.AvgAgg("price")),
//	)
//
// Parameters:
//   - aggs: A variadic list of sub-aggregations.
//
// Returns:
//
//	A modified es.termsAggType containing the nested aggregations.
func (terms termsAggType) Aggs(aggs ...aggsType) termsAggType {
	if len(aggs) == 1 && aggs[0] == nil {
		return terms
	}
	terms["aggs"] = reduceAggs(aggs...)
	return terms
}

func (terms termsAggType) putInTheField(key string, value any) termsAggType {
	return genericPutInTheField(terms, "terms", key, value)
}
