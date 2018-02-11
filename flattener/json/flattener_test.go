package json_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"jsonflatterner/flattener/json"
)

func TestFlatten(t *testing.T) {
	testCases := []struct {
		in  []byte
		out map[string]interface{}
	}{
		{
			in: []byte(
				`{
					"first_name": "Joe",
					"last_name": "Doe",
					"age": 25,
					"hobbies": ["travel", "sport", "books"],
					"address": {
						"street": "123 Main St.",
						"city": "Berlin",
						"zip_code": "11129"
					}
				}`,
			),
			out: map[string]interface{}{
				"age":              float64(25),
				"first_name":       "Joe",
				"last_name":        "Doe",
				"hobbies[0]":       "travel",
				"hobbies[1]":       "sport",
				"hobbies[2]":       "books",
				"address.street":   "123 Main St.",
				"address.city":     "Berlin",
				"address.zip_code": "11129",
			},
		},
		{
			in: []byte(
				`{
					"format": "long",
					"type": "objects",
					"offset": 0,
					"limit": 10,
					"objecttypes": ["objekte"],
					"language": "de - DE",
					"count": 2,
					"took": 165,
					"objects": [{
						"_acl": [{
							"_id": 1335359,
							"date_created": "2016 - 08 - 22 T19: 26: 10 + 02: 00",
							"who": [{
									"first_name": "Joe",
									"last_name": "Doe"
								},
								{
									"first_name": "Anna",
									"last_name": "Smith"
								}
							]
						}]
					}]
				}`,
			),
			out: map[string]interface{}{
				"format":         "long",
				"type":           "objects",
				"offset":         float64(0),
				"limit":          float64(10),
				"objecttypes[0]": "objekte",
				"language":       "de - DE",
				"count":          float64(2),
				"took":           float64(165),
				"objects[0]._acl[0]._id":               float64(1335359),
				"objects[0]._acl[0].date_created":      "2016 - 08 - 22 T19: 26: 10 + 02: 00",
				"objects[0]._acl[0].who[0].first_name": "Joe",
				"objects[0]._acl[0].who[0].last_name":  "Doe",
				"objects[0]._acl[0].who[1].first_name": "Anna",
				"objects[0]._acl[0].who[1].last_name":  "Smith",
			},
		},
	}

	f := json.NewJsonFlattener()

	for _, tc := range testCases {
		result, err := f.Flatten(tc.in)

		assert.NoError(t, err)
		assert.Equal(t, tc.out, result)
	}
}
