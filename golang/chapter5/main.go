package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// Keep the same struct definitions...

const (
	DBHost  = "localhost"
	DBPort  = ":5432"
	DBUser  = "postgres"
	DBPass  = "Pyari@123"
	DBDbase = "mydb"
	PORT    = "8080"
)

// Struct definitions
type Page struct {
	Id         int
	Title      string
	RawContent string
	Content    template.HTML
	Date       string
	Comments   []Comment
	GUID       string
}

type Comment struct {
	Id          int
	Name        string
	Email       string
	CommentText string
}

type JSONResponse struct {
	Fields map[string]string
}

var database *sql.DB

// http://localhost:8080/page/hello-world -> Access this URL

func main() {
	// Fix database connection initialization
	var err error
	dbURL := "host=localhost port=5432 user=postgres password=Pyari@123 sslmode=disable dbname=mydb"
	database, err = sql.Open("postgres", dbURL) // Remove := to use global database variable
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	// Test the connection
	err = database.Ping()
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	// Initialize router
	routes := mux.NewRouter()

	// API Routes
	routes.HandleFunc("/api/pages", APIPage).
		Methods("GET") // Remove Schemes("https") since we're not using TLS

	routes.HandleFunc("/api/pages/{guid:[0-9a-zA\\-]+}", APIPage).
		Methods("GET") // Remove Schemes("https")

	routes.HandleFunc("/api/comments", APICommentPost).
		Methods("POST")

	routes.HandleFunc("/api/comments/{id:[\\w\\d\\-]+}", APICommentPut).
		Methods("PUT")

	// Web Routes
	routes.HandleFunc("/page/{guid:[0-9a-zA\\-]+}", ServePage)

	log.Printf("Server starting on port %s", PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, routes))
}

func ServePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ServePage")
	vars := mux.Vars(r)
	pageGUID := vars["guid"]
	thisPage := Page{}

	// Modified query for PostgreSQL
	err := database.QueryRow("SELECT id, page_title, page_content, page_date FROM pages WHERE page_guid=$1",
		pageGUID).Scan(&thisPage.Id, &thisPage.Title, &thisPage.RawContent, &thisPage.Date)
	if err != nil {
		http.Error(w, http.StatusText(404), http.StatusNotFound)
		log.Println("Database error:", err)
		return
	}
	thisPage.Content = template.HTML(thisPage.RawContent)

	// Modified query for PostgreSQL
	comments, err := database.Query("SELECT id, comment_name, comment_email, comment_text FROM comments WHERE page_id=$1",
		thisPage.Id)
	if err != nil {
		log.Println("Comments query error:", err)
	}
	defer comments.Close()

	thisPage.Comments = []Comment{} // Initialize the slice
	for comments.Next() {
		var comment Comment
		err := comments.Scan(&comment.Id, &comment.Name, &comment.Email, &comment.CommentText)
		if err != nil {
			log.Println("Error scanning comment:", err)
			continue
		}
		thisPage.Comments = append(thisPage.Comments, comment)
	}

	// Render template
	t, err := template.ParseFiles("blog.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println("Template error:", err)
		return
	}
	err = t.Execute(w, thisPage)
	if err != nil {
		log.Println("Template execution error:", err)
	}
}

// Update APIPage for PostgreSQL
func APIPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pageGUID := vars["guid"]
	thisPage := Page{}

	err := database.QueryRow("SELECT page_title, page_content, page_date FROM pages WHERE page_guid=$1",
		pageGUID).Scan(&thisPage.Title, &thisPage.RawContent, &thisPage.Date)
	if err != nil {
		http.Error(w, http.StatusText(404), http.StatusNotFound)
		log.Println(err)
		return
	}
	thisPage.Content = template.HTML(thisPage.RawContent)

	APIOutput, err := json.Marshal(thisPage)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, string(APIOutput))
}

// Update APICommentPost for PostgreSQL
func APICommentPost(w http.ResponseWriter, r *http.Request) {
	var commentAdded bool
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}

	name := r.FormValue("name")
	email := r.FormValue("email")
	comments := r.FormValue("comments")
	pageID := r.FormValue("guid")

	var id int64
	err = database.QueryRow("INSERT INTO comments (comment_name, comment_email, comment_text, page_id) VALUES ($1, $2, $3, $4) RETURNING id",
		name, email, comments, pageID).Scan(&id)
	if err != nil {
		log.Println(err)
		commentAdded = false
	} else {
		commentAdded = true
	}

	resp := JSONResponse{
		Fields: map[string]string{
			"id":    strconv.FormatInt(id, 10),
			"added": strconv.FormatBool(commentAdded),
		},
	}

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, string(jsonResp))
}

// Update APICommentPut for PostgreSQL
func APICommentPut(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}

	vars := mux.Vars(r)
	id := vars["id"]
	name := r.FormValue("name")
	email := r.FormValue("email")
	comments := r.FormValue("comments")

	_, err = database.Exec("UPDATE comments SET comment_name=$1, comment_email=$2, comment_text=$3 WHERE id=$4",
		name, email, comments, id)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := JSONResponse{
		Fields: map[string]string{
			"updated": "true",
		},
	}

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, string(jsonResp))
}
