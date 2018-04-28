package models

// Schema ...
type Schema struct {
	Base
	Config     SchemaConfig `json:",omitempty"`
	Views      []View       `json:",omitempty"`
	Tables     []Table      `json:",omitempty"`
	Sequences  []Sequence   `json:",omitempty"`
	Namespaces []string     `json:",omitempty"`
}

//NewSchema ...
func NewSchema() Schema {
	return Schema{}
}
