package web

import "time"

type WebResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ServiceResponse struct {
	Id         string    `json:"id"`
	Name       string    `json:"name"`
	Created_At time.Time `json:"created_at"`
}

type ServiceUpdateResponse struct {
	Id         string    `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Updated_At time.Time `json:"updated_at"`
}
