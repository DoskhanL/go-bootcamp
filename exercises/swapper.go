package exercises

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Swapper
//
//  1. Change `color` to "orange"
//     and `color2` to "green" at the same time
//
//     (use multiple-assignment)
//
//  2. Print the variables
//
// EXPECTED OUTPUT
//  orange green
// ---------------------------------------------------------

// Swapper function
func Swapper() {
	// UNCOMMENT THE CODE BELOW:

	color, color2 := "red", "blue"

	// ?
	color, color2 = "orange", "green"

	fmt.Println(color, color2)

	swapperSecond()
}

// ---------------------------------------------------------
// EXERCISE: Swapper #2
//
//  1. Swap the values of `red` and `blue` variables
//
//  2. Print them
//
// EXPECTED OUTPUT
//  blue red
// ---------------------------------------------------------

// SwapperSecond function
func swapperSecond() {
	// UNCOMMENT THE CODE BELOW:

	red, blue := "red", "blue"

	// ?
	red, blue = blue, red
	fmt.Println(red, blue)
}
