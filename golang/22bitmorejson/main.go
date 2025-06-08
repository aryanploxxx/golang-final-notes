package main

import (
	"encoding/json"
	"fmt"
)

// c here is lowercaes indicating that we will not be able to access it outside the package
// type course struct {
// 	Name     string
// 	Price    int
// 	Platform string
// 	Password string
// 	Tags     []string
// }

type course struct {
	Name     string   `json:"coursename"`
	Price    int      `json:"price"`
	Platform string   `json:"website"`
	Password string   `json:"-"`              // - means that this field will not be included/reflected/shown in the JSON to whoever is consuming our API
	Tags     []string `json:"tags,omitempty"` // omitempty means that if the value is nil, then it will not be included in the JSON
	// Tags     []string `json:"tags, omitempty"` // if there is space between tags and omitempty, then it will not work and throw an error
}

// In this video we are going to work on the encoding of JSON
// It means that we have some data, it can be arrays, slices, maps, structs, etc. and we want to convert it into valid JSON format

func main() {
	fmt.Println("More JSON")
	// EncodeJSON()
	DecodeJSON()
}

func EncodeJSON() {
	// slice of type course(struct) which we created earlier
	// slice of multiple structs
	lcoCourses := []course{
		{"ReactJS Bootcamp", 299, "LearnCodeOnline.in", "abc123", []string{"web-dev", "js", "react"}},
		{"MERN Bootcamp", 199, "LearnCodeOnline.in", "bcd123", []string{"full-stack", "js", "react"}},
		{"Angular Bootcamp", 299, "LearnCodeOnline.in", "mjnad123", nil},
	}

	// Package this data into JSON format
	finalJSON, err := json.Marshal(lcoCourses)
	// json.Marshal() takes an interface and returns a slice of bytes and an error
	// it is the interface implemted by types that can marshal themselves into valid JSON
	// Interface is like the alternative version of struct

	if err != nil {
		panic(err)
	}

	fmt.Printf("JSON: %s", finalJSON)
	// JSON: [{"Name":"ReactJS Bootcamp","Price":299,"Platform":"LearnCodeOnline.in","Password":"abc123","Tags":["web-dev","js","react"]},{"Name":"MERN Bootcamp","Price":199,"Platform":"LearnCodeOnline.in","Password":"bcd123","Tags":["full-stack","js","react"]},{"Name":"Angular Bootcamp","Price":299,"Platform":"LearnCodeOnline.in","Password":"mjnad123","Tags":null}]
	// array of objects
	// However this is not readble format

	finalJSONIndented, _ := json.MarshalIndent(lcoCourses, "", "\t")
	// json.MarshalIndent() is used to indent the JSON data
	// lcoCourses, "", "\t" -> lcoCourses is the data we want to indent, "" is the prefix for each line, "\t" is the indentation marker
	fmt.Println("Indented JSON: ", string(finalJSONIndented))
	/*
		Indented JSON:
		[
			{
					"Name": "ReactJS Bootcamp",
					"Price": 299,
					"Platform": "LearnCodeOnlinein",
					"Password": "abc123",
					"Tags": [
							"web-dev",
							"js",
							"react"
					]
			},
			{
					"Name": "MERN Bootcamp",
					"Price": 199,
					"Platform": "LearnCodeOnlinein",
					"Password": "bcd123",
					"Tags": [
							"full-stack",
							"js",
							"react"
					]
			},
			{
					"Name": "Angular Bootcamp",
					"Price": 299,
					"Platform": "LearnCodeOnlinein",
					"Password": "mjnad123",
					"Tags": null
			}
		]
	*/

	// This is still not our ideal desired data format, as password is still visible, and the keys are in capital case, which is not recommended in case of JSON
	// To solve the above issue, we can use aliases against keys in struct and specify what their names should be in JSON format
	// We can also use the omitempty tag to avoid printing the key if the value is nil

	finalJSONCorrected, _ := json.MarshalIndent(lcoCourses, "", "\t")
	fmt.Println("Corrected JSON: ", string(finalJSONCorrected))
	/*
		[
			{
					"coursename": "ReactJS Bootcamp",
					"price": 299,
					"website": "LearnCodeOnlinein",
					"tags": [
							"web-dev",
							"js",
							"react"
					]
			},
			{
					"coursename": "MERN Bootcamp",
					"price": 199,
					"website": "LearnCodeOnlinein",
					"tags": [
							"full-stack",
							"js",
							"react"
					]
			},
			{
					"coursename": "Angular Bootcamp",
					"price": 299,
					"website": "LearnCodeOnlinein"
			}
		]
	*/
}

func DecodeJSON() {
	jsonDataFromWeb := []byte(`
	 	{
			"coursename": "ReactJS Bootcamp",
			"price": 299,
			"website": "LearnCodeOnline.in",
			"tags": ["web-dev","js","react"]
        }
	`)

	// thought behind the below operation is that whatevr the data we are getting from the web, we will arrange them in struct in GO
	var lcoCourse course
	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		// json.Unmarshal() takes a slice of bytes and a pointer to an interface
		// jsonDataFromWeb is the data we want to unmarshal, &lcoCourse is the pointer to the interface/struct where we want to store the destructured data
		fmt.Printf("%#v\n", lcoCourse)
		// Notice the use of %#v instead of %v, this will print the struct with field names, %v will just print the keys
		// %+v will print the struct with field names in key-value pairs -> {Name:ReactJS Bootcamp Price:299 Platform:LearnCodeOnline.in Password: Tags:[web-dev js react]}
		// main.course{Name:"ReactJS Bootcamp", Price:299, Platform:"LearnCodeOnline.in", Password:"", Tags:[]string{"web-dev", "js", "react"}}
		// It also did reverse mapping, coursename in JSON got mapped to Name in struct and vice-versa will happen when we again marhsal this struct
	} else {
		fmt.Println("JSON was not valid")
	}

	// Some cases where you want to add data into key value pairs - when we do not want to create a structure everytime, so that we can extract values based on keys
	// For this we can use map[string]interface{} instead of struct

	var myOnlineData map[string]interface{} // map with string as key and interface as value, we used interface as we do not know what type of data we will get
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)
	// map[string]interface {}{"coursename":"ReactJS Bootcamp", "price":299, "tags":[]interface {}{"web-dev", "js", "react"}, "website":"LearnCodeOnline.in"}

	for key, value := range myOnlineData {
		// fmt.Println("Key is: ", key, "Value is: ", value)
		fmt.Printf("Key is: %v, Value is: %v\n and Type of value is: %T\n", key, value, value)
	}

	// Key is: coursename, Value is: ReactJS Bootcamp and Type of value is: string
	// Key is: price, Value is: 299 and Type of value is: float64
	// Key is: website, Value is: LearnCodeOnline.in and Type of value is: string
	// Key is: tags, Value is: [web-dev js react] and Type of value is: []interface {}

	// Golang is treating price as float64 because this has the maximum precision value
}

/*
	In Go, JSON data is often represented as a byte slice ([]byte) because this format is convenient for reading and processing raw data. Here are a few reasons why you see JSON data represented as []byte:
	Efficiency: A byte slice is a raw representation of data, making it efficient to work with for various I/O operations, such as reading from a file, a network connection, or an HTTP response.
	Compatibility: Functions in Go's encoding/json package, like json.Unmarshal(), accept []byte as input. This makes it easier to decode JSON data into Go structs.
	Flexibility: By using a byte slice, you can handle JSON data as a stream of bytes, which provides flexibility in how you process or manipulate the data.
*/
