package main

import (
	tool "GT/Tool"
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {

	app := &cli.App{
		Name:    "GTool",
		Usage:   "Idong",
		Version: "0.0.1",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "duplicateFile",
				Aliases: []string{"df"},
				Usage:   "去重",
				Value:   " ",
				//Category: "文件操作",
				Required: false,
				Action: func(c *cli.Context, s string) error {
					return tool.Duplicate(c)
				},
			},
			&cli.BoolFlag{
				Name:    "env",
				Aliases: []string{"e"},
				//Value:    false,
				Usage: "读取环境变量",
				//Category: "敏感信息收集",
				Action: func(c *cli.Context, b bool) error {
					return tool.EnvSearch(c)
				},
			},
			&cli.StringFlag{
				Name:     "output",
				Usage:    "保存输出路径",
				Aliases:  []string{"o"},
				Required: false,
				//Category: "文件操作",
			},
		},
	}

	err := app.Run(os.Args)
	//config.ShowBanner()
	if err != nil {
		fmt.Println(err)
	}
}
