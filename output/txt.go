package output

import (
	"fmt"
	"os"

	"github.com/thanakritks/sbom-2-doc-go.git/main"
)

// GenerateTXT generates a TXT file from the SBOM
func GenerateTXT(sbom main.SBOM, outputFile string) error {
	file, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer file.Close()

	// Write the SBOM data to the file
	_, err = file.WriteString("Software Bill of Materials (SBOM)\n")
	_, err = file.WriteString("===============================\n\n")
	for _, component := range sbom.Components {
		line := fmt.Sprintf("* %s - Version: %s, License: %s\n", component.Name, component.Version, component.License)
		_, err = file.WriteString(line)
		if err != nil {
			return err
		}
	}

	return nil
}
