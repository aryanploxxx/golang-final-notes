package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")

	presentTime := time.Now()
	// time.Now() returns the current local time

	fmt.Println("Present time is: ", presentTime)
	// Present time is:  2024-12-04 15:53:01.5067992 +0530 IST m=+0.000523701

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))
	// 12-04-2024 15:55:34 Wednesday
	// Interestingly, 01-02-2006 15:04:05 Monday -> this is the exact value we need to enter to get the desired output
	// Go has this date-time as reference

	createdDate := time.Date(2024, time.December, 4, 15, 55, 34, 0, time.UTC)
	// createdDate := time.Date(year, time.month, day, hour, minute, second, nanosecond, time.location)
	// year, day, hour, minute, second all are type on int
	// month is of type time.month
	fmt.Println("Created date is: ", createdDate)
	fmt.Println("Created date is: ", createdDate.Format("01-02-2006 15:04:05 Monday"))

}
