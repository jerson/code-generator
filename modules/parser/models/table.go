package models

// Table ...
type Table struct {
	Name           string   `yaml:",omitempty"`
	Namespace      string   `yaml:",omitempty"`
	PrimaryKeyName bool     `yaml:",omitempty"`
	Columns        []Column `yaml:",omitempty"`
	//ImplicitIndexes []Index `yaml:",omitempty"`
	Indexes    []Index           `yaml:",omitempty"`
	ForeignKey []ForeignKey      `yaml:",omitempty"`
	Options    map[string]string `yaml:",omitempty"`
	Config     SchemaConfig      `yaml:",omitempty"`
}

//NewTable ...
func NewTable() Table {
	return Table{}
}
