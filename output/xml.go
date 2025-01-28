package output

import (
	"encoding/xml"
	"os"
	"sbom2doc-go/main"
)

type XMLSBOM struct {
	XMLName    xml.Name `xml:"sbom"`
	Components []struct {
		Name    string `xml:"name"`
		Version string `xml:"version"`
		License string `xml:"license"`
	} `xml:"component"`
}

func GenerateXML(sbom main.SBOM, outputFile string) error {
	file, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer file.Close()

	xmlSBOM := XMLSBOM{}
	for _, component := range sbom.Components {
		xmlSBOM.Components = append(xmlSBOM.Components, struct {
			Name    string `xml:"name"`
			Version string `xml:"version"`
			License string `xml:"license"`
		}{
			Name:    component.Name,
			Version: component.Version,
			License: component.License,
		})
	}

	encoder := xml.NewEncoder(file)
	encoder.Indent("", "  ")
	return encoder.Encode(xmlSBOM)
}