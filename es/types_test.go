package es_test

import (
	"encoding/json"
	"github.com/GokselKUCUKSAHIN/es-query-builder/es"
	"reflect"
	"testing"
)

func Test_New_should_creates_a_new_Object(t *testing.T) {
	// Given When
	bodyA := es.New()
	bodyB := es.New()

	// Then
	assertNotNil(t, bodyA)
	assertNotNil(t, bodyB)
	assertEqual(t, bodyA, bodyB)
	assertNotEqualReference(t, bodyA, bodyB)
	assertMarshalWithoutError(t, bodyA)
	assertMarshalWithoutError(t, bodyB)
}

func Test_New_should_return_type_of_Object(t *testing.T) {
	// Given
	body := es.New()

	// When
	bodyType := reflect.TypeOf(body).String()

	// Then
	assertNotNil(t, body)
	assertEqual(t, "es.Object", bodyType)
	assertMarshalWithoutError(t, body)
}

func Test_New_should_creates_empty_Object(t *testing.T) {
	// Given
	body := es.New()

	// When Then
	assertNotNil(t, body)
	bodyJSON := assertMarshalWithoutError(t, body)

	assertEqual(t, bodyJSON, "{}")
}

// Testing functions

func assertMarshalWithoutError(t *testing.T, body any) string {
	marshal, err := json.Marshal(body)
	if err != nil {
		t.Errorf("marshal error: '%s'", err.Error())
	}
	return string(marshal)
}

func assertEqual(t *testing.T, expected, actual any, messages ...string) {
	message := getMessage(messages)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected: '%v' type: <%s>, Actual: '%v' type: <%s>. %s", expected, reflect.TypeOf(expected).String(), actual, reflect.TypeOf(actual).String(), message)
	}
}

func assertEqualReference(t *testing.T, expected, actual any, messages ...string) {
	message := getMessage(messages)
	expectedPointer, actualPointer := reflect.ValueOf(expected).Pointer(), reflect.ValueOf(actual).Pointer()
	if !(expectedPointer == actualPointer) {
		t.Errorf("Expected: '%v', Actual: '%v'. %s", expectedPointer, actualPointer, message)
	}
}

func assertNotEqualReference(t *testing.T, expected, actual any, messages ...string) {
	message := getMessage(messages)
	expectedPointer, actualPointer := reflect.ValueOf(expected).Pointer(), reflect.ValueOf(actual).Pointer()
	if expectedPointer == actualPointer {
		t.Errorf("Expected and Actual have the same reference: '%v'. %s", expectedPointer, message)
	}
}

func assertTrue(t *testing.T, condition bool, messages ...string) {
	message := getMessage(messages)
	if !condition {
		t.Errorf("Expected: true, Actual: false. %s", message)
	}
}

func assertFalse(t *testing.T, condition bool, messages ...string) {
	message := getMessage(messages)
	if condition {
		t.Errorf("Expected: false, Actual: true. %s", message)
	}
}

func assertNil(t *testing.T, value any, messages ...string) {
	message := getMessage(messages)
	if value != nil {
		t.Errorf("Expected: nil, Actual: '%v'. %s", value, message)
	}
}

func assertNotNil(t *testing.T, value any, messages ...string) {
	message := getMessage(messages)
	if value == nil {
		t.Errorf("Expected: not nil, Actual: nil. %s", message)
	}
}

func getMessage(messages []string) string {
	if len(messages) > 0 {
		return messages[0]
	}
	return ""
}
