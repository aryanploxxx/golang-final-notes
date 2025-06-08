package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

const (
	DBHost  = "127.0.0.1"
	DBPort  = ":3306"
	DBUser  = "root"
	DBPass  = "Pyari@123"
	DBDbase = "mydb"
	PORT    = ":8080"
)

var database *sql.DB

// global variable database is defined to store the reference to the open database connection.

type Page struct {
	Title   string
	Content string
	Date    string
}

func ServePage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)  // Extracts parameters from the URL
	pageID := vars["id"] // The "id" parameter from the URL (this will be the GUID)
	thisPage := Page{}   // Initialize an empty Page struct
	fmt.Println(pageID)

	// Query the database for the page using the GUID
	err := database.QueryRow("SELECT page_title,page_content,page_date FROM pages WHERE id=?", pageID).Scan(&thisPage.Title, &thisPage.Content, &thisPage.Date)

	// err := database.QueryRow("SELECT page_title,page_content,page_date FROM pages WHERE id=?", pageID)
	// 		.Scan(&thisPage.Title, &thisPage.Content, &thisPage.Date)
	// Scan: Maps the result of the query (title, content, and date) to the struct fields (Title, Content, and Date).

	if err != nil {
		log.Println(err)
		log.Println("Couldn't get page!")
	}

	// Generate an HTML response
	html := `<html><head><title>` + thisPage.Title + `</title></head><body><h1>` + thisPage.Title + `</h1><div>` + thisPage.Content + `</div></body></html>`

	fmt.Fprintln(w, html)
	// The function fmt.Fprintln in Go is used to write formatted output to an io.Writer, such as a file, a network connection, or any other writable output stream
}

func main() {
	dbConn := fmt.Sprintf("%s:%s@/%s", DBUser, DBPass, DBDbase) // Sprintf only returns the formatted string, it does not print it.
	// The sql.Open function opens a connection to the database using the mysql driver.
	fmt.Println(dbConn)

	db, err := sql.Open("mysql", dbConn)
	// sql.Open function is used to connect to the MySQL database. The first argument is the driver name, and the second argument is the data source name.
	if err != nil {
		log.Println("Couldn't connect!")
		log.Println(err.Error())
	}

	database = db

	routes := mux.NewRouter()
	routes.HandleFunc("/page/{id:[0-9]+}", ServePage)
	http.Handle("/", routes)
	http.ListenAndServe(PORT, nil)

}
