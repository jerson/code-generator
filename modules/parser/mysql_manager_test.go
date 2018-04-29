package parser

import (
	"github.com/BurntSushi/toml"
	"github.com/jerson/code-generator/modules/context"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
	"testing"
)

func TestNewMySQLManager(t *testing.T) {
	ctx := context.NewSingle("test")
	defer ctx.Close()

	parser, err := NewMySQLManager(ctx, "mysql", "root:123456@/setbeat")
	if err != nil {
		t.Fatal(err)
	}

	schema, err := parser.Schema()
	if err != nil {
		t.Fatal(err)
	}

	toml.NewEncoder(os.Stdout).Encode(schema)
	//data, _ := json.MarshalIndent(schema, "", " ")
	//t.Log(string(data))

}
