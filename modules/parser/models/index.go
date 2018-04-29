package models

// Index ...
type Index struct {
	Name      string       `yaml:",omitempty"`
	Namespace string       `yaml:",omitempty"`
	Columns   []Identifier `yaml:",omitempty"`
	IsUnique  bool         `yaml:",omitempty"`
	IsPrimary bool         `yaml:",omitempty"`
	Flags     []string     `yaml:",omitempty"`
	//Options   map[string]string
}

//NewIndex ...
func NewIndex() Index {
	return Index{}
}
