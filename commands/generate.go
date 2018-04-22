//Package commands ...
package commands

import (
	"github.com/jerson/code-generator/modules/config"
	"github.com/jerson/code-generator/modules/context"
	"github.com/urfave/cli"
)

//GenerateCommand ...
var GenerateCommand cli.Command

func init() {
	GenerateCommand = cli.Command{
		Name:     "generate",
		Category: "default",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "config, c",
				Usage: "accept: toml, yaml files",
			},
			cli.StringFlag{
				Name:  "output_dir, o",
				Value: "./output",
			},
			cli.StringFlag{
				Name:  "database_source, s",
				Usage: "example: user:password@/dbname?charset=utf8",
			},
			cli.StringFlag{
				Name:  "database_driver, d",
				Value: "mysql",
			},
			cli.StringFlag{
				Name:  "project_name, n",
				Value: "project sample",
			},
			cli.StringFlag{
				Name:  "project_package, p",
				Value: "github.com/account/project",
			},
			cli.StringFlag{
				Name:  "template_name, t",
				Value: "go",
			},
		},
		Action: func(c *cli.Context) error {
			ctx := context.NewSingle("command")
			defer ctx.Close()

			err := fillConfig(c)
			if err != nil {
				return err
			}

			log := ctx.GetLogger("test")
			log.Info(config.Vars)

			return nil
		},
	}

}

func fillConfig(c *cli.Context) error {
	configFile := c.String("config")
	if configFile != "" {
		err := config.Read(configFile)
		if err != nil {
			return err
		}
	}
	if value := c.String("output_dir"); value != "" {
		config.Vars.Output.Dir = value
	}
	if value := c.String("output_dir"); value != "" {
		config.Vars.Output.Dir = value
	}
	if value := c.String("database_source"); value != "" {
		config.Vars.Database.Source = value
	}
	if value := c.String("database_driver"); value != "" {
		config.Vars.Database.Driver = value
	}
	if value := c.String("project_name"); value != "" {
		config.Vars.Project.Name = value
	}
	if value := c.String("project_package"); value != "" {
		config.Vars.Project.Package = value
	}
	if value := c.String("template_name"); value != "" {
		config.Vars.Template.Name = value
	}

	return nil
}
