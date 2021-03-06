package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"github.com/zenazn/goji/web/middleware"
)

func header(c web.C, w http.ResponseWriter, r *http.Request) {
	key := c.URLParams["header"]
	id := r.Header.Get(key)
	if strings.ToLower(key) == "host" {
		id = r.Host
	}
	w.Header().Add(key, id)
	w.Header().Add("X-Host", os.Getenv("X-HOST"))
	fmt.Fprintf(w, "%s=>%s", key, id)
}

func ping(c web.C, w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Host", r.Host)
	w.Header().Add("X-Host", os.Getenv("X-HOST"))
	fmt.Fprintf(w, "pong")
}

func main() {
	goji.Get("/header/:header", header)
	goji.Get("/*", ping)
	goji.Abandon(middleware.Logger)
	goji.Serve()
}
