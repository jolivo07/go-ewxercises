package models

type Earth struct {
	ID             string   `json:"id"`
	Date           string   `json:"date"`
	ServiceVersion string   `json:"service_version"`
	Url            string   `json:"url" `
	Resource       Resource `json:"resource" `
}

type Resource struct {
	Planet  string `json:"planet"`
	Dataset string `json:"dataset"`
}
