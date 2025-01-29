package output

import (
	"fmt"
	"os"

	"github.com/thanakritks/sbom-2-doc-go/sbom"
)

func GenerateTXT(sbom sbom.SBOM, outputFile string) error {
	file, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer file.Close()

	fmt.Fprintln(file, "╭──────────────╮")
	fmt.Fprintln(file, "│ SBOM Summary │")
	fmt.Fprintln(file, "╰──────────────╯")

	for _, component := range sbom.Components {
		_, err := fmt.Fprintf(file, "Name: %s, Version: %s, License: %s\n", component.Name, component.Version, component.License)
		if err != nil {
			return err
		}
	}

	return nil
}
