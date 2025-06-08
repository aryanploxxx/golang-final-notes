package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var applicationStatus bool // Controls program execution
var urls []string          // Stores URLs to scrape
var urlsProcessed int      // Counter for processed URLs
// var foundUrls []string     // Stores URLs found during scraping
var fullText string   // Concatenated text from all URLs
var totalURLCount int // Total number of URLs to process
// var wg sync.WaitGroup      // WaitGroup (unused in main logic)

// Adds scraped text to fullText variable
func addToScrapedText(textChannel chan string, processChannel chan bool) {
	for {
		select {
		case pC := <-processChannel:
			if !pC {
				// Close channels and exit if processing is complete
				close(textChannel)
				close(processChannel)
			}
		case tC := <-textChannel:
			// Add received text to fullText
			fullText += tC
		}
	}
}

// Monitors scraping progress
func evaluateStatus(statusChannel chan int, processChannel chan bool) {
	for {
		select {
		case status := <-statusChannel:
			fmt.Print(urlsProcessed, totalURLCount)
			urlsProcessed++

			if status == 0 {
				fmt.Println("Got url")
			}
			if status == 1 {
				close(statusChannel)
			}
			// Check if all URLs are processed
			if urlsProcessed == totalURLCount {
				fmt.Println("Read all top-level URLs")
				processChannel <- false
				applicationStatus = false
			}
		}
	}
}

// Performs the actual URL reading
func readURLs(statusChannel chan int, textChannel chan string) {
	time.Sleep(time.Millisecond * 1)
	fmt.Println("Grabbing", len(urls), "urls")
	for i := 0; i < totalURLCount; i++ {
		fmt.Println("Url", i, urls[i])
		resp, _ := http.Get(urls[i])           // Get URL content
		text, err := ioutil.ReadAll(resp.Body) // Read response body
		textChannel <- string(text)            // Send text to channel
		if err != nil {
			fmt.Println("No HTML body")
		}
		statusChannel <- 0 // Signal completion
	}
}

func main() {
	// Initialize channels and variables
	applicationStatus = true
	statusChannel := make(chan int)
	textChannel := make(chan string)
	processChannel := make(chan bool)

	// Add URLs to scrape
	urls = append(urls, "https://example.com/")
	urls = append(urls, "https://golang.org/")
	totalURLCount = len(urls)

	// Launch goroutines
	go evaluateStatus(statusChannel, processChannel)
	go readURLs(statusChannel, textChannel)
	go addToScrapedText(textChannel, processChannel)

	// Main loop waiting for completion
	for {
		if !applicationStatus {
			fmt.Println(fullText)
			fmt.Println("Done!")
			break
		}
		select {
		case sC := <-statusChannel:
			fmt.Println("Message on StatusChannel", sC)
		}
	}
}
