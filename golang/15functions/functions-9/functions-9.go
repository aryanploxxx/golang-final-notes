package main

import "fmt"

// Named return values

func main() {
	sum, avg := sum_avg(4, 2)
	fmt.Println(sum)
	fmt.Println(avg)
}

func sum_avg(a, b int) (sum, avg int) {
	sum = a + b
	avg = (a + b) / 2
	return
}
