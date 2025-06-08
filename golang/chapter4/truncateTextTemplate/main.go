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
var templates *template.Template

type Page struct {
	Title      string
	RawContent string
	Content    template.HTML
	Date       string
	GUID       string
}

func (p Page) TruncatedText() template.HTML {
	chars := 0
	for i := range p.Content {
		chars++
		if chars > 20 {
			return p.Content[:i]
		}
	}
	return p.Content
}

func ServePage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pageGUID := vars["guid"]
	thisPage := Page{}

	err := database.QueryRow("SELECT page_title, page_content, page_date FROM pages WHERE page_guid=?", pageGUID).Scan(&thisPage.Title, &thisPage.RawContent, &thisPage.Date)
	if err != nil {
		http.Error(w, http.StatusText(404), http.StatusNotFound)
		log.Println(err)
		return
	}

	thisPage.Content = template.HTML(thisPage.RawContent)
	thisPage.GUID = pageGUID

	err = templates.ExecuteTemplate(w, "single.html", thisPage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
	}
}

func HomePage(w http.ResponseWriter, r *http.Request) {
	var pages []Page

	rows, err := database.Query("SELECT page_title, page_content, page_date, page_guid FROM pages")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var p Page
		err = rows.Scan(&p.Title, &p.RawContent, &p.Date, &p.GUID)
		if err != nil {
			continue
		}
		p.Content = template.HTML(p.RawContent)
		pages = append(pages, p)
	}

	err = templates.ExecuteTemplate(w, "home.html", pages)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
	}
}

func main() {
	// Parse all templates at startup
	var err error
	templates, err = template.ParseFiles("home.html", "single.html")
	if err != nil {
		log.Fatal("Couldn't parse templates:", err)
	}

	dbConn := fmt.Sprintf("%s:%s@/%s", DBUser, DBPass, DBDbase)
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		log.Fatal("Couldn't connect to database:", err)
	}
	database = db
	defer db.Close()

	routes := mux.NewRouter()
	routes.HandleFunc("/", HomePage)
	routes.HandleFunc("/page/{guid:[0-9a-zA-Z\\-]+}", ServePage)

	http.Handle("/", routes)
	log.Printf("Server starting on port %s\n", PORT)
	log.Fatal(http.ListenAndServe(PORT, nil))
}
