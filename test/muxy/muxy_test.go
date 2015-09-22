package main

import (
	"fmt"
	"net/http"
	"strings"
	"sync"
	"testing"
)

func Test100calls(t *testing.T) {
	host := fmt.Sprintf("http://api.foo.com/header/x-host")
	wait := &sync.WaitGroup{}
	wait.Add(100)
	for i := 0; i < 100; i++ {
		go func() {
			defer wait.Done()
			resp, err := http.Get(host)
			fmt.Println(resp)
			checkErr(err, false, t)

			if resp.StatusCode != 200 {
				t.Fatalf("Expected 200 response code, but got %d", resp.StatusCode)
			}

			_body := make([]byte, 9)
			resp.Body.Read(_body)
			body := strings.TrimSpace(string(_body))
		}()
	}
	fmt.Println("Waiting for all requests to finish...")
	wait.Wait()
}

func checkErr(err error, expected bool, t *testing.T) {
	if err != nil && !expected {
		t.Fatalf("Error not expected: %s", err.Error())

	} else if err == nil && expected {
		t.Fatalf("Error expected, but did not get one")
	}
}
