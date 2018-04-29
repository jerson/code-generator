package models

// ForeignKey ...
type ForeignKey struct {
	Name               string            `yaml:",omitempty"`
	Namespace          string            `yaml:",omitempty"`
	LocalTable         string            `yaml:",omitempty"`
	LocalColumnName    []Identifier      `yaml:",omitempty"`
	ForeignTableName   Identifier        `yaml:",omitempty"`
	ForeignColumnNames []Identifier      `yaml:",omitempty"`
	Options            ForeignKeyOptions `yaml:",omitempty"`
}

// ForeignKeyOptions ...
type ForeignKeyOptions struct {
	OnUpdate string `yaml:",omitempty"`
	OnDelete string `yaml:",omitempty"`
}

//NewForeignKey ...
func NewForeignKey() ForeignKey {
	return ForeignKey{}
}

//NewForeignKeyOptions ...
func NewForeignKeyOptions() ForeignKeyOptions {
	return ForeignKeyOptions{}
}
