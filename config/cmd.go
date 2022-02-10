package config

import (
	"github.com/gourds/cdn_refresh/version"
	"github.com/urfave/cli"
	"github.com/wonderivan/logger"
	"os"
)

func cliConfig() {
	app := &cli.App{
		Name:    "cdn_refresh",
		Usage:   "refresh CDN",
		Version: version.GetVersion(),
		Author:  version.Author,
		Action: func(c *cli.Context) error { //该命令的执行动作函数
			return nil
		},
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:        "type,t",
				Value:       "Refresh",
				Usage:       "support{Refresh|Preheating}",
				Destination: &InputConfig.Type,
			},
			cli.StringFlag{
				Name:        "endpoint,e",
				Value:       "https://cdn.myhuaweicloud.com",
				Usage:       "endpoint address",
				Destination: &InputConfig.EndPoint,
			},
			cli.StringFlag{
				Name:        "accessKeyID,i",
				Value:       "",
				Usage:       "AccessKeyID",
				Destination: &InputConfig.AccessKeyID,
			},
			cli.StringFlag{
				Name:        "accessKeySecret,k",
				Value:       "",
				Usage:       "AccessKeySecret",
				Destination: &InputConfig.AccessKeySecret,
			},
			cli.StringFlag{
				Name:        "readFile,r",
				Value:       "urls.txt",
				Usage:       "urls file location",
				Destination: &InputConfig.ReadFile,
				Required:    true,
			},
			cli.StringFlag{
				Name:        "domainId,d",
				Value:       "",
				Usage:       "domainId",
				Destination: &InputConfig.DomainId,
			},
			cli.StringFlag{
				Name:        "ProjectId,p",
				Value:       "",
				Usage:       "ProjectID",
				Destination: &InputConfig.ProjectId,
			},
			cli.StringFlag{
				Name:		 "Platform",
				Value:       "HuaWei",
				Usage: 	     "Platform",
				Destination: &InputConfig.Platform,
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Debug(InputConfig)
}


func initConfig() {
	cliConfig()
}

