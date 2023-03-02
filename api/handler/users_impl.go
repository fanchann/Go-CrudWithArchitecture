package handler

import (
	"fanchann/Go-APIWithArchitecture2/internals/models/web"
	"fanchann/Go-APIWithArchitecture2/pkg/service"
	"fanchann/Go-APIWithArchitecture2/utils"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type UsersHandler struct {
	Service service.IUsersService
}

func NewUsersHandler(service service.IUsersService) IUsersHandler {
	return &UsersHandler{Service: service}
}

func (handler *UsersHandler) CreateUser(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	createNewUserFrom := web.UsersRegister{}
	utils.ReadRequestFromBody(request, createNewUserFrom)

	response := handler.Service.CreateUser(request.Context(), createNewUserFrom)

	utils.WriteToResponse(writer, response)

}

func (handler *UsersHandler) UpdateUser(writer http.ResponseWriter, request *http.Request, userId httprouter.Params) {
	updateForm := web.UsersEdit{}
	utils.ReadRequestFromBody(request, updateForm)

	response := handler.Service.UpdateUser(request.Context(), updateForm)

	utils.WriteToResponse(writer, response)
}

func (handler *UsersHandler) DeleteUser(writer http.ResponseWriter, request *http.Request, userId httprouter.Params) {

	id := userId.ByName("userId")
	idToInt, err := strconv.Atoi(id)
	utils.ErrorWithPanic(err)

	response := handler.Service.DeleteUser(request.Context(), idToInt)

	utils.WriteToResponse(writer, response)
}

func (handler *UsersHandler) FindUserById(writer http.ResponseWriter, request *http.Request, userId httprouter.Params) {

	id := userId.ByName("userId")
	idToInt, err := strconv.Atoi(id)
	utils.ErrorWithPanic(err)

	response := handler.Service.FindUserById(request.Context(), idToInt)

	utils.WriteToResponse(writer, response)
}

func (handler *UsersHandler) FindAllUser(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {

	response := handler.Service.FindAllUsers(request.Context())

	utils.WriteToResponse(writer, response)
}
