package repository

import (
	"bytes"
	"encoding/json"
	"io"

	"github.com/Trendyol/es-query-builder/es"
)

type ElasticsearchIndex interface {
	GetMappings() es.Object
	GetSettings() es.Object
}

func CreateIndexBody(index ElasticsearchIndex) (io.Reader, error) {
	esIndex := es.Object{
		"settings": index.GetSettings(),
		"mappings": index.GetMappings(),
	}

	data, err := json.Marshal(esIndex)
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(data), nil
}
