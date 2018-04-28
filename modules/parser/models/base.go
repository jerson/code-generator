package models

// Base ...
type Base struct {
	Name      string `json:",omitempty"`
	Namespace string `json:",omitempty"`
	//Quoted    bool `json:",omitempty"`
}

//NewBase ...
func NewBase() Base {
	return Base{}
}
