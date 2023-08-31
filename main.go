package main

import (
	tool "GT/Tool"
	"fmt"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"os"
	"strings"
)

func readFile(filePath string) (string, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
func removeDuplicates(lines []string) []string {
	seen := make(map[string]bool)
	uniqueLines := []string{}

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" && !seen[line] {
			seen[line] = true
			uniqueLines = append(uniqueLines, line)
		}
	}
	return uniqueLines
}
func saveToFile(filePath string, content string) error {
	//content := strings.Join(lines, "\n")
	return os.WriteFile(filePath, []byte(content), 0644)
}
func countLines(filepath string) (int, error) {
	content, err := os.ReadFile(filepath)
	if err != nil {
		return 0, err
	}
	lines := strings.Split(string(content), "\n")
	return len(lines), nil
}
func main() {
	app := &cli.App{
		Name:    "GTool",
		Usage:   "Idong",
		Version: "0.0.1",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "file",
				Aliases:  []string{"f"},
				Usage:    "去重",
				Value:    "文件地址",
				Category: "文件操作：",
			},
			// &cli.StringFlag{
			// 	Name:    "OutPut",
			// 	Aliases: []string{"o"},
			// 	Usage:   "Save file",
			// },
		},
		Action: func(c *cli.Context) error {
			filePath := c.String("file")
			//检查是否提供了文件路径
			if filePath == "" {
				return fmt.Errorf("Please provide a valid file path using --file flag")
			}
			//统计原始行数
			originalLineCount, err := countLines(filePath)
			if err != nil {
				return fmt.Errorf("无法统计行数：%s", err)
			}
			//读取文件
			content, err := readFile(filePath)
			if err != nil {
				return fmt.Errorf("Error reading file: %s", err)
			}
			//去重复
			uniqueLines := removeDuplicates(strings.Split(content, "\n"))
			//统计去重后的行数
			uniqueLineCount := len(uniqueLines)
			//outputPath := c.String("OutPut")
			processedContent := strings.Join(uniqueLines, "\n")
			err = saveToFile(filePath, processedContent)
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
	tool.Test()
}
