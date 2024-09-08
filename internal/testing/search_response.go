package testing

type FooDocument struct {
	Id  string `json:"id"`
	Foo string `json:"foo"`
}

type Hit struct {
	Source FooDocument `json:"_source"`
}

type Hits struct {
	Hits []Hit `json:"hits"`
}

type SearchResponse struct {
	Hits Hits `json:"hits"`
}
