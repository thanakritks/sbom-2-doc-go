package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"text/template"
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
	// Check if the user provided an SBOM file
	if len(os.Args) < 2 {
		fmt.Println("Usage: sbom2doc <sbom-file>")
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

	// Generate the documentation
	generateDoc(sbom)
}

func generateDoc(sbom SBOM) {
	// Define a template for the documentation
	const docTemplate = `Software Bill of Materials (SBOM)
===============================

{{range .Components}}
* {{.Name}} - Version: {{.Version}}, License: {{.License}}
{{end}}
`

	// Parse the template
	tmpl, err := template.New("doc").Parse(docTemplate)
	if err != nil {
		fmt.Printf("Error creating template: %s\n", err)
		os.Exit(1)
	}

	// Execute the template with the SBOM data
	err = tmpl.Execute(os.Stdout, sbom)
	if err != nil {
		fmt.Printf("Error executing template: %s\n", err)
		os.Exit(1)
	}
}
