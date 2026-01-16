package es

import "unsafe"

type Object map[string]any

type Array []any

type GenericObject[T any] map[string]T

type signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type integer interface {
	signed | unsigned
}

type float interface {
	~float32 | ~float64
}

type complexNumber interface {
	~complex64 | ~complex128
}

type number interface {
	integer | float | complexNumber
}

type primitive interface {
	number | ~string | ~bool | ~rune
}

func correctType(b any) (any, bool) {
	if b == nil || (*[2]uintptr)(unsafe.Pointer(&b))[1] == 0 {
		return nil, false
	}
	if _, ok := b.(BoolType); ok {
		return Object{"bool": b}, true
	}
	return b, true
}

func putInTheNestedField(o Object, topKey, fieldKey string, value any) {
	if top, ok := o[topKey].(Object); ok {
		for _, fieldObj := range top {
			if fieldObject, foOk := fieldObj.(Object); foOk {
				fieldObject[fieldKey] = value
				break
			}
		}
	}
}

func deleteFromNestedField(o Object, topKey, fieldKey string) {
	if top, ok := o[topKey].(Object); ok {
		for _, fieldObj := range top {
			if fieldObject, foOk := fieldObj.(Object); foOk {
				delete(fieldObject, fieldKey)
				break
			}
		}
	}
}

func putInTheField(o Object, topKey, fieldKey string, value any) {
	if top, ok := o[topKey].(Object); ok {
		top[fieldKey] = value
	}
}

func reduceAggs(aggs ...aggsType) Object {
	aggregates := Object{}
	for _, agg := range aggs {
		for key, value := range agg {
			aggregates[key] = value
		}
	}
	return aggregates
}
