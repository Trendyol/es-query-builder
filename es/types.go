package es

import "unsafe"

type Object map[string]any

type Array []any

type GenericObject[T any] map[string]T

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
