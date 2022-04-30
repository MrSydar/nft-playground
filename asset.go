package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

// getTemplatePlaceholderRegex returns a regexp that matches all placeholders in the given resource template
var templatePlaceholderRegex = regexp.MustCompile(`<!--\s*([a-z]+)\s*-->`)

// readAllFiles reads all files from the given directory and returns their content as a slice of byte slices
func readAllFiles(path string) ([][]byte, error) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, fmt.Errorf("can't open directory: %v", err)
	}

	variants := make([][]byte, 0, len(files))
	for _, file := range files {
		variant, err := ioutil.ReadFile(path + "/" + file.Name())
		if err != nil {
			return nil, fmt.Errorf("can't read file: %v", err)
		}
		variants = append(variants, variant)
	}

	return variants, nil
}

// getAssetResources reads resource template and all its assets from the subdirectories and
// returns their content as a slice of byte slices in the same order, as they are listed in the template
func getAssetResources(assetPath string) ([][][]byte, error) {
	template, err := os.ReadFile(assetPath + "/template.svg")
	if err != nil {
		return nil, fmt.Errorf("can't read template file: %v", err)
	}

	var resources [][][]byte
	for _, match := range templatePlaceholderRegex.FindAllStringSubmatch(string(template), -1) {
		variants, err := readAllFiles(assetPath + "/" + match[1])
		if err != nil {
			return nil, fmt.Errorf("can't read asset part variants: %v", err)
		}

		resources = append(resources, variants)
	}

	return resources, nil
}
