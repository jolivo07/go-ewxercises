package letters

func ScoreFor() int {
	var sum int
	for _, valueArrayLetters := range LettersArray() {
		sum = sum + LettersMap[valueArrayLetters]
	}
	return sum
}
