package exercises

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// LabelsMain function of main
func LabelsMain() {
	// crunchThePrimes()
	// crunchThePrimesSrc()

	pathSearcher()
}

// ---------------------------------------------------------
// EXERCISE: Crunch the primes
//
//  1. Get numbers from the command-line.
//  2. `for range` over them.
//  4. Convert each one to an int.
//  5. If one of the numbers isn't an int, just skip it.
//  6. Print the ones that are only the prime numbers.
//
// RESTRICTION
//  The user can run the program with any number of
//  arguments.
//
// HINT
//  Find whether a number is prime using this algorithm:
//  https://stackoverflow.com/a/1801446/115363
//
// EXPECTED OUTPUT
//  go run main.go 2 4 6 7 a 9 c 11 x 12 13
//    2 7 11 13
//
//  go run main.go 1 2 3 5 7 A B C
//    2 3 5 7
// ---------------------------------------------------------
func crunchThePrimes() {
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("Enter the numbers")
		return
	}

	var (
		n   int
		err error
	)

	for _, num := range args {
		n, err = strconv.Atoi(num)
		if err != nil {
			continue
		}

		if isPrime(n) {
			fmt.Printf("%d ", n)
		}
	}
	fmt.Println("")
}

func isPrime(n int) bool {
	if n == 2 || n == 3 {
		return true
	}

	if n <= 1 || n%2 == 0 || n%3 == 0 {
		return false
	}

	i := 5
	w := 2

	for i*i <= n {
		if n%i == 0 {
			return false
		}
		i += w
		w = 6 - w
	}

	return true

}

func crunchThePrimesSrc() {

main:
	for _, arg := range os.Args[1:] {
		n, err := strconv.Atoi(arg)
		if err != nil {
			// skip non-numerics
			continue
		}

		switch {
		// prime
		case n == 2 || n == 3:
			fmt.Print(n, " ")
			continue

		// not a prime
		case n <= 1 || n%2 == 0 || n%3 == 0:
			continue
		}

		for i, w := 5, 2; i*i <= n; {
			// not a prime
			if n%i == 0 {
				continue main
			}

			i += w
			w = 6 - w
		}

		// all checks ok: it's a prime
		fmt.Print(n, " ")
	}

	fmt.Println()
}

// ---------------------------------------------------------
// EXERCISE: Path Searcher
//
//  Your program should search inside the path environment
//  variable.
//
//  Remove the corpus constant then get the corpus from the
//  environment variable "Path" or "PATH".
//
// HINTS
//  1. Search the web to find out what is an environment
//     variable and how to use it (if you don't know
//     what it is already).
//
//  2. Look up for the necessary functions for getting
//     an environment variable. It's in the "os" package.
//
//     Search for it on the Go online documentation.
//
//  3. Look up for the necessary function for splitting
//     the path variable into directory names.
//
//     It's in the "path/filepath" package.
//
// EXAMPLE
//  For example, on my Mac, my PATH environment variable
//  looks like this:
//
//    "/usr/local/bin:/sbin:/Users/inanc/go/bin"
//
//  So, if the user runs the program like this:
//
//    go run main.go /sbin
//
//  It should print this:
//
//    #2 : "/sbin"
// ---------------------------------------------------------

// ---------------------------------------------------------
// BONUS EXERCISE
//  Make your program cross platform. So, it can search
//  the path environment variable when you run it on
//  a Windows or on a Mac (OS X) or on a Linux.
//
// BONUS EXERCISE #2
//  Also find the directories for the directories that
//  contains the searched query.
//
//  And let it match in a case-insensitive manner. See the examples.
//
//  EXAMPLE:
//    Let's say:
//      PATH="/usr/local/bin:/sbin:/Users/inanc/go/bin"
//
//  So, if the user runs the program like this:
//    go run main.go bin
//
//  It should print this:
//    #1 : "/usr/local/bin"
//    #2 : "/sbin"
//    #3 : "/Users/inanc/go/bin"
//
//  Or like this (case insensitive):
//    go run main.go Us
//
//  Output:
//    #1 : "/usr/local/bin"
//    #2 : "/Users/inanc/go/bin"
// ---------------------------------------------------------

func pathSearcher() {

	query := os.Args[1:]
	var (
		paths []string
	)
	path := os.Getenv("PATH")

	paths = filepath.SplitList(path)
	// paths = strings.Split(path, string(os.PathListSeparator))

	// if strings.Contains(path, ":") {
	// 	paths = strings.Split(path, ":")
	// } else if strings.Contains(path, ";") {
	// 	paths = strings.Split(path, ";")
	// }

	for _, q := range query {
		for i, w := range paths {
			q, w = strings.ToLower(q), strings.ToLower(w)
			if strings.Contains(w, q) {
				fmt.Printf("#%-2d: %q\n", i+1, w)
			}
		}
	}
}
