package tests

import (
	"context"
	"fmt"

	"integration-tests/repository"

	"github.com/bayraktugrul/go-await"
)

func OmitError[T any](value T, err error) T {
	if err != nil {
		fmt.Printf("error omitted for value: %v. error: %s", value, err.Error())
	}
	return value
}

func MapKeys[K comparable, V any](dict map[K]V) []K {
	result := make([]K, 0, len(dict))
	for key := range dict {
		result = append(result, key)
	}
	return result
}

func MapValues[K comparable, V any](dict map[K]V) []V {
	result := make([]V, 0, len(dict))
	for _, value := range dict {
		result = append(result, value)
	}
	return result
}

func WaitExists[ID comparable, T any](ctx context.Context, repository repository.BaseGenericRepository[ID, T], id ID) {
	await.New().Await(func() bool { return OmitError(repository.Exists(ctx, id)) })
}
