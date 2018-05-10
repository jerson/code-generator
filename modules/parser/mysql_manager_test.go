package parser

import (
	"github.com/jerson/code-generator/modules/context"
	"github.com/jerson/code-generator/modules/dump"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"testing"
)

func TestNewMySQLManager(t *testing.T) {
	ctx := context.NewSingle("test")
	defer ctx.Close()

	parser, err := NewMySQLManager(ctx, "mysql", "root:123456@/movie_app")
	if err != nil {
		t.Fatal(err)
	}

	schema, err := parser.Schema()
	if err != nil {
		t.Fatal(err)
	}

	data, _ := dump.NewMySQL(ctx, *schema, dump.MySQLOptions{WithDrop: true, Alter: true}).Dump()
	//data, _ := json.MarshalIndent(schema, "", " ")
	t.Log(data)

}
