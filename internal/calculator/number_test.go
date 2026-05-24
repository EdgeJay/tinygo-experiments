package calculator

import "testing"

func TestNumberAdd(t *testing.T) {
	got := Number(10).Add(Number(5))
	want := Number(15)

	if got != want {
		t.Fatalf("Add() = %v, want %v", got, want)
	}
}

func TestNumberSubtract(t *testing.T) {
	got := Number(10).Subtract(Number(5))
	want := Number(5)

	if got != want {
		t.Fatalf("Subtract() = %v, want %v", got, want)
	}
}

func TestNumberMultiply(t *testing.T) {
	got := Number(10).Multiply(Number(5))
	want := Number(50)

	if got != want {
		t.Fatalf("Multiply() = %v, want %v", got, want)
	}
}

func TestNumberDivide(t *testing.T) {
	got := Number(10).Divide(Number(5))
	want := Number(2)

	if got != want {
		t.Fatalf("Divide() = %v, want %v", got, want)
	}
}

func TestNumberDivideByZeroPanics(t *testing.T) {
	defer func() {
		r := recover()
		if r == nil {
			t.Fatal("expected panic on division by zero, got nil")
		}
		if r != "division by zero" {
			t.Fatalf("panic = %v, want %v", r, "division by zero")
		}
	}()

	_ = Number(10).Divide(Number(0))
}
