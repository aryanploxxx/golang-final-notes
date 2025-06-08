package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Web Server Requests")
	// PerformGetRequest()

	// We can either send the data to a server using JSON or url-encoded forms
	// PerfromPostJSONRequest()

	PerformPostFormRequest()
}

func PerformGetRequest() {
	// Get Request
	const myUrl = "http://localhost:8000/get"
	fmt.Println("URL is: ", myUrl)

	response, err := http.Get(myUrl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status Code: ", response.StatusCode)
	fmt.Println("Content Length: ", response.ContentLength)
	/*
		// To read the response body, we will need to take help of ioutil or io package
		content, err := ioutil.ReadAll(response.Body)

		fmt.Println("Content: ", string(content))
		// without wrapping content around string would result in byte array
		// Content:  {"message":"Hello from learnCodeonline.in"}
		// -> the above output is in string format as of now

		// There is nothing wrong with the above code, but we can make it more efficient by using the strings package
	*/

	content, err := ioutil.ReadAll(response.Body)
	var responseString strings.Builder // strings.Builder is a more efficient way to concatenate strings
	byteCount, _ := responseString.Write(content)
	// Write returns the number of bytes written; same as content length
	fmt.Println("Byte  is: ", byteCount)
	fmt.Println("Same Content: ", responseString.String())

	// -> This method is considered more powerful because we are always holding the raw bytes data and can do any operation on it at any given time. as opposed to the above method where we converted the entire data once and for all

}

func PerfromPostJSONRequest() {
	const myUrl = "http://localhost:8000/post"

	// Fake JSON payload
	// requestBody := strings.NewReader(``) -> we can write anything inside `` and create any type of data
	requestBody := strings.NewReader(`
		{
			"coursename": "Golang",
			"price": 0,
			"platform": "youtube.com"
		}
	`)

	// Post Request
	// http.Post(url, contentType, body)
	response, err := http.Post(myUrl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println("Content: ", string(content))

}

func PerformPostFormRequest() {
	const myUrl = "http://localhost:8000/postform"

	// Fake Form Data
	// Whenever we post any data, it can be accessed using the 'url' package
	data := url.Values{} // initially, we will keep it empty, bus as we get data we will inject in it using .Add() method
	data.Add("firstname", "aryan")
	data.Add("lastname", "gupta")
	data.Add("email", "aryan@mail.com")
	// somehow the data is getting automatically sorted

	response, err := http.PostForm(myUrl, data)
	//  we don't have to specify the content type as it is already set to form data, it's  a specially designed function to handle form data

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)
	fmt.Println("Content: ", string(content))

}
