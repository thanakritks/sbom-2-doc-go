package output

import (
	"encoding/xml"
	"os"

	"github.com/thanakritks/sbom-2-doc-go.git/main"
)

// SBOMXML represents the XML structure of the SBOM
type SBOMXML struct {
	XMLName    xml.Name       `xml:"sbom"`
	Components []ComponentXML `xml:"component"`
}

// ComponentXML represents a single component in the XML
type ComponentXML struct {
	Name    string `xml:"name"`
	Version string `xml:"version"`
	License string `xml:"license"`
}

// GenerateXML generates an XML file from the SBOM
func GenerateXML(sbom main.SBOM, outputFile string) error {
	file, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer file.Close()

	// Convert SBOM to XML structure
	var sbomXML SBOMXML
	for _, component := range sbom.Components {
		sbomXML.Components = append(sbomXML.Components, ComponentXML{
			Name:    component.Name,
			Version: component.Version,
			License: component.License,
		})
	}

	// Write XML to file
	encoder := xml.NewEncoder(file)
	encoder.Indent("", "  ")
	if err := encoder.Encode(sbomXML); err != nil {
		return err
	}

	return nil
}
