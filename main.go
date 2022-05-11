package main

import (
	"fmt"
	"os"

	"github.com/doskhanl/go-bootcamp/exercises"
)

func main() {
	// printArgs()
	// exercises.MakeItBlue()
	// exercises.VarsToVars()
	// exercises.AssignWithExp()
	// exercises.RectanglePerimeter()
	// exercises.MultiAssign()
	// exercises.MultiAssignSecond()
	// exercises.MultiShortFunc()
	// exercises.Swapper()
	// exercises.SwapperSecond()
	// exercises.DiscardTheFile()
	// exercises.FixIt()
	// exercises.CountArgs()
	// exercises.PrintThePath()
	// exercises.PrintMe()
	// exercises.GreetMore()

	// tutorials.ShowString()
	// tutorials.Banger()

	//exercises.StringsTask()
	//tutorials.ConstantsTask()
	//exercises.IotaMain()

	// tutorials.MainPrint()
	// exercises.Formattings()

	// exercises.MainIf()

	// tutorials.ErrorHandlingMain()
	// tutorials.ShortIf()

	// exercises.ErrorHandlingMain()

	// tutorials.SwitchMain()

	// exercises.SwitchMain()
	// helper.OperatePoints()

	// exercises.StringManipulator()

	//exercises.DaysInMonthSwitch()

	//tutorials.MainLoop()
	exercises.LoopMain()

}

func printArgs() {
	fmt.Printf("%#v\n", os.Args)

	for i, val := range os.Args {
		fmt.Printf("Args[%d] argument: %v \n", i, val)
	}

	fmt.Println("Number of arguments in os.Args:", len(os.Args))
}
