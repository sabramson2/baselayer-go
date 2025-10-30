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

func TestStringToJson(t *testing.T) {
	s := `{
		"a": "foo",
		"b": "bar",
		"c": [
			"1", "2", "3"
		]
	}`
	r, _ := StringToJson(s)
	t.Logf("a = %s\n", r["a"].(string))
	js, _ := JsonToString(r)
	t.Logf("%s\n", js)
}

func TestStringToJsonArray(t *testing.T) {
	s := `[ "1", "2", "3" ]`
	r, _ := StringToJsonArray(s)
	t.Logf("elem 0 = %s\n", r[0].(string))
	js, _ := JsonToString(r)
	t.Logf("%s\n", js)
}

func TestArrays(t *testing.T) {
	s := "a b c"
	t.Logf("first = %s, second = %s, third = %s\n", StringToSliceAny(s)...)
}

func TestRandNumString(t *testing.T) {
	num := RandNumString(20)
	t.Logf("%s\n", num)
}