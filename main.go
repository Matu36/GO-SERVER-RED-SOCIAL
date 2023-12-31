package main

import (
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/Matu36/RED-SOCIAL/awsgo"
	"github.com/Matu36/RED-SOCIAL/bd"
	"github.com/Matu36/RED-SOCIAL/handlers"
	"github.com/Matu36/RED-SOCIAL/models"

	"github.com/Matu36/RED-SOCIAL/secretmanager"
	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
)

func main() {
	// Inicia la función lambda, que se activa cuando ocurre un evento
	lambda.Start(EjecutoLambda)

}

func EjecutoLambda(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var res *events.APIGatewayProxyResponse
	// Inicializa la configuración de AWS, como la región
	awsgo.InicializoAWS()

	// Valida si se proporcionan los parámetros necesarios en las variables de entorno
	if !ValidoParametros() {
		res = &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Error en las variables de entorno. Deben incluir 'SecretName', 'BucketName', 'UrlPrefix'",
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return res, nil
	}

	headers := map[string]string{

		"Access-Control-Allow-Origin":      "*", // Usar el comodín para permitir cualquier origen
		"Access-Control-Allow-Methods":     "OPTIONS, GET, POST, PUT, DELETE",
		"Access-Control-Allow-Headers":     "Content-Type, Authorization, X-Amz-Date, X-Api-Key, X-Amz-Security-Token",
		"Access-Control-Allow-Credentials": "true",
	}

	// Si la solicitud es una OPTIONS (preflight), responde con encabezados CORS sin procesar la solicitud.
	if request.HTTPMethod == http.MethodOptions {
		return &events.APIGatewayProxyResponse{
			StatusCode: http.StatusOK,
			Headers:    headers,
			Body:       "",
		}, nil
	}

	// Obtiene el secreto almacenado en AWS Secrets Manager
	SecretModel, err := secretmanager.GetSecret(os.Getenv("SecretName"))

	if err != nil {
		res = &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Error en la lectura de Secret" + err.Error(),
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return res, nil
	}

	// Procesa la ruta eliminando el prefijo y configura el contexto para la solicitud
	path := strings.Replace(request.PathParameters["twittergo"], os.Getenv("UrlPrefix"), "", -1) //Esto lo que hace es eliminar el prefijo de la ruta que le pasamos en el array para que la ruta quede limpia
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("path"), path)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("method"), request.HTTPMethod)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("user"), SecretModel.Username)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("password"), SecretModel.Password)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("host"), SecretModel.Host)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("database"), SecretModel.Database)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("jwtSign"), SecretModel.JWTSign)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("body"), request.Body)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("bucketName"), os.Getenv("BucketName"))

	//Chequeo de conexion a base de datos

	err = bd.ConectarBD(awsgo.Ctx)
	if err != nil {
		res = &events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Error Conectando la base de datos" + err.Error(),
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return res, nil
	}

	// Llama a la función Manejadores para procesar la solicitud y obtener una respuesta
	respAPI := handlers.Manejadores(awsgo.Ctx, request)
	if respAPI.CustomResp == nil {
		// Si la función Manejadores no proporciona una respuesta personalizada,
		//crea una respuesta predeterminada
		res = &events.APIGatewayProxyResponse{
			StatusCode: respAPI.Status,
			Body:       respAPI.Message,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}
		return res, nil
	} else {
		// Devuelve la respuesta personalizada proporcionada por la función Manejadores
		return respAPI.CustomResp, nil
	}

}

func ValidoParametros() bool {
	// Valida si se han configurado las variables de entorno necesarias
	_, traeParametro := os.LookupEnv("SecretName")
	if !traeParametro {
		return traeParametro
	}

	_, traeParametro = os.LookupEnv("BucketName")
	if !traeParametro {
		return traeParametro
	}

	_, traeParametro = os.LookupEnv("UrlPrefix")
	if !traeParametro {
		return traeParametro
	}

	return traeParametro
}
