package models

import "github.com/jerson/code-generator/modules/parser/types"

// Column ...
type Column struct {
	Base
	SpecialType     *Type             `json:",omitempty"`
	Type            types.Value       `json:",omitempty"`
	Length          int               `json:",omitempty"`
	Precision       int               `json:",omitempty"`
	Scale           int               `json:",omitempty"`
	Unsigned        bool              `json:",omitempty"`
	Fixed           bool              `json:",omitempty"`
	NotNull         bool              `json:",omitempty"`
	Default         string            `json:",omitempty"`
	AutoIncrement   bool              `json:",omitempty"`
	PlatformOptions PlatformOptions   `json:",omitempty"`
	Comment         string            `json:",omitempty"`
	SchemaOptions   map[string]string `json:",omitempty"`
}

// PlatformOptions ...
type PlatformOptions struct {
	Collation    string `json:",omitempty"`
	CharacterSet string `json:",omitempty"`
}

//NewColumn ...
func NewColumn() Column {
	return Column{}
}

//NewPlatformOptions ...
func NewPlatformOptions() PlatformOptions {
	return PlatformOptions{}
}
