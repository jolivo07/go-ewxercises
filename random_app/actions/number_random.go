package actions

import (
	"math/rand"
	"net/http"
	"path"
	"text/template"
	"time"
	_ "github.com/lib/pq"

)

func randomNumber() int {
	rand.Seed(time.Now().UnixNano())
	return 1 + rand.Intn(199-1)
}

func RandomNumberList(w http.ResponseWriter, r *http.Request) {

	type Data struct {
		Numbers []int
	}

	data := Data{Numbers: []int{
		randomNumber(),
		randomNumber(),
		randomNumber(),
		randomNumber(),
		randomNumber(),
		randomNumber(),
		randomNumber(),
		randomNumber(),
		randomNumber(),
		randomNumber(),
	}}

	fp := path.Join("html/number_random.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
