package es

type innerHitsType Object

type fieldAndFormatType Object

type fieldCollapseType Object

func InnerHits() innerHitsType {
	return innerHitsType{}
}

func (ih innerHitsType) Collapse(fieldCollapse fieldCollapseType) innerHitsType {
	ih["collapse"] = fieldCollapse
	return ih
}

func (ih innerHitsType) DocvalueFields(fieldAndFormat ...fieldAndFormatType) innerHitsType {
	ih["docvalue_fields"] = fieldAndFormat
	return ih
}

func (ih innerHitsType) Explain(explain bool) innerHitsType {
	ih["explain"] = explain
	return ih
}

func (ih innerHitsType) Fields(fields ...string) innerHitsType {
	ih["fields"] = fields
	return ih
}

func (ih innerHitsType) From(from int) innerHitsType {
	ih["from"] = from
	return ih
}

func (ih innerHitsType) IgnoreUnmapped(ignoreUnmapped bool) innerHitsType {
	ih["ignore_unmapped"] = ignoreUnmapped
	return ih
}

func (ih innerHitsType) Name(name string) innerHitsType {
	ih["name"] = name
	return ih
}

func (ih innerHitsType) SeqNoPrimaryTerm(seqNoPrimaryTerm bool) innerHitsType {
	ih["seq_no_primary_term"] = seqNoPrimaryTerm
	return ih
}

func (ih innerHitsType) Size(size int) innerHitsType {
	ih["size"] = size
	return ih
}

func (ih innerHitsType) Sort(sorts ...sortType) innerHitsType {
	ih["sort"] = sorts
	return ih
}

func (ih innerHitsType) SourceFalse() innerHitsType {
	ih["_source"] = false
	return ih
}

func (ih innerHitsType) SourceIncludes(fields ...string) innerHitsType {
	if len(fields) == 0 {
		return ih
	}
	source, ok := ih["_source"].(Object)
	if !ok {
		source = Object{}
	}
	includes, ok := source["includes"].(Array)
	if !ok {
		includes = Array{}
	}
	for i := 0; i < len(fields); i++ {
		includes = append(includes, fields[i])
	}
	source["includes"] = includes
	ih["_source"] = source
	return ih
}

func (ih innerHitsType) SourceExcludes(fields ...string) innerHitsType {
	if len(fields) == 0 {
		return ih
	}
	source, ok := ih["_source"].(Object)
	if !ok {
		source = Object{}
	}
	excludes, exists := source["excludes"].(Array)
	if !exists {
		excludes = Array{}
	}
	for i := 0; i < len(fields); i++ {
		excludes = append(excludes, fields[i])
	}
	source["excludes"] = excludes
	ih["_source"] = source
	return ih
}

func (ih innerHitsType) StoredFields(fields ...string) innerHitsType {
	ih["stored_fields"] = fields
	return ih
}

func (ih innerHitsType) TrackScores(trackScores bool) innerHitsType {
	ih["track_scores"] = trackScores
	return ih
}

func (ih innerHitsType) Version(version bool) innerHitsType {
	ih["version"] = version
	return ih
}

func FieldCollapse(field string) fieldCollapseType {
	return fieldCollapseType{
		"field": field,
	}
}

func (fc fieldCollapseType) Collapse(fieldCollapse fieldCollapseType) fieldCollapseType {
	fc["collapse"] = fieldCollapse
	return fc
}

func (fc fieldCollapseType) InnerHits(innerHits ...innerHitsType) fieldCollapseType {
	fc["inner_hits"] = innerHits
	return fc
}

func (fc fieldCollapseType) MaxConcurrentGroupSearches(maxConcurrentGroupSearches int) fieldCollapseType {
	fc["max_concurrent_group_searches"] = maxConcurrentGroupSearches
	return fc
}

func FieldAndFormat(field string) fieldAndFormatType {
	return fieldAndFormatType{
		"field": field,
	}
}

func (fnf fieldAndFormatType) Format(format string) fieldAndFormatType {
	fnf["format"] = format
	return fnf
}

func (fnf fieldAndFormatType) IncludeUnmapped(includeUnmapped bool) fieldAndFormatType {
	fnf["include_unmapped"] = includeUnmapped
	return fnf
}
