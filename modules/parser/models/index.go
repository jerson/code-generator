package models

// Index ...
type Index struct {
	Base
	Columns   []Identifier
	IsUnique  bool
	IsPrimary bool
	Flags     []string
	Options   map[string]string
}
