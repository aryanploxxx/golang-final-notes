package main

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

// Define a type for the context key to avoid collisions
type contextKey string

const msgIDKey contextKey = "msgId"

func main() {
	helloWorldHandler := http.HandlerFunc(HelloWorld)
	http.Handle("/welcome", injectMsgID(helloWorldHandler))
	http.ListenAndServe(":8080", nil)
}

// HelloWorld is the handler that returns "Hello, world" with a msgId in the response header
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	msgID := ""
	if m := r.Context().Value(msgIDKey); m != nil {
		if value, ok := m.(string); ok {
			msgID = value
		}
	}
	w.Header().Add("msgId", msgID)
	w.Write([]byte("Hello, world"))
}

// injectMsgID is a middleware that adds a unique msgId to each request's context
func injectMsgID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		msgID := uuid.New().String()
		ctx := context.WithValue(r.Context(), msgIDKey, msgID)
		req := r.WithContext(ctx)
		next.ServeHTTP(w, req)
	})
}

// Expected Behavior:
// When a request is made to http://localhost:8080/welcome, the response includes a unique msgId in the headers.
// The injectMsgID middleware generates a new UUID for each request and passes it down via the request context.
// The HelloWorld handler retrieves the msgId from the context and adds it to the response headers.
