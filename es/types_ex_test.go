package es_test

import (
	"encoding/json"
	"github.com/GokselKUCUKSAHIN/es-query-builder/es"
	"github.com/GokselKUCUKSAHIN/es-query-builder/es/test/assert"
	"testing"
)

func TestX(t *testing.T) {
	id := 293
	q1 := map[string]interface{}{
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"must": []interface{}{
					map[string]interface{}{
						"bool": map[string]interface{}{
							"should": []interface{}{
								map[string]interface{}{
									"term": map[string]interface{}{
										"doc.id": id,
									},
								},
								map[string]interface{}{
									"term": map[string]interface{}{
										"file.fileId": id,
									},
								},
							},
						},
					},
				},
				"filter": []interface{}{
					map[string]interface{}{
						"terms": map[string]interface{}{
							"type": []string{
								"DOC", "FILE",
							},
						},
					},
				},
			},
		},
	}
	Q1, _ := json.Marshal(q1)

	q2 := es.NewQuery(
		es.Bool().
			Must(
				es.Bool().
					Should(
						es.Term("doc.id", id),
						es.Term("file.fileId", id),
					),
			).
			Filter(
				es.Terms("type", "DOC", "FILE"),
			),
	)
	Q2, _ := json.Marshal(q2)

	assert.Equal(t, string(Q1), string(Q2))
}
