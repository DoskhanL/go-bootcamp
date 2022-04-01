package tutorials

import "fmt"

func MainPrint() {
	var brand = "Google"
	fmt.Printf("%q\n", brand)

	printVarTypes()
}

func printVarTypes() {
	var speed int
	var heat float64
	var off bool
	var brand string

	fmt.Printf("%d %f %v %s\n", speed, heat, off, brand)

	var (
		planet   = "venus"
		distance = 261
		orbital  = 224.701
		hasLife  = false
	)

	fmt.Printf("Planet: %s\n", planet)
	fmt.Printf("Distance: %d millions kms\n", distance)
	fmt.Printf("Orbital period: %f days\n", orbital)
	fmt.Printf("Does %s have life? %t\n", planet, hasLife)

	fmt.Printf("%s is %d away. Think! %[2]d kms! %[1]s OMG!\n",
		planet, distance,
	)

	fmt.Printf("Orbital Period: %f days\n", orbital)
	fmt.Printf("Orbital Period: %.0f days\n", orbital)
	fmt.Printf("Orbital Period: %0.1f days\n", orbital)
	fmt.Printf("Orbital Period: %0.2f days\n", orbital)
	fmt.Printf("Orbital Period: %0.3f days\n", orbital)
	fmt.Printf("Orbital Period: %0.4f days\n", orbital)
	fmt.Printf("Orbital Period: %0.5f days\n", orbital)
}
