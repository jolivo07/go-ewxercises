package models

import (
	"math/rand"
	"time"
)

func RandomNumber() int {
	rand.Seed(time.Now().UnixNano())
	return 1 + rand.Intn(11-1)
}
