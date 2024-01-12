package Appliances

type Appliances struct {
	Name         string
	Color        string
	Manufacturer string
}

func InitAppliances() []Appliances {
	myFridge := Appliances{Name: "Fridge", Color: "Green", Manufacturer: "LG"}
	myMicrowave := Appliances{Name: "Microwave", Color: "White", Manufacturer: "LG"}
	myStove := Appliances{Name: "Stove", Color: "White", Manufacturer: "Artel"}
	myTelevisor := Appliances{Name: "Televisor", Color: "Black", Manufacturer: "Samsung"}
	myVacuumCleaner := Appliances{Name: "VacuumCleaner", Color: "Pink", Manufacturer: "LG"}
	return []Appliances{myFridge, myMicrowave, myStove, myTelevisor, myVacuumCleaner}
}
