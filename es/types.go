package es

import (
	Mode "github.com/GokselKUCUKSAHIN/es-query-builder/es/enums/sort/mode"
	Order "github.com/GokselKUCUKSAHIN/es-query-builder/es/enums/sort/order"
)

type Object map[string]any

type Array []any

type queryType Object

type boolType Object

type filterType Array

type mustType Array

type mustNotType Array

type shouldType Array

type termType Object

type termsType Object

type existsType Object

type rangeType Object

type sortType Object

type sourceType Object

type includesType Array

type excludesType Array

func New() Object {
	return Object{}
}

func (o Object) Query() queryType {
	q := queryType{}
	o["query"] = q
	return q
}

func (q queryType) Bool() boolType {
	b := boolType{}
	q["bool"] = b
	return b
}

func Term(key string, value any) termType {
	return termType{
		"term": Object{
			key: value,
		},
	}
}

func Terms(key string, values ...any) termsType {
	return termsType{
		"terms": Object{
			key: values,
		},
	}
}

func TermsArray(key string, values Array) termsType {
	return termsType{
		"terms": Object{
			key: values,
		},
	}
}

func Exists(key string) existsType {
	return existsType{
		"exists": Object{
			"field": key,
		},
	}
}

func Range(key string, lte any, gte any) rangeType {
	o := Object{}
	if lte != nil {
		o["lte"] = lte
	}
	if gte != nil {
		o["gte"] = gte
	}
	return rangeType{
		key: o,
	}
}

func (b boolType) Filter(items ...any) boolType {
	filter, exists := b["filter"]
	if !exists {
		filter = filterType{}
	}
	filter = append(filter.(filterType), items...)
	b["filter"] = filter
	return b
}

func (b boolType) Must(items ...any) boolType {
	must, exists := b["must"]
	if !exists {
		must = mustType{}
	}
	must = append(must.(mustType), items...)
	b["must"] = must
	return b
}

func (b boolType) MustNot(items ...any) boolType {
	mustNot, exists := b["must_not"]
	if !exists {
		mustNot = mustNotType{}
	}
	mustNot = append(mustNot.(mustNotType), items...)
	b["must_not"] = mustNot
	return b
}

func (b boolType) Should(items ...any) boolType {
	should, exists := b["should"]
	if !exists {
		should = shouldType{}
	}
	should = append(should.(shouldType), items...)
	b["should"] = should
	return b
}

func (o Object) SetTrackTotalHits(value bool) Object {
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

func SortWithMode(field string, order Order.Order, mode Mode.Mode) sortType {
	o := Object{}
	if order != Order.Default {
		o["order"] = order
	}
	if mode != Mode.Default {
		o["mode"] = mode
	}
	return sortType{
		field: o,
	}
}

func Sort(field string, order Order.Order) sortType {
	return SortWithMode(field, order, Mode.Default)
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
	for _, field := range fields {
		includes = append(includes.(includesType), field)
	}
	s["includes"] = includes
	return s
}

func (s sourceType) Excludes(fields ...string) sourceType {
	excludes, exists := s["excludes"]
	if !exists {
		excludes = excludesType{}
	}
	for _, field := range fields {
		excludes = append(excludes.(excludesType), field)
	}
	s["excludes"] = excludes
	return s
}
