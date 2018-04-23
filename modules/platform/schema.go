package platform

// Schema ...
type Schema struct {
	Base
	Config     SchemaConfig
	Views      []View
	Tables     []Table
	Secuence   []Secuence
	Namespaces map[string]string
}
