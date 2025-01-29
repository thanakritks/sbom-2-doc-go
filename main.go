package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/corp-ais/sbom-to-doc/output"
	"github.com/corp-ais/sbom-to-doc/sbom"
)

func main() {
	// Check if the user provided an SBOM file and output format
	if len(os.Args) < 3 {
		fmt.Println("Usage: sbom2doc <sbom-file> <output-format> [output-file]")
		fmt.Println("Supported formats: txt, csv, xml, pdf")
		os.Exit(1)
	}

	sbomFile := os.Args[1]
	outputFormat := os.Args[2]
	outputFile := "output." + outputFormat // Default output file name

	// Allow custom output file name
	if len(os.Args) > 3 {
		outputFile = os.Args[3]
	}

	// Read the SBOM file
	data, err := ioutil.ReadFile(sbomFile)
	if err != nil {
		fmt.Printf("Error reading SBOM file: %s\n", err)
		os.Exit(1)
	}

	// Parse the SBOM JSON
	var sbom sbom.SBOM
	err = json.Unmarshal(data, &sbom)
	if err != nil {
		fmt.Printf("Error parsing SBOM JSON: %s\n", err)
		os.Exit(1)
	}

	// Generate the output based on the chosen format
	switch outputFormat {
	case "txt":
		err = output.GenerateTXT(sbom, outputFile)
	case "csv":
		err = output.GenerateCSV(sbom, outputFile)
	case "xml":
		err = output.GenerateXML(sbom, outputFile)
	case "pdf":
		err = output.GeneratePDF(sbom, outputFile)
	default:
		fmt.Printf("Unsupported output format: %s\n", outputFormat)
		os.Exit(1)
	}

	if err != nil {
		fmt.Printf("Error generating %s: %s\n", outputFormat, err)
		os.Exit(1)
	}

	fmt.Printf("Successfully generated %s\n", outputFile)
}
