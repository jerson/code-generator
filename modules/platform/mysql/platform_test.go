package mysql

import (
	"github.com/jerson/code-generator/modules/context"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"testing"
)

func TestMySQL_Columns(t *testing.T) {
	ctx := context.NewSingle("test")
	defer ctx.Close()

	parser, err := NewPlatform(ctx, "mysql", "root:@/movies")
	if err != nil {
		t.Fatal(err)
	}

	database, err := parser.Database()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(database)

	columns, err := parser.Columns("media_link_watch", database)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(columns)

	columns2, err := parser.Columns("media_link_watch", "")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(columns2)

	indexes2, err := parser.Indexes("media_link_watch", database)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(indexes2)

	indexes, err := parser.Indexes("media_link_watch", "")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(indexes)

	foreignKeys, err := parser.ForeignKeys("media_link_watch", database)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(foreignKeys)

	foreignKeys2, err := parser.ForeignKeys("media_link_watch", "")
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
