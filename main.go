package main

import (
	"errors"
	"os"

	"github.com/codegangsta/cli"
	_ "github.com/shopspring/decimal"
)

var version = "1.0.0"
var soft_name = "model-auto"

/*
	./model_auto -import=github.com/model_auto/model -dbName=db_name -user=user_name -password=user_password -host=localhost -port=3306
*/
func main() {
	app := cli.NewApp()
	app.Name = soft_name
	app.Usage = "provide " + soft_name + " tools for koalaone@163.com "
	app.Version = version

	app.Flags = []cli.Flag{
		cli.StringFlag{Name: "import", Value: "", Usage: "model package import name", EnvVar: "APP_CONFIG"},
		cli.StringFlag{Name: "dbName", Value: "", Usage: "db name", EnvVar: "APP_CONFIG"},
		cli.StringFlag{Name: "user", Value: "", Usage: "db user name", EnvVar: "APP_CONFIG"},
		cli.StringFlag{Name: "password", Value: "", Usage: "db user password", EnvVar: "APP_CONFIG"},
		cli.StringFlag{Name: "host", Value: "", Usage: "db server host", EnvVar: "APP_CONFIG"},
		cli.StringFlag{Name: "port", Value: "", Usage: "db server port", EnvVar: "APP_CONFIG"},
		cli.StringFlag{Name: "tableName", Value: "", Usage: "db server port", EnvVar: "APP_CONFIG"},
	}

	app.Action = func(c *cli.Context) error {
		importName := c.GlobalString("import")
		if importName == "" {
			return errors.New("config package name is empty")
		}

		dbName := c.GlobalString("dbName")
		user := c.GlobalString("user")
		password := c.GlobalString("password")
		host := c.GlobalString("host")
		port := c.GlobalString("port")
		tableName := c.GlobalString("tableName")

		err := DBInit(dbName, user, password, host, port)
		if err != nil {
			return err
		}

		err = ModelGenerate(importName, tableName)
		if err != nil {
			return err
		}

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
		return
	}
}
