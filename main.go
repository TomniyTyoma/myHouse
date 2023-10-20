package main

import (
	"fmt"
	"myHouse/Floor/Furniture"
)

func main() {
	bed1 := Furniture.Bed{
		Size:     "1920X1080",
		Material: "wood",
		Color:    "blue",
	}
	fmt.Println(bed1, Furniture.Bed2())
}
