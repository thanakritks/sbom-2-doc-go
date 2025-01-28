package output

import (
    "encoding/xml"
    "os"
    "github.com/thanakritks/sbom-2-doc-go/sbom"
)

func GenerateXML(sbom sbom.SBOM, outputFile string) error {
    file, err := os.Create(outputFile)
    if err != nil {
        return err
    }
    defer file.Close()

    encoder := xml.NewEncoder(file)
    encoder.Indent("", "  ")
    err = encoder.Encode(sbom)
    if err != nil {
        return err
    }

    return nil
}