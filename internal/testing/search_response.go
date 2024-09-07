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
	Id  string `json:"id"`
	Foo string `json:"foo"`
}
