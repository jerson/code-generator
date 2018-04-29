package models

// View ...
type View struct {
	Name      string `yaml:",omitempty"`
	Namespace string `yaml:",omitempty"`
	SQL       string `yaml:",omitempty"`
}

//NewView ...
func NewView() View {
	return View{}
}
