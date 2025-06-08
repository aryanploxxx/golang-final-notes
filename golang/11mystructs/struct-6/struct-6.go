package main

import "fmt"

// Define the User struct
type User struct {
	ID    int
	Name  string
	Email string
}

func main() {
	// Create a new instance of User using the `new` keyword
	user := new(User)

	// Insert values using curly bracket notation with dereferencing
	*user = User{
		ID:    1,
		Name:  "John Doe",
		Email: "john.doe@example.com",
	}

	// Print the struct to verify
	fmt.Printf("User Details:\n ID: %d\n Name: %s\n Email: %s\n", user.ID, user.Name, user.Email)
}
