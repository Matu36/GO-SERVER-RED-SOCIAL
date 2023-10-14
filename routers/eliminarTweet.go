package routers

import (
	"github.com/Matu36/RED-SOCIAL/bd"
	"github.com/Matu36/RED-SOCIAL/models"
	"github.com/aws/aws-lambda-go/events"
)

func EliminarTweet(request events.APIGatewayProxyRequest, claim models.Claim) models.ResApi {

	var r models.ResApi
	r.Status = 400

	ID := request.QueryStringParameters["id"]
	if len(ID) < 1 {
		r.Message = "El paràmetro ID es obligatorio"
		return r
	}

	err := bd.BorroTweet(ID, claim.ID.Hex())
	if err != nil {
		r.Message = "Ocurrió un error al intentar eliminar el tweet " + err.Error()
		return r
	}

	r.Message = "Eliminar tweet Ok !"
	r.Status = 200
	return r

}
