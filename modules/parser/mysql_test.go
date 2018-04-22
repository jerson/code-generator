package parser

import (
	"github.com/jerson/code-generator/modules/context"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"testing"
)

func init() {
}

func TestMySQL_Columns(t *testing.T) {
	ctx := context.NewSingle("test")
	defer ctx.Close()

	parser, err := NewMySQL(ctx, "mysql", "root:123456@/setbeat")
	if err != nil {
		t.Fatal(err)
	}

	database, err := parser.Database()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(database)

	columns, err := parser.Columns("album_song", database)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(columns)

	columns2, err := parser.Columns("album_song", "")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(columns2)

	indexes2, err := parser.Indexes("album_song", database)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(indexes2)

	indexes, err := parser.Indexes("album_song", "")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(indexes)

	foreignKeys, err := parser.ForeignKeys("album_song", database)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(foreignKeys)

	foreignKeys2, err := parser.ForeignKeys("album_song", "")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(foreignKeys2)

	views, err := parser.Views(database)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(views)

	tables, err := parser.Tables()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tables)

}
