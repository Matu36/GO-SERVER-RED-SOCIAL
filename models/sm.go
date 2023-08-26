package models

type Secret struct {
	Host     string `json:"host"`
	Username string `string:"username"`
	Password string `string:"password"`
	JWTSign  string `string:"jwtsign"`
	Database string `string:"database"`
}
