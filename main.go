package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/beevik/etree"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: svg2json directory_name")
		return
	}

	dirName := os.Args[1]

	err := convertSVGFiles(dirName)
	if err != nil {
		fmt.Printf("\033[31mError: %v\033[0m\n", err)
		return
	}
}

func convertSVGFiles(dirName string) error {
	return filepath.Walk(dirName, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) == ".svg" {
			jsonFile := strings.TrimSuffix(path, ".svg") + ".json"
			if err := convertSVGToJSON(path, jsonFile); err != nil {
				return err
			}
			fmt.Printf("Converted %s to %s\n", path, jsonFile)
		}
		return nil
	})
}

func convertSVGToJSON(inputFile, outputFile string) error {
	doc := etree.NewDocument()
	if err := doc.ReadFromFile(inputFile); err != nil {
		return err
	}

	svgData := parseSVG(doc.Root())

	wrappedJSON := map[string]interface{}{
		"api_version": "v2",
		"kind":        "document",
		"type":        "svg",
		"spec":        svgData,
	}

	jsonData, err := json.MarshalIndent(wrappedJSON, "", "  ")
	if err != nil {
		return err
	}

	return ioutil.WriteFile(outputFile, jsonData, 0644)
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
