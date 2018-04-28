package models

// Table ...
type Table struct {
	Base
	PrimaryKeyName bool     `json:",omitempty"`
	Columns        []Column `json:",omitempty"`
	//ImplicitIndexes []Index `json:",omitempty"`
	Indexes    []Index           `json:",omitempty"`
	ForeignKey []ForeignKey      `json:",omitempty"`
	Options    map[string]string `json:",omitempty"`
	Config     SchemaConfig      `json:",omitempty"`
}
