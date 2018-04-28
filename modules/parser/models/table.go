package models

// Table ...
type Table struct {
	Base
	PrimaryKeyName bool
	Columns        []Column
	//ImplicitIndexes []Index
	Indexes    []Index
	ForeignKey []ForeignKey
	Options    map[string]string
	Config     SchemaConfig
}
