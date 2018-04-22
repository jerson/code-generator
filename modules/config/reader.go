//Package config ...
package config

import (
	"github.com/jinzhu/configor"
)

//Output ...
type Output struct {
	Dir string `toml:"dir" default:"./output"`
}

//Database ...
type Database struct {
	Driver string `toml:"driver"`
	Source string `toml:"source"`
}

//Project ...
type Project struct {
	Name    string `toml:"name"`
	Package string `toml:"package"`
}

//Template ...
type Template struct {
	Name string `toml:"name"`
}

//Vars ...
var Vars = struct {
	Debug    bool     `toml:"debug" default:"false"`
	Version  string   `toml:"version" default:"latest"`
	Output   Output   `toml:"output"`
	Database Database `toml:"database"`
	Project  Project  `toml:"project"`
	Template Template `toml:"template"`
}{}

//ReadDefault ...
func ReadDefault() error {
	return Read("./config.toml")
}

//Read ...
func Read(file string) error {
	return configor.New(&configor.Config{ENVPrefix: "APP", Debug: false, Verbose: false}).Load(&Vars, file)
}
