package model_repository

import "encoding/json"

////    COUNT     ////

type CountResponse struct {
	Shards *ShardsInfo `json:"_shards,omitempty"`
	Count  uint64      `json:"count"`
}

type ShardsInfo struct {
	Failures   []ShardFailure `json:"failures,omitempty"`
	Failed     uint           `json:"failed"`
	Skipped    uint           `json:"skipped,omitempty"`
	Successful uint           `json:"successful"`
	Total      uint           `json:"total"`
}

type ShardFailure struct {
	Reason  map[string]any `json:"reason,omitempty"`
	Index   string         `json:"_index,omitempty"`
	Node    string         `json:"_node,omitempty"`
	Status  string         `json:"status,omitempty"`
	Shard   uint           `json:"_shard,omitempty"`
	Primary bool           `json:"primary,omitempty"`
}

////    SEARCH    ////

type SearchResponse[T any] struct {
	Hits         *SearchHits[T]      `json:"hits,omitempty"`
	Shards       *ShardsInfo         `json:"_shards,omitempty"`
	Aggregations AggregateDictionary `json:"aggregations,omitempty"`
	ScrollId     string              `json:"_scroll_id,omitempty"`
	TookInMillis uint64              `json:"took,omitempty"`
	TimedOut     bool                `json:"timed_out,omitempty"`
}

type SearchHits[T any] struct {
	Hits     []SearchHit[T] `json:"hits"`
	Total    *Total         `json:"total,omitempty"`
	MaxScore *float64       `json:"max_score,omitempty"`
}

type SearchHit[T any] struct {
	Version *uint   `json:"_version,omitempty"`
	Id      string  `json:"_id"`
	Routing string  `json:"_routing"`
	Source  T       `json:"_source"`
	Score   float32 `json:"_score"`
	Found   bool    `json:"found"`
}

type Total struct {
	Relation string `json:"relation"`
	Value    int64  `json:"value"`
}

type AggregateDictionary map[string]json.RawMessage

////    EXISTS    ////

type ExistsDocument struct {
	Id      string `json:"id"`
	Routing string `json:"routing"`
}
