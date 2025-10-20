package util

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func UtilTest() string {
	return "hello from util package"
}

func LoadConfig(envVarName string) (map[string]any, error) {
	filePath := os.Getenv(envVarName)
	if filePath == "" {
		return nil, fmt.Errorf("environment variable %s is not set or empty", envVarName)
	}

	data, e := os.ReadFile(filePath)
	if e != nil {
		return nil, fmt.Errorf("failed to read file at path '%s': %w", filePath, e)
	}
	var config map[string]any
	e = json.Unmarshal(data, &config)
	if e != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON content: %w", e)
	}

	return config, nil
}

func WriteFile(path string, contents string) error {
	return os.WriteFile(path, []byte(contents), 0644)
}

func ReadFileToString(path string) (string, error) {
	data, e := os.ReadFile(path)
	if e != nil { return "", e }
	return string(data), nil
}

func ReadFileToLines(path string) ([]string, error) {
	file, e := os.Open(path)
	if e != nil { return nil, e }
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if e := scanner.Err(); e != nil {
		return nil, e
	}
	return lines, nil
}