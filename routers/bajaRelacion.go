package routers

import (
	"github.com/Matu36/RED-SOCIAL/bd"
	"github.com/Matu36/RED-SOCIAL/models"
	"github.com/aws/aws-lambda-go/events"
)

func BajaRelacion(request events.APIGatewayProxyRequest, claim models.Claim) models.ResApi {
	var r models.ResApi
	r.Status = 400

	ID := request.QueryStringParameters["id"]
	if len(ID) < 1 {
		r.Message = "El parámetro ID es obligatorio"
		return r
	}

	var t models.Relacion
	t.UsuarioID = claim.ID.Hex()
	t.UsuarioRelacionID = ID

	status, err := bd.BorroRelacion(t)
	if err != nil {
		r.Message = "Ocurriò un error al intentar borrar la relaciòn"
		return r
	}

	if !status {
		r.Message = "No se ha logrado borrar la relaciòn"
		return r
	}

	r.Status = 200
	r.Message = "Baja relaciòn OK !"
	return r

}
