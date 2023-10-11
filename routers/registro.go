package routers

import (
	"context"
	"encoding/json"
	"fmt"

"github.com/Matu36/RED-SOCIAL/bd"
"github.com/Matu36/RED-SOCIAL/models"

)

func Registro (ctx context.Context) models.RespApi {

var t models.Usuario
var r models.RespApi
r.Status=400

fmt.Println ("Entre a Registro")

body := ctx.Value (models.Key("body"))
err := json.Unmarshal([] byte(body), &t)

if err !nil {
	r.Message = err.Error()
	fmt Println(r.Message)
	return r
}



}