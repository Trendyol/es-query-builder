package es_test

import (
	"github.com/GokselKUCUKSAHIN/es-query-builder/es"
	"github.com/GokselKUCUKSAHIN/es-query-builder/es/test/assert"
	"reflect"
	"testing"
)

func Test_New_should_creates_a_new_Object(t *testing.T) {
	// Given When
	bodyA := es.New()
	bodyB := es.New()

	// Then

	assert.NotNil(t, bodyA)
	assert.NotNil(t, bodyB)
	assert.Equal(t, bodyA, bodyB)
	assert.NotEqualReference(t, bodyA, bodyB)
	assert.MarshalWithoutError(t, bodyA)
	assert.MarshalWithoutError(t, bodyB)
}

func Test_New_should_return_type_of_Object(t *testing.T) {
	// Given
	body := es.New()

	// When
	bodyType := reflect.TypeOf(body).String()

	// Then
	assert.NotNil(t, body)
	assert.Equal(t, "es.Object", bodyType)
	assert.MarshalWithoutError(t, body)
}

func Test_New_should_creates_empty_Object(t *testing.T) {
	// Given
	body := es.New()

	// When Then
	assert.NotNil(t, body)
	bodyJSON := assert.MarshalWithoutError(t, body)

	assert.Equal(t, bodyJSON, "{}")
}
