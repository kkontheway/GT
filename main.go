package main

import (
	tool "GT/Tool"
	"GT/config"
	"fmt"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
)

func main() {
	config.ShowBanner()
	app := &cli.App{
		Name:    "GTool",
		Usage:   "Idong",
		Version: "0.0.1",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "file",
				Aliases:  []string{"f"},
				Usage:    "去重",
				Value:    "",
				Category: "文件操作：",
			},
		},
		Action: func(c *cli.Context) error {
			filePath := c.String("file")
			//检查是否提供了文件路径
			if filePath == "" {
				return fmt.Errorf("Please provide a valid file path using --file flag")
			}
			//统计原始行数
			originalLineCount, err := tool.CountLines(filePath)
			if err != nil {
				return fmt.Errorf("无法统计行数：%s", err)
			}
			//读取文件
			content, err := tool.ReadFile(filePath)
			if err != nil {
				return fmt.Errorf("Error reading file: %s", err)
			}
			//去重复
			uniqueLines := tool.RemoveDuplicates(strings.Split(content, "\n"))
			//统计去重后的行数
			uniqueLineCount := len(uniqueLines)
			//outputPath := c.String("OutPut")
			processedContent := strings.Join(uniqueLines, "\n")
			err = tool.SaveToFile(filePath, processedContent)
			if err != nil {
				return fmt.Errorf("Error saving processed content to file: %s", err)
			}

			fmt.Printf("Processed content saved to %s\n", filePath)
			fmt.Println("去重前的行数是：", originalLineCount)
			fmt.Println("去重后的行数是：", uniqueLineCount)

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
