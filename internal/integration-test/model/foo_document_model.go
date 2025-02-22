package model

import "github.com/Trendyol/es-query-builder/es"

type FooDocument struct {
	ID  string `json:"id"`
	Foo string `json:"foo"`
}

func (foo *FooDocument) Copy() FooDocument {
	return FooDocument{
		ID:  foo.ID,
		Foo: foo.Foo,
	}
}

func (foo *FooDocument) GetMappings() es.Object {
	return es.Object{
		"properties": es.Object{
			"foo": es.Object{
				"type": "keyword",
			},
			"meta": es.Object{
				"properties": es.Object{
					"id": es.Object{
						"type": "keyword",
					},
				},
			},
		},
	}
}

func (foo *FooDocument) GetSettings() es.Object {
	return es.Object{
		"index": es.Object{
			"refresh_interval":   "1s",
			"number_of_shards":   1,
			"number_of_replicas": 1,
			"max_result_window":  10_000,
			"max_terms_count":    1024,
		},
	}
}
