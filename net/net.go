package net

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

//----------------------------------------
var Headersjj = map[string]string {
	"Content-Type": "application/json",
	"Accepts": "application/json",
}

var Headersj = map[string]string {
	"Accepts": "application/json",
}

//----------------------------------------
type Req struct {
	Url string
	Body string
	HeaderUpdater func(*http.Header)
}

type Response struct {
	Raw *http.Response
	Data map[string]any
}

type ArrayResponse struct {
	Raw *http.Response
	Data []any
}

type StringResponse struct {
	R *http.Response
	Data string
}

//----------------------------------------
// refactor post/get to allow for plugging in other response format processors

//----------------------------------------
func Post(url string, body string, headers map[string]string) (*StringResponse, error) {
	return DoRequest("POST", url, body, headers)
}

//----------------------------------------
func Get(url string, headers map[string]string) (*StringResponse, error) {
	return DoRequest("GET", url, "", headers)
}

func Put(url string, body string, headers map[string]string) (*StringResponse, error) {
	return DoRequest("POST", url, body, headers)
}

//----------------------------------------
func DoRequest(method string, url string, body string, headers map[string]string) (*StringResponse, error) {
	var payloadReader io.Reader
	payloadReader = nil
	if body != "" {
		payloadReader = bytes.NewBuffer([]byte(body))
	}
	client := http.Client {
		Timeout: 5 * time.Second,
	}
	req, e := http.NewRequest(method, url, payloadReader)
	if e != nil { return nil, e }

	for k,v := range headers {
		req.Header.Set(k, v)
	}

	r, e := client.Do(req)
	if e != nil { return nil, e }
	defer r.Body.Close()
	bodyBytes, e := io.ReadAll(r.Body)
	if e != nil { return nil, e }
	bodyString := string(bodyBytes)
	return &StringResponse{r, bodyString}, nil
}

//----------------------------------------
func Postjj(reqData *Req) (*Response, error) {
	payloadReader := bytes.NewBuffer([]byte(reqData.Body))
	client := http.Client {
		Timeout: 5 * time.Second,
	}
	req, _ := http.NewRequest("POST", reqData.Url, payloadReader)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accepts", "application/json")

	if reqData.HeaderUpdater != nil {
		reqData.HeaderUpdater(&req.Header)
	}

	r, e := client.Do(req)
	if e != nil { return nil, e }
	defer r.Body.Close()
	if !isOk(r) {
		handleNon200(r)
		return nil, errors.New("error, non 200 response")
	}

	result := bodyToMap(r)
	return &Response{r, result}, nil
}

//----------------------------------------
func Postja(reqData *Req) (*ArrayResponse, error) {
	payloadReader := bytes.NewBuffer([]byte(reqData.Body))
	client := http.Client {
		Timeout: 5 * time.Second,
	}
	req, _ := http.NewRequest("POST", reqData.Url, payloadReader)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accepts", "application/json")

	if reqData.HeaderUpdater != nil {
		reqData.HeaderUpdater(&req.Header)
	}

	r, e := client.Do(req)
	if e != nil { return nil, e }
	defer r.Body.Close()
	if !isOk(r) {
		handleNon200(r)
		return nil, errors.New("error, non 200 response")
	}

	result := bodyToList(r)
	return &ArrayResponse{r, result}, nil
}

//----------------------------------------
func Getj(reqData *Req) (*Response, error) {
	client := http.Client {
		Timeout: 5 * time.Second,
	}
	req, _ := http.NewRequest("GET", reqData.Url, nil)
	req.Header.Set("Accepts", "application/json")
	if reqData.HeaderUpdater != nil {
		reqData.HeaderUpdater(&req.Header)
	}
	r, e := client.Do(req)
	if e != nil { return nil, e }
	defer r.Body.Close()
	if !isOk(r) {
		handleNon200(r)
		return nil, fmt.Errorf("error, non 200 response: %d", r.StatusCode)
	}

	result := bodyToMap(r)
	return &Response{r, result}, nil
}

//----------------------------------------
func handleNon200(r *http.Response) {
	fmt.Printf("Response not ok - status = %d\n", r.StatusCode)
	bodyBytes, _ := io.ReadAll(r.Body)
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)
}

//----------------------------------------
func isOk(r *http.Response) bool {
	return r.StatusCode == http.StatusCreated || r.StatusCode == http.StatusOK
}

//----------------------------------------
func bodyToMap(r *http.Response) map[string]any {
	var result map[string]any
	json.NewDecoder(r.Body).Decode(&result)
	return result
}

//----------------------------------------
func bodyToList(r *http.Response) []any {
	var result []any
	json.NewDecoder(r.Body).Decode(&result)
	return result
}
