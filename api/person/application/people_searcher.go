package person_application

import (
	person_domain "github.com/MeliCGS/go-project-api-arq-hexa/api/person/domain"
)

type PeopleSearcher struct {
	Repo person_domain.PersonRepository
}

func (p *PeopleSearcher) SearchAll() []*person_domain.Person {
	return p.Repo.GetPeople()
}
