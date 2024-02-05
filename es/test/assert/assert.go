package assert

import (
	"encoding/json"
	"reflect"
	"testing"
)

func MarshalWithoutError(t *testing.T, body any) string {
	marshal, err := json.Marshal(body)
	if err != nil {
		t.Errorf("marshal error: '%s'", err.Error())
	}
	return string(marshal)
}

func Equal(t *testing.T, expected, actual any, messages ...string) {
	message := getMessage(messages)
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected: '%v' type: <%s>, Actual: '%v' type: <%s>. %s", expected, reflect.TypeOf(expected).String(), actual, reflect.TypeOf(actual).String(), message)
	}
}

func EqualReference(t *testing.T, expected, actual any, messages ...string) {
	message := getMessage(messages)
	expectedPointer, actualPointer := reflect.ValueOf(expected).Pointer(), reflect.ValueOf(actual).Pointer()
	if !(expectedPointer == actualPointer) {
		t.Errorf("Expected: '%v', Actual: '%v'. %s", expectedPointer, actualPointer, message)
	}
}

func NotEqualReference(t *testing.T, expected, actual any, messages ...string) {
	message := getMessage(messages)
	expectedPointer, actualPointer := reflect.ValueOf(expected).Pointer(), reflect.ValueOf(actual).Pointer()
	if expectedPointer == actualPointer {
		t.Errorf("Expected and Actual have the same reference: '%v'. %s", expectedPointer, message)
	}
}

func True(t *testing.T, condition bool, messages ...string) {
	message := getMessage(messages)
	if !condition {
		t.Errorf("Expected: true, Actual: false. %s", message)
	}
}

func False(t *testing.T, condition bool, messages ...string) {
	message := getMessage(messages)
	if condition {
		t.Errorf("Expected: false, Actual: true. %s", message)
	}
}

func Nil(t *testing.T, value any, messages ...string) {
	message := getMessage(messages)
	if value != nil {
		t.Errorf("Expected: nil, Actual: '%v'. %s", value, message)
	}
}

func NotNil(t *testing.T, value any, messages ...string) {
	message := getMessage(messages)
	if value == nil {
		t.Errorf("Expected: not nil, Actual: nil. %s", message)
	}
}

func IsType(t *testing.T, expected, actual any, messages ...string) {
	message := getMessage(messages)
	expectedType := reflect.TypeOf(expected)
	actualValue := reflect.TypeOf(actual)
	if expectedType != actualValue {
		t.Errorf("Expected type '%v', but got type '%v'. %s", expectedType, actualValue, message)
	}
}

func getMessage(messages []string) string {
	if len(messages) > 0 {
		return messages[0]
	}
	return ""
}
