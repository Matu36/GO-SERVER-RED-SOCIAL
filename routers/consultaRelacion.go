package routers

import (
	"encoding/json"

	"github.com/Matu36/RED-SOCIAL/bd"
	"github.com/Matu36/RED-SOCIAL/models"
	"github.com/aws/aws-lambda-go/events"
)

func ConsultaRelacion(request events.APIGatewayProxyRequest, claim models.Claim) models.ResApi {

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

	var resp models.RespuestaConsultaRelacion

	hayRelacion := bd.ConsultoRelacion(t)

	if !hayRelacion {
		resp.Status = false
	} else {
		resp.Status = true
	}

	respJson, err := json.Marshal(hayRelacion)
	if err != nil {
		r.Status = 500
		r.Message = "Error al formatear los datos del usuario como json " + err.Error()
		return r
	}

	r.Status = 200
	r.Message = string(respJson)
	return r

}
