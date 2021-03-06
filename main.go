package main

import (
	person_application "github.com/MeliCGS/go-project-api-arq-hexa/api/person/application"
	person_infrastructure "github.com/MeliCGS/go-project-api-arq-hexa/api/person/infrastructure"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	personController := &person_infrastructure.PersonController{
		PeopleSearcher: &person_application.PeopleSearcher{
			Repo: &person_infrastructure.FakePersonRepository{},
		},
	}

	personAddController := &person_infrastructure.PersonAddController{
		PeopleAdd: &person_application.PeopleAdd{
			Repo: &person_infrastructure.FakePersonRepository{},
		},
	}
	personDeleteController := &person_infrastructure.PersonDeleteController{
		PeopleDelete: &person_application.PeopleDelete{
			Repo: &person_infrastructure.FakePersonRepository{},
		},
	}

	personUpdateController := &person_infrastructure.PersonUpdateController{
		PeopleUpdate: &person_application.PeopleUpdate{
			Repo: &person_infrastructure.FakePersonRepository{},
		},
	}

	r.GET("/people", personController.GetAllHandler)
	r.POST("/people", personAddController.GetAllHandler)
	r.DELETE("/people/:id", personDeleteController.GetAllHandler)
	r.PATCH("/people/:id", personUpdateController.GetAllHandler)
	r.Run(":8000")
}
