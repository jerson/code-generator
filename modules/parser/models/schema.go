package models

// Schema ...
type Schema struct {
	Name       string       `yaml:",omitempty"`
	Namespace  string       `yaml:",omitempty"`
	Config     SchemaConfig `yaml:",omitempty"`
	Views      []View       `yaml:",omitempty"`
	Tables     []Table      `yaml:",omitempty"`
	Sequences  []Sequence   `yaml:",omitempty"`
	Namespaces []string     `yaml:",omitempty"`
}

//NewSchema ...
func NewSchema() Schema {
	return Schema{}
}
