package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Usuario struct {
	ID              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Nombre          string             `bson:"nombre" json:"nombre,omitempy"`
	Apellidos       string             `bson:"apellidos" json:"apellidos,omitempy"`
	FechaNacimiento time.Time          `bson:"fechaNacimiento" json:"fechaNacimiento,omitempy"`
	Email           string             `bson:"email" json:"email"`
	Password        string             `bson:"password" json:"password,omitempy"`
	Avatar          string             `bson:"avatar" json:"avatar,omitempy"`
	Banner          string             `bson:"banner" json:"banner,omitempy"`
	Biografia       string             `bson:"biografia" json:"biografia,omitempy"`
	Ubicacion       string             `bson:"ubicacion" json:"ubicacion,omitempy"`
	SitioWeb        string             `bson:"sitioweb" json:"sitioweb,omitempy"`
}
