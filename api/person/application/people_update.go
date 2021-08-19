package person_application

import (
	person_domain "github.com/MeliCGS/go-project-api-arq-hexa/api/person/domain"
)

type PeopleUpdate struct {
	Repo person_domain.PersonRepository
}

func (p *PeopleUpdate) UpdatePerson(Person *person_domain.Person) {
	p.Repo.UpdatePeople(Person)
}
