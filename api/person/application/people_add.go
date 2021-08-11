package person_application

import (
	person_domain "github.com/MeliCGS/go-project-api-arq-hexa/api/person/domain"
)

type PeopleAdd struct {
	Repo person_domain.PersonRepository
}

func (p *PeopleAdd) AddPerson(Person *person_domain.Person) {
	p.Repo.AddPeople(Person)
}
