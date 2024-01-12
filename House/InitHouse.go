package House

import (
	"fmt"
	"myHouse/House/Appliances"
	"myHouse/House/Family"
	"myHouse/House/Furniture"
	"myHouse/House/Rooms"
	//"myHouse/House/Rooms"
)

func InitHouse() {
	fmt.Println("\n Furniture:", Furniture.InitFurniture(),
		"\n Family:", Family.InitFamily(),
		"\n Appliances:", Appliances.InitAppliances(),
		"\n Rooms", Rooms.InitRooms())
}
