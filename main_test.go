package main

import (
	"fmt"
	"net/http"
	"os"
	"testing"
)

func Test_healthCheck(t *testing.T) {
	requestURL := "http://localhost:1234/_health"
	res, err := http.Get(requestURL)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		os.Exit(1)
	}
	if res.StatusCode != 200 {
		t.Errorf("no response status code : %v", res.StatusCode)
	}
	if res.StatusCode == 200 {
		fmt.Println("ok")
	}
}
