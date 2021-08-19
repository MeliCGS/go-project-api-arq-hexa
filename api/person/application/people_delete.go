package person_application

import (
	person_domain "github.com/MeliCGS/go-project-api-arq-hexa/api/person/domain"
)

type PeopleDelete struct {
	Repo person_domain.PersonRepository
}

func (p *PeopleDelete) DeletePerson(id string) {
	p.Repo.DeletePeople(id)
}
