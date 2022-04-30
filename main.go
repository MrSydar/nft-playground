package main

import (
	"flag"
	"fmt"
	"os"
)

var assetsPath = flag.String("a", "", "path to resource assets directory")
var outputPath = flag.String("o", "", "path to output directory where the resources will be created")

func main() {
	flag.Parse()

	if err := createResources(*assetsPath, *outputPath); err != nil {
		fmt.Fprintf(os.Stderr, "can't create NFTs: %v\n", err)
		os.Exit(1)
	}
}
