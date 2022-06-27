package tutorials

import (
	"fmt"
	"os"
	"strings"
)

const corpus = "lazy cat jumps again and again and again"

// WordFinder is a little search project
func WordFinder() {
	//	continueBreak()
	gotoLabel()
}

func gotoLabel() {
	var i int

loop:
	if i < 3 {
		fmt.Printf("#%d looping\n", i+1)
		i++
		goto loop
	}

	fmt.Println("done")
}

func continueBreak() {

	query := os.Args[1:]

	if len(query) < 1 {
		fmt.Println("Enter the search words")
		return
	}
	words := strings.Fields(corpus)

queries:
	for _, q := range query {
		// queries:
	search:
		for i, w := range words {
			switch q {
			case "or", "and", "the":
				break search
			}

			if w == strings.ToLower(q) {
				fmt.Printf("#%-2d: %q\n", i+1, w)
				// break queries
				continue queries
			}
		}
	}
}
