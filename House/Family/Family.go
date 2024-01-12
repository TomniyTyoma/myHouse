package Family

type Family struct {
	Name    string
	Surname string
	Age     int
	Member  string
}

func InitFamily() []Family {
	myBrother := Family{Name: "Aslanbek", Surname: "Ivanovich", Age: 12, Member: "Brother"}
	myMother := Family{Name: "Eleonora", Surname: "Nikitichna", Age: 23, Member: "Mother"}
	myFather := Family{Name: "Ivan", Surname: "Jhamshitovich", Age: 56, Member: "Father"}
	myGrandmother := Family{Name: "Valeriya", Surname: "Stepanovna", Age: 36, Member: "Grandmother"}
	myGrandfather := Family{Name: "Jhamshit", Surname: "Ignatovich", Age: 99, Member: "Grandfather"}
	return []Family{myMother, myFather, myBrother, myGrandmother, myGrandfather}
}
