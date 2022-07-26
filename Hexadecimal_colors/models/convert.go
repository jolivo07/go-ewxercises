package models

import (
	"bufio"
	"fmt"
	"strconv"

	"os"
)

type Hex string

type RGB struct {
	Red   uint8
	Green uint8
	Blue  uint8
}

func ConvertRgb() RGB {
	var rgb RGB
	hx := fmt.Sprintf("%x", letters())
	values, err := strconv.ParseUint(hx, 16, 64)
	if err != nil {
		return RGB{}
	}
	rgb = RGB{
		Red:   uint8(values >> 8),
		Green: uint8((values >> 16) & 0xFF),
		Blue:  uint8(values & 0xFF),
	}

	return rgb

}

func letters() string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter word")
	scanner.Scan()
	str := scanner.Text()

	return str
}
