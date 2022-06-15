package letters

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ScoreFor() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	letterScanner := scanner.Text()
	arrayLetters := strings.Split(strings.ToLower(letterScanner), "")

	sum := 0
	for _, valueArrayLetters := range arrayLetters {
		for key, v := range LettersMap() {
			if valueArrayLetters == key {
				sum = sum + v
			}
		}
	}
	fmt.Println(sum)

}
