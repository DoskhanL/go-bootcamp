package maths

import "testing"

type payload struct {
	x, y, expected int
}

type payloadFloat struct {
	x, y     int
	expected float64
}

func TestAdd(t *testing.T) {

	payloads := []payload{
		{2, 3, 5},
		{3, 7, 10},
		{100, 25, 125},
	}

	for _, data := range payloads {
		if result := Add(data.x, data.y); result != data.expected {
			t.Errorf("FAILED: Add(%d, %d) expected %d, got %d\n",
				data.x,
				data.y,
				data.expected,
				result,
			)
		} else {
			t.Logf("PASSED: Add(%d, %d) expected %d, got %d\n",
				data.x,
				data.y,
				data.expected,
				result,
			)
		}
	}

}

func TestDivide(t *testing.T) {
	payloads := []payloadFloat{
		{1, 0, 0},
		{21, 3, 7},
		{3, 4, 0.75},
		{100, 25, 4},
	}

	for _, data := range payloads {
		if result := Divide(data.x, data.y); result != data.expected {
			t.Errorf("FAILED: Divide(%d, %d) expected %f, got %f",
				data.x,
				data.y,
				data.expected,
				result,
			)
		} else {
			t.Logf("PASSED: : Divide(%d, %d) expected %f, got %f",
				data.x,
				data.y,
				data.expected,
				result,
			)
		}

	}
}

func TestSubtract(t *testing.T) {
	payloads := []payload{
		{2, 3, -1},
		{30, 7, 23},
		{100, 25, 75},
	}

	for _, data := range payloads {
		if result := Subtract(data.x, data.y); result != data.expected {
			t.Errorf("FAILED: Subtract(%d, %d) expected %d, got %d\n",
				data.x,
				data.y,
				data.expected,
				result,
			)
		} else {
			t.Logf("PASSED: Subtract(%d, %d) expected %d, got %d\n",
				data.x,
				data.y,
				data.expected,
				result,
			)
		}
	}
}

func TestMultiple(t *testing.T) {

	payloads := []payload{
		{2, 3, 6},
		{3, 7, 21},
		{100, 25, 2500},
	}

	for _, data := range payloads {
		if result := Multiple(data.x, data.y); result != data.expected {
			t.Errorf("FAILED: Multiple(%d, %d) expected %d, got %d\n",
				data.x,
				data.y,
				data.expected,
				result,
			)
		} else {
			t.Logf("PASSED: Multiple(%d, %d) expected %d, got %d\n",
				data.x,
				data.y,
				data.expected,
				result,
			)
		}
	}
}
