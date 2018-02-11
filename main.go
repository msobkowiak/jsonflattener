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

	input := []byte(`{"name":"Joe", "last_name": "Doe", "age": 25, "hobby": ["travel", "sport", "books"], "address":{"street":"123 Main St."}}`)

	jf := json.NewJsonFlattener()
	output, err := jf.Flatten(input, ".")
	if err != nil {
		fmt.Printf("Something went wrong: %s\n", err.Error())
	}

	fmt.Println(string(output))
}
