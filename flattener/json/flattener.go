package json

import (
	"encoding/json"
	"errors"
	"jsonflatterner/flattener"
	"strconv"
)

const (
	deliminator_dot = "."
	deliminator_left_bracket = "["
	deliminator_right_bracket = "]"
)

type jsonflattener struct {
}

// NewStateMachine returns a new state machine
func NewJsonFlattener() flattener.Flattener {
	return &jsonflattener{}
}

func (f *jsonflattener) Flatten(data []byte, delimiter string) ([]byte, error) {
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
			nested, err := f.flattenMap(t, k, deliminator_dot)
			if err != nil {
				return nil, err
			}
			out = joinMaps(out, nested)
		case []interface{}:
			nested, err := f.flattenArray(t, k, "")
			if err != nil {
				return nil, err
			}
			out = joinMaps(out, nested)
		default:
			// do nothing
		}
	}

	return json.Marshal(out)
}

func (f *jsonflattener) flattenMap(in map[string]interface{}, parent string, delimiter string) (map[string]interface{}, error) {
	out := make(map[string]interface{})

	for k, v := range in {
		if len(parent) > 0 {
			k = parent + delimiter + k
		}

		switch t := v.(type) {
		case string, int, float64, bool, nil:
			out[k] = v
		case map[string]interface{}:
			out, err := f.flattenMap(t, k, delimiter)
			if err != nil {
				return nil, err
			}
			for key, value := range out {
				out[key] = value
			}
		case []interface{}:
			out, err := f.flattenArray(t, k, delimiter)
			if err != nil {
				return nil, err
			}
			for key, value := range out {
				out[key] = value
			}
		}
	}

	return out, nil
}

func (f *jsonflattener) flattenArray(in []interface{}, parent string, delimiter string) (map[string]interface{}, error) {
	var (
		out  = make(map[string]interface{})
		k   string
	)
	for i, v := range in {
		if len(parent) > 0 {
			k = parent + deliminator_left_bracket + strconv.Itoa(i) + deliminator_right_bracket
		}

		switch t := v.(type) {
		case string, int, float64, bool, nil:
			out[k] = v
		case map[string]interface{}:
			out, err := f.flattenMap(t, k, delimiter)
			if err != nil {
				return nil, err
			}
			for key, value := range out {
				out[key] = value
			}
		case []interface{}:
			out, err := f.flattenArray(t, k, delimiter)
			if err != nil {
				return nil, err
			}
			for key, value := range out {
				out[key] = value
			}
		}
	}

	return out, nil
}

// utils
func joinMaps(m1 map[string]interface{}, m2 map[string]interface{}) map[string]interface{} {
	for k, v := range m1 {
		m2[k] = v
	}

	return m2
}

