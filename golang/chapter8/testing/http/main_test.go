package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloWorldHandler(t *testing.T) {
	// t *testing.T: A pointer to the testing.T type, which provides methods for reporting errors and failures during the test.

	// Create a new request to the /hello route
	req, err := http.NewRequest("GET", "/hello", nil)
	/*
		http.NewRequest: Creates a new HTTP request.
		Method ("GET"): Specifies the HTTP method (GET, POST, etc.) for the request.
		URL ("/hello"): The URL path to test.
		Body (nil): The request body (e.g., for POST/PUT requests). Here, it's nil because GET requests usually don't have a body.
		req: Holds the created HTTP request.
	*/

	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	// Record the response
	rec := httptest.NewRecorder()
	// httptest.NewRecorder: Creates an in-memory http.ResponseWriter to record the HTTP handler's response.
	// rec: Captures the handler's response (status code, headers, and body).

	http.HandlerFunc(HelloWorldHandler).ServeHTTP(rec, req)
	// http.HandlerFunc: Wraps the HelloWorldHandler function into an http.Handler type, making it compatible with ServeHTTP.
	// ServeHTTP: Executes the HTTP handler (HelloWorldHandler) with the provided request (req) and response recorder (rec).

	// Validate the response status code
	if rec.Code != http.StatusOK {
		t.Errorf("Expected status OK, got %v", rec.Code)
		// t.Fatalf: Logs the error and stops the test immediately.
	}

	// Validate the response body
	expectedBody := "Hello, World!\n"
	if rec.Body.String() != expectedBody {
		t.Errorf("Expected body %q, got %q", expectedBody, rec.Body.String())
	}
}

// HTTP Route Testing (HelloWorldHandler):
// - Simulates an HTTP GET request to the /hello endpoint.
// - Uses httptest.NewRecorder to capture the response.
// - Validates the HTTP status code and the response body.

/*
	Create a GET request for the /hello route.
	Pass the request to the HelloWorldHandler.
	HelloWorldHandler is expected to:
	Respond with 200 OK as the status code.
	Respond with "Hello, World!\n" as the body.
	If either the status code or body doesn’t match expectations, an error is logged.
*/
