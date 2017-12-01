package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/Juniper/contrail-common/go/pkg/util"
	"github.com/urfave/cli"
)

func generateCode(c *cli.Context) {
	schemasDir := c.String("schema-dir")
	templateConfPath := c.String("template-config")
	api, err := util.MakeAPI(schemasDir)
	if err != nil {
		log.Fatal(err)
	}
	templateConf, err := util.LoadTemplates(templateConfPath)
	if err != nil {
		log.Fatal(err)
	}
	err = util.ApplyTemplates(api, filepath.Dir(templateConfPath), templateConf)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	app := cli.NewApp()
	app.Commands = []cli.Command{
		{
			Name:  "generate",
			Usage: "Generate Source code",
			Flags: []cli.Flag{
				cli.StringFlag{Name: "schema-dir", Value: "", Usage: "Schema Output dir"},
				cli.StringFlag{Name: "template-config", Value: "", Usage: "Template Config"},
			},
			Action: generateCode,
		},
	}
	app.Run(os.Args)
}
