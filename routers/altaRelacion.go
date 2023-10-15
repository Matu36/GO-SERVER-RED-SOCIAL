package routers

import (
	"context"

	"github.com/Matu36/RED-SOCIAL/bd"
	"github.com/Matu36/RED-SOCIAL/models"
	"github.com/aws/aws-lambda-go/events"
)

func AltaRelacion(ctx context.Context, request events.APIGatewayProxyRequest, claim models.Claim) models.ResApi {
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

	status, err := bd.InsertoRelacion(t)
	if err != nil {
		r.Message = "Ocurriò un error al intentar insertar relaciòn " + err.Error()
		return r
	}

	if !status {
		r.Message = "No se ha logrado insertar la relaciòn"
		return r
	}

	r.Status = 200
	r.Message = "Alta de relacion OK"
	return r

}
