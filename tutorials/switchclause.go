package tutorials

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

// SwitchMain switch condition main function
func SwitchMain() {

	printCityCountry()

	switchBool()

	fallthroughClause()

	shortSwitch()

	partOfDay()

	printSeason()
}

func printCityCountry() {
	var (
		args = os.Args
		l    = len(args)
	)

	if l < 2 {
		fmt.Println("Enter city name")
		return
	}
	city := args[1]
	switch city {
	case "Paris", "Lyon":
		fmt.Println("France")
	case "Tokyo":
		fmt.Println("Japan")
	default:
		fmt.Println("Where?")
	}

	// if city == "Paris" {
	// 	fmt.Println("France")
	// } else if city == "Tokyo" {
	// 	fmt.Println("Japan")
	// }
}

func switchBool() {

	var (
		args = os.Args
		l    = len(args)
		err  error
		n    int
	)
	if l < 2 {
		fmt.Println("Enter a number")
		return
	}

	if n, err = strconv.Atoi(args[1]); err != nil {
		fmt.Printf("%q is not a number\n", args[1])
		return
	}

	switch {
	case n > 0:
		fmt.Println("positive")
	case n < 0:
		fmt.Println("negative")
	default:
		fmt.Println("zero")
	}
}

func fallthroughClause() {
	var (
		args = os.Args
		l    = len(args)
		err  error
		i    int
	)
	if l < 2 {
		fmt.Println("Enter a number")
		return
	}

	if i, err = strconv.Atoi(args[1]); err != nil {
		fmt.Printf("%q is not a number\n", args[1])
		return
	}

	switch {
	case i > 100:
		fmt.Print("big ")
		fallthrough
	case i > 0:
		fmt.Print("positive ")
		fallthrough
	default:
		fmt.Print("number\n")
	}
}

func shortSwitch() {
	switch i := 10; {
	case i > 0:
		fmt.Println("positive")
	case i < 0:
		fmt.Println("negative")
	default:
		fmt.Println("zero")
	}
}

func partOfDay() {
	h := time.Now().Hour()
	fmt.Println(h)
	switch {
	case h >= 18:
		fmt.Println("Good evening!")
	case h >= 12:
		fmt.Println("Good afternoon!")
	case h >= 6:
		fmt.Println("Good morning")
	default:
		fmt.Println("Good night!")
	}
}

func printSeason() {
	if len(os.Args) < 2 {
		fmt.Println("Gimme a month")
		return
	}

	// if m := os.Args[1]; m == "Dec" || m == "Jan" || m == "Feb" {
	// 	fmt.Println("Winter")
	// } else if m == "Mar" || m == "Apr" || m == "May" {
	// 	fmt.Println("Spring")
	// } else if m == "Jun" || m == "Jul" || m == "Aug" {
	// 	fmt.Println("Summer")
	// } else if m == "Sep" || m == "Oct" || m == "Nov" {
	// 	fmt.Println("Fall")
	// } else {
	// 	fmt.Printf("%q is not a correct month\n", m)
	// }

	switch m := os.Args[1]; m {
	case "Dec", "Jan", "Feb":
		fmt.Println("Winter")
	case "Mar", "Apr", "May":
		fmt.Println("Spring")
	case "Jun", "Jul", "Aug":
		fmt.Println("Summer")
	case "Sep", "Oct", "Nov":
		fmt.Println("Fall")
	default:
		fmt.Printf("%q is not a correct month\n", m)
	}
}
