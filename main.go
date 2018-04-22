//Package main ...
package main

import (
	"github.com/jerson/code-generator/commands"
	"github.com/jerson/code-generator/modules/config"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/urfave/cli"
	"os"
)

func init() {
	err := config.ReadDefault()
	if err != nil {
		panic(err)
	}
}

func main() {
	initCommands()
}

func initCommands() {
	app := cli.NewApp()
	app.Name = "code-generator"
	app.Usage = "generate code with single command"
	app.Version = config.Vars.Version

	app.Commands = []cli.Command{
		commands.GenerateCommand,
	}

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
