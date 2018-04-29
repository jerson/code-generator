package parser

import (
	"github.com/jerson/code-generator/modules/context"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gopkg.in/yaml.v2"
	"os"
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

	yaml.NewEncoder(os.Stdout).Encode(schema)
	//data, _ := json.MarshalIndent(schema, "", " ")
	//t.Log(string(data))

}
