package person_infrastructure

import (
	"net/http"

	person_application "github.com/MeliCGS/go-project-api-arq-hexa/api/person/application"
	person_domain "github.com/MeliCGS/go-project-api-arq-hexa/api/person/domain"
	"github.com/gin-gonic/gin"
)

type PersonUpdateController struct {
	PeopleUpdate *person_application.PeopleUpdate
}

func (p *PersonUpdateController) GetAllHandler(ctx *gin.Context) {
	person := &person_domain.Person{}
	ctx.ShouldBindJSON(person)
	person.ID = ctx.Param("id")
	p.PeopleUpdate.UpdatePerson(person)
	ctx.JSON(http.StatusOK, person)
}
