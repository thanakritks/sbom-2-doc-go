package sbom

// SBOM represents the structure of the SBOM JSON
type SBOM struct {
	Components []Component `json:"components"`
}

// Component represents a single component in the SBOM
type Component struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	License string `json:"license"`
}
