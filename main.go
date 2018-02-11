package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"jsonflatterner/flattener/json"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Please provide path to json file")
		os.Exit(1)
	}

	input, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error on reading data from file" + err.Error())
		os.Exit(1)
	}

	f := json.NewJsonFlattener()
	output, err := f.Flatten(input)
	if err != nil {
		fmt.Printf("Something went wrong: %s\n", err.Error())
		os.Exit(1)
	}

	printMap(output)
}

func printMap(data map[string]interface{}) {
	for k, v := range data {
		fmt.Printf("%s: %#v\n", k, v)
	}
}
