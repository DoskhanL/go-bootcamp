package exercises

import (
	"fmt"
	"os"
	"strconv"
)

// SwitchMain is a main exercise function
func SwitchMain() {
	richterScale()

	richterScaleSecond()

	convertUserAccess()
}

// ---------------------------------------------------------
// STORY
//  You're curious about the richter scales. When reporters
//  say: "There's been a 5.5 magnitude earthquake",
//
//  You want to know what that means.
//
//  So, you've decided to write a program to do that for you.
//
// EXERCISE: Richter Scale
//
//  1. Get the earthquake magnitude from the command-line.
//  2. Display its corresponding description.
//
// ---------------------------------------------------------
// MAGNITUDE                    DESCRIPTION
// ---------------------------------------------------------
// Less than 2.0                micro
// 2.0 and less than 3.0        very minor
// 3.0 and less than 4.0        minor
// 4.0 and less than 5.0        light
// 5.0 and less than 6.0        moderate
// 6.0 and less than 7.0        strong
// 7.0 and less than 8.0        major
// 8.0 and less than 10.0       great
// 10.0 or more                 massive
//
// EXPECTED OUTPUT
//  go run main.go
//    Give me the magnitude of the earthquake.
//
//  go run main.go ABC
//    I couldn't get that, sorry.
//
//  go run main.go 0.5
//    0.50 is micro
//
//  go run main.go 2.5
//    2.50 is very minor
//
//  go run main.go 3
//    3.00 is minor
//
//  go run main.go 4.5
//    4.50 is light
//
//  go run main.go 5
//    5.00 is moderate
//
//  go run main.go 6
//    6.00 is strong
//
//  go run main.go 7
//    7.00 is major
//
//  go run main.go 8
//    8.00 is great
//
//  go run main.go 11
//    11.00 is massive
// ---------------------------------------------------------

func richterScale() {
	var (
		args = os.Args
		l    = len(args)
		m    float64
		err  error
	)

	if l < 2 {
		fmt.Println("Give me the magnitude of the earthquake.")
		return
	}

	if m, err = strconv.ParseFloat(args[1], 32); err != nil {
		fmt.Println("I couldn't get that, sorry.")
		return
	}

	switch {
	case m >= 10:
		fmt.Printf("%.2f is massive\n", m)
	case m >= 8:
		fmt.Printf("%.2f is great\n", m)
	case m >= 7:
		fmt.Printf("%.2f is major\n", m)
	case m >= 6:
		fmt.Printf("%.2f is strong\n", m)
	case m >= 5:
		fmt.Printf("%.2f is moderate\n", m)
	case m >= 4:
		fmt.Printf("%.2f is light\n", m)
	case m >= 3:
		fmt.Printf("%.2f is minor\n", m)
	case m >= 2:
		fmt.Printf("%.2f is very minor\n", m)
	default:
		fmt.Printf("%.2f is micro\n", m)
	}

	// Less than 2.0                micro
	// 2.0 and less than 3.0        very minor
	// 3.0 and less than 4.0        minor
	// 4.0 and less than 5.0        light
	// 5.0 and less than 6.0        moderate
	// 6.0 and less than 7.0        strong
	// 7.0 and less than 8.0        major
	// 8.0 and less than 10.0       great
	// 10.0 or more                 massive
}

// ---------------------------------------------------------
// EXERCISE: Richter Scale #2
//
//  Repeat the previous exercise.
//
//  But, this time, get the description and print the
//  corresponding richter scale.
//
//  See the expected outputs.
//
// ---------------------------------------------------------
// MAGNITUDE                    DESCRIPTION
// ---------------------------------------------------------
// Less than 2.0                micro
// 2.0 to less than 3.0         very minor
// 3.0 to less than 4.0         minor
// 4.0 to less than 5.0         light
// 5.0 to less than 6.0         moderate
// 6.0 to less than 7.0         strong
// 7.0 to less than 8.0         major
// 8.0 to less than 10.0        great
// 10.0 or more                 massive
//
// EXPECTED OUTPUT
//  go run main.go
//   Tell me the magnitude of the earthquake in human terms.
//
//  go run main.go alien
//   alien's richter scale is unknown
//
//  go run main.go micro
//   micro's richter scale is less than 2.0
//
//  go run main.go "very minor"
//   very minor's richter scale is 2 - 2.9
//
//  go run main.go minor
//   minor's richter scale is 3 - 3.9
//
//  go run main.go light
//   light's richter scale is 4 - 4.9
//
//  go run main.go moderate
//   moderate's richter scale is 5 - 5.9
//
//  go run main.go strong
//   strong's richter scale is 6 - 6.9
//
//  go run main.go major
//   major's richter scale is 7 - 7.9
//
//  go run main.go great
//   great's richter scale is 8 - 9.9
//
//  go run main.go massive
//   massive's richter scale is 10+
// ---------------------------------------------------------

func richterScaleSecond() {

	var (
		args    = os.Args
		l       = len(args)
		richter string
	)

	if l < 2 {
		fmt.Println("Tell me the magnitude of the earthquake in human terms.")
		return
	}
	desc := args[1]

	switch desc {
	case "micro":
		richter = "less than 2.0"
	case "very minor":
		richter = "2 - 2.9"
	case "minor":
		richter = "3 - 3.9"
	case "light":
		richter = "4 - 4.9"
	case "moderate":
		richter = "5 - 5.9"
	case "strong":
		richter = "6 - 6.9"
	case "major":
		richter = "7 - 7.9"
	case "great":
		richter = "8 - 9.9"
	case "massive":
		richter = "10+"
	default:
		richter = "unknown"
	}

	fmt.Printf("%s's richter scale is %s\n", desc, richter)
}

// ---------------------------------------------------------
// EXERCISE: Convert
//
//  Convert the if statement to a switch statement.
// ---------------------------------------------------------

func convertUserAccess() {
	const (
		usage       = "Usage: [username] [password]"
		errUser     = "Access denied for %q.\n"
		errPwd      = "Invalid password for %q.\n"
		accessOK    = "Access granted to %q.\n"
		user, user2 = "jack", "inanc"
		pass, pass2 = "1888", "1879"
	)

	args := os.Args

	if len(args) != 3 {
		fmt.Println(usage)
		return
	}

	u, p := args[1], args[2]

	//
	// REFACTOR THIS TO A SWITCH
	//
	switch {
	case u != user && u != user2:
		fmt.Printf(errUser, u)
	case u == user && p == pass:
		// fmt.Printf(accessOK, u)
		fallthrough
	case u == user2 && p == pass2:
		fmt.Printf(accessOK, u)
	default:
		fmt.Printf(errPwd, u)
	}
}
