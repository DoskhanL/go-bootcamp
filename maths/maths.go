package maths

// Add function
func Add(x, y int) int {
	return x + y
}

// Subtract function
func Subtract(x, y int) int {
	return x - y
}

// Divide function
func Divide(x, y int) float64 {
	if y == 0 {
		return 0.0
	}
	return float64(x) / float64(y)
}

// Multiple function
func Multiple(x, y int) int {
	return x * y
}
