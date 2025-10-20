package util

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func UtilTest() string {
	return "hello from util package"
}

func LoadConfig(envVarName string) (map[string]any, error) {
	// 1. Get the value of the environment variable
	filePath := os.Getenv(envVarName)
	if filePath == "" {
		return nil, fmt.Errorf("environment variable %s is not set or empty", envVarName)
	}
	log.Printf("Found config file path: %s", filePath)

	// 2. Open and read the file content
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file at path '%s': %w", filePath, err)
	}
	log.Printf("Successfully read %d bytes from file.", len(data))

	// 3. Parse the JSON content into a map[string]any
	var config map[string]any // Changed to use 'any'
	// Use json.Unmarshal to decode the JSON data
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON content: %w", err)
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