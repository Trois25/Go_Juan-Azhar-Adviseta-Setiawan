package praktikum

import "testing"

func TestAddition(t *testing.T) {
	if addition(1,2) != 3{
		t.Error("Expected 1 (+) 2 to equal 3")
	}

	if addition(1,-2) != -1{
		t.Error("Expected 1 (+) -2 to equal -1")
	}

	if addition(-1,2) != 1{
		t.Error("Expected -1 (+) 2 to equal 1")
	}

	if addition(-1,-2) != -3{
		t.Error("Expected -1 (+) -2 to equal -3")
	}
}

func TestSubstraction(t *testing.T) {
	if substraction(1,2) != -1{
		t.Error("Expected 1 (-) 2 to equal -1")
	}

	if substraction(1,-2) != 3{
		t.Error("Expected 1 (-) -2 to equal 3")
	}

	if substraction(-1,2) != -3{
		t.Error("Expected -1 (-) 2 to equal -3")
	}

	if substraction(-1,-2) != 1{
		t.Error("Expected -1 (-) -2 to equal 1")
	}
}

func TestMult(t *testing.T) {
	if mult(1,2) != 2{
		t.Error("Expected 1 (*) 2 to equal 2")
	}

	if mult(1,-2) != -2{
		t.Error("Expected 1 (*) -2 to equal -2")
	}

	if mult(-1,2) != -2{
		t.Error("Expected -1 (*) 2 to equal -2")
	}

	if mult(-1,-2) != 2{
		t.Error("Expected -1 (*) -2 to equal 2")
	}
}

func TestDiv(t *testing.T) {
	if div(1,2) != 0.5{
		t.Error("Expected 1 (/) 2 to equal 0.5")
	}

	if div(1,-2) != -0.5{
		t.Error("Expected 1 (/) -2 to equal -0.5")
	}

	if div(-1,2) != -0.5{
		t.Error("Expected -1 (/) 2 to equal -0.5")
	}

	if div(-1,-2) != 0.5{
		t.Error("Expected -1 (/) -2 to equal 0.5")
	}
}