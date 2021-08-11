package person_infrastructure

import person_domain "github.com/MeliCGS/go-project-api-arq-hexa/api/person/domain"

var fakePeople = []*person_domain.Person{
	{
		ID:       "1",
		Nombre:   "camilo",
		Apellido: "garces",
		Edad:     30,
	},
	{
		ID:       "2",
		Nombre:   "Mateo",
		Apellido: "garces",
		Edad:     18,
	},
}

type FakePersonRepository struct {
}

func (f *FakePersonRepository) GetPeople() []*person_domain.Person {
	return fakePeople
}

func (f *FakePersonRepository) AddPeople(Person *person_domain.Person) {
	fakePeople = append(fakePeople, Person)
}
