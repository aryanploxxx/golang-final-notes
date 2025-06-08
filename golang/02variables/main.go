package main

import "fmt"

// const value must be known at compile time.
const LoginToken string = "kjbsfbkcbwefj"

// capital 'L' has a significant importance in Go, if a variable starts with capital letter, it is public, else private
// correct way to declare anything outside the main function

// jwtToken := 300000
// This will give error, because we cannot use walrus operator outside of a function

func main() {
	// fmt.Println("Variables")
	var username string = "aryan"
	fmt.Println("Username: ", username)
	fmt.Printf("Type of Username: %T \n", username)
	// %T is used for type of the variable

	var isLoggedIn bool = true
	fmt.Println("isLoggedIn: ", isLoggedIn)
	fmt.Printf("Type of isLoggedIn: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println("smallVal: ", smallVal)
	fmt.Printf("Type of smallVal: %T \n", smallVal)

	var smallFloat float32 = 255.37868678678
	fmt.Println("smallFloat: ", smallFloat)
	fmt.Printf("Type of smallFloat: %T \n", smallFloat)
	// Difference between float32 and float64 is the precision of the floating point number,
	// float64 give more precise answers that is more decimal points

	// Default values and aliases
	var anotherValue bool
	fmt.Println("anotherValue: ", anotherValue)
	fmt.Printf("Type of anotherValue: %T \n", anotherValue)
	// Go will assign default values to variables if not assigned i.e 0 for int and float, false for bool, "" for string
	// It does not assign garbage values

	// Implicit type declaration
	var anotherUser = "aryan"
	fmt.Println("anotherUser: ", anotherUser)
	fmt.Printf("Type of anotherUser: %T \n", anotherUser)
	// anotherUser = 3 -> will give error, as lexer of Go would have already assigned a type to it

	// No var style
	numberOfUsers := 1000
	fmt.Println("numberOfUsers: ", numberOfUsers)

	fmt.Println("LoginToken: ", LoginToken)
	fmt.Printf("Type of LoginToken: %T \n", LoginToken)

	var i int
	var f float64
	var b bool
	var s string = "wow"
	fmt.Printf("%v %v %v %v\n", i, f, b, s)
	// 0 0 false wow
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
	// 0 0 false "wow"

	// -> %v is the default format, has the same effect as not specifying a format
	// -> %q is used to display a string with quotes

	// %g verb is used to format a floating-point number in a compact and convenient way

	// var x, y int = 3, 4
	// var f float64 = math.Sqrt(float64(x*x + y*y))
	// var z uint = uint(f)
	// fmt.Println(x, y, z)

	// Constants cannot be declared using the := syntax.

	var t = 123                  //Type Inferred will be int
	var u = "circle"             //Type Inferred will be string
	var v = 5.6                  //Type Inferred will be float64
	var w = true                 //Type Inferred will be bool
	var x = 'a'                  //Type Inferred will be rune
	var y = 3 + 5i               //Type Inferred will be complex128
	var z = sample{name: "test"} //Type Inferred will be main.sample

	fmt.Printf("Type: %T Value: %v\n", t, t)
	fmt.Printf("Type: %T Value: %v\n", u, u)
	fmt.Printf("Type: %T Value: %v\n", v, v)
	fmt.Printf("Type: %T Value: %v\n", w, w)
	fmt.Printf("Type: %T Value: %v\n", x, x)
	fmt.Printf("Type: %T Value: %v\n", y, y)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}

type sample struct {
	name string
}

// In Go, when you define a struct or any other named type within a package, it is always fully qualified by the package it is defined in. The fully qualified type name includes the package name as a prefix. This is because Go has a package-based namespace system, and the type's full name is needed to uniquely identify it, especially in larger projects where different packages may define types with the same name.
// In your code, the sample struct is defined in the main package, so its fully qualified name is main.sample. The type inference system of Go uses this full name when inferring the type of z.
// This is why the output for z's type is main.sample and not just sample. If the struct were defined in another package, for example, mypackage, and you imported that package, the type would be mypackage.sample.
// This approach ensures clarity and avoids naming conflicts across packages.
