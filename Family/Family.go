package Family

type Family struct {
	Name    string
	Surname string
	Age     int
	Member  string
}

func InitFamily() []Family {
	myBrother := Family{Name: "Асланбек", Surname: "Иванович", Age: 12, Member: "Brother"}
	myMother := Family{Name: "Элионора", Surname: "Никитична", Age: 23, Member: "Mother"}
	myFather := Family{Name: "Иван", Surname: "Жамшитович", Age: 56, Member: "Father"}
	myGrandmother := Family{Name: "Валерия", Surname: "Степановна", Age: 36, Member: "Grandmother"}
	myGrandfather := Family{Name: "Жамшит", Surname: "Игнатович", Age: 99, Member: "Grandfather"}
	return []Family{myMother, myFather, myBrother, myGrandmother, myGrandfather}
}
