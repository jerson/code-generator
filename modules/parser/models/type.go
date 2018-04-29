package models

import "github.com/jerson/code-generator/modules/parser/types"

// Type ...
type Type struct {
	Name      string            `yaml:",omitempty"`
	Namespace string            `yaml:",omitempty"`
	Length    int               `yaml:",omitempty"`
	Precision int               `yaml:",omitempty"`
	Scale     int               `yaml:",omitempty"`
	Fixed     bool              `yaml:",omitempty"`
	Value     types.Value       `yaml:",omitempty"`
	Options   *TypeExtraOptions `yaml:",omitempty"`
}

//NewType ...
func NewType() Type {
	return Type{}
}

// TypeExtraOptions ...
type TypeExtraOptions struct {
	Length int      `yaml:",omitempty"`
	Help   string   `yaml:",omitempty"`
	Values []string `yaml:",omitempty"`
}

//NewTypeExtraOptions ...
func NewTypeExtraOptions() TypeExtraOptions {
	return TypeExtraOptions{}
}
