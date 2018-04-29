package models

import (
	"github.com/jerson/code-generator/modules/parser/keys"
	"github.com/jerson/code-generator/modules/parser/types"
)

// Column ...
type Column struct {
	Name            string            `yaml:",omitempty"`
	Namespace       string            `yaml:",omitempty"`
	SpecialType     *Type             `yaml:",omitempty"`
	Type            types.Value       `yaml:",omitempty"`
	Length          int               `yaml:",omitempty"`
	Precision       int               `yaml:",omitempty"`
	Scale           int               `yaml:",omitempty"`
	Unsigned        bool              `yaml:",omitempty"`
	Fixed           bool              `yaml:",omitempty"`
	NotNull         bool              `yaml:",omitempty"`
	Default         string            `yaml:",omitempty"`
	AutoIncrement   bool              `yaml:",omitempty"`
	Key             keys.Value        `yaml:",omitempty"`
	PlatformOptions PlatformOptions   `yaml:",omitempty"`
	Comment         string            `yaml:",omitempty"`
	SchemaOptions   map[string]string `yaml:",omitempty"`
}

// PlatformOptions ...
type PlatformOptions struct {
	Collation    string `yaml:",omitempty"`
	CharacterSet string `yaml:",omitempty"`
}

//NewColumn ...
func NewColumn() Column {
	return Column{}
}

//NewPlatformOptions ...
func NewPlatformOptions() PlatformOptions {
	return PlatformOptions{}
}
