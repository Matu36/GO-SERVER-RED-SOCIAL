package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DevuelvoTweetsSeguidores struct {
	ID               primitive.ObjectID `bson:"_id" json:"_id,omitempy"`
	UsuarioID        string             `bson:"usuarioid" json:"userId,omitempy"`
	UsuarioRelacioID string             `bson:"usuariorelacionid" json:"userRelationId,omitempy"`
	Tweet            struct {
		Mensaje string    `bson:"mensaje" json:"mensaje,omitempy"`
		Fecha   time.Time `bson:"fecha" json:"fecha,omitempy"`
		ID      string    `bson:"_id" json:"_id,omitempy"`
	}
}
