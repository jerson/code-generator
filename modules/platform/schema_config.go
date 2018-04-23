package platform

// SchemaConfig ...
type SchemaConfig struct {
	Name                         string
	MaxIdentifierLength          int
	HasExplicitForeignKeyIndexes bool
	DefaultTableOptions          map[string]string
}

func NewSchemaConfig() SchemaConfig {
	return SchemaConfig{MaxIdentifierLength: 63}
}
