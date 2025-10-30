package net

import (
	"testing"

	u "github.com/sabramson2/baselayer-go/util"
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

	t.Log(r)
	for k, v := range r.Data {
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

	t.Logf("a = %s\n", r.Data["a"].(string))
	t.Logf("id = %d\n", int(r.Data["id"].(float64)))
}

//----------------------------------------
func TestPostString(t *testing.T) {
	url := "https://jsonplaceholder.typicode.com/posts"
	body := `{
		"id": 1,
		"title", "title2",
		"body", "body2",
		"userId": 102
	}`

	r, e := Post(url, body, Headersjj)
	if e != nil { u.Pe(e); return }

	t.Logf("response status: %d %s\n", r.R.StatusCode, r.R.Status)
	t.Logf("response body: %s\n", r.Data)
}

func TestPutString(t *testing.T) {
	url := "https://jsonplaceholder.typicode.com/posts/1"
	body := `{
		"id": 1,
		"title": "title2",
		"body": "body2",
		"userId": 102
	}`

	r, e := Put(url, body, Headersjj)
	if e != nil { u.Pe(e); return }

	t.Logf("response status: %d %s\n", r.R.StatusCode, r.R.Status)
	t.Logf("response body: %s\n", r.Data)
}

func TestGetString(t *testing.T) {
	url := "https://jsonplaceholder.typicode.com/posts/1"
	r, e := Get(url, Headersj)
	if e != nil { u.Pe(e); return }

	t.Logf("response status: %d %s\n", r.R.StatusCode, r.R.Status)
	t.Logf("response body: %s\n", r.Data)
}
