package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/thanakritks/sbom-2-doc-go.git/output"
)

// SBOM represents the structure of the SBOM JSON
type SBOM struct {
	Components []Component `json:"components"`
}

// Component represents a single component in the SBOM
type Component struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	License string `json:"license"`
}

func main() {
	// Check if the user provided an SBOM file and output format
	if len(os.Args) < 3 {
		fmt.Println("Usage: sbom2doc <sbom-file> <output-format> [output-file]")
		fmt.Println("Supported formats: txt, csv, xml, pdf")
		os.Exit(1)
	}

	// Read the SBOM file
	sbomFile := os.Args[1]
	data, err := ioutil.ReadFile(sbomFile)
	if err != nil {
		fmt.Printf("Error reading SBOM file: %s\n", err)
		os.Exit(1)
	}

	// Parse the SBOM JSON
	var sbom SBOM
	err = json.Unmarshal(data, &sbom)
	if err != nil {
		fmt.Printf("Error parsing SBOM JSON: %s\n", err)
		os.Exit(1)
	}

	// Determine the output format
	outputFormat := os.Args[2]
	outputFile := "output." + outputFormat // Default output file name
	if len(os.Args) > 3 {
		outputFile = os.Args[3] // Custom output file name
	}

	// Generate the output based on the format
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
		fmt.Printf("Error generating %s output: %s\n", outputFormat, err)
		os.Exit(1)
	}

	fmt.Printf("Output successfully written to %s\n", outputFile)
}
