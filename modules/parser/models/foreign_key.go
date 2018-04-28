package models

// ForeignKey ...
type ForeignKey struct {
	Base
	LocalTable         string            `json:",omitempty"`
	LocalColumnName    []Identifier      `json:",omitempty"`
	ForeignTableName   Identifier        `json:",omitempty"`
	ForeignColumnNames []Identifier      `json:",omitempty"`
	Options            ForeignKeyOptions `json:",omitempty"`
}

// ForeignKeyOptions ...
type ForeignKeyOptions struct {
	OnUpdate string `json:",omitempty"`
	OnDelete string `json:",omitempty"`
}

//NewForeignKey ...
func NewForeignKey() ForeignKey {
	return ForeignKey{}
}

//NewForeignKeyOptions ...
func NewForeignKeyOptions() ForeignKeyOptions {
	return ForeignKeyOptions{}
}
