package exercises

import (
	"fmt"
	"os"
	"path"
)

// ---------------------------------------------------------
// EXERCISE: Print the Path
//
//  Print the path of the running program by getting it
//  from `os.Args` variable.
//
// HINT
//  Use `go build` to build your program.
//  Then run it using the compiled executable program file.
//
// EXPECTED OUTPUT SHOULD INCLUDE THIS
//  myprogram
// ---------------------------------------------------------

func PrintThePath() {
	_, file := path.Split(os.Args[0])
	fmt.Println(file)
}
