package tool

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"os"
	"strings"
)

func Duplicate(c *cli.Context) error {
	//处理去重
	filePath := c.String("duplicateFile")
	//检查是否提供了文件路径
	if filePath == "" {
		return fmt.Errorf("Please provide a valid file path using --file flag")
	}
	//统计原始行数
	originalLineCount, err := CountLines(filePath)
	if err != nil {
		return fmt.Errorf("无法统计行数：%s", err)
	}
	//读取文件
	content, err := ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("Error reading file: %s", err)
	}
	//去重复
	uniqueLines := RemoveDuplicates(strings.Split(content, "\n"))
	//统计去重后的行数
	uniqueLineCount := len(uniqueLines)
	//outputPath := c.String("OutPut")
	processedContent := strings.Join(uniqueLines, "\n")
	outputPath := c.String("output")
	if outputPath == "" {
		outputPath = filePath
	}
	err = SaveToFile(outputPath, processedContent)
	if err != nil {
		return fmt.Errorf("Error saving processed content to file: %s", err)
	}
	fmt.Printf("Processed content saved to %s\n", filePath)
	fmt.Println("去重前的行数是：", originalLineCount)
	fmt.Println("去重后的行数是：", uniqueLineCount)
	return nil
}
func ReadFile(filePath string) (string, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}
func RemoveDuplicates(lines []string) []string {
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

func CountLines(filepath string) (int, error) {
	content, err := os.ReadFile(filepath)
	if err != nil {
		return 0, err
	}
	lines := strings.Split(string(content), "\n")
	return len(lines), nil
}
