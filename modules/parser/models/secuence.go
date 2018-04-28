package models

// Sequence ...
type Sequence struct {
	Base
	AllocationSize int
	InitialValue   int
	Cache          *int
}
