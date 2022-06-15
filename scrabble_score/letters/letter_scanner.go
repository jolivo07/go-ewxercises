package letters

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func LettersArray() []string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter word")
	scanner.Scan()
	letterScanner := scanner.Text()
	arrayLetters := strings.Split(strings.ToLower(letterScanner), "")
	
	return arrayLetters
}
