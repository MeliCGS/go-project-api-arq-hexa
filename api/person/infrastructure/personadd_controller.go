package person_infrastructure

import (
	"net/http"

	person_application "github.com/MeliCGS/go-project-api-arq-hexa/api/person/application"
	person_domain "github.com/MeliCGS/go-project-api-arq-hexa/api/person/domain"
	"github.com/gin-gonic/gin"
)

type PersonAddController struct {
	PeopleAdd *person_application.PeopleAdd
}

func (p *PersonAddController) GetAllHandler(ctx *gin.Context) {
	person := &person_domain.Person{}
	ctx.ShouldBindJSON(person)
	p.PeopleAdd.AddPerson(person)
	ctx.JSON(http.StatusOK, person)
}
