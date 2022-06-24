package main

import (
	"conections/actions"
	

	"net/http"
)

func main() {
	// Show Earth Information
	http.HandleFunc("/earth-information/show", actions.ShowEarthInformation)

	// Show Book Information
	http.HandleFunc("/book/show", actions.ShowBook)

	// Generic Information
	http.HandleFunc("/form/show", actions.ShowForm)
	http.HandleFunc("/form/send", actions.SendInformation)

	// Random Numbers
	http.HandleFunc("/random-numbers/show", actions.RandomNumberList)

	// 
	http.HandleFunc("/person-form/show", actions.ShowPersonForm)
	http.HandleFunc("/create-person", actions.CreatePerson)


	//
	
	http.ListenAndServe(":8080", nil)
}
