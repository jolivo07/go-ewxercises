package company_data


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

func (e Employes) NumberEmployees()int {

	return len(e)
}

func (e Employes) NumberMen() int{
	cont := 0
	
	for _, v := range e {
		if v.Grender == "M" {
			cont++
		}
	}
	return cont
}

func (e Employes) NumberWomen() int {
	cont := 0
	for _, v := range e {
		if v.Grender == "F" {
			cont++
		}
	}
	return cont
}

func (e Employes) Avg() int {
	sum := 0

	for _, v := range e {
		sum += v.Age
	}
	return sum / len(e)
}

func (e Employes) AboveAvg() int{
	cont := 0

	for _, v := range e {
		if v.Age >= e.Avg() {
			cont ++
		}
	}
	return cont
}


func (e Employes) BelowAvg() int{
	cont := 0

	for _, v := range e {
		if v.Age <= e.Avg() {
			cont ++
		}
	}
	return cont
}
