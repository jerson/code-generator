package models

import "github.com/jerson/code-generator/modules/parser/types"

// Type ...
type Type struct {
	Base
	Length    int               `json:",omitempty"`
	Precision int               `json:",omitempty"`
	Scale     int               `json:",omitempty"`
	Fixed     bool              `json:",omitempty"`
	Value     types.Value       `json:",omitempty"`
	Options   *TypeExtraOptions `json:",omitempty"`
}

//NewType ...
func NewType() Type {
	return Type{}
}

// TypeExtraOptions ...
type TypeExtraOptions struct {
	Length int    `json:",omitempty"`
	Help   string `json:",omitempty"`
}

//NewTypeExtraOptions ...
func NewTypeExtraOptions() TypeExtraOptions {
	return TypeExtraOptions{}
}
