package Furniture

type Furniture struct {
	Name     string
	Material string
	Color    string
	Height   float64
	Width    float64
	Lang     float64
}

func InitFurniture() []Furniture {
	myBed := Furniture{Name: "Bed", Material: "wood", Color: "white", Height: 60.3, Width: 300.0, Lang: 250.4}
	myChair := Furniture{Name: "Chair", Material: "wood", Color: "white", Height: 50.0, Width: 50.1, Lang: 50.1}
	myTable := Furniture{Name: "Table", Material: "wood", Color: "white", Height: 94.1, Width: 300.0, Lang: 250.0}
	myWardrobe := Furniture{Name: "Wardrobe", Material: "wood", Color: "white", Height: 220.0, Width: 300.0, Lang: 100.0}
	mySofa := Furniture{Name: "Sofa", Material: "Polyurethane foam", Color: "white", Height: 130.0, Width: 400.0, Lang: 120.4}
	return []Furniture{myBed, myChair, myTable, myWardrobe, mySofa}
}
