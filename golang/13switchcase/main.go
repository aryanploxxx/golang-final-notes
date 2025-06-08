package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	fmt.Println("Switch and Case")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	// Without +1 it would have generated 0 to 5

	fmt.Println("Value of dice is ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 and you can open")
	case 2:
		fmt.Println("You can move 2 places")
	case 3:
		fmt.Println("You can move 3 places")
		fallthrough
		// fallthrough makes sure when this case is executed, the case just below it also gets executed
		/*
			Value of dice is  3
			You can move 3 places
			You can move 4 places
			-? for ex. if 3 and 4 both had fallthorugh, then when die rolled 3, we would exceute 3 and 4case 4, which also has fallthrough, then case 5 would also get executed
		*/
	case 4:
		fmt.Println("You can move 4 places")
	case 5:
		fmt.Println("You can move 5 places")
	case 6:
		fmt.Println("You can move 6 places and roll the dice again")
	default:
		fmt.Println("What was that!")
	}

	fmt.Print("Go is running on on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
	// In Go, weekdays are represented by integers starting from 0 (Sunday) to 6 (Saturday). We can use this information to calculate the day of the week for Saturday and compare it with the current day to determine how many days are left until the weekend.

	// Switch without a condition is the same as switch true. This construct can be a clean way to write long if-then-else chains.

}

/*
	rand.Seed()
	rand.Seed is a function from the math/rand package used to initialize the random number generator.
	By setting a seed value, you ensure that the sequence of random numbers generated is different each time your program runs.
	time.Now().UnixNano()
	time.Now() retrieves the current local time.
	.UnixNano() converts this time to an integer representing the number of nanoseconds since January 1, 1970 (Unix epoch).
	Using nanoseconds ensures a high level of granularity, making it an excellent seed value for randomness.
	Putting it all together, rand.Seed(time.Now().UnixNano()) initializes the random number generator with a unique seed value based on the current time in nanoseconds, ensuring that the sequence of random numbers is different every time the program is executed. Here's a quick example to illustrate:

	rand.Seed(time.Now().UnixNano()) fmt.Println(rand.Intn(100)) // Prints a random number between 0 and 99
*/
