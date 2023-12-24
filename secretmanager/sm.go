package secretmanager

import (
	"encoding/json"
	"fmt"

	"github.com/Matu36/RED-SOCIAL/awsgo"
	"github.com/Matu36/RED-SOCIAL/models"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

func GetSecret(secretName string) (models.Secret, error) {
	var datosSecret models.Secret
	fmt.Println("> Pido Secreto " + secretName)

	// Crea un cliente para interactuar con AWS Secrets Manager
	svc := secretsmanager.NewFromConfig(awsgo.Cfg)

	// Llama a AWS Secrets Manager para obtener el secreto
	clave, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	})

	if err != nil {
		// Si se produce un error al obtener el secreto, se registra y se devuelve un error
		fmt.Println(err.Error())
		return datosSecret, err
	}
	// Decodifica el secreto obtenido (en formato JSON) y lo almacena en datosSecret
	json.Unmarshal([]byte(*clave.SecretString), &datosSecret)
	// Registro del éxito al leer el secreto
	fmt.Println("> Lectura de Secret Ok " + secretName)
	// Devuelve el secreto descifrado y sin errores
	return datosSecret, nil
}
