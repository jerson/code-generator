package models

// SchemaConfig ...
type SchemaConfig struct {
	Name                         string
	MaxIdentifierLength          int
	HasExplicitForeignKeyIndexes bool
	DefaultTableOptions          map[string]string
}

//NewSchemaConfig ...
func NewSchemaConfig() SchemaConfig {
	return SchemaConfig{MaxIdentifierLength: 63}
}
