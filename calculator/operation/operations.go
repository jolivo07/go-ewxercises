package operation

import "fmt"

type Operation struct {
	Num1 int
	Num2 int
}

func (c Operation) Addition() int {
	return c.Num1 + c.Num2
}
func (c Operation) Subtraction() int {
	return c.Num1 - c.Num2
}
func (c Operation) Multiplication() int {
	return c.Num1 * c.Num2
}
func (c Operation) Division() int {
	return c.Num1 / c.Num2
}

func (c Operation) Result() {
	fmt.Printf("Addition: %d \nSubtraction: %d \nMultiplication: %d \nDivision: %d \n", c.Addition(), c.Subtraction(), c.Multiplication(), c.Division())

}
