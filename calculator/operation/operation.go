package operation

import "fmt"

type Numbers struct {
	Num1 int
	Num2 int
}

func (c Numbers) Addition() int {
	return c.Num1 + c.Num2
}
func (c Numbers) Subtraction() int {
	return c.Num1 - c.Num2
}
func (c Numbers) Multiplication() int {
	return c.Num1 * c.Num2
}
func (c Numbers) Division() int {
	return c.Num1 / c.Num2
}

func (c Numbers) ResultMessage() {
	fmt.Printf("Addition: %d \nSubtraction: %d \nMultiplication: %d \nDivision: %d \n", c.Addition(), c.Subtraction(), c.Multiplication(), c.Division())

}
