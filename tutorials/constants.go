package tutorials

import "fmt"

func ConstantsTask() {
	const (
		monday = iota + 1
		tuesday
		wednesday
		thursday
		friday
		saturday
		sunday
	)

	fmt.Println(monday, tuesday, wednesday, thursday, friday, saturday, sunday)
	showTimeZone()
}

func showTimeZone() {
	const (
		EST = -(5 + iota)
		_
		MST
		PST
	)

	fmt.Println(EST, MST, PST)
}
