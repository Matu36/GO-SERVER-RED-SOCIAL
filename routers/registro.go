package routers

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Matu36/RED-SOCIAL/bd"
	"github.com/Matu36/RED-SOCIAL/models"
)

func Registro(ctx context.Context) models.ResApi {

	var t models.Usuario
	var r models.ResApi
	r.Status = 400

	fmt.Println("Entre a Registro")

	body := ctx.Value(models.Key("body")).(string)
	err := json.Unmarshal([]byte(body), &t)

	if err != nil {
		r.Message = err.Error()
		fmt.Println(r.Message)
		return r
	}

	if len(t.Email) == 0 {
		r.Message = "Debe especificar el Email"
		fmt.Print(r.Message)
		return r
	}

	if len(t.Password) < 6 {
		r.Message = "Debe especificar una contraseña de al menos 6 caracteres"
		fmt.Print(r.Message)
		return r
	}

	_, encontrado, _ := bd.ChequeoYaExisteUsuario(t.Email)
	if encontrado {
		r.Message = "Ya existe un usuario registrado con ese Email"
		fmt.Print(r.Message)
		return r
	}
	_, status, err := bd.InsertoRegistro(t)
	if err != nil {
		r.Message = "Ocurrió un error al intentar realizar el registro del usuario " + err.Error()
		fmt.Print(r.Message)
		return r
	}

	if !status {
		r.Message = "No se ha logrado intentar el registro de usuario"
		fmt.Print(r.Message)
		return r
	}

	r.Status = 200
	r.Message = "Registro Ok"
	fmt.Print(r.Message)
	return r
}
