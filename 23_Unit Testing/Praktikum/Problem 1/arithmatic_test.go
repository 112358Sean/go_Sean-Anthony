package arithmatic

import (
	"testing"
)

func TestAddition(t *testing.T) {
	result := Addition(10, 5)
	
	if result == 15 {
		t.Log("Test passed")
	} else {
		t.Log("Test fail")
		t.Fail()
	}
}

func TestSubtraction(t *testing.T) {
	result := Subtraction(10, 5)

	if result == 5 {
		t.Log("Test passed")
	} else {
		t.Log("Test fail")
		t.Fail()
	}
}

func TestMultiplication(t *testing.T) {
	result := Multiplication(10, 5)

	if result == 50 {
		t.Log("Test passed")
	} else {
		t.Log("Test fail")
		t.Fail()
	}
}

func TestDivision(t *testing.T) {
	result := Division(10, 5)

	if result == 2 {
		t.Log("Test passed")
	} else {
		t.Log("Test fail")
		t.Fail()
	}
}
