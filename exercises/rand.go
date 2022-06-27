package exercises

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// RandMain this game main function
func RandMain() {
	// firstTurnWinner()

	// doubleGuesses()

	// verboseWin()

	// enoughPicks()

	dynamicDifficulty()
}

// ---------------------------------------------------------
// EXERCISE: First Turn Winner
//
//  If the player wins on the first turn, then display
//  a special bonus message.
//
// RESTRICTION
//  The first picked random number by the computer should
//  match the player's guess.
//
// EXAMPLE SCENARIO
//  1. The player guesses 6
//  2. The computer picks a random number and it happens
//     to be 6
//  3. Terminate the game and print the bonus message
// ---------------------------------------------------------

// ---------------------------------------------------------
// EXERCISE: Random Messages
//
//  Display a few different won and lost messages "randomly".
//
// HINTS
//  1. You can use a switch statement to do that.
//  2. Set its condition to the random number generator.
//  3. I would use a short switch.
//
// EXAMPLES
//  The Player wins: then randomly print one of these:
//
//  go run main.go 5
//    YOU WON
//  go run main.go 5
//    YOU'RE AWESOME
//
//  The Player loses: then randomly print one of these:
//
//  go run main.go 5
//    LOSER!
//  go run main.go 5
//    YOU LOST. TRY AGAIN?
// ---------------------------------------------------------

func firstTurnWinner() {
	rand.Seed(time.Now().UnixNano())
	const (
		maxTurns = 5
		usage    = `Welcome to the Luck Number Game! 🍀
The program will pick %d random numbers.
Your mission is to guess on those numbers.

The greater you numbers is, the harder it gets.

Wanna play?
`
	)

	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Printf(usage, maxTurns)
		return
	}
	guess, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Not a number")
		return
	}

	if guess < 0 {
		fmt.Println("Please pick a positive number")
		return
	}

	for turn := 0; turn < maxTurns; turn++ {
		n := rand.Intn(guess + 1)
		if n != guess {
			continue
		}

		if turn == 1 {
			fmt.Println("🥇 FIRST TIME WINNER!!!")
		} else {
			switch rand.Intn(3) {
			case 0:
				fmt.Println("🎉 YOU WON!")
			case 1:
				fmt.Println("🎉 YOU'RE AWESOME")
			case 2:
				fmt.Println("🎉 PERFECT!")
			}
		}
		return
	}

	switch rand.Intn(3) {
	case 0:
		fmt.Println("💀 LOSER!")
	case 1:
		fmt.Println("💀 YOU LOST... Try again?")
	case 2:
		fmt.Println("💀 JUST A BAD LUCK...")
	}
}

// ---------------------------------------------------------
// EXERCISE: Double Guesses
//
//  Let the player guess two numbers instead of one.
//
// HINT:
//  Generate random numbers using the greatest number
//  among the guessed numbers.
//
// EXAMPLES
//  go run main.go 5 6
//  Player wins if the random number is either 5 or 6.
// ---------------------------------------------------------

func doubleGuesses() {

	rand.Seed(time.Now().UnixNano())
	const (
		maxTurns = 5
		usage    = `Welcome to the Luck Number Game! 🍀
The program will pick %d random numbers.
Your mission is to guess on those numbers.

The greater you numbers is, the harder it gets.

Wanna play?
Enter two numbers...
`
	)

	args := os.Args[1:]

	if len(args) < 2 {
		fmt.Printf(usage, maxTurns)
		return
	}
	first, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Not a number")
		return
	}

	second, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("Not a number")
		return
	}

	if first < 0 || second < 0 {
		fmt.Println("Please pick positive numbers")
		return
	}

	var guess int
	if first > second {
		guess = first
	} else {
		guess = second
	}

	for turn := 0; turn < maxTurns; turn++ {
		n := rand.Intn(guess + 1)
		if n != first && n != second {
			continue
		}

		if turn == 1 {
			fmt.Println("🥇 FIRST TIME WINNER!!! Number is", n)
		} else {
			switch rand.Intn(3) {
			case 0:
				fmt.Println("🎉 YOU WON! Number is", n)
			case 1:
				fmt.Println("🎉 YOU'RE AWESOME Number is", n)
			case 2:
				fmt.Println("🎉 PERFECT! Number is", n)
			}
		}
		return
	}

	switch rand.Intn(3) {
	case 0:
		fmt.Println("💀 LOSER!")
	case 1:
		fmt.Println("💀 YOU LOST... Try again?")
	case 2:
		fmt.Println("💀 JUST A BAD LUCK...")
	}

}

func verboseWin() {
	const (
		maxTurns = 5 // less is more difficult
		usage    = `Welcome to the Lucky Number Game! 🍀
The program will pick %d random numbers.
Your mission is to guess one of those numbers.
The greater your number is, harder it gets.
Wanna play?
(Provide -v flag to see the picked numbers.)
`
	)

	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Printf(usage, maxTurns)
		return
	}

	var (
		guess   int
		err     error
		verbose = false
	)

	if args[0] == "-v" {
		verbose = true
	}

	if guess, err = strconv.Atoi(args[len(args)-1]); err != nil {
		fmt.Println("Not a number.")
		return
	}

	if guess <= 0 {
		fmt.Println("Pick a positive number.")
		return
	}

	rand.Seed(time.Now().UnixNano())

	for turn := 1; turn <= maxTurns; turn++ {
		n := rand.Intn(guess) + 1

		if verbose {
			fmt.Printf("%d ", n)
		}

		if n == guess {
			if turn == 1 {
				fmt.Println("🥇 FIRST TIME WINNER!!! Number is", n)
			} else {
				switch rand.Intn(3) {
				case 0:
					fmt.Println("🎉 YOU WON! Number is", n)
				case 1:
					fmt.Println("🎉 YOU'RE AWESOME Number is", n)
				case 2:
					fmt.Println("🎉 PERFECT! Number is", n)
				}
			}
			return
		}
	}

	switch rand.Intn(3) {
	case 0:
		fmt.Println("💀 LOSER!")
	case 1:
		fmt.Println("💀 YOU LOST... Try again?")
	case 2:
		fmt.Println("💀 JUST A BAD LUCK...")
	}

}

// ---------------------------------------------------------
// EXERCISE: Enough Picks
//
//  If the player's guess number is below 10;
//  then make sure that the computer generates a random
//  number between 0 and 10.
//
//  However, if the guess number is above 10; then the
//  computer should generate the numbers
//  between 0 and the guess number.
//
// WHY?
//  This way the game will be harder.
//
//  Because, in the current version of the game, if
//  the player's guess number is for example 3; then the
//  computer picks a random number between 0 and 3.
//
//  So, currently a player can easily win the game.
//
// EXAMPLE
//  Suppose that the player runs the game like this:
//    go run main.go 9
//
//  Or like this:
//    go run main.go 2
//
//    Then the computer should pick a random number
//    between 0-10.
//
//  Or, if the player runs it like this:
//    go run main.go 15
//
//    Then the computer should pick a random number
//    between 0-15.
// ---------------------------------------------------------

func enoughPicks() {
	var (
		maxTurns = 5 // less is more difficult
		usage    = `Welcome to the Lucky Number Game! 🍀
The program will pick %d random numbers.
Your mission is to guess one of those numbers.
The greater your number is, harder it gets.
Wanna play?
`
		guess    int
		err      error
		maxValue int
	)

	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Printf(usage, maxTurns)
		return
	}

	guess, err = strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Not a number")
		return
	}

	if guess < 0 {
		fmt.Println("Pick a positive number")
		return
	}

	maxValue = guess
	if guess < 10 {
		maxValue = 10
	}

	rand.Seed(time.Now().UnixNano())
	for item := 0; item < maxTurns; item++ {
		r := rand.Intn(maxValue) + 1
		fmt.Printf("%d ", r)
		if r == guess {
			fmt.Println("YOU WIN!!!")
			return
		}
	}

	fmt.Println("YOU LOST! TRY AGAIN!")
}

// ---------------------------------------------------------
// EXERCISE: Dynamic Difficulty
//
//  Current game picks only 5 numbers (5 turns).
//
//  Make sure that the game adjust its own difficulty
//  depending on the guess number.
//
// RESTRICTION
//  Do not make the game too easy. Only adjust the
//  difficulty if the guess is above 10.
//
// EXPECTED OUTPUT
//  Suppose that the player runs the game like this:
//    go run main.go 5
//
//    Then the computer should pick 5 random numbers.
//
//  Or, if the player runs it like this:
//    go run main.go 25
//
//    Then the computer may pick 11 random numbers
//    instead.
//
//  Or, if the player runs it like this:
//    go run main.go 100
//
//    Then the computer may pick 30 random numbers
//    instead.
//
//  As you can see, greater guess number causes the
//  game to increase the game turns, which in turn
//  adjust the game's difficulty dynamically.
//
//  Because, greater guess number makes it harder to win.
//  But, this way, game's difficulty will be dynamic.
//  It will adjust its own difficulty depending on the
//  guess number.
// ---------------------------------------------------------

func dynamicDifficulty() {

	var (
		maxTurns = 5 // less is more difficult
		usage    = `Welcome to the Lucky Number Game! 🍀
The program will pick %d random numbers.
Your mission is to guess one of those numbers.
The greater your number is, harder it gets.
Wanna play?
`
		guess int
		err   error
	)

	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println(usage, maxTurns)
		return
	}

	guess, err = strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Not a number")
		return
	}

	if guess <= 0 {
		fmt.Println("Chose a positive number")
		return
	}

	balancer := guess / 4
	fmt.Println(balancer)
	maxTurns += balancer

	fmt.Printf("Max turns is %d\n", maxTurns)
	rand.Seed(time.Now().UnixNano())
	for item := maxTurns; maxTurns > 0; item-- {

		n := rand.Intn(guess) + 1
		fmt.Printf("%d ", n)
		if guess == n {
			fmt.Println("YOU WIN!")
			return
		}
	}

	fmt.Println("YOU LOST! TRY AGAIN...")
}
