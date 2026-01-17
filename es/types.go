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

// correctType normalizes query values for Elasticsearch DSL. It returns (nil, false)
// for nil or typed-nil values, wraps es.BoolType in {"bool": ...} structure, and returns
// other types unchanged. Uses unsafe pointer inspection to detect typed-nil cases.
func correctType(b any) (any, bool) {
	if b == nil || (*[2]uintptr)(unsafe.Pointer(&b))[1] == 0 {
		return nil, false
	}
	if _, ok := b.(BoolType); ok {
		return Object{"bool": b}, true
	}
	return b, true
}
