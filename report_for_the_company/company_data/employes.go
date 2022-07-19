package company_data

import "fmt"

type Employee struct {
	Name    string
	Grender string
	Age     int
}

type Employes []Employee

func Init() Employes {
	E1 := Employee{Name: "joaquin", Grender: "M", Age: 19}
	E2 := Employee{Name: "Carlos", Grender: "M", Age: 20}
	E3 := Employee{Name: "Sofia", Grender: "F", Age: 21}
	E4 := Employee{Name: "Ana", Grender: "F", Age: 22}
	E5 := Employee{Name: "Ángela", Grender: "F", Age: 40}

	ES := []Employee{E1, E2, E3, E4, E5}

	return ES

}

func (e Employes) NumberEmployees() {

	fmt.Printf("Number of employees: %v \n", len(e))

}

func (e Employes) NumberMen() {
	count := 0
	for _, v := range e {
		if v.Grender == "M" {
			count++
		}
	}
	fmt.Printf("Number Men: %v \n", count)
}

func (e Employes) NumberWomen() {
	count := 0
	for _, v := range e {
		if v.Grender == "F" {
			count++
		}
	}
	fmt.Printf("Number Women: %v \n", count)
}

func (e Employes) Avg() int {
	sum := 0

	for _, v := range e {
		sum += v.Age
	}

	return sum / len(e)
}

func (e Employes) AboveAvg(){

	cont := 0
	for _, v := range e {
		if v.Age >= e.Avg() {
			cont ++
		}
	}
	fmt.Println("employees above the average age:",cont)
}


func (e Employes) BelowAvg(){

	cont := 0
	for _, v := range e {
		if v.Age <= e.Avg() {
			cont ++
		}
	}
	fmt.Println("employees below the average age:",cont)
}
