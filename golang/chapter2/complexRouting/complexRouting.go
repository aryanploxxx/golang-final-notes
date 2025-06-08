package complexRouting

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

const (
	PORT = ":8080"
)

func pageHandler(w http.ResponseWriter, r *http.Request) {
	// Extracting dynamic parameters from the URL
	vars := mux.Vars(r)
	// mux.Vars(r) extracts route variables into a map
	// {id} -> The parameter name (id) is extracted using mux.Vars().
	// [0-9]+ -> A regular expression constraint ensures only numeric IDs are accepted.
	pageID := vars["id"] // extracting value of the corresponding key from the map
	// vars["id"] fetches the id parameter from the URL
	fileName := "files/" + pageID + ".html"
	// Dynamically constructs the file path based on the id parameter.
	// files/1.html, files/2.html, files/3.html, etc.
	// these files path are with respect to the main.go file, not the complexRouting.go package as ultimately we are running the main.go file

	_, err := os.Stat(fileName)
	if err != nil {
		fileName = "files/404.html"
	}

	// Serve the requested file
	http.ServeFile(w, r, fileName)
}

func ComplexRouting() {
	rtr := mux.NewRouter()

	// Define routes
	rtr.HandleFunc("/pages/{id:[0-9]+}", pageHandler) // Accepts only numeric IDs

	// Attach the router to the default HTTP server
	http.Handle("/", rtr)

	// Start the server
	log.Println("Server started at http://localhost" + PORT)
	http.ListenAndServe(PORT, nil)

}
