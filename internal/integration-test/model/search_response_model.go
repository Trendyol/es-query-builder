package model

type GenericHit[T any] struct {
	Source T `json:"_source"`
}

type GenericHits[T any] struct {
	Hits []GenericHit[T] `json:"hits"`
}

type GenericSearchResponse[T any] struct {
	Hits GenericHits[T] `json:"hits"`
}
