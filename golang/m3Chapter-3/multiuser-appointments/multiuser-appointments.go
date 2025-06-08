// Package main implements a simple web-based scheduling system
package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"sync"

	"github.com/gorilla/mux"
)

// User represents a person in the scheduling system
type User struct {
	Name     string        // User's name
	Times    map[int]bool  // Map of available time slots (hour -> availability)
	DateHTML template.HTML // HTML representation of user's schedule
}

// Page represents the data structure used for template rendering
type Page struct {
	Title string          // Page title
	Body  template.HTML   // Page content
	Users map[string]User // Map of all users in the system
}

// Global variables
var (
	Users     map[string]User                                      // Store of all users
	userIndex int                                                  // Counter for users (unused in current implementation)
	mutex     sync.Mutex                                           // Mutex for thread-safe operations
	templates = template.Must(template.New("template").ParseFiles( // Load and parse HTML templates
		"view_users.html",
		"register.html",
	))
)

// register handles user registration
func register(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request to /register")
	params := mux.Vars(r)
	name := params["name"]

	// Check if user already exists
	if _, exists := Users[name]; exists {
		t, _ := template.ParseFiles("generic.txt")
		page := &Page{
			Title: "User already exists",
			Body:  template.HTML("User " + name + " already exists"),
		}
		t.Execute(w, page)
		return
	}

	// Create new user
	newUser := User{Name: name}
	initUser(&newUser)
	Users[name] = newUser

	// Render success page
	t, _ := template.ParseFiles("generic.txt")
	page := &Page{
		Title: "User created!",
		Body:  template.HTML("You have created user " + name),
	}
	t.Execute(w, page)
}

// formatTime converts 24-hour time to 12-hour format with am/pm
func formatTime(hour int) string {
	hourText := hour
	ampm := "am"

	if hour > 11 {
		ampm = "pm"
	}
	if hour > 12 {
		hourText = hour - 12
	}

	return strconv.FormatInt(int64(hourText), 10) + ampm
}

// FormatAvailableTimes creates HTML representation of a user's available times
func (u User) FormatAvailableTimes() template.HTML {
	HTML := "<b>" + u.Name + "</b> - "

	for hour, available := range u.Times {
		if available {
			formattedTime := formatTime(hour)
			HTML += fmt.Sprintf(
				`<a href='/schedule/%s/%d' class='button'>%s</a> `,
				u.Name,
				hour,
				formattedTime,
			)
		}
	}

	return template.HTML(HTML)
}

// users handles the display of all users and their schedules
func users(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request to /users")
	t, _ := template.ParseFiles("users.txt")
	page := &Page{
		Title: "View Users",
		Users: Users,
	}
	t.Execute(w, page)
}

// schedule handles the scheduling of appointments
func schedule(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request to /schedule")
	params := mux.Vars(r)
	name := params["name"]
	time := params["hour"]

	timeVal, _ := strconv.ParseInt(time, 10, 0)
	intTimeVal := int(timeVal)
	createURL := "/register/" + name

	// Check if user exists
	if user, exists := Users[name]; exists {
		// Check if timeslot is available
		if user.Times[intTimeVal] {
			// Thread-safe update of availability
			mutex.Lock()
			Users[name].Times[intTimeVal] = false
			mutex.Unlock()

			t, _ := template.ParseFiles("generic.txt")
			page := &Page{
				Title: "Successfully Scheduled!",
				Body:  template.HTML("This appointment has been scheduled. <a href='/users'>Back to users</a>"),
			}
			t.Execute(w, page)
		} else {
			// Timeslot is already booked
			t, _ := template.ParseFiles("generic.txt")
			page := &Page{
				Title: "Booked!",
				Body:  template.HTML("Sorry, " + name + " is booked for " + time + " <a href='/users'>Back to users</a>"),
			}
			t.Execute(w, page)
		}
	} else {
		// User doesn't exist
		t, _ := template.ParseFiles("generic.txt")
		page := &Page{
			Title: "User Does Not Exist!",
			Body:  template.HTML("Sorry, that user does not exist. Click <a href='" + createURL + "'>here</a> to create it. <a href='/users'>Back to users</a>"),
		}
		t.Execute(w, page)
	}
}

// defaultPage handles the root route (currently empty)
func defaultPage(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement landing page
}

// initUser initializes a new user's available time slots (9am-5pm)
func initUser(user *User) {
	user.Times = make(map[int]bool)
	for i := 9; i < 18; i++ {
		user.Times[i] = true
	}
}

func main() {
	// Initialize application state
	Users = make(map[string]User)
	userIndex = 0

	// Create default user "Bill"
	bill := User{Name: "Bill"}
	initUser(&bill)
	Users["Bill"] = bill
	userIndex++

	// Set up routing
	r := mux.NewRouter()
	r.HandleFunc("/", defaultPage)
	r.HandleFunc("/users", users)
	r.HandleFunc("/register/{name:[A-Za-z]+}", register)
	r.HandleFunc("/schedule/{name:[A-Za-z]+}/{hour:[0-9]+}", schedule)

	// Start server
	http.Handle("/", r)
	if err := http.ListenAndServe(":1900", nil); err != nil {
		panic(fmt.Sprintf("ListenAndServe: %v", err))
	}
}
