package output

import (
	"fmt"

	"github.com/jung-kurt/gofpdf"
	"github.com/thanakritks/sbom-2-doc-go.git/main"
)

// GeneratePDF generates a PDF file from the SBOM
func GeneratePDF(sbom main.SBOM, outputFile string) error {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Software Bill of Materials (SBOM)")

	// Add SBOM data
	pdf.SetFont("Arial", "", 12)
	for _, component := range sbom.Components {
		line := fmt.Sprintf("* %s - Version: %s, License: %s", component.Name, component.Version, component.License)
		pdf.Ln(10)
		pdf.Cell(40, 10, line)
	}

	// Save to file
	return pdf.OutputFileAndClose(outputFile)
}
