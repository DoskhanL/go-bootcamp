package tutorials

import (
	"fmt"
	"os"
	"strconv"
)

func ErrorHandlingMain() {

	convertAge()
	feetToMeter()
}

func convertAge() {
	if len(os.Args) != 2 {
		return
	}

	age := os.Args[1]
	n, err := strconv.Atoi(age)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	fmt.Printf("SUCCESS: Converted %q to %d\n", age, n)
}

func feetToMeter() {
	if len(os.Args) != 2 {
		return
	}
	age := os.Args[1]

	n, err := strconv.ParseFloat(age, 64)
	if err != nil {
		fmt.Printf("error: %q is not a number.\n ", age)
		return
	}
	fmt.Printf("%g feet is a %g meters.\n", n, n*0.3048)
}
