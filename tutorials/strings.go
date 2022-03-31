package tutorials

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func ShowString() {
	var s string

	s = "Hello there"
	fmt.Println(s)

	s = `Hello there`
	fmt.Println(s)

	s = "<html>\n\t<body>\"Hello\"</body>\n</html>"
	fmt.Println(s)
	s = `
<html>
	<body>"Hello"</body>
</html>`
	fmt.Println(s)

	fmt.Println("c:\\my\\dir\\file")
	fmt.Println(`c:\my\dir\file`)

	var txt = "Әлһәмдуллиләһ"
	fmt.Println("Lenth of", txt, "is with len", len(txt))
	fmt.Println("Lenth of", txt, "is with RuneCountInString from encoding/utf8",
		utf8.RuneCountInString(txt))
}

func Banger() {
	args := os.Args
	var word string
	if len(args) >= 2 {
		word = args[1]
	}

	l := utf8.RuneCountInString(word)
	exMark := strings.Repeat("!", l)
	fmt.Println(exMark + strings.ToUpper(word) + exMark)
}
