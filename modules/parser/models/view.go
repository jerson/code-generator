package models

// View ...
type View struct {
	Base
	SQL string `json:",omitempty"`
}

//NewView ...
func NewView() View {
	return View{}
}
