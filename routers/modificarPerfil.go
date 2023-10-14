package routers

import (
	"context"
	"encoding/json"

	"github.com/Matu36/RED-SOCIAL/bd"
	"github.com/Matu36/RED-SOCIAL/models"
)

func ModificarPerfil(ctx context.Context, claim models.Claim) models.ResApi {
	var r models.ResApi
	r.Status = 400

	var t models.Usuario

	body := ctx.Value(models.Key("body")).(string)
	err := json.Unmarshal([]byte(body), &t)

	if err != nil {
		r.Message = "Datos incorrectos " + err.Error()
	}

	status, err := bd.ModificoRegistro(t, claim.ID.Hex())

	if err != nil {
		r.Message = "Ocurriò un error al querer modificar el registro " + err.Error()
		return r
	}

	if !status {
		r.Message = "No se ha logrado modificar el registro del usuario"
		return r
	}

	r.Status = 200
	r.Message = "Modificaciòn Perfil Ok!"
	return r
}
