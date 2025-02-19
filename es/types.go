package es

import "unsafe"

type Object map[string]any

type Array []any

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

func unsafeIsNil(x any) bool {
	return (*[2]uintptr)(unsafe.Pointer(&x))[1] == 0
}

func correctType(b any) (any, bool) {
	if b == nil || unsafeIsNil(b) {
		return Object{}, false
	}
	if _, ok := b.(BoolType); ok {
		return Object{"bool": b}, true
	}
	return b, true
}
