package models

// SchemaConfig ...
type SchemaConfig struct {
	Name                         string            `yaml:",omitempty"`
	MaxIdentifierLength          int               `yaml:",omitempty"`
	HasExplicitForeignKeyIndexes bool              `yaml:",omitempty"`
	DefaultTableOptions          map[string]string `yaml:",omitempty"`
}

//NewSchemaConfig ...
func NewSchemaConfig() SchemaConfig {
	return SchemaConfig{MaxIdentifierLength: 63}
}
