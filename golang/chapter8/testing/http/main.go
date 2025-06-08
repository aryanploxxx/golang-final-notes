package main

import (
	"fmt"
	"net/http"
)

// HelloWorldHandler returns a "Hello, World!" message
func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")
}

func main() {
	http.HandleFunc("/hello", HelloWorldHandler)
	http.ListenAndServe(":8080", nil)
}

// go test -v -> command to run test
// the file which contains the testing functions should have a name ending with _test.go
