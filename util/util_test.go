package util

import (
	"testing"
)

func TestLoadFile1(t *testing.T) {
	config, e := LoadConfig("TEST_GO_JSON")
	if e != nil {
		t.Errorf("error getting config: %s", e)
	}
	t.Logf("value a from config: %s", config["a"].(string))
}

func TestWriteFile(t *testing.T) {
	testData :=
`{
	"this": "is",
	"some": "test",
	"data": "."
}
`
	e := WriteFile("/Users/stephen/tmp/test123.json", testData)
	if e != nil {
		t.Errorf("there was an error writing the file: %s", e)
	}
}

func TestReadFileToString(t *testing.T) {
	fileData, e := ReadFileToString("/Users/stephen/tmp/test123.json")
	if e != nil { t.Errorf("error: %s", e); return }
	t.Log("here is the file data:")
	t.Log(fileData)
}

func TestReadFileToLines(t *testing.T) {
	lines, e := ReadFileToLines("/Users/stephen/tmp/test123.json")
	if e != nil { t.Errorf("error: %s", e); return }
	t.Log("here is the file data:")
	for i := range(len(lines)) {
		t.Logf("line %d: %s\n", i, lines[i])
	}
}