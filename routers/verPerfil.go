package routers

import (
	"encoding/json"
	"fmt"

	"github.com/Matu36/RED-SOCIAL/bd"
	"github.com/Matu36/RED-SOCIAL/models"
	"github.com/aws/aws-lambda-go/events"
)

func VerPerfil(request events.APIGatewayProxyRequest) models.ResApi {
	var r models.ResApi
	r.Status = 400

	fmt.Println("Entre en VerPerfil")

	ID := request.QueryStringParameters["id"]
	if len(ID) < 1 {
		r.Message = "El parámetro ID es obligatorio"
		return r
	}

	perfil, err := bd.BuscoPerfil(ID)
	if err != nil {
		r.Message = "Ocurrió un error al intentar buscar el registro " + err.Error()
		return r
	}
	respJson, err := json.Marshal(perfil)
	if err != nil {
		r.Status = 500
		r.Message = "Error al formatear los datos del usuario como Json " + err.Error()
		return r
	}

	r.Status = 200
	r.Message = string(respJson)
	return r
}

//MINUTO 1:50 DE VER PERFIL EXPLICAR EL TEMA DEL BEARER (TOKEN). VIDEO 58
