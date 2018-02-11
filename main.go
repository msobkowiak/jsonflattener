package main

import (
	"fmt"
	"github.com/msobkowiak/jsonflatterner/flattener/json"
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

	input := []byte(`{"name":"Joe", "address":{"street":"123 Main St."}}`)
	output, err :=  json.Flatten(input, ".")
	if err != nil {
		fmt.Printf("Something went wrong: %s\n", err.Error())
	}

	fmt.Print(string(output))
}
