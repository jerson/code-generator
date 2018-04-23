package platform

// Index ...
type Index struct {
	Base
	Columns   []Identifier
	IsUnique  bool
	IsPrimary bool
	Flags     map[string]string
	Options   map[string]string
}
