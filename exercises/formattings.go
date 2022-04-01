package exercises

import (
	"fmt"
	"os"
)

// ---------------------------------------------------------
// EXERCISE: Print Your Age
//
//  Print your age using Printf
//
// EXPECTED OUTPUT
//  I'm 30 years old.
//
// NOTE
//  You should change 30 to your age, of course.
// ---------------------------------------------------------

func Formattings() {
	// ?
	fmt.Printf("I'm %d years old.\n", 37)

	printYourName()

	falseClaims()

	printTemperature()

	doubleQuotes()

	printType()

	printTypeSecond()

	printTypeThird()

	printTypeFourth()

	printFullName()
}

// ---------------------------------------------------------
// EXERCISE: Print Your Name and LastName
//
//  Print your name and lastname using Printf
//
// EXPECTED OUTPUT
//  My name is Inanc and my lastname is Gumus.
//
// BONUS
//  Store the formatting specifier (first argument of Printf)
//    in a variable.
//  Then pass it to printf
// ---------------------------------------------------------

func printYourName() {
	// BONUS: Use a variable for the format specifier
	var (
		msg      = "My name is %s and my lastname is %s.\n"
		name     = "Doskhan"
		lastName = "Lesbekov"
	)
	fmt.Printf(msg, name, lastName)
}

// ---------------------------------------------------------
// EXERCISE: False Claims
//
//  Use printf to print the expected output using a variable.
//
// EXPECTED OUTPUT
//  These are false claims.
// ---------------------------------------------------------

func falseClaims() {
	// UNCOMMENT THE FOLLOWING CODE
	// AND DO NOT CHANGE IT AFTERWARDS
	tf := false

	// TYPE YOUR CODE HERE
	// ?

	fmt.Printf("These are %t claims.\n", tf)
}

// ---------------------------------------------------------
// EXERCISE: Print the Temperature
//
//  Print the current temperature in your area using Printf
//
// NOTE
//  Do not use %v verb
//  Output "shouldn't" be like 29.500000
//
// EXPECTED OUTPUT
//  Temperature is 29.5 degrees.
// ---------------------------------------------------------

func printTemperature() {
	fmt.Printf("Temperature is %.1f degrees.\n", 29.5)
}

// ---------------------------------------------------------
// EXERCISE: Double Quotes
//
//  Print "hello world" with double-quotes using Printf
//  (As you see here)
//
// NOTE
//  Output "shouldn't" be just: hello world
//
// EXPECTED OUTPUT
//  "hello world"
// ---------------------------------------------------------

func doubleQuotes() {
	fmt.Printf("%q\n", "hello world")
}

// ---------------------------------------------------------
// EXERCISE: Print the Type
//
//  Print the type and value of 3 using fmt.Printf
//
// EXPECTED OUTPUT
//  Type of 3 is int
// ---------------------------------------------------------

func printType() {
	fmt.Printf("Type of %d is %[1]T\n", 3)
}

// ---------------------------------------------------------
// EXERCISE: Print the Type #2
//
//  Print the type and value of 3.14 using fmt.Printf
//
// EXPECTED OUTPUT
//  Type of 3.14 is float64
// ---------------------------------------------------------

func printTypeSecond() {
	pi := 3.14
	fmt.Printf("Type of %.2f is %[1]T\n", pi)
}

// ---------------------------------------------------------
// EXERCISE: Print the Type #3
//
//  Print the type and value of "hello" using fmt.Printf
//
// EXPECTED OUTPUT
// 	Type of hello is string
// ---------------------------------------------------------

func printTypeThird() {
	fmt.Printf("Type of %s is %[1]T\n", "hello")
}

// ---------------------------------------------------------
// EXERCISE: Print the Type #4
//  Print the type and value of true using fmt.Printf
//
// EXPECTED OUTPUT
//  Type of true is bool
// ---------------------------------------------------------

func printTypeFourth() {
	fmt.Printf("Type of %t is %[1]T\n", true)
}

// ---------------------------------------------------------
// EXERCISE: Print Your Fullname
//
//  1. Get your name and lastname from the command-line
//  2. Print them using Printf
//
// EXAMPLE INPUT
//  Inanc Gumus
//
// EXPECTED OUTPUT
//  Your name is Inanc and your lastname is Gumus.
// ---------------------------------------------------------

func printFullName() {
	// BONUS: Use a variable for the format specifier
	var msg = "Your name is %s and your lastname is %s.\n"
	args := os.Args
	if len(args) < 3 {
		return
	}

	var (
		firstName = args[1]
		lastName  = args[2]
	)
	fmt.Printf(msg, firstName, lastName)
}
