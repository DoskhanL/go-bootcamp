package exercises

import (
	"fmt"
	"os"
)

// ---------------------------------------------------------
// EXERCISE: Print Your Name
//
//  Get it from the first command-line argument
//
// INPUT
//  Call the program using your name
//
// EXPECTED OUTPUT
//  It should print your name
//
// EXAMPLE
//  go run main.go inanc
//
//    inanc
//
// BONUS: Make the output like this:
//
//  go run main.go inanc
//    Hi inanc
//    How are you?
// ---------------------------------------------------------

func PrintMe() {
	// get your name from the command-line
	// and print it
	args := os.Args
	if len(args) >= 2 {
		fmt.Println("Hello", args[1])
		fmt.Println("How are you?")
	} else {
		fmt.Println("Enter your name, please?")
	}
}
