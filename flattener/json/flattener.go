package json

import (
	"encoding/json"
	"errors"
	"strconv"

	"jsonflatterner/flattener"
)

const (
	deliminatorDot          = "."
	deliminatorLeftBracket  = "["
	deliminatorRightBracket = "]"
)

type jsonflattener struct {
}

// NewStateMachine returns a new state machine
func NewJsonFlattener() flattener.Flattener {
	return &jsonflattener{}
}

func (f *jsonflattener) Flatten(data []byte) (map[string]interface{}, error) {
	var (
		in  map[string]interface{}
		out = make(map[string]interface{})
	)

	err := json.Unmarshal(data, &in)
	if err != nil {
		return nil, errors.New("Invalid input:" + err.Error())
	}

	for k, v := range in {
		switch t := v.(type) {
		case string, int, float64, bool, nil:
			out[k] = v
		case map[string]interface{}:
			nested, err := f.flattenMap(t, k)
			if err != nil {
				return nil, err
			}
			out = joinMaps(out, nested)
		case []interface{}:
			nested, err := f.flattenArray(t, k)
			if err != nil {
				return nil, err
			}
			out = joinMaps(out, nested)
		}
	}

	return out, nil
}

func (f *jsonflattener) flattenMap(in map[string]interface{}, parent string) (map[string]interface{}, error) {
	out := make(map[string]interface{})

	for k, v := range in {
		if len(parent) > 0 {
			k = parent + deliminatorDot + k
		}

		switch t := v.(type) {
		case string, int, float64, bool, nil:
			out[k] = v
		case map[string]interface{}:
			nested, err := f.flattenMap(t, k)
			if err != nil {
				return nil, err
			}
			for key, value := range out {
				nested[key] = value
			}
			out = joinMaps(out, nested)
		case []interface{}:
			nested, err := f.flattenArray(t, k)
			if err != nil {
				return nil, err
			}
			for key, value := range out {
				nested[key] = value
			}
			out = joinMaps(out, nested)
		}
	}

	return out, nil
}

func (f *jsonflattener) flattenArray(in []interface{}, parent string) (map[string]interface{}, error) {
	var (
		out = make(map[string]interface{})
		k   string
	)

	for i, v := range in {
		if len(parent) > 0 {
			k = parent + deliminatorLeftBracket + strconv.Itoa(i) + deliminatorRightBracket
		}

		switch t := v.(type) {
		case string, int, float64, bool, nil:
			out[k] = v
		case map[string]interface{}:
			nested, err := f.flattenMap(t, k)
			if err != nil {
				return nil, err
			}
			for key, value := range out {
				nested[key] = value
			}
			out = joinMaps(out, nested)
		case []interface{}:
			nested, err := f.flattenArray(t, k)
			if err != nil {
				return nil, err
			}
			for key, value := range out {
				nested[key] = value
			}
			out = joinMaps(out, nested)
		}
	}

	return out, nil
}

func joinMaps(m1 map[string]interface{}, m2 map[string]interface{}) map[string]interface{} {
	for k, v := range m1 {
		m2[k] = v
	}

	return m2
}
