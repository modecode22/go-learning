package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

type Item struct {
	Name string `json:"name"`
	Code string `json:"dial"`
}

type ResultItem struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

func processAndWriteToFile(inputFile string) error {
	// Read input JSON file
	inputData, err := os.ReadFile(inputFile)
	if err != nil {
		return err
	}

	var list []Item

	// Unmarshal JSON data into the list
	err = json.Unmarshal(inputData, &list)
	if err != nil {
		return err
	}

	result := make([]ResultItem, len(list))

	for i, item := range list {
		result[i] = ResultItem{Name: item.Name, Code: item.Code}
	}

	// Convert the result to JSON
	resultJSON, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		return err
	}

	// Determine the output file path
	outputFilePath := filepath.Join(filepath.Dir(inputFile), "output.json")

	// Write JSON to the output file
	err = os.WriteFile(outputFilePath, resultJSON, 0644)
	if err != nil {
		return err
	}

	fmt.Printf("Result written to %s\n", outputFilePath)
	return nil
}

func main() {
	// Parse command-line arguments
	var inputFile string
	flag.StringVar(&inputFile, "input", "", "Input JSON file path")
	flag.Parse()

	if inputFile == "" {
		fmt.Println("Usage: go run main.go -input <input JSON file>")
		os.Exit(1)
	}

	// Process the input file and write to an output file
	err := processAndWriteToFile(inputFile)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
