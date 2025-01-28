package output

import (
	"encoding/csv"
	"os"

	"github.com/thanakritks/sbom-2-doc-go/sbom"
)

func GenerateCSV(sbom sbom.SBOM, outputFile string) error {
	file, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	 // Write header with box
	 writer.Write([]string{"╭──────────────╮"})
	 writer.Write([]string{"│ SBOM Summary │"})
	 writer.Write([]string{"╰──────────────╯"})
	 writer.Write([]string{}) // Empty line for spacing


	// Write header
	writer.Write([]string{"Name", "Version", "License"})

	// Write components
	for _, component := range sbom.Components {
		writer.Write([]string{component.Name, component.Version, component.License})
	}

	return nil
}
