package models

// ForeignKey ...
type ForeignKey struct {
	Base
	LocalTable         string
	LocalColumnName    []Identifier
	ForeignTableName   Identifier
	ForeignColumnNames []Identifier
	Options            ForeignKeyOptions
}

// ForeignKeyOptions ...
type ForeignKeyOptions struct {
	OnUpdate string
	OnDelete string
}
