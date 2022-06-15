package models

func RandomColor() string {
	colors := map[int]string{1: "black", 2: "blue", 3: "brown", 4: "gray", 5: "green", 6: "orange", 7: "yellow", 8: "purple", 9: "red", 10: "white"}
	return colors[RandomNumber()]
}
