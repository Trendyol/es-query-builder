package es

import (
	"unsafe"

	Operator "github.com/GokselKUCUKSAHIN/es-query-builder/es/enums/match/operator"
	ScoreMode "github.com/GokselKUCUKSAHIN/es-query-builder/es/enums/nested/score-mode"
	Mode "github.com/GokselKUCUKSAHIN/es-query-builder/es/enums/sort/mode"
	Order "github.com/GokselKUCUKSAHIN/es-query-builder/es/enums/sort/order"
)

type Object map[string]any

type Array []any

type boolType Object

type filterType Array

type mustType Array

type mustNotType Array

type shouldType Array

type matchType Object

type matchAllType Object

type matchNoneType Object

type termType Object

type termsType Object

type existsType Object

type rangeType Object

type sortType Object

type sourceType Object

type includesType Array

type excludesType Array

type nestedType Object

type aggsType Object

type aggTermType Object

func unsafeIsNil(x any) bool {
	return (*[2]uintptr)(unsafe.Pointer(&x))[1] == 0
}

func correctType(b any) (any, bool) {
	if b == nil || unsafeIsNil(b) {
		return Object{}, false
	}
	switch b.(type) {
	case boolType:
		return Object{"bool": b}, true
	case rangeType:
		return Object{"range": b}, true
	}
	return b, true
}

func NewQuery(c any) Object {
	if field, ok := correctType(c); ok {
		return Object{
			"query": field,
		}
	}
	return Object{
		"query": Object{},
	}
}

func (o Object) putInQuery(key string, value any) Object {
	if query, exists := o["query"]; exists {
		query.(Object)[key] = value
	}
	return o
}

func Bool() boolType {
	return boolType{}
}

func (b boolType) MinimumShouldMatch(minimumShouldMatch int) boolType {
	b["minimum_should_match"] = minimumShouldMatch
	return b
}

func (b boolType) Boost(boost float64) boolType {
	b["boost"] = boost
	return b
}

func (b boolType) Filter(items ...any) boolType {
	filter, exists := b["filter"]
	if !exists {
		filter = filterType{}
	}
	for i := 0; i < len(items); i++ {
		if field, ok := correctType(items[i]); ok {
			filter = append(filter.(filterType), field)
		}
	}
	b["filter"] = filter
	return b
}

func (b boolType) Must(items ...any) boolType {
	must, exists := b["must"]
	if !exists {
		must = mustType{}
	}
	for i := 0; i < len(items); i++ {
		if field, ok := correctType(items[i]); ok {
			must = append(must.(mustType), field)
		}
	}
	b["must"] = must
	return b
}

func (b boolType) MustNot(items ...any) boolType {
	mustNot, exists := b["must_not"]
	if !exists {
		mustNot = mustNotType{}
	}
	for i := 0; i < len(items); i++ {
		if field, ok := correctType(items[i]); ok {
			mustNot = append(mustNot.(mustNotType), field)
		}
	}
	b["must_not"] = mustNot
	return b
}

func (b boolType) Should(items ...any) boolType {
	should, exists := b["should"]
	if !exists {
		should = shouldType{}
	}
	for i := 0; i < len(items); i++ {
		if field, ok := correctType(items[i]); ok {
			should = append(should.(shouldType), field)
		}
	}
	b["should"] = should
	return b
}

func (o Object) TrackTotalHits(value bool) Object {
	o["track_total_hits"] = value
	return o
}

func (o Object) Size(size int) Object {
	o["size"] = size
	return o
}

func (o Object) From(from int) Object {
	o["from"] = from
	return o
}

func Sort(field string) sortType {
	return sortType{
		field: Object{},
	}
}

func (s sortType) putInTheField(key string, value any) sortType {
	for field := range s {
		if fieldObject, ok := s[field].(Object); ok {
			fieldObject[key] = value
		}
	}
	return s
}

func (s sortType) Order(order Order.Order) sortType {
	return s.putInTheField("order", order)
}

func (s sortType) Mode(mode Mode.Mode) sortType {
	return s.putInTheField("mode", mode)
}

func (o Object) Sort(sorts ...sortType) Object {
	o["sort"] = sorts
	return o
}

func (o Object) Source() sourceType {
	s := sourceType{}
	o["_source"] = s
	return s
}

func (o Object) SourceFalse() Object {
	o["_source"] = false
	return o
}

func (s sourceType) Includes(fields ...string) sourceType {
	includes, exists := s["includes"]
	if !exists {
		includes = includesType{}
	}
	for i := 0; i < len(fields); i++ {
		includes = append(includes.(includesType), fields[i])
	}
	s["includes"] = includes
	return s
}

func (s sourceType) Excludes(fields ...string) sourceType {
	excludes, exists := s["excludes"]
	if !exists {
		excludes = excludesType{}
	}
	for i := 0; i < len(fields); i++ {
		excludes = append(excludes.(excludesType), fields[i])
	}
	s["excludes"] = excludes
	return s
}

func Term[T any](key string, value T) termType {
	return termType{
		"term": Object{
			key: value,
		},
	}
}

func TermFunc[T any](key string, value T, f func(key string, value T) bool) termType {
	if !f(key, value) {
		return nil
	}
	return Term(key, value)
}

func Terms(key string, values ...any) termsType {
	return termsType{
		"terms": Object{
			key: values,
		},
	}
}

func TermsArray[T any](key string, values []T) termsType {
	return termsType{
		"terms": Object{
			key: values,
		},
	}
}

func TermsArrayFunc[T any](key string, values []T, f func(key string, values []T) bool) termsType {
	if !f(key, values) {
		return nil
	}
	return TermsArray(key, values)
}

func Exists(key string) existsType {
	return existsType{
		"exists": Object{
			"field": key,
		},
	}
}

func ExistsFunc(key string, f func(key string) bool) existsType {
	if !f(key) {
		return nil
	}
	return Exists(key)
}

func Match[T any](key string, query T) matchType {
	return matchType{
		"match": Object{
			key: Object{
				"query": query,
			},
		},
	}
}

func (m matchType) putInTheField(key string, value any) matchType {
	if match, exists := m["match"]; exists {
		if matchObject, moOk := match.(Object); moOk {
			for field := range matchObject {
				if fieldObject, foOk := matchObject[field].(Object); foOk {
					fieldObject[key] = value
				}
			}
		}
	}
	return m
}

func (m matchType) Operator(operator Operator.Operator) matchType {
	return m.putInTheField("operator", operator)
}

func (m matchType) Boost(boost float64) matchType {
	return m.putInTheField("boost", boost)
}

func MatchNone[T any](key string, query T) matchNoneType {
	return matchNoneType{
		"match_none": Object{
			key: Object{
				"query": query,
			},
		},
	}
}

func (m matchNoneType) putInTheField(key string, value any) matchNoneType {
	if matchNone, exists := m["match_none"]; exists {
		if matchNoneObject, mnoOk := matchNone.(Object); mnoOk {
			for field := range matchNoneObject {
				if fieldObject, foOk := matchNoneObject[field].(Object); foOk {
					fieldObject[key] = value
				}
			}
		}
	}
	return m
}

func (m matchNoneType) Operator(operator Operator.Operator) matchNoneType {
	return m.putInTheField("operator", operator)
}

func (m matchNoneType) Boost(boost float64) matchNoneType {
	return m.putInTheField("boost", boost)
}

func MatchAll() matchAllType {
	return matchAllType{
		"match_all": Object{},
	}
}

func (m matchAllType) putInTheField(key string, value any) matchAllType {
	if matchAll, exists := m["match_all"]; exists {
		matchAll.(Object)[key] = value
	}
	return m
}

func (m matchAllType) Boost(boost float64) matchAllType {
	return m.putInTheField("boost", boost)
}

func Range(key string) rangeType {
	return rangeType{
		key: Object{},
	}
}

func (o Object) Range(key string) rangeType {
	r := Range(key)
	o.putInQuery("range", r)
	return r
}

func (r rangeType) LesserThan(lt any) rangeType {
	for key := range r {
		if rangeObject, ok := r[key].(Object); ok {
			rangeObject["lt"] = lt
			delete(rangeObject, "lte")
		}
	}
	return r
}

func (r rangeType) LesserThanOrEqual(lte any) rangeType {
	for key := range r {
		if rangeObject, ok := r[key].(Object); ok {
			rangeObject["lte"] = lte
			delete(rangeObject, "lt")
		}
	}
	return r
}

func (r rangeType) GreaterThan(gt any) rangeType {
	for key := range r {
		if rangeObject, ok := r[key].(Object); ok {
			rangeObject["gt"] = gt
			delete(rangeObject, "gte")
		}
	}
	return r
}

func (r rangeType) GreaterThanOrEqual(gte any) rangeType {
	for key := range r {
		if rangeObject, ok := r[key].(Object); ok {
			rangeObject["gte"] = gte
			delete(rangeObject, "gt")
		}
	}
	return r
}

func (r rangeType) Format(format string) rangeType {
	for key := range r {
		if rangeObject, ok := r[key].(Object); ok {
			rangeObject["format"] = format
		}
	}
	return r
}

func (r rangeType) Boost(boost float64) rangeType {
	for key := range r {
		if rangeObject, ok := r[key].(Object); ok {
			rangeObject["boost"] = boost
		}
	}
	return r
}

func Nested[T any](path string, nestedQuery T) nestedType {
	o := NewQuery(nestedQuery)
	o["path"] = path
	return nestedType{
		"nested": o,
	}
}

func (n nestedType) putInNested(key string, value any) nestedType {
	if nested, exists := n["nested"]; exists {
		nested.(Object)[key] = value
	}
	return n
}

func (n nestedType) InnerHits(innerHits Object) nestedType {
	return n.putInNested("inner_hits", innerHits)
}

func (n nestedType) ScoreMode(scoreMode ScoreMode.ScoreMode) nestedType {
	return n.putInNested("score_mode", scoreMode)
}

func AggTerm(field string) aggTermType {
	return aggTermType{
		"field": field,
	}
}

func (aggTerm aggTermType) Missing(missing string) aggTermType {
	aggTerm["missing"] = missing
	return aggTerm
}

func AggTerms() aggsType {
	return aggsType{
		"terms": Object{},
	}
}

func AggMultiTerms() aggsType {
	return aggsType{
		"multi_terms": Object{},
	}
}

func AggCustom(agg Object) aggsType {
	return aggsType(agg)
}

func (agg aggsType) putInTheField(key string, value any) aggsType {
	for field := range agg {
		if fieldObject, ok := agg[field].(Object); ok {
			fieldObject[key] = value
		}
	}
	return agg
}

func (agg aggsType) Aggs(name string, nestedAgg aggsType) aggsType {
	agg["aggs"] = Object{
		name: nestedAgg,
	}
	return agg
}

func (agg aggsType) Field(field string) aggsType {
	return agg.putInTheField("field", field)
}

func (agg aggsType) Size(size int) aggsType {
	return agg.putInTheField("size", size)
}

func (agg aggsType) Order(field string, order Order.Order) aggsType {
	return agg.putInTheField("order",
		Object{
			field: order,
		},
	)
}

func (agg aggsType) Include(include string) aggsType {
	return agg.putInTheField("include", include)
}

func (agg aggsType) Exclude(exclude string) aggsType {
	return agg.putInTheField("exclude", exclude)
}

func (agg aggsType) Terms(terms ...aggTermType) aggsType {
	return agg.putInTheField("terms", terms)
}

func (o Object) Aggs(name string, agg aggsType) Object {
	aggs, exists := o["aggs"]
	if !exists {
		aggs = Object{}
	}
	aggs.(Object)[name] = agg
	o["aggs"] = aggs
	return o
}
