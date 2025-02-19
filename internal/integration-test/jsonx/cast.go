package jsonx

import "encoding/json"

func Cast[Model any](raw []byte, err error) (*Model, error) {
	if err != nil {
		return nil, err
	}
	var result Model
	if err = json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func CastSlice[Model any](raw []byte, err error) ([]Model, error) {
	if err != nil {
		return nil, err
	}
	var result []Model
	if err = json.Unmarshal(raw, &result); err != nil {
		return nil, err
	}
	return result, nil
}
