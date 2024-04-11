package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/GokselKUCUKSAHIN/es-query-builder/es"
	"github.com/GokselKUCUKSAHIN/es-query-builder/es/enums/sort/order"
)

func mockGetDocumentsEs(ctx context.Context, query string) (string, error) {
	context.WithValue(ctx, "query", query)
	return fmt.Sprintf("query result for '%v'", query), nil
}

func main() {
	ctx := context.Background()

	id := 42
	queryString, err := json.Marshal(buildQuery(id))
	if err != nil {
		log.Fatal(err.Error())
	}

	documentsEs, err := mockGetDocumentsEs(ctx, string(queryString))
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("query result: %s\n", documentsEs)
}

func buildQuery(id int) es.Object {
	query := es.NewQuery(
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
	query.Size(45)
	query.Sort(es.Sort("name", order.Asc))
	query.Source().
		Includes("id", "type", "indexedAt", "chapters")
	return query
}
