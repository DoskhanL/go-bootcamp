package exercises

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// ---------------------------------------------------------
// STORY
//  You want to write a program that will manipulate a
//  given string to uppercase, lowercase, and title case.
//
// EXERCISE: String Manipulator
//
//  1. Get the operation as the first argument.
//
//  2. Get the string to be manipulated as the 2nd argument.
//
// HINT
//  You can find the manipulation functions in the strings
//  package Go documentation (ToLower, ToUpper, Title).
//
// EXPECTED OUTPUT
//
//  go run main.go
//    [command] [string]
//
//    Available commands: lower, upper and title
//
//  go run main.go lower 'OMG!'
//    omg!
//
//  go run main.go upper 'omg!'
//    OMG!
//
//  go run main.go title "mr. charles darwin"
//    Mr. Charles Darwin
//
//  go run main.go genius "mr. charles darwin"
//    Unknown command: "genius"
// ---------------------------------------------------------

//StringManipulator main function
func StringManipulator() {
	const usage = `[command] [string]
	Available commands: lower, upper and title`

	args := os.Args

	if len(args) < 3 {
		fmt.Println(usage)
		return
	}

	cmd, val := args[1], args[2]

	switch strings.ToLower(cmd) {
	case "lower":
		fmt.Println(strings.ToLower(val))
	case "upper":
		fmt.Println(strings.ToUpper(val))
	case "title":
		fmt.Println(strings.Title(val))
	default:
		fmt.Printf("Unknown command: %q\n", cmd)
	}
}

// ---------------------------------------------------------
// EXERCISE: Days in a Month
//
//  Refactor the previous exercise from the if statement
//  section.
//
//  "Print the number of days in a given month."
//
//  Use a switch statement instead of if statements.
// ---------------------------------------------------------

//DaysInMonthSwitch function check
func DaysInMonthSwitch() {
	if len(os.Args) != 2 {
		fmt.Println("Give me a month name")
		return
	}

	year := time.Now().Year()
	leap := year%4 == 0 && (year%100 != 0 || year%400 == 0)

	days, month := 28, os.Args[1]

	switch m := strings.ToLower(month); {
	case m == "april",
		m == "june",
		m == "september",
		m == "november":
		days = 30
	case m == "january",
		m == "march",
		m == "may",
		m == "july",
		m == "august",
		m == "october",
		m == "december":
		days = 31
	case m == "february":
		if leap {
			days = 29
		}
	default:
		fmt.Printf("%q is not a month.\n", month)
		return
	}

	fmt.Printf("%q has %d days.\n", month, days)
}
