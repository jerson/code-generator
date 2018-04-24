package mysql

import (
	"encoding/json"
	"github.com/jerson/code-generator/modules/context"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"testing"
)

func TestNewPlatform(t *testing.T) {
	ctx := context.NewSingle("test")
	defer ctx.Close()

	parser, err := NewPlatform(ctx, "mysql", "root:123456@/movie_app")
	if err != nil {
		t.Fatal(err)
	}

	database, err := parser.Database()
	if err != nil {
		t.Fatal(err)
	}

	data, _ := json.MarshalIndent(database, "", " ")
	t.Log(string(data))

	columns, err := parser.Columns("movie_link", database)
	if err != nil {
		t.Fatal(err)
	}
	data, _ = json.MarshalIndent(columns, "", " ")
	t.Log(string(data))

	columns2, err := parser.Columns("movie_link", "")
	if err != nil {
		t.Fatal(err)
	}
	data, _ = json.MarshalIndent(columns2, "", " ")
	t.Log(string(data))

	indexes2, err := parser.Indexes("movie_link", database)
	if err != nil {
		t.Fatal(err)
	}
	data, _ = json.MarshalIndent(indexes2, "", " ")
	t.Log(string(data))

	indexes, err := parser.Indexes("movie_link", "")
	if err != nil {
		t.Fatal(err)
	}
	data, _ = json.MarshalIndent(indexes, "", " ")
	t.Log(string(data))

	foreignKeys, err := parser.ForeignKeys("movie_link", database)
	if err != nil {
		t.Fatal(err)
	}
	data, _ = json.MarshalIndent(foreignKeys, "", " ")
	t.Log(string(data))

	foreignKeys2, err := parser.ForeignKeys("movie_link", "")
	if err != nil {
		t.Fatal(err)
	}
	data, _ = json.MarshalIndent(foreignKeys2, "", " ")
	t.Log(string(data))

	views, err := parser.Views(database)
	if err != nil {
		t.Fatal(err)
	}
	data, _ = json.MarshalIndent(views, "", " ")
	t.Log(string(data))

	tables, err := parser.Tables()
	if err != nil {
		t.Fatal(err)
	}
	data, _ = json.MarshalIndent(tables, "", " ")
	t.Log(string(data))

}
