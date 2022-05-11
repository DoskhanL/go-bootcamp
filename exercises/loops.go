package exercises

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
	"time"
)

// LoopMain main entry for loop exercises
func LoopMain() {
	// loopTable()

	//mathOperations()
	//fmt.Println(strings.Repeat("-", 30))
	//mathOperationsSrc()

	// sumTheNumbers()

	// sumTheNumbersVerbose()
	// fmt.Println(strings.Repeat("-", 45))
	// sumTheNumbersVerboseSrc()

	// sumUpToN()

	// sumUpToNSrc()

	// sumUpToNEven()
	// fmt.Println(strings.Repeat("-", 50))
	// sumUpToNEvenSrc()

	// breakUp()

	// infiniteKill()

	infiniteKillSrc()
}

// ---------------------------------------------------------
// EXERCISE: Dynamic Table
//
//  Get the size of the table from the command-line
//    Passing 5 should create a 5x5 table
//    Passing 10 for a 10x10 table
//
// RESTRICTION
//  Solve this exercise without looking at the original
//  multiplication table exercise.
//
// HINT
//  There was a max constant in the original program.
//  That determines the size of the table.
//
// EXPECTED OUTPUT
//
//  go run main.go
//    Give me the size of the table
//
//  go run main.go -5
//    Wrong size
//
//  go run main.go ABC
//    Wrong size
//
//  go run main.go 1
//    X    0    1
//    0    0    0
//    1    0    1
//
//  go run main.go 2
//    X    0    1    2
//    0    0    0    0
//    1    0    1    2
//    2    0    2    4
//
//  go run main.go 3
//    X    0    1    2    3
//    0    0    0    0    0
//    1    0    1    2    3
//    2    0    2    4    6
//    3    0    3    6    9
// ---------------------------------------------------------

func loopTable() {

	args := os.Args

	if len(args) < 2 {
		fmt.Println("Give me the size of the table")
		return
	}

	max, err := strconv.Atoi(args[1])
	if err != nil || max <= 0 {
		fmt.Println("Wrong size")
		return
	}

	fmt.Printf("%5s", "X")
	for i := 0; i <= max; i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Println()

	for i := 0; i <= max; i++ {
		fmt.Printf("%5d", i)

		for j := 0; j <= max; j++ {
			fmt.Printf("%5d", i*j)
		}
		fmt.Println()
	}
}

// ---------------------------------------------------------
// EXERCISE: Math Tables
//
//  Create division, addition and subtraction tables
//
//  1. Get the math operation and
//     the size of the table from the user
//
//  2. Print the table accordingly
//
//  3. Correctly handle the division by zero error
//
//
// BONUS #1
//
//  Use strings.IndexAny function to detect
//    the valid operations.
//
//  Search on Google for: golang pkg strings IndexAny
//
// BONUS #2
//
//  Add remainder operation as well (remainder table using %).
//
//
// EXPECTED OUTPUT
//
//  go run main.go
//    Usage: [op=*/+-] [size]
//
//  go run main.go "*"
//    Size is missing
//    Usage: [op=*/+-] [size]
//
//  go run main.go "%" 4
//    Invalid operator.
//    Valid ops one of: */+-
//
//  go run main.go "*" 4
//    *    0    1    2    3    4
//    0    0    0    0    0    0
//    1    0    1    2    3    4
//    2    0    2    4    6    8
//    3    0    3    6    9   12
//    4    0    4    8   12   16
//
//  go run main.go "/" 4
//    /    0    1    2    3    4
//    0    0    0    0    0    0
//    1    0    1    0    0    0
//    2    0    2    1    0    0
//    3    0    3    1    1    0
//    4    0    4    2    1    1
//
//  go run main.go "+" 4
//    +    0    1    2    3    4
//    0    0    1    2    3    4
//    1    1    2    3    4    5
//    2    2    3    4    5    6
//    3    3    4    5    6    7
//    4    4    5    6    7    8
//
//  go run main.go "-" 4
//    -    0    1    2    3    4
//    0    0   -1   -2   -3   -4
//    1    1    0   -1   -2   -3
//    2    2    1    0   -1   -2
//    3    3    2    1    0   -1
//    4    4    3    2    1    0
//
// BONUS:
//
//  go run main.go "%" 4
//    %    0    1    2    3    4
//    0    0    0    0    0    0
//    1    0    0    1    1    1
//    2    0    0    0    2    2
//    3    0    0    1    0    3
//    4    0    0    0    1    0
//
// NOTES:
//
//   When running the program in Windows Git Bash, you might need
//   to escape the characters like this.
//
//   This happens because those characters have special meaning.
//
//   Division:
//     go run main.go // 4
//
//   Multiplication and others:
//   (this is also necessary for non-Windows bashes):
//
//     go run main.go "*" 4
// ---------------------------------------------------------

func mathOperations() {

	validOps := "%*/+-"
	usage := "Usage: [op=" + validOps + "] [size]"
	args := os.Args

	if len(args) < 2 {
		fmt.Println(usage)
		return
	}

	op := args[1]
	mIdx := strings.IndexAny(op, "*")
	dIdx := strings.IndexAny(op, "/")
	aIdx := strings.IndexAny(op, "+")
	sIdx := strings.IndexAny(op, "-")
	rIdx := strings.IndexAny(op, "%")

	if mIdx == -1 && dIdx == -1 && aIdx == -1 && sIdx == -1 && rIdx == -1 {
		fmt.Println(`Invalid operator. 
Valid ops one of: ` + validOps)
		return
	}

	if len(args) < 3 {
		fmt.Println("Size is missing")
		fmt.Println(usage)
		return
	}

	max, err := strconv.Atoi(args[2])
	if err != nil || max <= 0 {
		fmt.Println("Wrong size")
		return
	}

	fmt.Printf("%5s", op)
	for i := 0; i <= max; i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Println()

	for i := 0; i <= max; i++ {
		fmt.Printf("%5d", i)

		for j := 0; j <= max; j++ {
			if mIdx != -1 {
				fmt.Printf("%5d", i*j)
			} else if dIdx != -1 {
				if j != 0 {
					fmt.Printf("%5d", i/j)
				} else {
					fmt.Printf("%5d", 0)
				}
			} else if rIdx != -1 {
				if j != 0 {
					fmt.Printf("%5d", i%j)
				} else {
					fmt.Printf("%5d", 0)
				}
			} else if aIdx != -1 {
				fmt.Printf("%5d", i+j)
			} else if sIdx != -1 {
				fmt.Printf("%5d", i-j)
			}
		}
		fmt.Println()
	}
}

func mathOperationsSrc() {

	const (
		validOps       = "%*/+-"
		usageMsg       = "Usage: [op=" + validOps + "] [size]"
		sizeMissingMsg = "Size is missing"
		invalidOpMsg   = `Invalid operator.
Valid ops one of: ` + validOps
		invalidOp = -1
	)

	args := os.Args[1:]

	switch l := len(args); {
	case l == 1:
		fmt.Println(sizeMissingMsg)
		fallthrough
	case l < 1:
		fmt.Println(usageMsg)
		return
	}

	op := args[0]
	if strings.IndexAny(op, validOps) == invalidOp {
		fmt.Println(invalidOpMsg)
		return
	}

	size, err := strconv.Atoi(args[1])
	if err != nil || size <= 0 {
		fmt.Println("Wrong size")
		return
	}

	fmt.Printf("%5s", op)
	for i := 0; i <= size; i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Println()

	for i := 0; i <= size; i++ {
		fmt.Printf("%5d", i)

		for j := 0; j <= size; j++ {
			var res int

			switch op {
			case "*":
				res = i * j
			case "%":
				if j != 0 {
					res = i % j
				}
			case "/":
				if j != 0 {
					res = i / j
				}
			case "+":
				res = i + j
			case "-":
				res = i - j
			}

			fmt.Printf("%5d", res)
		}
		fmt.Println()
	}
}

// ---------------------------------------------------------
// EXERCISE: Sum the Numbers
//
//  1. By using a loop, sum the numbers between 1 and 10.
//  2. Print the sum.
//
// EXPECTED OUTPUT
//  Sum: 55
// ---------------------------------------------------------

func sumTheNumbers() {
	var sum int
	for i := 1; i <= 10; i++ {
		sum += i
	}

	fmt.Printf("Sum: %d\n", sum)
}

// ---------------------------------------------------------
// EXERCISE: Sum the Numbers: Verbose Edition
//
//  By using a loop, sum the numbers between 1 and 10.
//
// HINT
//  1. For printing it as in the expected output, use Print
//     and Printf functions. They don't print a newline
//     automatically (unlike a Println).
//
//  2. Also, you need to use an if statement to prevent
//     printing the last plus sign.
//
// EXPECTED OUTPUT
//    1 + 2 + 3 + 4 + 5 + 6 + 7 + 8 + 9 + 10 = 55
// ---------------------------------------------------------

func sumTheNumbersVerbose() {
	var (
		res string
		sum int
	)
	for i := 1; i <= 10; i++ {
		res += fmt.Sprintf("%d", i)
		sum += i
		if i == 10 {
			res += fmt.Sprintf(" = %d", sum)
		} else {
			res += " + "
		}
	}

	fmt.Println(res)
}

func sumTheNumbersVerboseSrc() {
	const min, max = 1, 10

	var sum int
	for i := min; i <= max; i++ {
		sum += i

		fmt.Print(i)
		if i < max {
			fmt.Print(" + ")
		}
	}
	fmt.Printf(" = %d\n", sum)
}

// ---------------------------------------------------------
// EXERCISE: Sum up to N
//
//  1. Get two numbers from the command-line: min and max
//  2. Convert them to integers (using Atoi)
//  3. By using a loop, sum the numbers between min and max
//
// RESTRICTIONS
//  1. Be sure to handle the errors. So, if a user doesn't
//     pass enough arguments or she passes non-numerics,
//     then warn the user and exit from the program.
//
//  2. Also, check that, min < max.
//
// HINT
//  For converting the numbers, you can use `strconv.Atoi`.
//
// EXPECTED OUTPUT
//  Let's suppose that the user runs it like this:
//
//    go run main.go 1 10
//
//  Then it should print:
//
//    1 + 2 + 3 + 4 + 5 + 6 + 7 + 8 + 9 + 10 = 55
// ---------------------------------------------------------

func sumUpToN() {

	args := os.Args[1:]
	var (
		usage         = "Enter the min and max numbers"
		min, max, sum int
	)

	if len(args) < 1 {
		fmt.Println(usage)
		return
	}

	min, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Enter the correct min number")
		return
	}

	if len(args) < 2 {
		fmt.Println("Enter the max number")
		return
	}
	max, err = strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Enter the correct max number")
		return
	}

	if min > max {
		fmt.Println("Enter the correct min and max numbers. Min number must be less than the max")
		return
	}

	for i := min; i <= max; i++ {
		sum += i
		fmt.Print(i)

		if i < max {
			fmt.Print(" + ")
		}
	}

	fmt.Printf(" = %d\n", sum)
}

func sumUpToNSrc() {
	if len(os.Args) != 3 {
		fmt.Println("gimme two numbers")
		return
	}

	min, err1 := strconv.Atoi(os.Args[1])
	max, err2 := strconv.Atoi(os.Args[2])
	if err1 != nil || err2 != nil || min >= max {
		fmt.Println("wrong numbers")
		return
	}

	var sum int
	for i := min; i <= max; i++ {
		sum += i

		fmt.Print(i)
		if i < max {
			fmt.Print(" + ")
		}
	}
	fmt.Printf(" = %d\n", sum)
}

// ---------------------------------------------------------
// EXERCISE: Only Evens
//
//  1. Extend the "Sum up to N" exercise
//  2. Sum only the even numbers
//
// RESTRICTIONS
//  Skip odd numbers using the `continue` statement
//
// EXPECTED OUTPUT
//  Let's suppose that the user runs it like this:
//
//    go run main.go 1 10
//
//  Then it should print:
//
//    2 + 4 + 6 + 8 + 10 = 30
// ---------------------------------------------------------

func sumUpToNEven() {

	args := os.Args[1:]
	var (
		usage         = "Enter the min and max numbers"
		min, max, sum int
	)

	if len(args) < 1 {
		fmt.Println(usage)
		return
	}

	min, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Enter the correct min number")
		return
	}

	if len(args) < 2 {
		fmt.Println("Enter the max number")
		return
	}
	max, err = strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Enter the correct max number")
		return
	}

	if min > max {
		fmt.Println("Enter the correct min and max numbers. Min number must be less than the max")
		return
	}

	for i := min; i <= max; i++ {
		if i%2 != 0 {
			continue
		}
		sum += i
		fmt.Print(i)

		if i < max-1 {
			fmt.Print(" + ")
		}
	}

	fmt.Printf(" = %d\n", sum)
}

func sumUpToNEvenSrc() {
	if len(os.Args) < 3 {
		fmt.Println("gimme two numbers")
		return
	}

	min, err1 := strconv.Atoi(os.Args[1])
	max, err2 := strconv.Atoi(os.Args[2])
	if err1 != nil || err2 != nil || min >= max {
		fmt.Println("wrong numbers")
		return
	}

	var sum int
	for i := min; i <= max; i++ {
		if i%2 != 0 {
			continue
		}
		sum += i

		fmt.Print(i)
		if i < max-1 {
			fmt.Print(" + ")
		}
	}
	fmt.Printf(" = %d\n", sum)
}

// ---------------------------------------------------------
// EXERCISE: Break Up
//
//  1. Extend the "Only Evens" exercise
//  2. This time, use an infinite loop.
//  3. Break the loop when it reaches to the `max`.
//
// RESTRICTIONS
//  You should use the `break` statement once.
//
// HINT
//  Do not forget incrementing the `i` before the `continue`
//  statement and at the end of the loop.
//
// EXPECTED OUTPUT
//    go run main.go 1 10
//    2 + 4 + 6 + 8 + 10 = 30
// ---------------------------------------------------------

func breakUp() {

	var (
		min, max, sum int
	)
	args := os.Args[1:]

	if len(args) < 2 {
		fmt.Println("Enter the min and max")
		return
	}

	min, minErr := strconv.Atoi(args[0])
	max, maxErr := strconv.Atoi(args[1])

	if minErr != nil || maxErr != nil {
		fmt.Println("Wrong numbers")
		return
	}

	if min > max {
		fmt.Println("Min must be less than max")
		return
	}
	i := min
	for {
		if i > max {
			break
		} else if i%2 != 0 {
			i++
			continue
		}
		fmt.Print(i)
		if i < max-1 {
			fmt.Print(" + ")
		}
		sum += i
		i++
	}

	fmt.Printf(" = %d\n", sum)
}

// ---------------------------------------------------------
// EXERCISE: Infinite Kill
//
//  1. Create an infinite loop
//  2. On each step of the loop print a random character.
//  3. And, sleep for 250 milliseconds.
//  4. Run the program and wait a couple of seconds
//     then kill it using CTRL+C keys
//
// RESTRICTIONS
//  1. Print one of those characters randomly: \ / | -
//  2. Before printing a character print also this
//     escape sequence: \r
//
//     Like this: "\r/", or this: "\r|", and so on...
//
//  3. NOTE : If you're using Go Playground, use "\f" instead of "\r"
//
// HINTS
//  1. Use `time.Sleep` to sleep.
//  2. You should pass a `time.Duration` value to it.
//  3. Check out the Go online documentation for the
//     `Millisecond` constant.
//  4. When printing, do not use a newline! Or a Println!
//     Use Print or Printf instead.
//
// NOTE
//  If this exercise is hard for you now, wait until the
//  lucky number lecture. Even then so, then just skip it.
//
//  You can return back to it afterwards.
//
// EXPECTED OUTPUT
//  - The program should display the following messages in any order.
//  - And, the first character (\, -, /, or |) should be randomly
//    displayed.
//  - \r or \f sequence will clear the line.
//
//  \ Please Wait. Processing....
//  - Please Wait. Processing....
//  / Please Wait. Processing....
//  | Please Wait. Processing....
// ---------------------------------------------------------

func infiniteKill() {
	const (
		warnMsg = "Please Wait. Processing...."
	)

	for {
		rnChar, err := generateRandomChar(1)
		if err != nil {
			fmt.Printf("err while generating random char %s", err.Error())
		}
		fmt.Printf("\r%s %s", rnChar, warnMsg)

		time.Sleep(time.Millisecond * 250)
	}

}

func infiniteKillSrc() {
	for {
		var c string
		var n *big.Int
		var err error

		if n, err = rand.Int(rand.Reader, big.NewInt(int64(4))); err != nil {
			n = big.NewInt(int64(0))
		}
		switch n.Int64() {
		case 0:
			c = "\\"
		case 1:
			c = "/"
		case 2:
			c = "|"
		case 3:
			c = "-"
		}
		fmt.Printf("\r%s Please Wait. Processing....", c)
		time.Sleep(time.Millisecond * 250)
	}
}

func generateRandomChar(n int) (string, error) {
	const letters = "/\\-|"
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		ret[i] = letters[num.Int64()]
	}

	return string(ret), nil
}
