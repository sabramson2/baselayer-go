package net

import (
	"testing"
)

//----------------------------------------
func TestGet(t *testing.T) {
	req := Req {
		"https://jsonplaceholder.typicode.com/posts/1",
		"",
		nil,
	}
	r, e := Getj(&req)
	if e != nil { t.Errorf("error: %s", e); return }

	for k, v := range r.data {
		t.Logf("key: %s, value: %v\n", k, v)
	}
}

//----------------------------------------
func TestPost(t *testing.T) {
	req := Req {
		"https://jsonplaceholder.typicode.com/posts",
		`
		{
			"a": "val0",
			"b": "val1"
		}
		`,
		nil,
	}
	r, e := Postjj(&req)
	if e != nil { t.Errorf("error: %s", e); return }

	t.Logf("a = %s\n", r.data["a"].(string))
	t.Logf("id = %d\n", int(r.data["id"].(float64)))
}
