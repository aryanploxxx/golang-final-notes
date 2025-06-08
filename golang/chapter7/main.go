package main

import (
	"bytes"
	"crypto/rand"
	"crypto/sha1"
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	_ "github.com/lib/pq"
	"github.com/streadway/amqp"
)

const (
	DBHost  = "localhost"
	DBPort  = ":5432"
	DBUser  = "postgres"
	DBPass  = "Pyari@123"
	DBDbase = "mydb"
	PORT    = ":8080"
	MQHost  = "127.0.0.1"
	MQPort  = ":5672"
)

var database *sql.DB
var sessionStore = sessions.NewCookieStore([]byte("our-social-network-application"))
var UserSession Session

var WelcomeTitle = "You've succcessfully registered!"
var WelcomeEmail = "Welcome to our CMS, {{Email}}!  We're glad you could join us."

type RegistrationData struct {
	Email   string `json:"email"`
	Message string `json:"message"`
}

type Comment struct {
	Id          int
	Name        string
	Email       string
	CommentText string
}

type Page struct {
	Id         int
	Title      string
	RawContent string
	Content    template.HTML
	Date       string
	Comments   []Comment
	Session    Session
	GUID       string
}

type User struct {
	Id   int
	Name string
}

type Session struct {
	Id              string
	Authenticated   bool
	Unauthenticated bool
	User            User
}

type JSONResponse struct {
	Fields map[string]string
}

func getSessionUID(sid string) int {
	user := User{}
	err := database.QueryRow("SELECT user_id FROM sessions WHERE session_id=$1", sid).Scan(user.Id)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}
	return user.Id
}

func updateSession(sid string, uid int) {
	const timeFmt = "2006-01-02T15:04:05.999999999"
	tstamp := time.Now().Format(timeFmt)
	_, err := database.Exec("INSERT INTO sessions SET session_id=$1, user_id=$2, session_update=$3 ON DUPLICATE KEY UPDATE user_id=$4, session_update=$5", sid, uid, tstamp, uid, tstamp)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func generateSessionId() string {
	sid := make([]byte, 24)
	_, err := io.ReadFull(rand.Reader, sid)
	if err != nil {
		log.Fatal("Could not generate session id")
	}
	return base64.URLEncoding.EncodeToString(sid)
}

func validateSession(w http.ResponseWriter, r *http.Request) {
	session, _ := sessionStore.Get(r, "app-session")
	if sid, valid := session.Values["sid"]; valid {
		currentUID := getSessionUID(sid.(string))
		updateSession(sid.(string), currentUID)
		UserSession.Id = strconv.Itoa(currentUID)
	} else {
		newSID := generateSessionId()
		session.Values["sid"] = newSID
		session.Save(r, w)
		UserSession.Id = newSID
		updateSession(newSID, 0)
	}
	fmt.Println(session.ID)
}

func ServePage(w http.ResponseWriter, r *http.Request) {
	validateSession(w, r)
	vars := mux.Vars(r)
	pageGUID := vars["guid"]
	thisPage := Page{}
	thisPage.GUID = pageGUID
	err := database.QueryRow("SELECT id,page_title,page_content,page_date FROM pages WHERE page_guid=$1", pageGUID).Scan(&thisPage.Id, &thisPage.Title, &thisPage.RawContent, &thisPage.Date)
	thisPage.Content = template.HTML(thisPage.RawContent)
	if err != nil {
		http.Error(w, http.StatusText(404), http.StatusNotFound)
		log.Println(err)
		return
	}

	comments, err := database.Query("SELECT id, comment_name as Name, comment_email, comment_text FROM comments WHERE page_id=$1", thisPage.Id)
	if err != nil {
		log.Println(err)
	}
	for comments.Next() {
		var comment Comment
		comments.Scan(&comment.Id, &comment.Name, &comment.Email, &comment.CommentText)
		thisPage.Comments = append(thisPage.Comments, comment)
	}
	thisPage.Session.Authenticated = false
	t, _ := template.ParseFiles("blog.html")
	t.Execute(w, thisPage)
}

func APIPage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	pageGUID := vars["guid"]
	thisPage := Page{}
	fmt.Println(pageGUID)
	err := database.QueryRow("SELECT page_title,page_content,page_date FROM pages WHERE page_guid=?", pageGUID).Scan(&thisPage.Title, &thisPage.RawContent, &thisPage.Date)
	thisPage.Content = template.HTML(thisPage.RawContent)
	if err != nil {
		http.Error(w, http.StatusText(404), http.StatusNotFound)
		log.Println(err)
		return
	}
	APIOutput, _ := json.Marshal(thisPage)
	fmt.Println(APIOutput)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, thisPage)
}

func APICommentPost(w http.ResponseWriter, r *http.Request) {
	var commentAdded string
	err := r.ParseForm()
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Println(r.FormValue("name"))
	name := r.FormValue("name")
	email := r.FormValue("email")
	comments := r.FormValue("comments")

	res, err := database.Exec("INSERT INTO comments SET comment_name=$1, comment_email=$2, comment_text=$3", name, email, comments)

	if err != nil {
		log.Println(err.Error())
	}

	id, err := res.LastInsertId()
	if err != nil {
		commentAdded = "false"
	} else {
		commentAdded = "true"
	}

	var resp JSONResponse
	resp.Fields["id"] = strconv.FormatInt(id, 10)
	resp.Fields["added"] = commentAdded
	jsonResp, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, jsonResp)
}

func APICommentPut(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Heyyyyyyy")
	err := r.ParseForm()
	if err != nil {
		log.Println(err.Error())
	}
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Println(id)
	name := r.FormValue("name")
	email := r.FormValue("email")
	comments := r.FormValue("comments")
	fmt.Println("UPDATE comments SET comment_name=$1, comment_email=$2, comment_text=$3 WHERE comment_id=$4", name, email, comments, id)
	res, err := database.Exec("UPDATE comments SET comment_name=$1, comment_email=$2, comment_text=$3 WHERE comment_id=$4", name, email, comments, id)
	fmt.Println(res)
	if err != nil {
		log.Println(err.Error())
	}

	var resp JSONResponse

	jsonResp, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, jsonResp)
}

func weakPasswordHash(password string) []byte {
	hash := sha1.New()
	io.WriteString(hash, password)
	return hash.Sum(nil)
}

func MQConnect() (*amqp.Connection, *amqp.Channel, error) {
	url := "amqp://" + MQHost + MQPort
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, nil, err
	}
	channel, err := conn.Channel()
	if err != nil {
		return nil, nil, err
	}
	if _, err := channel.QueueDeclare("", false, true, false, false, nil); err != nil {
		return nil, nil, err
	}
	return conn, channel, nil
}

func MQPublish(channel *amqp.Channel, message []byte) {
	channel.Publish(
		"email", // exchange
		"",      // routing key
		false,   // mandatory
		false,   // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
}

func RegisterPOST(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err.Error())
	}
	name := r.FormValue("user_name")
	email := r.FormValue("user_email")
	pass := r.FormValue("user_password")
	pageGUID := r.FormValue("referrer")
	// pass2 := r.FormValue("user_password2")
	gure := regexp.MustCompile("[^A-Za-z0-9]+")
	guid := gure.ReplaceAllString(name, "")
	password := weakPasswordHash(pass)

	res, err := database.Exec("INSERT INTO users SET user_name=$1, user_guid=$2, user_email=$3, user_password=$4", name, guid, email, password)
	fmt.Println(res)
	if err != nil {
		fmt.Fprintln(w, err.Error())
	} else {
		Email := RegistrationData{Email: email, Message: ""}
		message, _ := template.New("email").Parse(WelcomeEmail)
		var mbuf bytes.Buffer
		message.Execute(&mbuf, Email)
		messageBytes, err := json.Marshal(mbuf.String())
		if err != nil {
			log.Println(err)
			return
		}
		conn, channel, err := MQConnect()
		if err != nil {
			log.Println(err)
			return
		}
		defer conn.Close()
		defer channel.Close()
		MQPublish(channel, messageBytes)
		http.Redirect(w, r, "/page/"+pageGUID, http.StatusMovedPermanently)
	}
}
func LoginPOST(w http.ResponseWriter, r *http.Request) {
	validateSession(w, r)
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err.Error())
	}
	u := User{}
	name := r.FormValue("user_name")
	pass := r.FormValue("user_password")
	password := weakPasswordHash(pass)
	err = database.QueryRow("SELECT user_id, user_name FROM users WHERE user_name=$1 and user_password=$2", name, password).Scan(&u.Id, &u.Name)
	if err != nil {
		fmt.Fprintln(w, err.Error())
		u.Id = 0
		u.Name = ""
	} else {
		updateSession(UserSession.Id, u.Id)
		fmt.Fprintln(w, u.Name)
	}
}

func main() {
	// dbConn := fmt.Sprintf("%s:%s@/%s", DBUser, DBPass, DBDbase)
	// fmt.Println(dbConn)
	// db, err := sql.Open("mysql", dbConn)

	dbURL := "host=localhost port=5432 user=postgres password=Pyari@123 sslmode=disable dbname=mydb"
	db, err := sql.Open("postgres", dbURL) // Remove := to use global database variable

	if err != nil {
		log.Println("Couldn't connect!")
		log.Println(err.Error())
	}
	database = db

	routes := mux.NewRouter()
	routes.HandleFunc("/register", RegisterPOST).Methods("POST")
	routes.HandleFunc("/login", LoginPOST).
		Methods("POST")
	routes.HandleFunc("/api/pages", APIPage).
		Methods("GET").
		Schemes("https")
	routes.HandleFunc("/api/page/{id:[\\w\\d\\-]+}", APIPage).
		Methods("GET").
		Schemes("https")
	routes.HandleFunc("/api/comments", APICommentPost).
		Methods("POST")
	routes.HandleFunc("/api/comments/{id:[\\w\\d\\-]+}", APICommentPut).
		Methods("PUT")
	routes.HandleFunc("/page/{guid:[0-9a-zA\\-]+}", ServePage)
	http.Handle("/", routes)
	http.ListenAndServe(PORT, nil)

}
