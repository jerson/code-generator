package platform

// ForeignKey ...
type ForeignKey struct {
	Base
	LocalTable         string
	LocalColumnName    string
	ForeignTableName   Table
	ForeignColumnNames []Identifier
	Options            map[string]string
}
