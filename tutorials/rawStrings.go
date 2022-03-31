package tutorials

import "fmt"

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
}
