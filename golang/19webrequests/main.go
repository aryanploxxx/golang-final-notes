package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://go.dev/dl/"

func main() {
	fmt.Println("Web Requests")
	// This effectively behaves as a web scrapper, we can use this to get the data from the web

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Println("Response: ", response.Body)
	// &{[] {0xc000280180} <nil> <nil>}

	fmt.Println("Response: ", response)
	/*
		Response:  &{200 OK 200 HTTP/2.0 2 0
			map[
				Access-Control-Allow-Credentials:[true]
				Age:[26041]
				Alt-Svc:[h3=":443"; ma=86400]
				Cache-Control:[max-age=43200]
				Cf-Cache-Status:[HIT]
				Cf-Ray:[8fddf73d1dfbedb8-MXP]
				Content-Type:[application/json; charset=utf-8]
				Date:[Mon, 06 Jan 2025 19:03:01 GMT]
				Etag:[W/"6b80-Ybsq/K6GwwqrYkAsFxqDXGC7DoM"]
				Expires:[-1]
				Nel:[{"report_to":"heroku-nel","max_age":3600,"success_fraction":0.005,"failure_fraction":0.05,"response_headers":["Via"]}]
				Pragma:[no-cache]
				Report-To:[{"group":"heroku-nel","max_age":3600,"endpoints":[{"url":"https://nel.heroku.com/reports?ts=1733228488&sid=e11707d5-02a7-43ef-b45e-2cf4d2036f7d&s=hsM7ZP4mYihRgZe6Gwc6TbV8Vf2TAc8wXWpPUtmWP0U%3D"}]}]
				Reporting-Endpoints:[heroku-nel=https://nel.heroku.com/reports?ts=1733228488&sid=e11707d5-02a7-43ef-b45e-2cf4d2036f7d&s=hsM7ZP4mYihRgZe6Gwc6TbV8Vf2TAc8wXWpPUtmWP0U%3D] Server:[cloudflare] Server-Timing:[cfL4;desc="?proto=TCP&rtt=230052&min_rtt=198320&rtt_var=84091&sent=10&recv=9&lost=0&retrans=0&sent_bytes=4288&recv_bytes=1739&delivery_rate=27361&cwnd=253&unsent_bytes=0&cid=b2bbf8c7457cbe1d&ts=274&x=0"] Vary:[Origin, Accept-Encoding] Via:[1.1 vegur] X-Content-Type-Options:[nosniff] X-Powered-By:[Express] X-Ratelimit-Limit:[1000] X-Ratelimit-Remaining:[999] X-Ratelimit-Reset:[1733228528]] 0xc00019f380 -1 [] false true map[] 0xc00009c280 0xc00022a0c0}
	*/

	fmt.Printf("Response type is %T", response)
	// Response type is *http.Response
	// We get a pointer to the response
	// This guarantees that the response is not copied, it is passed by reference

	defer response.Body.Close()
	// It is the caller's responsibilty to close the connection

	databytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println("Data: ", string(databytes))

	/*
		The function ioutil.ReadAll() returns two values:
		A byte slice ([]byte) containing the data read from the io.Reader.
		An error value that indicates if there was an error while reading.
		In your example, the databytes variable will be of type []byte, and err will be of type error. Here's a quick rundown:
			databytes: A byte slice that holds the content read from the response.Body.
			err: An error value that is nil if the read operation was successful or contains an error message if something went wrong.
	*/
}
