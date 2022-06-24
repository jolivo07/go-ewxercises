package actions

import (
	"conections/models"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"path"
	"text/template"
	_ "github.com/lib/pq"

)

func ShowEarthInformation(w http.ResponseWriter, r *http.Request) {
	response, err := http.Get("https://api.nasa.gov/planetary/earth/assets?lon=100.75&lat=1.5&date=2014-02-01&dim=0.15&api_key=DEMO_KEY")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var earth models.Earth
	if err := json.Unmarshal(responseData, &earth); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fp := path.Join("html/nasa_api.html")
	tmpl, err := template.ParseFiles(fp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := tmpl.Execute(w, earth); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
