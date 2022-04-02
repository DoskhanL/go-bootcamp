package exercises

import (
	"fmt"
	"os"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Age Seasons
//
//  Let's start simple. Print the expected outputs,
//  depending on the age variable.
//
// EXPECTED OUTPUT
//  If age is greater than 60, print:
//    Getting older
//  If age is greater than 30, print:
//    Getting wiser
//  If age is greater than 20, print:
//    Adulthood
//  If age is greater than 10, print:
//    Young blood
//  Otherwise, print:
//    Booting up
// ---------------------------------------------------------

func MainIf() {
	// Change this accordingly to produce the expected outputs
	age := 10

	// Type your if statement here.

	if age > 60 {
		fmt.Println("Getting older")
	} else if age > 30 {
		fmt.Println("Getting wiser")
	} else if age > 20 {
		fmt.Println("Adulthood")
	} else if age > 10 {
		fmt.Println("Young blood")
	} else {
		fmt.Println("Booting up")
	}

	simplifyIt()

	argCount()

	vowelOrCons()

	challengeFirst()

	challengeSecond()
}

// ---------------------------------------------------------
// EXERCISE: Simplify It
//
//  Can you simplify the if statement inside the code below?
//
//  When:
//    isSphere == true and
//    radius is equal or greater than 200
//
//    It will print "It's a big sphere."
//
//    Otherwise, it will print "I don't know."
//
// EXPECTED OUTPUT
//  It's a big sphere.
// ---------------------------------------------------------

func simplifyIt() {
	// DO NOT TOUCH THIS
	isSphere, radius := true, 200

	if isSphere && radius >= 200 {
		fmt.Println("It's a big sphere.")
	} else {
		fmt.Println("I don't know.")
	}
}

// ---------------------------------------------------------
// EXERCISE: Arg Count
//
//  1. Get arguments from command-line.
//  2. Print the expected outputs below depending on the number
//     of arguments.
//
// EXPECTED OUTPUT
//  go run main.go
//    Give me args
//
//  go run main.go hello
//    There is one: "hello"
//
//  go run main.go hi there
//    There are two: "hi there"
//
//  go run main.go I wanna be a gopher
//    There are 5 arguments
// ---------------------------------------------------------

func argCount() {
	var (
		args = os.Args
		l    = len(args) - 1
	)
	nums := make(map[int]string)
	var values string

	nums[0] = "one"
	nums[1] = "two"
	nums[2] = "free"
	nums[3] = "four"
	nums[4] = "five"
	nums[5] = "six"
	nums[6] = "seven"
	nums[7] = "eight"
	nums[8] = "nine"
	nums[9] = "ten"

	if l == 0 {
		fmt.Println("Give me args")
	} else if l == 1 {
		fmt.Printf("There is %s: %q\n", nums[l-1], args[1])
	} else {
		values = strings.Join(args[1:], " ")
		fmt.Printf("There are %s: %q\n", nums[l-1], values)
	}

}

// ---------------------------------------------------------
// EXERCISE: Vowel or Consonant
//
//  Detect whether a letter is vowel or consonant.
//
// NOTE
//  y or w is called a semi-vowel.
//  Check out: https://en.oxforddictionaries.com/explore/is-the-letter-y-a-vowel-or-a-consonant/
//
// HINT
//  + You can find the length of an argument using the len function.
//
//  + len(os.Args[1]) will give you the length of the 1st argument.
//
// BONUS
//  Use strings.IndexAny function to detect the vowels.
//  Search on Google for: golang pkg strings IndexAny
//
// Furthermore, you can also use strings.ContainsAny. Check out: https://golang.org/pkg/strings/#ContainsAny
//
// EXPECTED OUTPUT
//  go run main.go
//    Give me a letter
//
//  go run main.go hey
//    Give me a letter
//
//  go run main.go a
//    "a" is a vowel.
//
//  go run main.go y
//    "y" is sometimes a vowel, sometimes not.
//
//  go run main.go w
//    "w" is sometimes a vowel, sometimes not.
//
//  go run main.go x
//    "x" is a consonant.
// ---------------------------------------------------------

func vowelOrCons() {
	var (
		args = os.Args
		l    = len(args) - 1
	)
	if l >= 1 && len(args[1]) == 1 {
		char := args[1]
		if char == "y" || char == "w" {
			fmt.Printf("%q is sometimes a vowel, sometimes not.\n", char)
		} else if strings.IndexAny(char, "aeiou") != -1 || isVowel(char) {
			fmt.Printf("%q is a vowel.\n", char)
		} else {
			fmt.Printf("%q is a consonant.\n", char)
		}
	} else {
		fmt.Println("Give me a letter")
	}
}

func isVowel(char string) bool {
	var vowels = []string{"a", "e", "i", "o", "u"}
	for _, v := range vowels {
		if v == char {
			return true
		}
	}
	return false
}

// ---------------------------------------------------------
// CHALLENGE #1
//  Create a user/password protected program.
//
// EXAMPLE USER
//  username: jack
//  password: 1888
//
// EXPECTED OUTPUT
//  go run main.go
//    Usage: [username] [password]
//
//  go run main.go albert
//    Usage: [username] [password]
//
//  go run main.go hacker 42
//    Access denied for "hacker".
//
//  go run main.go jack 6475
//    Invalid password for "jack".
//
//  go run main.go jack 1888
//    Access granted to "jack".
// ---------------------------------------------------------

func challengeFirst() {
	var (
		args     = os.Args
		l        = len(args) - 1
		username = "jack"
		password = "1888"
	)

	if l < 2 {
		fmt.Println("Usage: [username] [password]")
		return
	}

	if args[1] != username {
		fmt.Printf("Access denied for %q\n", args[1])
	} else if args[2] != password {
		fmt.Printf("Invalid password for %q\n", username)
	} else {
		fmt.Printf("Access granted to %q\n", username)
	}

}

// ---------------------------------------------------------
// CHALLENGE #2
//  Add one more user to the PassMe program below.
//
// EXAMPLE USERS
//  username: jack
//  password: 1888
//
//  username: inanc
//  password: 1879
//
// EXPECTED OUTPUT
//  go run main.go
//    Usage: [username] [password]
//
//  go run main.go hacker 42
//    Access denied for "hacker".
//
//  go run main.go jack 1888
//    Access granted to "jack".
//
//  go run main.go inanc 1879
//    Access granted to "inanc".
//
//  go run main.go jack 1879
//    Invalid password for "jack".
//
//  go run main.go inanc 1888
//    Invalid password for "inanc".
// ---------------------------------------------------------

const (
	usage        = "Usage: [username] [password]"
	errUser      = "Access denied for %q.\n"
	errPwd       = "Invalid password for %q.\n"
	accessOK     = "Access granted to %q.\n"
	user1, user2 = "jack", "inanc"
	pass1, pass2 = "1888", "1879"
)

func challengeSecond() {
	args := os.Args

	if len(args) != 3 {
		fmt.Println(usage)
		return
	}

	u, p := args[1], args[2]

	if u != user1 && u != user2 {
		fmt.Printf(errUser, u)
	} else if p != pass1 && p != pass2 {
		fmt.Printf(errPwd, u)
	} else {
		fmt.Printf(accessOK, u)
	}
}
