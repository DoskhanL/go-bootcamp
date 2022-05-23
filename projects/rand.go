package projects

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

// RandMain function for simple projects
func RandMain() {

	// rand.Seed(time.Now().UnixNano())
	// guess := 10

	// for n := 0; n != guess; {
	// 	n = rand.Intn(guess + 1)

	// 	fmt.Printf("%d ", n)
	// }
	// fmt.Println()

	luckyNumber()
}

func luckyNumber() {
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

	for i := 0; i < maxTurns; i++ {
		n := rand.Intn(guess + 1)
		if guess == n {
			fmt.Println("🎁 YOU WIN")
			return
		}
	}
	fmt.Println("💀 YOU LOST... Try again?")
}
