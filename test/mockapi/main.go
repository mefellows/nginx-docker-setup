package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"github.com/quipo/statsd"
	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
	"github.com/zenazn/goji/web/middleware"
)

var s *statsd.StatsdClient

func header(c web.C, w http.ResponseWriter, r *http.Request) {
	key := c.URLParams["header"]
	id := r.Header.Get(key)
	if strings.ToLower(key) == "host" {
		id = r.Host
	}
	w.Header().Add(key, id)
	w.Header().Add("X-Host", os.Getenv("X-HOST"))
	fmt.Fprintf(w, "%s=>%s", key, id)
	s.Incr(os.Getenv("X-HOST"), 1)
}

func ping(c web.C, w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Host", r.Host)
	w.Header().Add("X-Host", os.Getenv("X-HOST"))
	s.Incr(os.Getenv("X-HOST"), 1)
	fmt.Fprintf(w, "pong")
}

func main() {
	fmt.Println("Connecting to Statsd server...")

	prefix := "muxy."
	statsdclient := statsd.NewStatsdClient("statsd:8125", prefix)
	s = statsdclient
	s.Incr("test", 100)
	s.Gauge("something", 500)
	statsdclient.CreateSocket()
	//interval := time.Second * 2 // aggregate stats and flush every 2 seconds
	//s = statsd.NewStatsdBuffer(interval, statsdclient)

	if s == nil {
		fmt.Println("Could not connect to statsd server, exiting")
		os.Exit(1)
	}

	goji.Get("/header/:header", header)
	goji.Get("/*", ping)
	goji.Abandon(middleware.Logger)
	goji.Serve()
}
