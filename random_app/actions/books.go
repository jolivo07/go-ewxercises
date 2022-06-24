package actions

import (
	"conections/models"
	"net/http"
	"path"
	"text/template"
	"time"
	_ "github.com/lib/pq"

)


func ShowBook(w http.ResponseWriter, r *http.Request) {
	book := models.Book{Title: "harry Potter", Creation: time.Now(), Autor: "Joaquin Olivo"}

	fp := path.Join("html/book.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, book); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}