package testing

type SearchResponse struct {
	Hits struct {
		Hits []Hit `json:"hits"`
	} `json:"hits"`
}

type Hit struct {
	Source FooDocument `json:"_source"`
}

type FooDocument struct {
	Foo string `json:"foo"`
}
