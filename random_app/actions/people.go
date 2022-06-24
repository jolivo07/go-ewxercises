package actions

import (
	"conections/db_conn"
	"conections/models"
	"fmt"
	"net/http"
	"path"
	"text/template"
	_ "github.com/lib/pq"

)

func addPerson(name, description string) error {
	db := db_conn.Setup()
	defer db.Close()
	insertStm := fmt.Sprintf(`insert into "Person" ("name", "description") values ('%s' , '%s')`, name, description)
	_, e := db.Exec(insertStm)
	return e
}

func getPeople() ([]models.Person, error) {
	db := db_conn.Setup()
	defer db.Close()

	rows, err := db.Query(`SELECT *  FROM public."Person"`)
	if err != nil {
		return []models.Person{}, err
	}
	defer rows.Close()

	var people []models.Person
	for rows.Next() {
		var person models.Person
		err = rows.Scan(&person.Name, &person.Description)
		if err != nil {
			panic(err)
		}
		people = append(people, person)
	}
	return people, err
}

func ShowPersonForm(w http.ResponseWriter, r *http.Request) {

	people, err := getPeople()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	pp := models.PeopleData{People: people}

	fp := path.Join("html/person.html")
	tmpl, err := template.ParseFiles(fp)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, pp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	
	r.ParseForm()
	name := r.Form["name"][0]
	description := r.Form["description"][0]

	if err := addPerson(name, description); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	http.Redirect(w, r, "http://localhost:8080/person-form/show", http.StatusMovedPermanently)
}


