package model

import "github.com/Trendyol/es-query-builder/es"

type GeoPoint struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type FooDocument struct {
	ID       string    `json:"id"`
	Foo      string    `json:"foo"`
	Location *GeoPoint `json:"location,omitempty"`
}

func (foo *FooDocument) Copy() FooDocument {
	copied := FooDocument{
		ID:  foo.ID,
		Foo: foo.Foo,
	}
	if foo.Location != nil {
		location := *foo.Location
		copied.Location = &location
	}
	return copied
}

func (foo *FooDocument) GetMappings() es.Object {
	return es.Object{
		"properties": es.Object{
			"foo": es.Object{
				"type": "keyword",
			},
			"location": es.Object{
				"type": "geo_point",
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
