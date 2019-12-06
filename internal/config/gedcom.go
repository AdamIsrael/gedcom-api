package config

// GedcomConfiguration stores the configuration related to GEDCOM files
type GedcomConfiguration struct {
	// Path to store GEDCOM files
	Path string

	MaxFileSize int64
}
