package actions

import (
	"fmt"
	"html/template"
	"net/http"
	_ "github.com/lib/pq"

	"path"
)

func ShowForm(w http.ResponseWriter, r *http.Request) {
	fp := path.Join("html/form.html")

	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, fp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

func SendInformation(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	fmt.Println("username:", r.Form["name"])
	fmt.Println("password:", r.Form["last-name"])
	fmt.Println("age:", r.Form["age"])

	http.Redirect(w, r, "http://localhost:8080//form/show", http.StatusMovedPermanently)
}
