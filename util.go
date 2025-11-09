package baselayergo

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

//----------------------------------------
func UtilTest() string {
	return "hello from util package"
}

func Pe(e error) {
	fmt.Printf("!!! error: %s\n", e)
}

func P(s string) {
	fmt.Println(s)
}

func Pf(s string, args ...any) {
	//fmt.Printf(s, StringSliceToSliceAny(args)...)
	fmt.Printf(s, args...)
}

func E(s string) error {
	return fmt.Errorf(s, nil)
}

func Ef(s string, args ...string) error {
	return fmt.Errorf(s, StringSliceToSliceAny(args)...)
}

func Sf(s string, args ...string) string {
	return fmt.Sprintf(s, StringSliceToSliceAny(args)...)
}

//----------------------------------------
/*
given an env var holding the path to a json file, load that json file
and return it as a map[string]any
*/
func LoadConfig(envVarName string) (map[string]any, error) {
	filePath := os.Getenv(envVarName)
	if filePath == "" {
		return nil, fmt.Errorf("environment variable %s is not set or empty", envVarName)
	}

	config, e := ReadFileToJson(filePath)
	if e != nil { return nil, e }

	return config, nil
}

//----------------------------------------
func WriteFile(path string, contents string) error {
	return os.WriteFile(path, []byte(contents), 0644)
}

//----------------------------------------
func ReadFileToString(path string) (string, error) {
	data, e := os.ReadFile(path)
	if e != nil { return "", e }
	return string(data), nil
}

//----------------------------------------
func ReadFileToJson(path string) (map[string]any, error) {
	data, e := os.ReadFile(path)
	if e != nil { return nil, e }

	var obj map[string]any
	e = json.Unmarshal(data, &obj)
	if e != nil { return nil, e }

	return obj, nil
}

//----------------------------------------
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

//----------------------------------------
func JsonToString(o any) (string, error) {
	r, e := json.MarshalIndent(o, "", "    ")
	if e != nil { return "", e }
	return string(r), nil
}

//----------------------------------------
func StringToJson(s string) (map[string]any, error) {
	var r map[string]any
	e := json.Unmarshal([]byte(s), &r)
	if e != nil { return nil, e }
	return r, nil
}

//----------------------------------------
func StringToJsonArray(s string) ([]any, error) {
	var r []any
	e := json.Unmarshal([]byte(s), &r)
	if e != nil { return nil, e }
	return r, nil
}

//----------------------------------------
func StringToSliceAny(s string) []any {
	parts := strings.Split(s, " ")
	return StringSliceToSliceAny(parts)
}

//----------------------------------------
func StringSliceToSliceAny(s []string) []any {
	partsAny := make([]any, len(s))
	for i := range s {
		partsAny[i] = any(s[i])
	}
	return partsAny
}