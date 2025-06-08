package main

import (
	"log"
	"os"
	"strconv"
)

type GoroutineLogger struct {
	UniqueLog *log.Logger
	Name      string
}

func main() {
	block := make(chan bool)

	for i := 0; i < 10; i++ {
		go func() {
			newLogger := GoroutineLogger{}

			newLogger.Name = "Goroutine " + strconv.Itoa(i+1)

			filename := "log" + strconv.Itoa(i+1) + ".txt"

			logFile, _ := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
			// logFile, _ := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

			newLogger.UniqueLog = log.New(logFile, "Log: ", log.Ldate|log.Ltime|log.Lshortfile)

			newLogger.UniqueLog.Printf("Hello from %s", newLogger.Name)

			block <- true
		}()

		<-block

	}

	// for i := 0; i < 10; i++ {
	// 	<-block
	// }
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
