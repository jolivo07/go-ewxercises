package main

import (
	"fmt"
	"scrabble/company_data"
)

func main() {
	employes := company_data.Init()

	employes.NumberEmployees()
	employes.NumberMen()
	employes.NumberWomen()
	fmt.Println("Age Average:", employes.Avg())
	employes.AboveAvg()
	employes.BelowAvg()

}
