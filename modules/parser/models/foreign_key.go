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

//ForeignColumns ...
func (f ForeignKey) ForeignColumns() []string {
	var names []string
	for _, column := range f.ForeignColumnNames {
		names = append(names, column.Name)
	}
	return names
}

//LocalColumns ...
func (f ForeignKey) LocalColumns() []string {
	var names []string
	for _, column := range f.LocalColumnName {
		names = append(names, column.Name)
	}
	return names
}
