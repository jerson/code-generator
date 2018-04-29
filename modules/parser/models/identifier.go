package models

// Identifier ...
type Identifier struct {
	Name      string `yaml:",omitempty"`
	Namespace string `yaml:",omitempty"`
}

//NewIdentifier ...
func NewIdentifier() Identifier {
	return Identifier{}
}
