package awsgo

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

var Ctx context.Context
var Cfg aws.Config
var err error

func InicializoAWS() {
	Ctx = context.TODO()
	Cfg, err = config.LoadDefaultConfig(Ctx, config.WithDefaultRegion("us-east-1"))
	if err != nil {
		panic("error al cargar la configuracion .aws/config" + err.Error())
	}
}

/*

Ctx y Cfg son variables que almacenan el contexto y la configuración de AWS, respectivamente.
Estas variables se utilizan en toda la aplicación para proporcionar la configuración
necesaria al interactuar con servicios de AWS.

InicializoAWS es una función que inicializa la configuración de AWS.
Aquí está lo que hace en detalle:

Inicializa un contexto vacío (context.TODO()): El contexto se utiliza
para gestionar los valores y tiempo de vida de los objetos AWS SDK y sus solicitudes.

Carga la configuración por defecto de AWS usando config.LoadDefaultConfig.
Esto carga la configuración desde varios lugares, incluido el archivo de
configuración .aws/config de tu entorno, las variables de entorno, entre otros. Además, se
especifica la región por defecto como "us-east-1".

Si se produce un error durante la carga de la configuración, la función lanza un
pánico, lo que significa que tu aplicación se detendrá. Esto suele ser una estrategia
para manejar errores críticos que no pueden recuperarse.

*/
