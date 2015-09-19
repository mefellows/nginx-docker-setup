package main

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"testing"
)

// const API_HOST = "api.foo.local"
const API_HOST = "api.foo.com"
const API_PORT = 80
const NGINX_HOST = "nginx"
const NGINX_PORT = 80

func TestResponseHeaders(t *testing.T) {
	resp, err := http.Get(fmt.Sprintf("http://%s:%d/header/host", API_HOST, API_PORT))
	checkErr(err, false, t)

	if resp.StatusCode != 200 {
		t.Fatalf("Expected 200 response code, but got %d", resp.StatusCode)
	}
}

func TestIncomingRequestHeaders(t *testing.T) {
	cases := []struct {
		Header  string
		Regex   string
		Message string
	}{
		// Expect X-Request-Id => 279B8A31-A6DB-4BED-A4D4-530D961CF792
		{
			"X-Request-Id",
			"^(([A-Z0-9]+)-){4}([A-Z0-9]+)$",
			"Expected X-Request-Id to be set, and a valid UUID",
		},
		// Expect X-Real-Ip=>192.168.0.1
		{
			"X-Real-Ip",
			`^(([0-9]+)\.){3}([0-9]+)$`,
			"Expected X-Real-Ip to be set and a valid IP Address",
		},
		// Expect X-Forwarded-For
		{
			"X-Forwarded-For",
			`\.*`,
			"Expected X-Forwarded-For header to be set",
		},
		// Expect X-Forwarded
		{
			"X-Forwarded",
			`proto=https`,
			"Expected X-Forwarded header to be set and match `proto=https`",
		},
		// Expect Host
		{
			"Host",
			`\.*`,
			"Expected Host header to be set",
		},
	}
	for _, c := range cases {
		checkResponseHeader(c.Header, t,
			checkResponseHeaderRegexFunc(c.Regex, c.Message))
	}
}

func TestExistingUUIDHeader(t *testing.T) {

	// Set X-Request-Id=blah and check that doesn't get changed
	header := "X-Request-Id"
	expectedHeaderVal := "279B8A31-A6DB-4BED-A4D4-530D961CF792"
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("http://%s:%d/header/%s", API_HOST, API_PORT, header), nil)
	checkErr(err, false, t)
	req.Header.Add(header, expectedHeaderVal)
	resp, err := client.Do(req)
	checkErr(err, false, t)

	if resp.Header.Get(header) != expectedHeaderVal {

	}
}

func TestCheckReadTimeout(t *testing.T)    {}
func TestCheckConnectTimeout(t *testing.T) {}
func TestLogFormat(t *testing.T)           {}

func TestHealth(t *testing.T) {
	resp, err := http.Get(fmt.Sprintf("http://%s:%d/health", NGINX_HOST, NGINX_PORT))
	checkErr(err, false, t)

	if resp.StatusCode != 200 {
		t.Fatalf("Expected 200 response code, but got %d", resp.StatusCode)
	}

	_body := make([]byte, 9)
	resp.Body.Read(_body)
	body := strings.TrimSpace(string(_body))

	if body != strings.TrimSpace("health ok") {
		t.Fatalf("Expected body to be 'health ok' but got '%s'", body)
	}
}

//
// Test helper utilities
//

// Checks that a response mody is present
func checkResponseBody(url string, t *testing.T) string {
	resp, err := http.Get(fmt.Sprintf("http://%s:%d%s", API_HOST, API_PORT, url))
	checkErr(err, false, t)

	_body := make([]byte, 64)
	resp.Body.Read(_body)
	body := strings.TrimSpace(string(_body))

	if body == "" {
		t.Fatalf("Expected body to be present")
	}
	return body
}

// Checks that a response header is present, optionally running a user-defined validation function
// on the result
func checkResponseHeader(header string, t *testing.T, validate func(header string) (bool, error)) string {
	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("http://%s:%d/header/%s", API_HOST, API_PORT, header), nil)
	checkErr(err, false, t)
	req.Header.Add("Host", API_HOST)
	resp, err := client.Do(req)
	checkErr(err, false, t)

	h := resp.Header.Get(header)
	if h == "" {
		t.Fatalf("Expected header '%s' to be present", header)
	}

	// Pass the header to the validate func
	if validate != nil {
		if res, err := validate(h); err != nil || !res {
			t.Fatalf("Validation failed on header %s. Context/Error: %s", header, err)
		}
	}
	return h
}

func checkResponseHeaderRegexFunc(regex string, err string) func(string) (bool, error) {
	return func(h string) (bool, error) {
		res, e := regexp.MatchString(regex, h)
		if e != nil {
			return false, errors.New(err)
		}
		return res, nil
	}
}

func checkErr(err error, expected bool, t *testing.T) {
	if err != nil && !expected {
		t.Fatalf("Error not expected: %s", err.Error())

	} else if err == nil && expected {
		t.Fatalf("Error expected, but did not get one")
	}
}
