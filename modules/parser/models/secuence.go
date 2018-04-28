package models

// Sequence ...
type Sequence struct {
	Base
	AllocationSize int  `json:",omitempty"`
	InitialValue   int  `json:",omitempty"`
	Cache          *int `json:",omitempty"`
}

//NewSequence ...
func NewSequence() Sequence {
	return Sequence{}
}
