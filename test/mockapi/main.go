package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
)

func header(c web.C, w http.ResponseWriter, r *http.Request) {
	key := c.URLParams["header"]
	id := r.Header.Get(key)
	if strings.ToLower(key) == "host" {
		id = r.Host
	}
	w.Header().Add(key, id)
	fmt.Fprintf(w, "%s=>%s", key, id)
}

func ping(c web.C, w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Host", r.Host)
	fmt.Fprintf(w, "pong")
}

func main() {
	goji.Get("/header/:header", header)
	goji.Get("/*", ping)
	goji.Serve()
}
