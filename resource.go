package main

import (
	"fmt"
	"os"
	"strings"
)

func createResources(assetDirPath string, outputDirPath string) error {
	parts, err := getAssetResources(assetDirPath)
	if err != nil {
		return fmt.Errorf("can't get asset resources: %v", err)
	}

	template, err := os.ReadFile(assetDirPath + "/template.svg")
	if err != nil {
		return fmt.Errorf("can't read template file: %v", err)
	}

	count := 0
	assemble := func(parts [][]byte) error {
		svg := string(template)

		for _, part := range parts {
			placeholder := templatePlaceholderRegex.FindString(svg)
			if placeholder == "" {
				return fmt.Errorf("can't find matched placeholder in template")
			}

			svg = strings.Replace(svg, placeholder, string(part), 1)
		}

		if err := os.WriteFile(outputDirPath+"/"+fmt.Sprintf("%d.svg", count), []byte(svg), 0644); err != nil {
			return fmt.Errorf("can't write file: %v", err)
		}

		// metadata := metadata{
		// 	Name:        babbler.Babble(),
		// 	Descriction: "Cool ice cream",
		// 	Image:       fmt.Sprintf("ipfs://YOUR_ASSET_CID"),
		// }

		count++
		return nil
	}

	if err := generateProduct(assemble, parts...); err != nil {
		return fmt.Errorf("can't generate product: %v", err)
	}

	return nil
}
