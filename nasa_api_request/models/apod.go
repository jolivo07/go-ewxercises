package models

import "fmt"

type Apod struct {
	Copyright       string `json:"copyright"`
	Date            string `json:"date"`
	Explanation     string `json:"explanation"`
	Hdurl           string `json:"hdurl"`
	Media_type      string `json:"media_type"`
	Service_version string `json:"service_version"`
	Title           string `json:"title"`
	Url             string `json:"url"`
}

func (c Apod) Message() {
	fmt.Printf("Title: %s \n -> Date: %s \nservice Vercion: %s \nExplanation: \n%s \nMedia_type: %s \nHdurl: %s \nUrl: %s \nCopyright : %s \n", c.Title, c.Date, c.Service_version, c.Explanation, c.Media_type, c.Hdurl, c.Url, c.Copyright)
}
