package models

import "time"

type Book struct {
	Title    string
	Creation time.Time
	Autor    string
}