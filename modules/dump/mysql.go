package dump

import (
	"fmt"
	"github.com/jerson/code-generator/modules/context"
	"github.com/jerson/code-generator/modules/parser/models"
	"github.com/jerson/code-generator/modules/parser/types"
	"strings"
)

// MySQL ...
type MySQL struct {
	ctx     context.Base
	schema  models.Schema
	options MySQLOptions
}

// MySQLOptions ...
type MySQLOptions struct {
	WithDrop bool
	Alter    bool
}

// NewMySQL ...
func NewMySQL(ctx context.Base, schema models.Schema, options MySQLOptions) MySQL {
	return MySQL{ctx: ctx, schema: schema, options: options}
}

//Dump ...
func (m MySQL) Dump() (string, error) {

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf(`-- database: %s`, m.schema.Name))
	sb.WriteRune('\n')

	for _, table := range m.schema.Tables {
		sb.WriteString(m.Table(table))
		sb.WriteRune('\n')

		for _, foreignKey := range table.ForeignKey {
			sb.WriteString(m.ForeignKey(foreignKey))
			sb.WriteRune('\n')
		}
		sb.WriteRune('\n')

		for _, index := range table.Indexes {
			sb.WriteString(m.Index(index, table))
			sb.WriteRune('\n')
		}
		sb.WriteRune('\n')
	}
	for _, view := range m.schema.Views {
		sb.WriteString(m.View(view))
		sb.WriteRune('\n')
	}

	return sb.String(), nil
}

//Quoted ...
func (m MySQL) Quoted(name string) string {
	return fmt.Sprintf("´%s´", name)
}

//View ...
func (m MySQL) View(view models.View) string {
	return fmt.Sprintf(`CREATE VIEW %s AS %s`, view.Name, view.SQL)
}

//Type ...
func (m MySQL) Type(column models.Column) string {
	columnType := "TEXT"
	useLength := false

	switch column.Type {
	case types.Unknown:
		columnType = "TEXT"
	case types.String, types.Array:
		useLength = true
		columnType = "VARCHAR"
	case types.JSON:
		columnType = "JSON"
	case types.Integer:
		useLength = true
		columnType = "INT"
	case types.BigInt:
		useLength = true
		columnType = "BIGINT"
	case types.SmallInt:
		columnType = "TINYINT"
	case types.Boolean:
		useLength = true
		columnType = "TINYINT"
	case types.Datetime:
		columnType = "DATETIME"
	case types.Timestamp:
		columnType = "TIMESTAMP"
	case types.Date:
		columnType = "DATE"
	case types.Time:
		columnType = "TIME"
	case types.Year:
		columnType = "YEAR"
	case types.Decimal:
		columnType = "DECIMAL"
	case types.Binary:
		columnType = "BINARY"
	case types.Blob:
		columnType = "BLOB"
	case types.Float:
		columnType = "FLOAT"
	case types.Text:
		columnType = "TEXT"

		if column.Length > 0 && column.Length <= 255 {
			columnType = "TINYTEXT"
		} else if column.Length > 255 && column.Length <= 65535 {
			columnType = "TEXT"
		} else if column.Length > 65535 && column.Length <= 16777215 {
			columnType = "MEDIUMTEXT"
		} else if column.Length > 16777215 {
			columnType = "LONGTEXT"
		}

	}
	length := ""
	if useLength {
		length = fmt.Sprintf("(%d)", column.Length)
	}

	return fmt.Sprintf(`%s%s`, columnType, length)
}

//ForeignKey ...
func (m MySQL) ForeignKey(key models.ForeignKey) string {
	return fmt.Sprintf(`ALTER TABLE %s ADD CONSTRAINT %s FOREIGN KEY (%s) REFERENCES %s (%s);`, key.LocalTable, key.Name, strings.Join(key.LocalColumns(), ","), key.ForeignTableName.Name, strings.Join(key.ForeignColumns(), ","))
}

//Index ...
func (m MySQL) Index(index models.Index, table models.Table) string {

	if index.IsPrimary {
		return fmt.Sprintf(`ALTER TABLE %s ADD PRIMARY KEY (%s);`, table.Name, strings.Join(index.ColumnsNames(), ","))
	} else if index.IsUnique {
		return fmt.Sprintf(`ALTER TABLE %s ADD UNIQUE (%s);`, table.Name, strings.Join(index.ColumnsNames(), ","))
	} else {
		return fmt.Sprintf(`CREATE INDEX %s ON %s (%s);`, index.Name, table.Name, strings.Join(index.ColumnsNames(), ","))
	}
}

//ColumnOptions ...
func (m MySQL) ColumnOptions(column models.Column) string {
	var options []string

	if column.AutoIncrement {
		options = append(options, "AUTO_INCREMENT")
	}
	if column.Default != "" {
		options = append(options, fmt.Sprintf("DEFAULT '%s'", column.Default))
	}
	if column.Comment != "" {
		options = append(options, fmt.Sprintf("COMMENT '%s'", column.Comment))
	}
	if column.NotNull {
		options = append(options, "NOT NULL")
	} else {
		options = append(options, "NULL")

	}
	return strings.Join(options, " ")
}

//Column ...
func (m MySQL) Column(column models.Column) string {
	var options []string

	options = []string{
		column.Name,
		m.Type(column),
		m.ColumnOptions(column),
	}

	return strings.Join(options, " ")
}

//Table ...
func (m MySQL) Table(table models.Table) string {

	var columns []string
	for _, column := range table.Columns {
		columns = append(columns, m.Column(column))
	}

	return fmt.Sprintf(`CREATE TABLE %s 
(
%s
);
`,
		table.Name,
		strings.Join(columns, ",\n"),
	)
}
