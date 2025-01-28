package output

import (
	"github.com/jung-kurt/gofpdf"
	"github.com/thanakritks/sbom-2-doc-go/sbom"
)

func GeneratePDF(sbom sbom.SBOM, outputFile string) error {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 12)
	pdf.CellFormat(0, 10, "╭──────────────╮", "", 1, "C", false, 0, "")
	pdf.CellFormat(0, 10, "│ SBOM Summary │", "", 1, "C", false, 0, "")
	pdf.CellFormat(0, 10, "╰──────────────╯", "", 1, "C", false, 0, "")
	pdf.Ln(10)
	for _, component := range sbom.Components {
		pdf.Cell(40, 10, "Name: "+component.Name)
		pdf.Ln(10)
		pdf.Cell(40, 10, "Version: "+component.Version)
		pdf.Ln(10)
		pdf.Cell(40, 10, "License: "+component.License)
		pdf.Ln(20)
	}

	err := pdf.OutputFileAndClose(outputFile)
	if err != nil {
		return err
	}

	return nil
}
