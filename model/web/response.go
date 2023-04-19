package web

import "mygram/model/entity"

// Response for swagger
type Response struct {
	Message string `json:"message"`
}

type ResponseLogin struct {
	Token string      `json:"token"`
	Data  entity.User `json:"data"`
}
