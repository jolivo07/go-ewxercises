package letters

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ScoreFor() {
	var letters = make(map[int][]string)
	letters[1] = []string{"a", "e", "i", "o", "u", "l", "n", "r", "s", "t"}
	letters[2] = []string{"d", "g"}
	letters[3] = []string{"b", "c", "m", "p"}
	letters[4] = []string{"f", "h", "w", "v", "y"}
	letters[5] = []string{"k"}
	letters[8] = []string{"j", "x"}
	letters[10] = []string{"q", "z"}
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	letterScanner := scanner.Text()
	arrayLetters := strings.Split(letterScanner, "")

	sum := 0
	for _, valueArrayLetters := range arrayLetters {

		for key, v := range letters {

			for _, value := range v {
				if value == valueArrayLetters {
					sum = sum + key
				}
			}

		}
	}
	fmt.Println(sum)

}
