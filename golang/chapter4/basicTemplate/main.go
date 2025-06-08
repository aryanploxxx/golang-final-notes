package main

import (
	"database/sql"
	"fmt"
	"html/template"
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

type Page struct {
	Title   string
	Content string
	Date    string
}

// http://localhost:8080/page/hello-world
// Hit the query on this url as we are using guid in the url and it will be mapped to the cooresponding page

func ServePage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pageGUID := vars["guid"]
	thisPage := Page{}
	fmt.Println(pageGUID)

	err := database.QueryRow("SELECT page_title,page_content,page_date FROM pages WHERE page_guid=?", pageGUID).Scan(&thisPage.Title, &thisPage.Content, &thisPage.Date)

	if err != nil {
		http.Error(w, http.StatusText(404), http.StatusNotFound) // This line is showing 'Not Found on the frontend'
		log.Println("Couldn't get page!")
		return
	}
	// html := `<html><head><title>` + thisPage.Title + `</title></head><body><h1>` + thisPage.Title + `</h1><div>` + thisPage.Content + `</div></body></html>`
	// fmt.Fprintln(w, html)

	t, _ := template.ParseFiles("templates/blog.html")
	t.Execute(w, thisPage)
	// thisPage is a struct populated with the data from the database and being passed to the template
}

func main() {
	dbConn := fmt.Sprintf("%s:%s@/%s", DBUser, DBPass, DBDbase)
	fmt.Println(dbConn)

	db, err := sql.Open("mysql", dbConn)

	if err != nil {
		log.Println("Couldn't connect!")
		log.Println(err.Error())
	}

	database = db

	routes := mux.NewRouter()
	routes.HandleFunc("/page/{guid:[0-9a-zA\\-]+}", ServePage)
	http.Handle("/", routes)
	http.ListenAndServe(PORT, nil)

}
