package output

import (
	"fmt"
	"sbom2doc-go/main"

	"github.com/jung-kurt/gofpdf"
)

func GeneratePDF(sbom main.SBOM, outputFile string) error {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Software Bill of Materials (SBOM)")

	pdf.Ln(12)
	pdf.SetFont("Arial", "", 12)

	for _, component := range sbom.Components {
		line := fmt.Sprintf("* %s - Version: %s, License: %s", component.Name, component.Version, component.License)
		pdf.Cell(40, 10, line)
		pdf.Ln(8)
	}

	return pdf.OutputFileAndClose(outputFile)
}
