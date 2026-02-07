package es

type reverseNestedAggType Object

// ReverseNestedAgg creates a reverse nested aggregation.
//
// A reverse nested aggregation is a special single bucket aggregation that enables
// aggregating on parent documents from nested documents. It can only be used inside
// a nested aggregation. It effectively "joins back" to the root/parent document level.
//
// Example usage:
//
//	agg := es.NestedAgg("comments").
//		Aggs(
//			es.Agg("top_usernames", es.TermsAgg("comments.username")),
//			es.Agg("comment_to_issue", es.ReverseNestedAgg().
//				Aggs(es.Agg("top_tags_per_comment", es.TermsAgg("tags")))),
//		)
//
// Returns:
//
//	An es.reverseNestedAggType object representing the reverse nested aggregation.
func ReverseNestedAgg() reverseNestedAggType {
	return reverseNestedAggType{
		"reverse_nested": Object{},
	}
}

// Path sets the "path" parameter in the reverse nested aggregation.
//
// This method defines which nested object field should be joined back to. The default
// is empty, which means it joins back to the root/parent document level. If a path is
// specified, it joins back to the specified nested object level.
//
// Example usage:
//
//	agg := es.ReverseNestedAgg().Path("comments")
//
// Parameters:
//   - path: A string representing the nested object path to join back to.
//
// Returns:
//
//	The updated es.reverseNestedAggType object with the "path" parameter set.
func (rn reverseNestedAggType) Path(path string) reverseNestedAggType {
	return rn.putInTheField("path", path)
}

// Aggs adds sub-aggregations to the reverse nested aggregation.
//
// This method allows performing additional aggregations on the parent documents
// after joining back from the nested context.
//
// Example usage:
//
//	agg := es.ReverseNestedAgg().
//		Aggs(es.Agg("top_tags", es.TermsAgg("tags")))
//
// Parameters:
//   - aggs: A variadic list of sub-aggregations.
//
// Returns:
//
//	An es.reverseNestedAggType object with the specified sub-aggregations added.
func (rn reverseNestedAggType) Aggs(aggs ...aggsType) reverseNestedAggType {
	return genericPutAggsInRoot(rn, aggs)
}

// Meta adds metadata to the reverse nested aggregation.
//
// Example usage:
//
//	agg := es.ReverseNestedAgg().Meta("description", "Back to parent")
//
// Parameters:
//   - key: Metadata key.
//   - value: Metadata value.
//
// Returns:
//
//	A modified es.reverseNestedAggType with the meta field set.
func (rn reverseNestedAggType) Meta(key string, value any) reverseNestedAggType {
	meta, ok := rn["meta"].(Object)
	if !ok {
		meta = Object{}
	}
	meta[key] = value
	rn["meta"] = meta
	return rn
}

func (rn reverseNestedAggType) putInTheField(key string, value any) reverseNestedAggType {
	return genericPutInTheField(rn, "reverse_nested", key, value)
}
