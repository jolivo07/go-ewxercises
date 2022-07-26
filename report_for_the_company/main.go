package main

import (
	"fmt"
	"scrabble/company_data"
)

func main() {
	employes := company_data.Init()

	fmt.Println("Number Employes:", employes.NumberEmployees())
	fmt.Println("Number Men:", employes.NumberMen())
	fmt.Println("Number Women:", employes.NumberWomen())
	fmt.Println("Age Average:", employes.Avg())
	fmt.Println("employees above the average age:",employes.AboveAvg())
	fmt.Println("employees below the average age:",employes.BelowAvg())

}
