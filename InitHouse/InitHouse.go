package InitHouse

import (
	"fmt"
	"myHouse/Family"
	"myHouse/Furniture"
)

func InitHouse() {
	fmt.Println("Мебель:", Furniture.InitFurniture(), "Семья:", Family.InitFamily())
}
