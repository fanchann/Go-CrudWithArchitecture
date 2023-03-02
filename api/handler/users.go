package handler

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type IUsersHandler interface {
	CreateUser(writer http.ResponseWriter, request *http.Request, _ httprouter.Params)
	UpdateUser(writer http.ResponseWriter, request *http.Request, userId httprouter.Params)
	DeleteUser(writer http.ResponseWriter, request *http.Request, userId httprouter.Params)
	FindUserById(writer http.ResponseWriter, request *http.Request, userId httprouter.Params)
	FindAllUser(writer http.ResponseWriter, request *http.Request, _ httprouter.Params)
}
