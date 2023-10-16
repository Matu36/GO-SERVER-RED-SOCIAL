package jwt

//  Esto envia las credenciales las cuales luego va a recibir el FrontEnd //

import (
	"errors"
	"strings"

	"github.com/Matu36/RED-SOCIAL/bd"
	"github.com/Matu36/RED-SOCIAL/models"
	jwt "github.com/golang-jwt/jwt/v5"
)

var Email string
var IDUsuario string

// ProcesoToken verifica y decodifica un token JWT, y realiza una verificación contra la
// base de datos.
// Recibe el token JWT en forma de cadena y la clave de firma para validar el token.

func ProcesoToken(tk string, JWTSign string) (*models.Claim, bool, string, error) {
	// Convierte la clave de firma en un byte slice.
	miClave := []byte(JWTSign)

	// Estructura para almacenar los reclamos (claims) del token JWT.
	var claims models.Claim

	// Divide el token en la parte "Bearer" y el token real.
	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return &claims, false, string(""), errors.New("formato de token inválido")
	}

	tk = strings.TrimSpace(splitToken[1])

	// Parsea el token con las reclamaciones y verifica la firma utilizando la clave de firma.
	tkn, err := jwt.ParseWithClaims(tk, &claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})

	if err == nil {
		// Rutina que verifica si el usuario ya existe en la base de datos.
		// Si se encuentra un usuario, almacena su email y ID en las variables
		// globales Email e IDUsuario.

		_, encontrado, _ := bd.ChequeoYaExisteUsuario(claims.Email)
		if encontrado {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}
		return &claims, encontrado, IDUsuario, nil
	}

	if !tkn.Valid {
		return &claims, false, string(""), errors.New("Token Inválido")
	}

	return &claims, false, string(""), err

}
