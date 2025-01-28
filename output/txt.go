package output

import (
	"fmt"
	"os"
	"sbom2doc-go/main"
)

func GenerateTXT(sbom main.SBOM, outputFile string) error {
	file, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer file.Close()

	header := "Software Bill of Materials (SBOM)\n===============================\n\n"
	file.WriteString(header)

	for _, component := range sbom.Components {
		line := fmt.Sprintf("* %s - Version: %s, License: %s\n", component.Name, component.Version, component.License)
		file.WriteString(line)
	}

	return nil
}