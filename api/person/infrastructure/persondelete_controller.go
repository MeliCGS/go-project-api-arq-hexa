package person_infrastructure

import (
	"net/http"

	person_application "github.com/MeliCGS/go-project-api-arq-hexa/api/person/application"
	person_domain "github.com/MeliCGS/go-project-api-arq-hexa/api/person/domain"
	"github.com/gin-gonic/gin"
)

type PersonDeleteController struct {
	PeopleDelete *person_application.PeopleDelete
}

func (p *PersonDeleteController) GetAllHandler(ctx *gin.Context) {
	person := &person_domain.Person{}
	ctx.ShouldBindJSON(person)
	p.PeopleDelete.DeletePerson(ctx.Param("id"))
	ctx.JSON(http.StatusOK, person)
}
