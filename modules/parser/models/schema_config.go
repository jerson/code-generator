package models

// SchemaConfig ...
type SchemaConfig struct {
	Name                         string            `json:",omitempty"`
	MaxIdentifierLength          int               `json:",omitempty"`
	HasExplicitForeignKeyIndexes bool              `json:",omitempty"`
	DefaultTableOptions          map[string]string `json:",omitempty"`
}

//NewSchemaConfig ...
func NewSchemaConfig() SchemaConfig {
	return SchemaConfig{MaxIdentifierLength: 63}
}
