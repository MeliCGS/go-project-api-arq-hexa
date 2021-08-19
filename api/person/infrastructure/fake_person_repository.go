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

func (f *FakePersonRepository) DeletePeople(id string) {
	foundIndex := -1
	lastIndex := len(fakePeople) - 1

	for i, person := range fakePeople {
		if person.ID == id {
			foundIndex = i
			break
		}
	}

	fakePeople[foundIndex] = fakePeople[lastIndex]
	fakePeople = fakePeople[:lastIndex]
	//esto es como un slice:
	//fakePeople = append(fakePeople[:foundIndex], fakePeople[foundIndex+1:]...)

	// esto seria en javascript
	//const arr = [1, 2, 3]
	//arr[0] = arr.pop()

}

func (f *FakePersonRepository) UpdatePeople(Person *person_domain.Person) {
	//fakePeople = append(fakePeople, Person)
	var foundTodo *person_domain.Person

	for _, t := range fakePeople {
		if t.ID == Person.ID {
			foundTodo = t
			break
		}
	}

	if foundTodo != nil {
		foundTodo.Nombre = Person.Nombre
		foundTodo.Apellido = Person.Apellido
	}

}
