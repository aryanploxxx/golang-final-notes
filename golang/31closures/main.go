package main

func main() {
	// Closures
	// A closure is a function value that references variables from outside its body.
	// The function may access and assign to the referenced variables; in this sense the function is "bound" to the variables.
	// For example, the adder function returns a closure. Each closure is bound to its own sum variable.
	// adder function returns a closure function that takes an int and returns the sum of the int and the value of the sum variable
	// The sum variable is shared between all the closures returned by the adder function
	adder := func() func(int) int {
		sum := 0
		return func(x int) int {
			sum += x
			return sum
		}
	}

	// Each closure returned by the adder function is bound to the sum variable
	// The sum variable is shared between all the closures returned by the adder function
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		println(
			pos(i),
			neg(-2*i),
		)
	}
}
