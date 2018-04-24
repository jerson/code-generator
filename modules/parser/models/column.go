package models

// Column ...
type Column struct {
	Base
	Type            Type
	Length          int
	Precision       int
	Scale           int
	Unsigned        bool
	Fixed           bool
	NotNull         bool
	Default         *string
	AutoIncrement   bool
	PlatformOptions PlatformOptions
	Comment         string
	SchemaOptions   map[string]string
}

// PlatformOptions ...
type PlatformOptions struct {
	Collation    string
	CharacterSet string
}
