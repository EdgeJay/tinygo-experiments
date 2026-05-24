package calculator

type Number float64

func (n Number) Add(other Number) Number {
	return n + other
}

func (n Number) Subtract(other Number) Number {
	return n - other
}

func (n Number) Multiply(other Number) Number {
	return n * other
}

func (n Number) Divide(other Number) Number {
	if other == 0 {
		panic("division by zero")
	}
	return n / other
}
