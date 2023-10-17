BACKEND EN GO

DESCARGAR GO

LUEGO EN EL REPOSITORIO EN VISUAL STUDIO INICIALIZAR EL PROYECTO CON:

go mod init example.com/"nombredelproyecto"

MONGODB!

EN MONGODB PONEMOS ESTA IP PARA QUE PUEDA RECIBIR SOLICITUDES DE TODOS LADOS
0.0.0.0/0

BAJAR EL COMPASS

ESTABLECER LA CONEXION CON NUESTRAS CREDENCIALES; COPIAR LA URL QUE NOS DA MONGODB EN
LA CONEXION QUE NOS PIDE COMPASS.


LAMBDA -> PERMITE EJECUTAR CODIGO SIN PENSAR EN LOS SERVIDORES.
ES UN SERVICIO SERVER LESS (NO HAY QUE PREOCUPARSE POR LA INFRAESTRUCTURA QUE HAY
DEBAJO DE ESE SERVICIO, YA QUE AMAZON VA A ADMINISTRAR NUESTROS RECURSOS Y PODAMOS SUBIR CODIGO
Y QUE ESE CODIGO SE PUEDA EJECUTAR)

AWS S3

SECRETS MANAGER -> SERVICIO DE AMAZON QUE NOS PERMITE ALOJAR SECRETOS (TOKENS, CREDENCIALES
DE ACCESO A BASE DE DATOS)

Nombre de secreto: twitter--go

API GATEWAY

AGREGARLE UN PROXY, DONDE TODOS LOS ENDPOINTS VAN A PEGAR AHI; SE DEBE COLOCAR
ENTRE LLAVES Y CON UN + AL FINAL: EJ: {twittergo+}
- HABILITAMOS LOS CORS
- PONEMOS LA FUNCION LAMBDA QUE HABIAMOS CREADO; EN ESTE CASO TWITTER

Esto significa que cuando llamemos a mi api gateway remotamente va a enviar el flujo de control
hacia la funcion lambda, y la lambda es la que va a tener la logica de proceso de cada uno
de los endpoints.
Va a conceder permiso a API Gateway para invocar la función Lambda:
arn:aws:lambda:us-east-1:964820568970:function:TWITTER

CREAMOS LA API Y LUEGO LA IMPLEMENTAMOS.
me da esta url: "GUARDADA EN POSTMAN! --> 
ruta del back para usar en postman

Una vez que tenemos la api completa desactivar en API GATEWAY - registros rastreo, cloudwatch
CLOUDWATCH: servicio de logs;
en grupo de registros vemos la lambda que hemos creado, etc.

En configuracion poner en medios binarios, agregar; y ahi poner: multipart/form-data 
(Esto acepta imagenes, etc) --> aca se pueden configurar para recibir pdf, etc.


Empezamos con el BACKEND:
importamos lambda "github.com/aws/aws-lambda-go/lambda"

En la terminal ponemos: go get github.com/aws/aws-lambda-go/lambda. 
Con esto descargamos para poder trabajar con lambda.

Para descargar paquetes en go comenzamos con el go get.

Las variables que se ponen en mayuscula se pueden usar en cualquier lado; es decir, con la mayuscula se hacen publicas.

En la carpeta models va todo lo referido a secrets manager.

Context es una capa global que envuelve todo el desarrollo.


COMPONENTES DEL BACK END:


-- AWSGO.go:

Importaciones:
El componente importa los paquetes necesarios para interactuar con AWS:

context: Proporciona la funcionalidad de contexto para la gestión de tiempo de vida de 
las operaciones.
github.com/aws/aws-sdk-go-v2/aws: Proporciona tipos y funciones comunes para trabajar 
con servicios de AWS.
github.com/aws/aws-sdk-go-v2/config: Permite la configuración de las credenciales y 
opciones de AWS.
Variables:
El componente declara las siguientes variables:

Ctx: Un objeto de contexto que se utiliza para gestionar el tiempo de vida de las operaciones.
Cfg: Un objeto de configuración de AWS que se utilizará para realizar las solicitudes a 
los servicios de AWS.
err: Una variable para capturar posibles errores.
Función InicializoAWS():
Esta función es responsable de inicializar la configuración de AWS. Aquí está lo que hace:

Inicializa el contexto Ctx con context.TODO().
Carga la configuración predeterminada de AWS utilizando config.LoadDefaultConfig().
Configura la región predeterminada para ser "us-east-1".
Si ocurre un error durante la carga de la configuración, se produce un pánico con un 
mensaje de error.
En resumen, este componente en Go inicializa la configuración de AWS utilizando la 
región "us-east-1" y captura cualquier error que ocurra durante este proceso. 
El contexto se inicializa pero no parece ser utilizado en este fragmento específico. 
Este código es una parte fundamental para la interacción exitosa con los servicios de 
AWS desde tu backend en Go.


-- conexion bd.GO

Importaciones:
El componente importa varios paquetes necesarios para trabajar con MongoDB y otros 
componentes relacionados:

context: Proporciona la funcionalidad de contexto para la gestión de tiempo de 
vida de las operaciones.
fmt: Proporciona funciones para formatear y mostrar cadenas.
github.com/Matu36/RED-SOCIAL/models: Importa un paquete que parece contener modelos
 o definiciones relacionadas con una red social.
go.mongodb.org/mongo-driver/mongo: Proporciona funcionalidades para interactuar con MongoDB.
go.mongodb.org/mongo-driver/mongo/options: Proporciona opciones de configuración para 
el cliente de MongoDB.
Variables:
El componente declara las siguientes variables:

MongoCN: Un puntero a un cliente de MongoDB, que se utilizará para realizar operaciones 
en la base de datos.
DatabaseName: Un nombre de base de datos, que parece ser extraído del contexto.
Función ConectarBD(ctx context.Context) error:
Esta función es responsable de conectar la base de datos MongoDB. Aquí está lo que hace:

Extrae información de contexto, como el usuario, contraseña, host y nombre de la base de datos, para construir una cadena de conexión.
Crea opciones de configuración del cliente con la cadena de conexión.
Conecta al cliente de MongoDB utilizando las opciones de configuración.
Verifica la conexión utilizando el método Ping.
Si la conexión es exitosa, guarda el cliente en la variable MongoCN y el nombre de la base 
de datos en DatabaseName.
Función BaseConectada() bool:
Esta función verifica si hay una conexión activa con la base de datos MongoDB. Aquí está 
lo que hace:

Realiza un ping a la base de datos utilizando el cliente MongoCN y un contexto vacío.
Retorna true si el ping es exitoso (no hay error) y false en caso contrario.
En resumen, este componente en Go se encarga de establecer y gestionar una conexión 
con una base de datos MongoDB. Proporciona funciones para conectar la base de datos y
 verificar la conexión.


-- handlers.go

Importaciones:
El componente importa varios paquetes necesarios para manejar solicitudes de eventos de 
AWS Lambda y modelos relacionados con una red social:

context: Proporciona la funcionalidad de contexto para la gestión de tiempo de vida de 
las operaciones.
fmt: Proporciona funciones para formatear y mostrar cadenas.
github.com/Matu36/RED-SOCIAL/models: Importa un paquete que parece contener modelos o 
definiciones relacionadas con una red social.
github.com/aws/aws-lambda-go/events: Importa paquetes relacionados con los eventos de 
AWS Lambda y la API Gateway.
Función Manejadores(ctx context.Context, request events.APIGatewayProxyRequest) 
models.ResApi:
Esta función es el manejador principal que procesa las solicitudes entrantes en función 
del contexto y la solicitud. Aquí está lo que hace:

Imprime en la consola detalles sobre la solicitud que se va a procesar, como la ruta 
y el método HTTP.
Declara una variable r de tipo models.ResApi que parece ser una estructura de respuesta.
Inicializa el campo Status de la respuesta con el valor 400.
La función utiliza un conjunto de instrucciones switch para manejar diferentes métodos 
HTTP (POST, GET, PUT, DELETE). En cada caso, el componente actualmente no tiene 
implementada ninguna lógica específica. Cada caso vacío 
(por ejemplo, switch ctx.Value(models.Key("path")).(string) {}) está esperando 
ser llenado con la lógica correspondiente para cada tipo de solicitud y ruta.

Si no coincide con ninguno de los casos anteriores, establece el campo Message 
de la respuesta en "Method Invalid".

Finalmente, la función retorna la estructura de respuesta r.

En resumen, este componente en Go es un manejador de funciones para solicitudes en 
una API Gateway a través de AWS Lambda. Actualmente, no contiene lógica específica 
para cada tipo de solicitud y ruta, pero está diseñado para ser completado con la 
lógica adecuada en cada caso.


-- procesoToken.go

Importaciones:
El componente importa varios paquetes necesarios para trabajar con tokens JWT y 
modelos relacionados con una red social:

errors: Proporciona funciones y tipos para manejar errores.
strings: Proporciona funciones para trabajar con cadenas.
github.com/Matu36/RED-SOCIAL/models: Importa un paquete que parece contener modelos
 o definiciones relacionadas con una red social.
github.com/golang-jwt/jwt/v5: Importa el paquete jwt para trabajar con tokens JWT.
Variables:
El componente declara las siguientes variables:

Email: Una variable que parece almacenar la dirección de correo electrónico asociada 
con el token.
IDUsuario: Una variable que parece almacenar el ID del usuario asociado con el token.
Función ProcesoToken(tk string, JWTSign string) (*models.Claim, bool, string, error):
Esta función es responsable de procesar un token JWT. Aquí está lo que hace:

Recibe un token JWT (tk) y la clave de firma (JWTSign) como parámetros.
Declara una variable para almacenar las reclamaciones del token (claims).
Divide el token recibido por el espacio para separar la parte "Bearer" del token real.
Si el token no tiene exactamente dos partes después de la división, devuelve un error de 
formato de token inválido.
Limpia espacios en blanco y asigna el token real a tk.
Intenta analizar el token JWT utilizando la clave de firma (miClave) y el tipo de 
reclamaciones (claims).
Si la analización tiene éxito, realiza una rutina que parece verificar el token en 
la base de datos (se espera que se implemente).
Si el token no es válido, devuelve un error de token inválido.
Finalmente, devuelve las reclamaciones, un valor booleano, una cadena y un posible error.
En resumen, este componente en Go es un manejador de funciones relacionado con la 
generación y validación de tokens JWT para autenticación. La implementación actual 
es una base que espera ser llenada con lógica personalizada, como la validación del 
token en la base de datos y la asignación de datos del token a las variables Email e 
IDUsuario.


-- sm.go

Importaciones:
El componente importa varios paquetes necesarios para trabajar con Secrets Manager, 
AWS SDK y modelos relacionados con una red social:

encoding/json: Proporciona funciones para codificar y decodificar datos JSON.
fmt: Proporciona funciones para formatear y mostrar cadenas.
github.com/Matu36/RED-SOCIAL/awsgo: Importa un paquete que parece relacionarse con la 
configuración de AWS.
github.com/Matu36/RED-SOCIAL/models: Importa un paquete que contiene definiciones de 
modelos relacionados con una red social.
github.com/aws/aws-sdk-go-v2/aws: Importa el paquete AWS SDK para trabajar con tipos y 
funciones comunes de AWS.
github.com/aws/aws-sdk-go-v2/service/secretsmanager: Importa el paquete AWS SDK para 
trabajar con el servicio AWS Secrets Manager.
Función GetSecret(secretName string) (models.Secret, error):
Esta función se encarga de obtener un secreto desde AWS Secrets Manager y decodificarlo. 
Aquí está lo que hace:

Declara una variable datosSecret del tipo models.Secret, que aparentemente se utilizará 
para almacenar el secreto decodificado.
Imprime en la consola que se está pidiendo un secreto con el nombre proporcionado.
Crea un cliente de Secrets Manager (SVC) utilizando la configuración de AWS.
Llama a GetSecretValue para obtener los detalles del secreto con el nombre especificado.
Si ocurre un error al obtener el secreto, lo imprime en la consola y retorna el valor de error.
Utiliza json.Unmarshal para decodificar el contenido del campo SecretString en el 
resultado en la estructura datosSecret.
Imprime en la consola que la lectura del secreto fue exitosa.
Retorna datosSecret y un valor nulo de error.
En resumen, este componente en Go interactúa con AWS Secrets Manager para obtener y 
decodificar secretos almacenados en el servicio. Los detalles del secreto son 
decodificados y almacenados en una estructura definida en models.Secret.


PAQUETES:

MAIN.GO
En resumen, estas funciones se utilizan para configurar y gestionar una función lambda 
que interactúa con AWS Secrets Manager, valida las variables de entorno y procesa 
solicitudes de API Gateway utilizando la función Manejadores del paquete handlers. 
También establecen el contexto y la configuración necesarios para realizar estas 
operaciones.

AWSGO:

Este paquete simplifica la inicialización de la configuración de AWS y proporciona 
un punto de entrada centralizado para la configuración de AWS SDK en tu aplicación. 
Esto puede ser útil para asegurarte de que todas las partes de tu aplicación estén 
configuradas de la misma manera y sigan las mejores prácticas al interactuar con 
servicios de AWS.

BD:

Un paquete "bd" en una aplicación Go es responsable de la interacción con la base 
de datos subyacente, permitiendo a la aplicación almacenar y recuperar datos. La 
funcionalidad específica del paquete dependerá de los requisitos de tu aplicación y
de la tecnología de base de datos que estés utilizando.

Ejemplo de una función en BD:

Esta es la función: 
package bd

import (
	"context"

	"github.com/Matu36/RED-SOCIAL/models"
)

func BorroRelacion(t models.Relacion) (bool, error) {
	ctx := context.TODO()
	db := MongoCN.Database(DatabaseName)
	col := db.Collection("relacion")

	_, err := col.DeleteOne(ctx, t)
	if err != nil {
		return false, err
	}

	return true, nil
}

Explicación:

Inicializa un contexto (ctx) usando context.TODO(). El contexto se utiliza para controlar 
el tiempo de vida y los valores asociados con las operaciones de la base de datos.

Obtiene una referencia a la base de datos (db) a partir de MongoCN. El nombre 
MongoCN sugiere que se está utilizando una base de datos MongoDB, y se asume que 
MongoCN contiene la configuración para la conexión a la base de datos.

Obtiene una referencia a una colección llamada "relacion" en la base de datos. Esto 
se hace con db.Collection("relacion"), lo que indica que se está trabajando con la 
colección "relacion" en MongoDB.

Utiliza el método DeleteOne para eliminar un documento de la colección. El documento 
que se va a eliminar se especifica como t. En este contexto, t representa un objeto de 
tipo models.Relacion que probablemente contiene información necesaria para identificar
el documento específico que se desea eliminar.

Verifica si se produjo un error durante la eliminación. Si se encuentra un error, 
la función devuelve false y el error.

Si la eliminación se realizó con éxito, la función devuelve true y un valor nil 
para indicar que la operación se realizó con éxito y no hubo errores.


HANDLERS:

El paquete "handlers" generalmente se utiliza para definir y gestionar las funciones 
que manejan las solicitudes entrantes en una aplicación web o servicio en Go. Estas 
funciones, conocidas como controladores o manejadores, son responsables de procesar 
las solicitudes HTTP, interactuar con la lógica de la aplicación y devolver respuestas 
adecuadas al cliente. El paquete "handlers" puede incluir funciones para diversas rutas y 
métodos HTTP.

Básicamente se ponen las rutas y se definen los controladores, los cuales se definen
luego en routers.

JWT:

El paquete jwt se usa para el manejo de tokens JWT (JSON Web Tokens).
Básicamente, el paquete jwt se utiliza para:

Generar tokens JWT: Puede generar tokens JWT para autenticar a los usuarios o 
proporcionar autorización para realizar ciertas acciones en tu aplicación.

Validar tokens JWT: Puede verificar si un token JWT proporcionado es válido, 
lo que implica verificar su firma digital y verificar su integridad.

Decodificar tokens JWT: Puede extraer la información contenida en un token JWT,
como los datos del usuario o cualquier reclamo (claim) específico que hayas incluido
en el token.

Procesar tokens JWT: Puede llevar a cabo la lógica de procesamiento necesaria para validar 
y gestionar los tokens JWT utilizados en la autenticación y autorización de los usuarios.


MODELS:

Los models se utilizan para definir la estructura de la base de datos.
EJEMPLO:

type Claim struct {
	Email string             `json:"email"`
	ID    primitive.ObjectID `bson:"_id" json:"_id,omitempy"`
	jwt.RegisteredClaims
}

El tipo de datos "BSON" se refiere a "Binary JSON" (JSON binario), que es 
un formato de serialización de datos utilizado comúnmente en bases de datos NoSQL, 
especialmente en bases de datos orientadas a documentos como MongoDB.


ROUTERS:

En resumen, el paquete "routers" se utiliza para definir cómo se deben manejar 
las solicitudes HTTP entrantes, cómo se deben enrutar a las funciones apropiadas 
y cómo se deben procesar. Esto facilita la organización y mantenimiento de una 
aplicación web o API al separar las responsabilidades de enrutamiento y lógica 
de manejo de solicitudes.

Se encarga de dirigir las solicitudes del cliente a las funciones o controladores 
correspondientes que deben manejar esas solicitudes.


EXPLICACIÓN BREVE DE ROUTERS, BD Y HANDLERS:

En HANDLERS se definen las rutas, los métodos y que función corresponde a cada ruta.
En BD hago la lógica de la función.
En ROUTERS hago el manejo de errores, etc.


SECRETMANAGER:

El paquete "secretManager" se usa para interactuar con el AWS Secrets Manager, 
un servicio de AWS que permite gestionar secretos y configuraciones en aplicaciones 
de forma segura. El Secrets Manager se utiliza para almacenar, recuperar y rotar 
secretos, como contraseñas, claves de API u otra información confidencial.


METODOS, VARIABLES, ETC:




