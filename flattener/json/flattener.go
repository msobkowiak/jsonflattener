package json

import (
	"encoding/json"
	"errors"
	"fmt"
)

func Flatten(data []byte, delimiter string) ([]byte, error) {
	var (
		input  map[string]interface{}
		output = make(map[string]interface{})
	)

	err := json.Unmarshal(data, &input)
	if err != nil {
		return nil, errors.New("Invalid input:" + err.Error())
	}

	for k, v := range input {
		switch v.(type) {
		case string, int, float64, bool, nil:
			fmt.Println(v, k)
			output[k] = v
		default:
			// do nothing
		}
	}

	fmt.Println(output)

	return json.Marshal(output)
}
