package person_infrastructure

import (
	"net/http"

	person_application "github.com/MeliCGS/go-project-api-arq-hexa/api/person/application"
	"github.com/gin-gonic/gin"
)

type PersonController struct {
	PeopleSearcher *person_application.PeopleSearcher
}

func (p *PersonController) GetAllHandler(ctx *gin.Context) {
	people := p.PeopleSearcher.SearchAll()
	ctx.JSON(http.StatusOK, people)
}
