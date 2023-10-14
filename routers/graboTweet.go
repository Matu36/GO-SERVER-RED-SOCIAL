package routers

import (
	"context"
	"encoding/json"
	"time"

	"github.com/Matu36/RED-SOCIAL/bd"
	"github.com/Matu36/RED-SOCIAL/models"
)

func GraboTweet(ctx context.Context, claim models.Claim) models.ResApi {
	var mensaje models.Tweet
	var r models.ResApi
	r.Status = 400
	IDUsuario := claim.ID.Hex()

	body := ctx.Value(models.Key("body")).(string)
	err := json.Unmarshal([]byte(body), &mensaje)

	if err != nil {
		r.Message = "Ocurrió un error al intentar decodificar el body " + err.Error()
		return r
	}

	registro := models.GraboTweet{
		UserID:  IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}

	_, status, err := bd.InsertoTweet(registro)
	if err != nil {
		r.Message = "Ocurriò un error al intentar insertar el registro " + err.Error()
		return r
	}

	if !status {
		r.Message = "No se ha logrado insertar el tweet " + err.Error()
		return r
	}

	r.Status = 200
	r.Message = "Tweet Creado Correctamente"
	return r

}
