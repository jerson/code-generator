package models

// Secuence ...
type Secuence struct {
	Base
	AllocationSize int
	InitialValue   int
	Cache          *int
}
