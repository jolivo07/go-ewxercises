// package main

// import (
// 	"fmt"
// 	"scrabble/models"
// 	"strings"
// )

// func main() {
// 	fmt.Println(strings.ToUpper(models.ConvertStrColorHex()))
// }

package main

import (
	"fmt"
	"scrabble/models"
)

func main() {


	rgb := models.ConvertRgb()

	fmt.Printf("#%02x%02x%02x \n", rgb.Red, rgb.Green, rgb.Blue)

}
