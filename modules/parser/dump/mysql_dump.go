package dump

import (
	"fmt"
	"github.com/jerson/code-generator/modules/context"
	"github.com/jerson/code-generator/modules/parser/models"
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

//Column ...
func (m MySQLDump) Column(column models.Column) string {
	return fmt.Sprintf(`%s %s`, column.Name, column.Type)
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
