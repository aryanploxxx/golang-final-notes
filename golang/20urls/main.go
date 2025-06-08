package main

import (
	"fmt"
	"net/url"
)

const myURL string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=hajaskas"

func main() {
	fmt.Println("URLs")
	fmt.Println("URL is: ", myURL)

	// URL Parsing - extracting information from the URL
	result, _ := url.Parse(myURL)

	fmt.Println("Scheme: ", result.Scheme)
	// Scheme:  https
	fmt.Println("Host: ", result.Host)
	// Host:  lco.dev:3000
	fmt.Println("Path: ", result.Path)
	// Path:  /learn
	fmt.Println("Port: ", result.Port())
	// Port:  3000
	fmt.Println("Query: ", result.RawQuery)
	// Query:  coursename=reactjs&paymentid=hajaskas

	// Query Parsing
	qparams := result.Query()
	fmt.Println("Query Params: ", qparams)
	// Query Params:  map[coursename:[reactjs] paymentid:[hajaskas]]
	fmt.Printf("Type of Query Params are: %T\n", qparams)
	// Type of Query Params are: url.Values
	// url.Values is a fancy name for key-value pairs

	fmt.Println(qparams["coursename"]) // [reactjs]

	for key, value := range qparams {
		fmt.Println("Key: ", key, " Value: ", value)
	}

	// Key:  coursename  Value:  [reactjs]
	// Key:  paymentid  Value:  [hajaskas]

	// Assuming we have all query parameters and now want to create a URL out of them
	// Notice how we passed as a reference
	partsOfURL := &url.URL{
		Scheme:   "https",
		Host:     "lco.dev:3000",
		Path:     "/learn",
		RawQuery: "coursename=reactjs&paymentid=hajaskas",
	}

	anotherURL := partsOfURL.String()
	fmt.Println("Another URL: ", anotherURL)

	queryParams, _ := url.ParseQuery(result.RawQuery)
	fmt.Println("Query Params: ", queryParams)
	// Query Params:  map[coursename:[reactjs] paymentid:[hajaskas]]
}
