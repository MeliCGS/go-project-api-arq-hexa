package person_domain

type PersonRepository interface {
	GetPeople() []*Person
	AddPeople(Person *Person)
	DeletePeople(id string)
	UpdatePeople(Person *Person)
}
