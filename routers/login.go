package routers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/Matu36/RED-SOCIAL/bd"
	"github.com/Matu36/RED-SOCIAL/jwt"
	"github.com/Matu36/RED-SOCIAL/models"
	"github.com/aws/aws-lambda-go/events"
)

func Login(ctx context.Context) models.ResApi {
	var t models.Usuario
	var r models.ResApi
	r.Status = 400

	body := ctx.Value(models.Key("body")).(string)
	err := json.Unmarshal([]byte(body), &t)

	if err != nil {
		r.Message = "Usuario y/o Contraseñas Inválidos " + err.Error()
		return r
	}
	if len(t.Email) == 0 {
		r.Message = "El email del usuario es requerido"
		return r
	}

	userData, existe := bd.IntentoLogin(t.Email, t.Password)
	if !existe {
		r.Message = "Usuario y/o contraseñas inválidos " + err.Error()
		return r
	}

	jwtKey, err := jwt.GeneroJWT(ctx, userData)
	if err != nil {
		r.Message = "Ocurrió un error al generar el token correspondiente " + err.Error()
		return r
	}

	resp := models.RespuestaLogin{
		Token: jwtKey,
	}

	token, err2 := json.Marshal(resp)
	if err2 != nil {
		r.Message = "Ocurrió un error al intentar formatear el token a JSON " + err.Error()
		return r
	}

	// La cookie que almacena el token del usuario logueado para que no tenga que volver
	//a ingresar sus credenciales en un determinado tiempo que le pasamos mas abajo

	cookie := &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: time.Now().Add(24 * time.Hour),
	}

	cookieString := cookie.String()

	res := &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(token),
		Headers: map[string]string{
			"Content-Type":                "application/json",
			"Access-Control-Allow-Origin": "*",
			"Set-Cookie":                  cookieString,
		},
	}

	r.Status = 200
	r.Message = string(token)
	r.CustomResp = res

	return r
}
