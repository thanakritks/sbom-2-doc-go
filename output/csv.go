package output

import (
	"encoding/csv"
	"os"

	"github.com/thanakritks/sbom-2-doc-go.git/main"
)

// GenerateCSV generates a CSV file from the SBOM
func GenerateCSV(sbom main.SBOM, outputFile string) error {
	file, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write the header
	header := []string{"Name", "Version", "License"}
	if err := writer.Write(header); err != nil {
		return err
	}

	// Write the SBOM data
	for _, component := range sbom.Components {
		record := []string{component.Name, component.Version, component.License}
		if err := writer.Write(record); err != nil {
			return err
		}
	}

	return nil
}
