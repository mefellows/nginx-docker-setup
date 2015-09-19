package main

import (
	"fmt"
	"net/http"
	"testing"
)

func TestFooProxyNoCookie(t *testing.T) {

	// Hit new bar.com homepage
	// No cookie     => should proxy old homepage (bar.com)
	// 'somecookie' cookie => should proxy new homepage (newbar.com)

	host := "myfandangledwebsite.com"
	origin := "bar.com"
	header := "host"

	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("http://%s:%d/header/%s", host, API_PORT, header), nil)
	checkErr(err, false, t)
	req.Header.Add("Host", host)
	resp, err := client.Do(req)
	checkErr(err, false, t)

	h := resp.Header.Get(header)
	if h != origin {
		t.Fatalf("Expected Host header to be %s, but got %s", origin, h)
	}
}
func TestFooProxyWithCookie(t *testing.T) {

	// Hit test new bar.com homepage
	// No cookie     => should proxy old homepage (bar.com)
	// 'somecookie' cookie => should proxy new homepage (newbar.com)

	host := "myfandangledwebsite.com"
	somecookie := "newbar.com"
	header := "host"

	client := &http.Client{}
	req, err := http.NewRequest("GET", fmt.Sprintf("http://%s:%d/header/%s", host, API_PORT, header), nil)
	checkErr(err, false, t)

	req.Header.Add("Host", host)
	req.AddCookie(&http.Cookie{Name: "somecookie", Value: "true"})
	resp, err := client.Do(req)
	checkErr(err, false, t)

	h := resp.Header.Get(header)
	if h != somecookie {
		t.Fatalf("Expected Host header to be %s, but got %s", somecookie, h)
	}
}
