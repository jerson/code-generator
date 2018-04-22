// Package parser ...
package parser

import ()

// TableInterface ...
type TableInterface interface {
	// GetName ...
	GetName() string
}

// Column ...
type Column struct {
	Field        string `gorm:"column:Field"`
	Type         string `gorm:"column:Type"`
	Null         string `gorm:"column:Null"`
	Key          string `gorm:"column:Key"`
	Default      string `gorm:"column:Default"`
	Extra        string `gorm:"column:Extra"`
	Comment      string `gorm:"column:Comment"`
	CharacterSet string `gorm:"column:CharacterSet"`
	Collation    string `gorm:"column:Collation"`
}

// Index ...
type Index struct {
	Table       string `gorm:"column:Table"`
	NonUnique   string `gorm:"column:Non_unique"`
	KeyName     string `gorm:"column:Key_name"`
	SeqInIndex  string `gorm:"column:Seq_in_index"`
	ColumnName  string `gorm:"column:Column_name"`
	Collation   string `gorm:"column:Collation"`
	Cardinality string `gorm:"column:Cardinality"`
	SubPart     string `gorm:"column:Sub_part"`
	Packed      string `gorm:"column:Packed"`
	Null        string `gorm:"column:Null"`
	IndexType   string `gorm:"column:Index_type"`
	Comment     string `gorm:"column:Comment"`
}

// View ...
type View struct {
	TableCatalog        string `gorm:"column:TABLE_CATALOG"`
	TableSchema         string `gorm:"column:TABLE_SCHEMA"`
	TableName           string `gorm:"column:TABLE_NAME"`
	ViewDefinition      string `gorm:"column:VIEW_DEFINITION"`
	CheckOption         string `gorm:"column:CHECK_OPTION"`
	IsUpdateable        string `gorm:"column:IS_UPDATABLE"`
	Definer             string `gorm:"column:DEFINER"`
	SecurityType        string `gorm:"column:SECURITY_TYPE"`
	CharacterSetClient  string `gorm:"column:CHARACTER_SET_CLIENT"`
	CollationConnection string `gorm:"column:COLLATION_CONNECTION"`
}

// ForeignKey ...
type ForeignKey struct {
	ConstraintName       string `gorm:"column:CONSTRAINT_NAME"`
	ColumnName           string `gorm:"column:COLUMN_NAME"`
	ReferencedTableName  string `gorm:"column:REFERENCED_TABLE_NAME"`
	ReferencedColumnName string `gorm:"column:REFERENCED_COLUMN_NAME"`
	UpdateRule           string `gorm:"column:update_rule"`
	DeleteRule           string `gorm:"column:delete_rule"`
}

// Table ...
type Table struct {
	Name      string
	TableType string `gorm:"column:Table_type"`
}
