package main

import "fmt"

func main() {
	fmt.Println("Maps")

	// languages := make(map[int]string)
	// here int is the type for key, and string is the type for value
	// new gives a lot of errors and since it is a zeroed value, we cannot store data in it

	languages := make(map[string]string)
	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of languages: ", languages)
	// List of languages:  map[JS:Javascript PY:Python RB:Ruby]
	//  Notice that the values in array, slices are not comma seperated

	fmt.Println("JS stands for: ", languages["JS"])

	// To delete any value:
	delete(languages, "JS")

	fmt.Println("Updated List of languages: ", languages)
	// Updated List of languages:  map[PY:Python RB:Ruby]

	// Iterating over maps
	for key, value := range languages {
		fmt.Printf("For key %v, the value is %v\n", key, value)
		// %v is a placeholder
	}
	// languages here is the thing we are iterating over
}
