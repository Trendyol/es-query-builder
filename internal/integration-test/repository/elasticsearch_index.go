package repository

import "github.com/Trendyol/es-query-builder/es"

type ElasticsearchIndex interface {
	GetMappings() string
	GetSettings() es.Object
}
