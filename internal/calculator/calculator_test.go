package calculator

import "testing"

func TestCalculatorCalculate(t *testing.T) {
	tests := []struct {
		name     string
		equation string
		want     Number
	}{
		{name: "addition", equation: "2+3", want: Number(5)},
		{name: "subtraction", equation: "9-4", want: Number(5)},
		{name: "multiplication", equation: "7*6", want: Number(42)},
		{name: "division", equation: "8/2", want: Number(4)},
		{name: "mixed precedence", equation: "2+3*4", want: Number(14)},
		{name: "left to right multiply divide", equation: "20/5*3", want: Number(12)},
		{name: "chained additions and subtraction", equation: "10+5-3+2", want: Number(14)},
		{name: "combined operations", equation: "40+20/5*3-5*2", want: Number(42)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			calc := NewCalculator(tt.equation)
			got, err := calc.Calculate()
			if err != nil {
				t.Fatalf("Calculate() returned an error: %v", err)
			}

			if got != tt.want {
				t.Fatalf("Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
