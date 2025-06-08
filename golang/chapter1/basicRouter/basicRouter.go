package basicRouter

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func TestHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Test endpoint is working!"))
}

func BasicRouter(x *mux.Router) {
	x.HandleFunc("/test", TestHandler)
	http.Handle("/", x)
	// http.ListenAndServe(":8080", nil)
	http.ListenAndServe(":8080", nil)
	fmt.Println("Everything is set up!")
}
