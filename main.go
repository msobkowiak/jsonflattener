package main

import (
	"fmt"

	"jsonflatterner/flattener/json"
)

func main() {
	// read a json file
	// flatten
	// print result

	// TODO: read json file path form console
	// TODO: add .json extension??
	//json, err := ioutil.ReadFile("testdata.json")
	//if err != nil {
	//	fmt.Println("Error on reading data from file" + err.Error())
	//	return
	//}

	input := []byte(
		`{
			"format": "long",
			"type": "objects",
			"offset": 0,
			"limit": 10,
			"facets": {},
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
	)

	jf := json.NewJsonFlattener()
	output, err := jf.Flatten(input)
	if err != nil {
		fmt.Printf("Something went wrong: %s\n", err.Error())
	}

	fmt.Println(output)
}
