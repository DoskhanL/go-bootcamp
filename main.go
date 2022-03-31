package main

import (
	"fmt"
	"os"

	"github.com/doskhanl/go-bootcamp/exercises"
)

func main() {
	printArgs()
	exercises.MakeItBlue()
	exercises.VarsToVars()
	exercises.AssignWithExp()
	exercises.RectanglePerimeter()
	exercises.MultiAssign()
	exercises.MultiAssignSecond()
	exercises.MultiShortFunc()
	exercises.Swapper()
	exercises.SwapperSecond()
	exercises.DiscardTheFile()
	exercises.FixIt()

}

func printArgs() {
	fmt.Printf("%#v\n", os.Args)

	for i, val := range os.Args {
		fmt.Printf("Args[%d] argument: %v \n", i, val)
	}

	fmt.Println("Number of arguments in os.Args:", len(os.Args))
}
