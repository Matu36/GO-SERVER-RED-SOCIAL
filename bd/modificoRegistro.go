package bd

import (
	"context"

	"github.com/Matu36/RED-SOCIAL/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ModificoRegistro(u models.Usuario, ID string) (bool, error) {

	ctx := context.TODO()
	db := MongoCN.Database(DatabaseName)
	col := db.Collection("Usuarios")

	registro := make(map[string]interface{})
	if len(u.Nombre) > 0 {
		registro["nombre"] = u.Nombre
	}

	if len(u.Apellidos) > 0 {
		registro["apellidos"] = u.Apellidos

		registro["fechaNacimiento"] = u.FechaNacimiento

		if len(u.Avatar) > 0 {
			registro["avatar"] = u.Avatar
		}

		if len(u.Banner) > 0 {
			registro["banner"] = u.Banner
		}

		if len(u.Biografia) > 0 {
			registro["biografia"] = u.Biografia
		}
		if len(u.SitioWeb) > 0 {
			registro["sitioweb"] = u.SitioWeb
		}

	}
	updtString := bson.M{
		"$set": registro,
	}

	objID, _ := primitive.ObjectIDFromHex(ID)
	filtro := bson.M{"_id": bson.M{"$eq": objID}}

	_, err := col.UpdateOne(ctx, filtro, updtString)
	if err != nil {
		return false, err
	}

	return true, nil
}

// VIDEO 61 MINUTO 7:00, VER QUE HACE CON EL TOKEN EN POSTMAN
