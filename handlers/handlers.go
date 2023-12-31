package handlers

import (
	"context"
	"fmt"

	"github.com/Matu36/RED-SOCIAL/jwt"
	"github.com/Matu36/RED-SOCIAL/models"
	"github.com/Matu36/RED-SOCIAL/routers"
	"github.com/aws/aws-lambda-go/events"
)

// Esta función es el controlador principal que maneja las solicitudes HTTP.
// Determina la acción a tomar en función de la ruta y el método HTTP.
// Devuelve una respuesta de API personalizada.

func Manejadores(ctx context.Context, request events.APIGatewayProxyRequest) models.ResApi {
	fmt.Println("Voy a procesar " + ctx.Value(models.Key("path")).(string) + " > " + ctx.Value(models.Key("method")).(string))

	var r models.ResApi
	r.Status = 400

	isOK, StatusCode, msg, claim := validoAuthorization(ctx, request)
	if !isOK {
		r.Status = StatusCode
		r.Message = msg
		fmt.Printf("Error en la autorización: %s\n", msg)
		return r
	}

	switch ctx.Value(models.Key("method")).(string) {
	case "POST":
		// Maneja las solicitudes POST para diferentes rutas. Llama a funciones específicas
		// para procesar cada tipo de solicitud y devuelve la respuesta correspondiente.
		switch ctx.Value(models.Key("path")).(string) {
		case "registro":
			return routers.Registro(ctx)

		case "login":
			return routers.Login(ctx)

		case "tweet":
			return routers.GraboTweet(ctx, claim)

		case "altaRelacion":
			return routers.AltaRelacion(ctx, request, claim)

		case "subirAvatar":
			return routers.UploadImage(ctx, "A", request, claim)
		case "subirBanner":
			return routers.UploadImage(ctx, "B", request, claim)

		default:
			fmt.Printf("Error: Ruta POST no manejada - %s\n", ctx.Value(models.Key("path")).(string))

		}

	case "GET":
		switch ctx.Value(models.Key("path")).(string) {
		case "verPerfil":
			return routers.VerPerfil(request)
		case "leoTweets":
			return routers.LeoTweets(request)
		case "obtenerAvatar":
			return routers.ObtenerImagen(ctx, "A", request, claim)
		case "obtenerBanner":
			return routers.ObtenerImagen(ctx, "B", request, claim)
		case "consultaRelacion":
			return routers.ConsultaRelacion(request, claim)
		case "listaUsuarios":
			return routers.ListaUsuarios(request, claim)
		case "leoTweetsSeguidores":
			return routers.LeoTweetsSeguidores(request, claim)

		default:
			fmt.Printf("Error: Ruta GET no manejada - %s\n", ctx.Value(models.Key("path")).(string))
		}

	case "PUT":
		switch ctx.Value(models.Key("path")).(string) {
		case "modificarPerfil":
			return routers.ModificarPerfil(ctx, claim)

		}

	case "DELETE":
		switch ctx.Value(models.Key("path")).(string) {
		case "eliminarTweet":
			return routers.EliminarTweet(request, claim)
		case "bajaRelacion":
			return routers.BajaRelacion(request, claim)

		}
	}

	r.Message = "Method Invalid"
	return r

}

// Esta función valida la autorización de una solicitud. Verifica si se proporcionó un token
// de autorización y si es válido. Devuelve información sobre la validez del token.

func validoAuthorization(ctx context.Context, request events.APIGatewayProxyRequest) (bool, int, string, models.Claim) {
	path := ctx.Value(models.Key("path")).(string)

	// Permite rutas sin token de autorización. Se utiliza para registro,
	//inicio de sesión y otras rutas públicas.

	if path == "registro" || path == "login" || path == "obtenerAvatar" || path == "obtenerBanner" {
		return true, 200, "", models.Claim{}
	}

	token := request.Headers["Authorization"]
	if len(token) == 0 {
		return false, 401, "Token requerido", models.Claim{}
	}

	claim, todoOK, msg, err := jwt.ProcesoToken(token, ctx.Value(models.Key("jwtSign")).(string))
	if !todoOK {
		if err != nil {
			fmt.Println("Error en el token " + err.Error())
			return false, 401, err.Error(), models.Claim{}
		} else {
			fmt.Println("Error en ele token " + msg)
			return false, 401, msg, models.Claim{}
		}
	}

	fmt.Println("Token OK")
	return true, 200, msg, *claim

}
