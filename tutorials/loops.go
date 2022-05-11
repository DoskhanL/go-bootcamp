package tutorials

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

//MainLoop function for main loop
func MainLoop() {
	var (
		sum int
		i   = 1
	)

	for {
		if i > 5 {
			break
		}
		sum += i
		fmt.Println(i, "-->", sum)
		i++
	}

	fmt.Println(sum)

	logOddNumbers()

	multiplicationTable()

	parseFields()
}

func logOddNumbers() {
	var (
		sum int
		i   = 1
	)
	for {
		if i > 10 {
			break
		}

		if i%2 != 0 {
			i++
			continue
		}
		sum += i
		fmt.Println(i, "-->", sum)
		i++
	}

	fmt.Println(sum)
}

func multiplicationTable() {
	args := os.Args
	if len(args) < 3 {
		fmt.Println("Give an number for size and type of table data multiply, divide, add, subtract")
		return
	}

	max, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println(fmt.Errorf("Error while parsing int %s", err.Error()))
		return
	}
	operation := strings.ToLower(args[2])
	operations := make(map[string]struct{})
	operations["add"] = struct{}{}
	operations["subtract"] = struct{}{}
	operations["multiply"] = struct{}{}
	operations["divide"] = struct{}{}

	if _, ok := operations[operation]; !ok {
		fmt.Println("Enter correct operation: add, subtract, multiply, divide")
		return
	}

	fmPattern := "%s%d%s"
	space := 5
	if operation == "divide" {
		space = 10
	}
	fmtStr := fmt.Sprintf(fmPattern, "%", space, "s")
	fmtInt := fmt.Sprintf(fmPattern, "%", space, "d")
	fmtFloat := fmt.Sprintf(fmPattern, "%", space, ".4f")

	fmt.Printf(fmtStr, "X")

	for i := 0; i <= max; i++ {
		fmt.Printf(fmtInt, i)
	}
	fmt.Println()
	for i := 0; i <= max; i++ {
		fmt.Printf(fmtInt, i)

		for j := 0; j <= max; j++ {
			switch operation {
			case "add":
				fmt.Printf(fmtInt, i+j)
			case "subtract":
				fmt.Printf(fmtInt, i-j)
			case "divide":
				if j != 0 {
					fmt.Printf(fmtFloat, float64(i)/float64(int64(j)))
				} else {
					fmt.Printf(fmtStr, "-")
				}
			case "multiply":
				fmt.Printf(fmtInt, i*j)
			}
		}

		fmt.Println()
	}
}

func parseFields() {
	words := strings.Fields("lazy cat jumps again and again and again")

	// for j := 0; j < len(words); j++ {
	// 	fmt.Printf("#%-2d: %q\n", j+1, words[j])
	// }

	var (
		i int
		v string
	)
	for i, v = range words {
		fmt.Printf("#%-2d: %q\n", i+1, v)
	}

	fmt.Printf("Last value of the i = %d v = %q\n", i, v)

	args := os.Args

	for i, v := range args {
		if i == 0 {
			continue
		}

		fmt.Printf("%q\n", v)
	}

	for _, v := range args[1:] {
		fmt.Printf("%q\n", v)
	}
}
