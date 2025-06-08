package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("File Handling")
	// Go has a built-in support for text files, other formats like pdf and csv require special libraries

	content := "This needs to go in a text file"

	file, err := os.Create("./text.txt")

	if err != nil {
		panic(err)
		// panic will shutdown the program and show us the error we are facing
		// this is a common way of handling the errors
	}

	// io package will be used for writing anything to the file
	length, err := io.WriteString(file, content)
	// will return the length of the file if the package is executed
	// this will also create a file text.txt

	checkNil(err)
	// Since checking error was a repitive task, we created a function for it

	fmt.Println("Length of the file is: ", length)
	// Length of the file is:  31

	defer file.Close()
	// defer keyword is used to make sure that the file is closed at the extreme end of the function, we can even write some code after this line and still it will execute the very bottom of the code

	readFile("./text.txt")
}

// Creation of files is an OS operation, but reading files and other manipulations, there is another utility provided to us

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	// ioutil is deprecated, new package is io
	// Whenever we are reading a file, it is NOT being read in string format, it is being read in bytes
	checkNil(err)

	fmt.Println("Text file data is: \n", databyte)
	// Text file data is: [84 104 105 115 32 110 101 101 100 115 32 116 111 32 103 111 32 105 110 32 97 32 116 101 120 116 32 102 105 108 101]
	fmt.Println("Text file data is: \n", string(databyte))
	// Text file data is: This needs to go in a text file
}

func checkNil(err error) {
	if err != nil {
		panic(err)
	}
}
