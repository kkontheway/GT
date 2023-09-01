package tool

import (
	"io/ioutil"
	"os"
	"strings"
)

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
func SaveToFile(filePath string, content string) error {
	//content := strings.Join(lines, "\n")
	return os.WriteFile(filePath, []byte(content), 0644)
}
func CountLines(filepath string) (int, error) {
	content, err := os.ReadFile(filepath)
	if err != nil {
		return 0, err
	}
	lines := strings.Split(string(content), "\n")
	return len(lines), nil
}
