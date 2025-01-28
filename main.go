package main

import (
	"encoding/csv"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jung-kurt/gofpdf" // For PDF generation
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
		fmt.Println("Supported formats: csv, xml, pdf")
		os.Exit(1)
	}

	// Read the SBOM file
	sbomFile := os.Args[1]
	outputFormat := os.Args[2]
	outputFile := "output." + outputFormat
	if len(os.Args) > 3 {
		outputFile = os.Args[3]
	}

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

	// Generate the output in the specified format
	switch outputFormat {
	case "csv":
		err = generateCSV(sbom, outputFile)
	case "xml":
		err = generateXML(sbom, outputFile)
	case "pdf":
		err = generatePDF(sbom, outputFile)
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

// generateCSV generates a CSV file from the SBOM
func generateCSV(sbom SBOM, outputFile string) error {
	file, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write CSV header
	writer.Write([]string{"Name", "Version", "License"})

	// Write each component to the CSV
	for _, component := range sbom.Components {
		writer.Write([]string{component.Name, component.Version, component.License})
	}

	return nil
}

// generateXML generates an XML file from the SBOM
func generateXML(sbom SBOM, outputFile string) error {
	file, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := xml.NewEncoder(file)
	encoder.Indent("", "  ")

	// Write XML header
	file.WriteString(`<?xml version="1.0" encoding="UTF-8"?>` + "\n")

	// Encode the SBOM as XML
	err = encoder.Encode(sbom)
	if err != nil {
		return err
	}

	return nil
}

// generatePDF generates a PDF file from the SBOM
func generatePDF(sbom SBOM, outputFile string) error {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)

	// Add a title
	pdf.Cell(40, 10, "Software Bill of Materials (SBOM)")
	pdf.Ln(12)

	// Set font for the table
	pdf.SetFont("Arial", "", 12)

	// Add table headers
	pdf.Cell(60, 10, "Name")
	pdf.Cell(40, 10, "Version")
	pdf.Cell(40, 10, "License")
	pdf.Ln(10)

	// Add each component to the PDF
	for _, component := range sbom.Components {
		pdf.Cell(60, 10, component.Name)
		pdf.Cell(40, 10, component.Version)
		pdf.Cell(40, 10, component.License)
		pdf.Ln(10)
	}

	// Save the PDF to a file
	return pdf.OutputFileAndClose(outputFile)
}