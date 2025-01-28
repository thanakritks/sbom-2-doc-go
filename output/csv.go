package output

import (
	"encoding/csv"
	"os"
	"sbom2doc-go/main"
)

func GenerateCSV(sbom main.SBOM, outputFile string) error {
	file, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header
	writer.Write([]string{"Name", "Version", "License"})

	// Write components
	for _, component := range sbom.Components {
		writer.Write([]string{component.Name, component.Version, component.License})
	}

	return nil
}
