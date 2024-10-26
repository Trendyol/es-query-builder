package es

import "unsafe"

type Object map[string]any

type Array []any

func unsafeIsNil(x any) bool {
	return (*[2]uintptr)(unsafe.Pointer(&x))[1] == 0
}

func correctType(b any) (any, bool) {
	if b == nil || unsafeIsNil(b) {
		return Object{}, false
	}
	switch b.(type) {
	case BoolType:
		return Object{"bool": b}, true
	case rangeType:
		return Object{"range": b}, true
	}
	return b, true
}
