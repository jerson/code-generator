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
	sb.WriteString(fmt.Sprintf(`# schema: %s`, m.schema.Name))
	sb.WriteRune('\n')

	for _, table := range m.schema.Tables {
		sb.WriteString(fmt.Sprintf(`# table: %s`, table.Name))
		sb.WriteRune('\n')
		for _, index := range table.Indexes {
			sb.WriteString(fmt.Sprintf(`# index: %s`, index.Name))
			sb.WriteRune('\n')
		}
		for _, key := range table.ForeignKey {
			sb.WriteString(fmt.Sprintf(`# key: %s`, key.Name))
			sb.WriteRune('\n')
		}
	}
	for _, view := range m.schema.Views {
		sb.WriteString(fmt.Sprintf(`# view: %s`, view.Name))
		sb.WriteRune('\n')
	}

	return sb.String(), nil
}

//View ...
func (m MySQLDump) View(view models.View) (string, error) {

	return "", nil
}

//Table ...
func (m MySQLDump) Table(table models.Table) (string, error) {

	return "", nil
}
