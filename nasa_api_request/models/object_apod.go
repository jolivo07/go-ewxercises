package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func GetApiApod() Apod {
	response, err := http.Get("https://api.nasa.gov/planetary/apod?api_key=DEMO_KEY")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	var apod Apod
	json.Unmarshal(responseData, &apod)

	return apod

}
