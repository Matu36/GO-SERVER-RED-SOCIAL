package models

import "time"

type GraboTweet struct {
	UserID  string    `bson:"userid" json:"userid,omitempy"`
	Mensaje string    `bson:"mensaje" json:"mensaje,omitempy"`
	Fecha   time.Time `bson:"fecha" json:"fecha,omitempy"`
}
