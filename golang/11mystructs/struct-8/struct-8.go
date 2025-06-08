package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type employee struct {
	Name   string
	Age    int
	salary int
}

func main() {
	emp := employee{Name: "Sam", Age: 31, salary: 2000}
	//Marshal
	empJSON, err := json.Marshal(emp)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Printf("Marshal funnction output %s\n", string(empJSON))

	//MarshalIndent
	empJSON, err = json.MarshalIndent(emp, "", "  ")
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Printf("MarshalIndent funnction output %s\n", string(empJSON))
}

/*
	Marshal funnction output {"Name":"Sam","Age":31}
	MarshalIndent funnction output {
	"Name": "Sam",
	"Age": 31
	}
*/

// The salary field is not printed in the output because it begins with a lowercase letter and is not exported. The Marshal function output is not formatted while the MarshalIndent function output is formatted.
// It is to be noted that both Marshal and MarshalIndent function can only access the exported fields of a struct,
// which means that only the capitalized fields can be accessed and encoded in JSON form.

// golang also allows the JSON encoded struct key name to be different by the use of struct meta fields as will see in the next section.
