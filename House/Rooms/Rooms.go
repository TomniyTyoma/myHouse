package Rooms

type Rooms struct {
	Name          string
	Width         float64
	Length        float64
	Height        float64
	WindowsNumber int
}

func InitRooms() []Rooms {
	myKitchen := Rooms{Name: "Kitchen", Width: 340, Length: 500, WindowsNumber: 1, Height: 250.3}
	myLivingRooms := Rooms{Name: "Living Room", Width: 400, Length: 600, WindowsNumber: 2, Height: 250.3}
	return []Rooms{myLivingRooms, myKitchen}
}
