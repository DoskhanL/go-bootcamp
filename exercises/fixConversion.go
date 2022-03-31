package exercises

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Convert and Fix
//
//  Fix the code by using the conversion expression.
//
// EXPECTED OUTPUT
//  15.5
// ---------------------------------------------------------

func FixIt() {
	a, b := 10, 5.5
	fmt.Println(float64(a) + b)

	second()

	third()

	fourth()

	fifth()
}

// ---------------------------------------------------------
// EXERCISE: Convert and Fix #2
//
//  Fix the code by using the conversion expression.
//
// EXPECTED OUTPUT
//  10.5
// ---------------------------------------------------------

func second() {
	a, b := 10, 5.5
	a = int(b)
	fmt.Println(float64(a) + b)
}

// ---------------------------------------------------------
// EXERCISE: Convert and Fix #3
//
//  Fix the code.
//
// EXPECTED OUTPUT
//  5.5
// ---------------------------------------------------------

func third() {
	fmt.Println(5.5)
}

// ---------------------------------------------------------
// EXERCISE: Convert and Fix #4
//
//  Fix the code.
//
// EXPECTED OUTPUT
//  9.5
// ---------------------------------------------------------

func fourth() {
	age := 2
	//fmt.Println(float64(7.5) + float64(int(age)))
	fmt.Println(7.5 + float64(int(age)))
}

// ---------------------------------------------------------
// EXERCISE: Convert and Fix #5
//
//  Fix the code.
//
// HINTS
//   maximum of int8  can be 127
//   maximum of int16 can be 32767
//
// EXPECTED OUTPUT
//  1127
// ---------------------------------------------------------

func fifth() {
	// DO NOT TOUCH THESE VARIABLES
	min := int8(127)
	max := int16(1000)

	// FIX THE CODE HERE
	fmt.Println(max + int16(min))
	//fmt.Println(int8(max), int8(max)+min)

	// EXPLANATION
	//
	// `int8(max)` destroys the information of max
	// It reduces it to 127
	// Which is the maximum value of int8
	//
	// Correct conversion is int16(min)
	// Because, int16 > int8
	// When you do so, min doesn't lose information
	//
	// You will learn more about this in
	// the "Go Type System" section.
}
