package models

// Index ...
type Index struct {
	Base      `json:",omitempty"`
	Columns   []Identifier `json:",omitempty"`
	IsUnique  bool         `json:",omitempty"`
	IsPrimary bool         `json:",omitempty"`
	Flags     []string     `json:",omitempty"`
	//Options   map[string]string
}
