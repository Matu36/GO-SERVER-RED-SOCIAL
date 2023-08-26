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

	SVC := secretsmanager.NewFromConfig(awsgo.Cfg)
	clave, err := SVC.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	})

	if err != nil {
		fmt.Println(err.Error())
		return datosSecret, err
	}

	json.Unmarshal([]byte(*clave.SecretString), &datosSecret)
	fmt.Println("> Lectura de Secret Ok " + secretName)
	return datosSecret, nil
}
