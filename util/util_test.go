package util

import (
	"testing"

	bl "github.com/sabramson2/baselayer-go"
)

func TestLoadFile1(t *testing.T) {
	config, e := bl.LoadConfig("TEST_GO_JSON")
	if e != nil {
		t.Errorf("error getting config: %s", e)
	}
	t.Logf("value a from config: %s", config["this"].(string))
}

func TestWriteFile(t *testing.T) {
	testData :=
`{
	"this": "is",
	"some": "test",
	"data": "."
}
`
	e := bl.WriteFile("/Users/stephen/tmp/test123.json", testData)
	if e != nil {
		t.Errorf("there was an error writing the file: %s", e)
	}
}

func TestReadFileToString(t *testing.T) {
	fileData, e := bl.ReadFileToString("/Users/stephen/tmp/test123.json")
	if e != nil { t.Errorf("error: %s", e); return }
	t.Log("here is the file data:")
	t.Log(fileData)
}

func TestReadFileToLines(t *testing.T) {
	lines, e := bl.ReadFileToLines("/Users/stephen/tmp/test123.json")
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
	r, _ := bl.StringToJson(s)
	t.Logf("a = %s\n", r["a"].(string))
	js, _ := bl.JsonToString(r)
	t.Logf("%s\n", js)
}

func TestStringToJsonArray(t *testing.T) {
	s := `[ "1", "2", "3" ]`
	r, _ := bl.StringToJsonArray(s)
	t.Logf("elem 0 = %s\n", r[0].(string))
	js, _ := bl.JsonToString(r)
	t.Logf("%s\n", js)
}

func TestArrays(t *testing.T) {
	s := "a b c"
	t.Logf("first = %s, second = %s, third = %s\n", bl.StringToSliceAny(s)...)
}

func TestRandNumString(t *testing.T) {
	num := bl.RandNumString(20)
	t.Logf("%s\n", num)
}

func TestFormatPrinting(t *testing.T) {
	bl.P("foo")
	bl.Pf("foo %s\n", "bar")
	bl.Pe(bl.Ef("some error %s", "errorfoo"))
}