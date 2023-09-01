package tool

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
	"strings"
)

func EnvSearch(c *cli.Context) error {
	outputPath := c.String("output")
	envVariables := os.Environ()
	outputContent := strings.Join(envVariables, "\n")
	err := SaveToFile(outputPath, outputContent)
	if err != nil {
		return fmt.Errorf("Error saving environment variables to file: %s", err)
	}
	fmt.Printf("Environment variables saved to %s\n", outputPath)
	return nil
	//处理环境变量
	//if c.Bool("env") {
	//	fmt.Println("环境变量：")
	//
	//	for _, envVar := range envVariables {
	//		fmt.Println(envVar)
	//	}
	//}
	return nil
}
