package repository

import (
	"context"
	"encoding/json"
	"github.com/Trendyol/es-query-builder/es"
	"integration-tests/model_repository"
	"io"
)

type BaseGenericRepository[ID comparable, T any] interface {
	Search(ctx context.Context, query es.Object) (*model_repository.SearchResponse, error)
	// GetSearchHits(ctx context.Context, query es.Object) (map[ID]T, error)
	Exists(ctx context.Context, documentID ID) (bool, error)
	Insert(ctx context.Context, document T) error
	BulkInsert(ctx context.Context, documents ...T) error
	Delete(ctx context.Context, documentID ID) error
	DeleteByQuery(ctx context.Context, query es.Object) error
	BulkDelete(ctx context.Context, documentIDs ...ID) error
}

func mapToReader[T any](object T) io.Reader {
	pipeReader, pipeWriter := io.Pipe()
	go func() {
		defer pipeWriter.Close()
		if err := json.NewEncoder(pipeWriter).Encode(object); err != nil {
			pipeWriter.CloseWithError(err)
		}
	}()
	return pipeReader
}

//GetCount(ctx context.Context, query es.Object) (*model_repository.CountResponse, error)
//GetByID(ctx context.Context, documentID ID, routingID string) (*T, error)
