package models

// Sequence ...
type Sequence struct {
	Name           string `yaml:",omitempty"`
	Namespace      string `yaml:",omitempty"`
	AllocationSize int    `yaml:",omitempty"`
	InitialValue   int    `yaml:",omitempty"`
	Cache          *int   `yaml:",omitempty"`
}

//NewSequence ...
func NewSequence() Sequence {
	return Sequence{}
}
