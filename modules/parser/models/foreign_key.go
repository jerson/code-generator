package models

// ForeignKey ...
type ForeignKey struct {
	Base
	LocalTable             string
	LocalColumnName        string
	ForeignTableName       *Table
	ForeignTableIdentifier *Identifier
	ForeignColumnNames     []Identifier
	Options                map[string]string
}
