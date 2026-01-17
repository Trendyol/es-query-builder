package es

import (
	CollectMode "github.com/Trendyol/es-query-builder/es/enums/collect-mode"
	ExecutionHint "github.com/Trendyol/es-query-builder/es/enums/execution-hint"
)

type multiTermsAggType Object

type termAggType Object

// MultiTermsAgg creates a multi-terms aggregation using the specified term aggregations.
//
// Multi-terms aggregation is similar to a terms aggregation but allows specifying multiple fields
// for bucketing documents.
//
// Example usage:
//
//	agg := es.MultiTermsAgg(es.TermAgg("category"), es.TermAgg("brand"))
//	// agg now contains an multiTermsAggType object that groups documents by category and brand.
//
// Parameters:
//   - terms: A variadic list of term aggregations defining the fields for bucketing.
//
// Returns:
//
//	An es.multiTermsAggType object representing the multi-terms aggregation.
func MultiTermsAgg(terms ...termAggType) multiTermsAggType {
	return multiTermsAggType{
		"multi_terms": Object{
			"terms": terms,
		},
	}
}

// TermAgg creates a term aggregation for the given field.
//
// A term aggregation is used to bucket documents based on unique values of a field.
//
// Example usage:
//
//	term := es.TermAgg("category")
//	// term now represents a field-based term aggregation.
//
// Parameters:
//   - field: The field to group documents by.
//
// Returns:
//
//	A termAggType object representing the term aggregation.
func TermAgg(field string) termAggType {
	return termAggType{
		"field": field,
	}
}

// Missing sets a default value to use for documents that do not contain the field.
//
// Example usage:
//
//	agg := es.MultiTermsAgg(es.TermAgg("category")).Missing("unknown")
//	// Documents without "category" will be assigned "unknown".
//
// Parameters:
//   - missing: The value to use when a document lacks the field.
//
// Returns:
//
//	An es.multiTermsAggType object with the "missing" field set.
func (multiTerms multiTermsAggType) Missing(missing any) multiTermsAggType {
	return multiTerms.putInTheField("missing", missing)
}

// Script sets a script to compute dynamic bucket values instead of using a field.
//
// Example usage:
//
//	agg := es.MultiTermsAgg(es.TermAgg("category")).Script(es.ScriptSource("doc['category'].value + '_modified'", ScriptLanguage.Painless))
//
// Parameters:
//   - script: A script for computing dynamic term values.
//
// Returns:
//
//	An es.multiTermsAggType object with the "script" field set.
func (multiTerms multiTermsAggType) Script(script scriptType) multiTermsAggType {
	return multiTerms.putInTheField("script", script)
}

// Size sets the maximum number of buckets to return.
//
// Example usage:
//
//	agg := es.MultiTermsAgg(es.TermAgg("category")).Size(10)
//	// Limits the number of buckets to 10.
//
// Parameters:
//   - size: The maximum number of buckets.
//
// Returns:
//
//	An es.multiTermsAggType object with the "size" field set.
func (multiTerms multiTermsAggType) Size(size int) multiTermsAggType {
	return multiTerms.putInTheField("size", size)
}

// IgnoreUnmapped specifies whether to ignore fields that are not mapped in the index.
//
// Example usage:
//
//	agg := es.MultiTermsAgg(es.TermAgg("category")).IgnoreUnmapped(true)
//
// Parameters:
//   - ignoreUnmapped: Whether to ignore unmapped fields.
//
// Returns:
//
//	An es.multiTermsAggType object with the "ignore_unmapped" field set.
func (multiTerms multiTermsAggType) IgnoreUnmapped(ignoreUnmapped bool) multiTermsAggType {
	return multiTerms.putInTheField("ignore_unmapped", ignoreUnmapped)
}

// ShardSize sets the number of term buckets per shard before merging.
//
// Example usage:
//
//	agg := es.MultiTermsAgg(es.TermAgg("category")).ShardSize(1000)
//
// Parameters:
//   - shardSize: The number of term buckets per shard.
//
// Returns:
//
//	An es.multiTermsAggType object with the "shard_size" field set.
func (multiTerms multiTermsAggType) ShardSize(shardSize float64) multiTermsAggType {
	return multiTerms.putInTheField("shard_size", shardSize)
}

// Include filters buckets by matching only specific terms.
//
// Example usage:
//
//	agg := es.MultiTermsAgg(es.TermAgg("category")).Include("electronics", "clothing")
//
// Parameters:
//   - include: A variadic list of terms to include.
//
// Returns:
//
//	An es.multiTermsAggType object with the "include" field set.
func (multiTerms multiTermsAggType) Include(include ...string) multiTermsAggType {
	return multiTerms.putInTheField("include", include)
}

// Exclude filters buckets by removing specific terms.
//
// Example usage:
//
//	agg := es.MultiTermsAgg(es.TermAgg("category")).Exclude("miscellaneous")
//
// Parameters:
//   - exclude: A variadic list of terms to exclude.
//
// Returns:
//
//	An es.multiTermsAggType object with the "exclude" field set.
func (multiTerms multiTermsAggType) Exclude(exclude ...string) multiTermsAggType {
	return multiTerms.putInTheField("exclude", exclude)
}

// MinDocCount sets the minimum document count required for a term to be included in the results.
//
// Example usage:
//
//	agg := es.MultiTermsAgg(es.TermAgg("category")).MinDocCount(5)
//	// Only includes terms that appear in at least 5 documents.
//
// Parameters:
//   - minDocCount: The minimum document count for a term to be included.
//
// Returns:
//
//	An es.multiTermsAggType object with the "min_doc_count" field set.
func (multiTerms multiTermsAggType) MinDocCount(minDocCount int) multiTermsAggType {
	return multiTerms.putInTheField("min_doc_count", minDocCount)
}

// ExecutionHint sets the execution strategy for terms aggregation.
//
// Example usage:
//
//	agg := es.MultiTermsAgg(es.TermAgg("category")).ExecutionHint(ExecutionHint.Map)
//
// Parameters:
//   - executionHint: The execution strategy (e.g., "map", "global_ordinals").
//
// Returns:
//
//	An es.multiTermsAggType object with the "execution_hint" field set.
func (multiTerms multiTermsAggType) ExecutionHint(executionHint ExecutionHint.ExecutionHint) multiTermsAggType {
	return multiTerms.putInTheField("execution_hint", executionHint)
}

// CollectMode sets the collection mode for the multi-terms aggregation.
//
// Example usage:
//
//	agg := es.MultiTermsAgg(es.TermAgg("category")).CollecMode(CollectMode.BreadthFirst)
//
// Parameters:
//   - collectMode: The collection mode (e.g., "depth_first", "breadth_first").
//
// Returns:
//
//	An es.multiTermsAggType object with the "collect_mode" field set.
func (multiTerms multiTermsAggType) CollectMode(collectMode CollectMode.CollectMode) multiTermsAggType {
	return multiTerms.putInTheField("collect_mode", collectMode)
}

// Order sets the sorting order for term buckets.
//
// Example usage:
//
//	agg := es.MultiTermsAgg(es.TermAgg("category")).Order(es.Order("_count", "desc"))
//
// Parameters:
//   - orders: A variadic list of ordering rules.
//
// Returns:
//
//	An es.multiTermsAggType object with the "order" field set.
func (multiTerms multiTermsAggType) Order(orders ...aggOrder) multiTermsAggType {
	if len(orders) == 1 && orders[0] == nil {
		return multiTerms
	}
	return multiTerms.putInTheField("order", orders)
}

// Aggs adds sub-aggregations to the multi-terms aggregation.
//
// Example usage:
//
//	agg := es.MultiTermsAgg(es.TermAgg("category")).Aggs(es.Agg("avg_price", es.AvgAgg("price")))
//
// Parameters:
//   - aggs: A variadic list of aggsType representing sub-aggregations.
//
// Returns:
//
//	An es.multiTermsAggType object with the specified sub-aggregations added.
func (multiTerms multiTermsAggType) Aggs(aggs ...aggsType) multiTermsAggType {
	if len(aggs) == 1 && aggs[0] == nil {
		return multiTerms
	}
	multiTerms["aggs"] = reduceAggs(aggs...)
	return multiTerms
}

func (multiTerms multiTermsAggType) putInTheField(key string, value any) multiTermsAggType {
	return genericPutInTheField(multiTerms, "multi_terms", key, value)
}
