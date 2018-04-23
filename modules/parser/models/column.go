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
	PlatformOptions map[string]string
	Comment         string
	SchemaOptions   map[string]string
}
