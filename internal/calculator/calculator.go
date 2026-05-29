package calculator

import (
	"regexp"
	"strconv"
)

type Calculator struct {
	input string
}

const (
	division       = "/"
	multiplication = "*"
	addition       = "+"
	subtraction    = "-"
)

func NewCalculator(input string) *Calculator {
	return &Calculator{input: input}
}

func (c *Calculator) Calculate() (Number, error) {
	// use regex to extract numbers with division and multiplication first, reading from left to right
	// meaning that "40+20÷5×3–5×2" will be evaluated as "40+(20÷5)×3–5×2"
	// (first pair of division/multiplication operation detected will be processed first)
	// which then translates to "40+4×3–5×2" and then "40+12–5×2" and then "40+12–10" and finally "42"

	// find first occurrence of division or multiplication with two numbers around it
	// then replace that occurrence with the result of the operation,
	// and repeat until no more division or multiplication is found
	re := regexp.MustCompile(`(\d+(?:\.\d+)?)([/\*])(\d+(?:\.\d+)?)`)
	for {
		matchIdx := re.FindStringSubmatchIndex(c.input)
		if len(matchIdx) == 0 {
			break
		}

		leftStr := c.input[matchIdx[2]:matchIdx[3]]
		operation := c.input[matchIdx[4]:matchIdx[5]]
		rightStr := c.input[matchIdx[6]:matchIdx[7]]

		left, err := strconv.ParseFloat(leftStr, 64)
		if err != nil {
			return Number(0), err
		}

		right, err := strconv.ParseFloat(rightStr, 64)
		if err != nil {
			return Number(0), err
		}

		var result Number
		switch operation {
		case division:
			result = Number(left).Divide(Number(right))
		case multiplication:
			result = Number(left).Multiply(Number(right))
		}

		resultStr := strconv.FormatFloat(float64(result), 'f', -1, 64)
		c.input = c.input[:matchIdx[0]] + resultStr + c.input[matchIdx[1]:]
	}

	// then find first occurrence of addition or subtraction with two numbers around it
	// then replace that occurrence with the result of the operation,
	// and repeat until no more addition or subtraction is found
	re = regexp.MustCompile(`(\d+(?:\.\d+)?)([+\-])(\d+(?:\.\d+)?)`)
	for {
		matchIdx := re.FindStringSubmatchIndex(c.input)
		if len(matchIdx) == 0 {
			break
		}

		leftStr := c.input[matchIdx[2]:matchIdx[3]]
		operation := c.input[matchIdx[4]:matchIdx[5]]
		rightStr := c.input[matchIdx[6]:matchIdx[7]]

		left, err := strconv.ParseFloat(leftStr, 64)
		if err != nil {
			return Number(0), err
		}

		right, err := strconv.ParseFloat(rightStr, 64)
		if err != nil {
			return Number(0), err
		}

		var result Number
		switch operation {
		case addition:
			result = Number(left).Add(Number(right))
		case subtraction:
			result = Number(left).Subtract(Number(right))
		}

		resultStr := strconv.FormatFloat(float64(result), 'f', -1, 64)
		c.input = c.input[:matchIdx[0]] + resultStr + c.input[matchIdx[1]:]
	}

	if result, err := strconv.ParseFloat(c.input, 64); err == nil {
		return Number(result), nil
	}

	return Number(0), nil
}
