package main

import (
	"log"
	"os"
)

var (
	Warn  *log.Logger
	Error *log.Logger
)

// err.Error() is a method that converts the error err into a human-readable string format

func main() {
	// Open warning log file
	warnFile, err := os.OpenFile("warnings.log", os.O_RDWR|os.O_APPEND, 0660)
	// os.OpenFile opens a file for logging.
	// os.O_RDWR: Read and write access.
	// os.O_APPEND: Append new logs to the file.
	// 0660: Read/write access for owner and group.

	defer warnFile.Close()
	if err != nil {
		log.Fatal(err) // Exit if unable to open file
	}

	// Open error log file
	errorFile, err := os.OpenFile("error.log", os.O_RDWR|os.O_APPEND, 0660)
	defer errorFile.Close()
	if err != nil {
		log.Fatal(err)
	}

	// Create loggers
	// A logger can be customized using the log.New() function.
	/*
		io.Writer: Destination for log output (e.g., file).
		Prefix string: Prepends each log message (e.g., "ERROR: ").
		Flags: Controls the date/time format (e.g., log.Ldate | log.Ltime).
	*/
	Warn = log.New(warnFile, "WARNING aryan: ", log.LstdFlags)
	Error = log.New(errorFile, "ERROR: ", log.Ldate|log.Ltime)

	numone := 5
	numtwo := 10
	if numone < numtwo {
		Warn.Println("numone is less than numtwo")
	} else {
		Error.Println("numone is not less than numtwo")
	}

	// Log messages
	Warn.Println("This is a warning message")
	Error.Println("This is an error message")

}

/*

# log.New() -> create a new logger, set output location and prefix for log file
file, _ := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
defer file.Close()

customLogger := log.New(file, "CUSTOM: ", log.Ldate|log.Ltime|log.Lshortfile)
customLogger.Println("This is a custom log message.")

# log.SetOutput() -> change the output position of the default logger
file, _ := os.OpenFile("default.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
defer file.Close()

log.SetOutput(file)
log.Println("This message will be written to default.log")

*/
