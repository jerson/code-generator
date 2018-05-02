package dump

import (
	"fmt"
	"github.com/jerson/code-generator/modules/context"
	"github.com/jerson/code-generator/modules/parser/models"
	"github.com/jerson/code-generator/modules/parser/types"
	"strings"
)

// MySQLDump ...
type MySQLDump struct {
	ctx    context.Base
	schema models.Schema
}

// NewMySQLDump ...
func NewMySQLDump(ctx context.Base, schema models.Schema) MySQLDump {
	return MySQLDump{ctx: ctx, schema: schema}
}

//Dump ...
func (m MySQLDump) Dump() (string, error) {

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf(`-- database: %s`, m.schema.Name))
	sb.WriteRune('\n')

	for _, table := range m.schema.Tables {
		sb.WriteString(m.Table(table))
		sb.WriteRune('\n')
		for _, index := range table.Indexes {
			sb.WriteString(fmt.Sprintf(`# index: %s`, index.Name))
			sb.WriteRune('\n')
		}
	}
	for _, view := range m.schema.Views {
		sb.WriteString(m.View(view))
		sb.WriteRune('\n')
	}

	return sb.String(), nil
}

//Quoted ...
func (m MySQLDump) Quoted(name string) string {
	return fmt.Sprintf("´%s´", name)
}

//View ...
func (m MySQLDump) View(view models.View) string {
	return fmt.Sprintf(`CREATE VIEW %s AS %s`, view.Name, view.SQL)
}

//ForeignKey ...
func (m MySQLDump) ForeignKey(key models.ForeignKey) string {
	return fmt.Sprintf(`FOREIGN KEY (%s) REFERENCES %s (%s)`, key.LocalTable, key.Namespace)
}

//Index ...
func (m MySQLDump) Index(index models.Index, table models.Table) string {
	return fmt.Sprintf(`CREATE INDEX %s ON %s (%s);`, index.Name, table.Name, strings.Join(index.ColumnsNames(), ","))
}

//Type ...
func (m MySQLDump) Type(column models.Column) string {
	columnType := "TEXT"
	useLength := false
	switch column.Type {
	case types.String:
	case types.Array:
		useLength = true
		columnType = "VARCHAR"
		break
	case types.Integer:
		useLength = true
		columnType = "INT"
		break
	case types.BigInt:
		useLength = true
		columnType = "BIGINT"
		break
	case types.SmallInt:
	case types.Boolean:
		useLength = true
		columnType = "TINYINT"
		break
	case types.Datetime:
		useLength = false
		columnType = "DATETIME"
		break
	case types.Timestamp:
		useLength = false
		columnType = "TIMESTAMP"
		break
	case types.Text:
		useLength = false
		columnType = "TEXT"
		break

	}
	length := ""
	if useLength {
		length = fmt.Sprintf("(%d)", column.Length)
	}

	return fmt.Sprintf(`%s%s`, columnType, length)
}

//Column ...
func (m MySQLDump) Column(column models.Column) string {
	return fmt.Sprintf(`%s %s`, column.Name, m.Type(column))
}

//Table ...
func (m MySQLDump) Table(table models.Table) string {

	var columns []string
	var indexes []string

	for _, column := range table.Columns {
		columns = append(columns, m.Column(column))
	}
	for _, foreignKey := range table.ForeignKey {
		columns = append(columns, m.ForeignKey(foreignKey))
	}
	for _, index := range table.Indexes {
		indexes = append(indexes, m.Index(index, table))
	}

	return fmt.Sprintf(`CREATE TABLE %s 
(
%s
);

%s
`,
		table.Name,
		strings.Join(columns, ",\n"),
		strings.Join(indexes, "\n"),
	)
}
