package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/beevik/etree"
)

func main() {
	// Define command-line flags
	inputFile := flag.String("input", "", "Input SVG file")
	outputFile := flag.String("output", "", "Output JSON file")
	flag.Parse()

	if *inputFile == "" {
		fmt.Println("Please provide an input SVG file using the -input flag.")
		return
	}

	if *outputFile == "" {
		// If the output file is not specified, create a default output file name.
		// For example, if input is "example.svg", output will be "example.json".
		inputFileName := *inputFile
		*outputFile = strings.TrimSuffix(inputFileName, ".svg") + ".json"
	}

	// Read the SVG file
	doc := etree.NewDocument()
	if err := doc.ReadFromFile(*inputFile); err != nil {
		fmt.Printf("Error reading input SVG file: %v\n", err)
		return
	}

	// Parse the SVG data and extract important information
	svgData := parseSVG(doc.Root())

	// Wrap the JSON data in the desired format
	wrappedJSON := map[string]interface{}{
		"api_version": "v2",
		"kind":        "document",
		"type":        "svg",
		"spec":        svgData,
	}

	// Convert the wrapped JSON data to a JSON string
	jsonData, err := json.MarshalIndent(wrappedJSON, "", "  ")
	if err != nil {
		fmt.Printf("Error converting wrapped JSON data to JSON string: %v\n", err)
		return
	}

	// Write JSON data to the output file
	err = ioutil.WriteFile(*outputFile, jsonData, 0644)
	if err != nil {
		fmt.Printf("Error writing JSON data to the output file: %v\n", err)
		return
	}

	fmt.Printf("SVG converted to JSON and saved to %s\n", *outputFile)
}

func parseSVG(element *etree.Element) map[string]interface{} {
	result := make(map[string]interface{})
	result["name"] = element.Tag

	attributes := make(map[string]string)
	for _, attr := range element.Attr {
		attributes[attr.Key] = attr.Value
	}
	result["attributes"] = attributes

	children := []map[string]interface{}{}
	for _, child := range element.ChildElements() {
		children = append(children, parseSVG(child))
	}

	if len(children) > 0 {
		result["children"] = children
	} else {
		result["content"] = element.Text()
	}

	return result
}
