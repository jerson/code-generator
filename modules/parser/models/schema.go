package models

// Schema ...
type Schema struct {
	Base
	Config     SchemaConfig
	Views      []View
	Tables     []Table
	Sequences  []Sequence
	Namespaces []string
}
