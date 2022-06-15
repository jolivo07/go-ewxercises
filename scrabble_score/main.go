package main

import (
	"fmt"
	"scrabble/letters"
)

func main() {
	sum := letters.ScoreFor()
	fmt.Println(sum)
}
