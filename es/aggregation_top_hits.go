package es

type topHitsAggType Object

// TopHitsAgg creates a top_hits aggregation that returns the most relevant documents
// for each bucket (commonly used as a sub-aggregation under terms aggregations).
//
// Example usage:
//
//	agg := es.TopHitsAgg().Size(3)
//	// agg now contains an es.topHitsAggType object that returns up to 3 hits.
//
// Returns:
//
//	An es.topHitsAggType representing the top_hits aggregation.
func TopHitsAgg() topHitsAggType {
	return topHitsAggType{
		"top_hits": Object{},
	}
}

// Size sets the maximum number of top matching hits to return per bucket.
//
// Example usage:
//
//	agg := es.TopHitsAgg().Size(3)
//
// Parameters:
//   - size: The maximum number of hits to return.
//
// Returns:
//
//	An es.topHitsAggType object with the "size" field set.
func (th topHitsAggType) Size(size int) topHitsAggType {
	return th.putInTheField("size", size)
}

// From sets the offset from the first result to return.
//
// Example usage:
//
//	agg := es.TopHitsAgg().From(0)
//
// Parameters:
//   - from: The starting offset for the hits.
//
// Returns:
//
//	An es.topHitsAggType object with the "from" field set.
func (th topHitsAggType) From(from int) topHitsAggType {
	return th.putInTheField("from", from)
}

// Sort sets the sort criteria for the top hits.
//
// Example usage:
//
//	agg := es.TopHitsAgg().Sort(es.Sort("date").Order(Order.Desc))
//
// Parameters:
//   - sorts: A variadic list of es.sortType objects defining the sorting criteria.
//
// Returns:
//
//	An es.topHitsAggType object with the "sort" field set.
func (th topHitsAggType) Sort(sorts ...sortType) topHitsAggType {
	topHits, ok := th["top_hits"].(Object)
	if !ok {
		return th
	}
	sort, ok := topHits["sort"].([]sortType)
	if !ok {
		sort = make([]sortType, 0, len(sorts))
	}
	topHits["sort"] = append(sort, sorts...)
	return th
}

// SourceFalse sets the "_source" field to false.
//
// Example usage:
//
//	agg := es.TopHitsAgg().SourceFalse()
//
// Returns:
//
//	An es.topHitsAggType object with "_source" set to false.
func (th topHitsAggType) SourceFalse() topHitsAggType {
	return th.putInTheField("_source", false)
}

// SourceIncludes sets the fields to include in the _source of top hits.
//
// Example usage:
//
//	agg := es.TopHitsAgg().SourceIncludes("title", "price")
//
// Parameters:
//   - fields: A variadic list of field names to include.
//
// Returns:
//
//	An es.topHitsAggType object with "_source.includes" set.
func (th topHitsAggType) SourceIncludes(fields ...string) topHitsAggType {
	if len(fields) == 0 {
		return th
	}
	topHits, ok := th["top_hits"].(Object)
	if !ok {
		return th
	}
	source, ok := topHits["_source"].(Object)
	if !ok {
		source = Object{}
	}
	includes, ok := source["includes"].(Array)
	if !ok {
		includes = make(Array, 0, len(fields))
	}
	for i := 0; i < len(fields); i++ {
		includes = append(includes, fields[i])
	}
	source["includes"] = includes
	topHits["_source"] = source
	return th
}

// SourceExcludes sets the fields to exclude from the _source of top hits.
//
// Example usage:
//
//	agg := es.TopHitsAgg().SourceExcludes("description")
//
// Parameters:
//   - fields: A variadic list of field names to exclude.
//
// Returns:
//
//	An es.topHitsAggType object with "_source.excludes" set.
func (th topHitsAggType) SourceExcludes(fields ...string) topHitsAggType {
	if len(fields) == 0 {
		return th
	}
	topHits, ok := th["top_hits"].(Object)
	if !ok {
		return th
	}
	source, ok := topHits["_source"].(Object)
	if !ok {
		source = Object{}
	}
	excludes, ok := source["excludes"].(Array)
	if !ok {
		excludes = make(Array, 0, len(fields))
	}
	for i := 0; i < len(fields); i++ {
		excludes = append(excludes, fields[i])
	}
	source["excludes"] = excludes
	topHits["_source"] = source
	return th
}

// Highlight sets the highlight configuration for top hits.
//
// Example usage:
//
//	agg := es.TopHitsAgg().Highlight(es.Highlight().Field(es.HighlightField("title")))
//
// Parameters:
//   - highlight: An es.highlightType object defining highlight settings.
//
// Returns:
//
//	An es.topHitsAggType object with the "highlight" field set.
func (th topHitsAggType) Highlight(highlight highlightType) topHitsAggType {
	return th.putInTheField("highlight", highlight)
}

// Explain sets whether to include explanation of how the score was computed.
//
// Example usage:
//
//	agg := es.TopHitsAgg().Explain(true)
//
// Parameters:
//   - explain: Whether to include score explanations.
//
// Returns:
//
//	An es.topHitsAggType object with the "explain" field set.
func (th topHitsAggType) Explain(explain bool) topHitsAggType {
	return th.putInTheField("explain", explain)
}

// Version sets whether to include the document version in top hits.
//
// Example usage:
//
//	agg := es.TopHitsAgg().Version(true)
//
// Parameters:
//   - version: Whether to include document versions.
//
// Returns:
//
//	An es.topHitsAggType object with the "version" field set.
func (th topHitsAggType) Version(version bool) topHitsAggType {
	return th.putInTheField("version", version)
}

// SeqNoPrimaryTerm sets whether to include sequence number and primary term.
//
// Example usage:
//
//	agg := es.TopHitsAgg().SeqNoPrimaryTerm(true)
//
// Parameters:
//   - seqNoPrimaryTerm: Whether to include seq_no and primary_term.
//
// Returns:
//
//	An es.topHitsAggType object with the "seq_no_primary_term" field set.
func (th topHitsAggType) SeqNoPrimaryTerm(seqNoPrimaryTerm bool) topHitsAggType {
	return th.putInTheField("seq_no_primary_term", seqNoPrimaryTerm)
}

// TrackScores sets whether to track scores even when sorting by a field.
//
// Example usage:
//
//	agg := es.TopHitsAgg().TrackScores(true)
//
// Parameters:
//   - trackScores: Whether to track scores.
//
// Returns:
//
//	An es.topHitsAggType object with the "track_scores" field set.
func (th topHitsAggType) TrackScores(trackScores bool) topHitsAggType {
	return th.putInTheField("track_scores", trackScores)
}

// Meta sets a custom metadata field on the top_hits aggregation.
//
// Example usage:
//
//	agg := es.TopHitsAgg().Meta("source", "sales_data")
//
// Parameters:
//   - key: The metadata key.
//   - value: The metadata value.
//
// Returns:
//
//	An es.topHitsAggType object with the "meta" field set.
func (th topHitsAggType) Meta(key string, value any) topHitsAggType {
	meta, ok := th["meta"].(Object)
	if !ok {
		meta = Object{}
	}
	meta[key] = value
	th["meta"] = meta
	return th
}

func (th topHitsAggType) putInTheField(key string, value any) topHitsAggType {
	return genericPutInTheField(th, "top_hits", key, value)
}
