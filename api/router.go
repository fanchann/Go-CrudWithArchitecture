package api

import (
	"fanchann/Go-APIWithArchitecture2/api/handler"
	"fanchann/Go-APIWithArchitecture2/internals/repository"
	"fanchann/Go-APIWithArchitecture2/pkg/db"
	"fanchann/Go-APIWithArchitecture2/pkg/service"

	"github.com/julienschmidt/httprouter"
)

func NewRouter() httprouter.Router {
	router := httprouter.New()
	db := db.MysqlConnect()
	repo := repository.NewUsersRepositoryImpl()
	service := service.NewUsersServiceImpl(db, repo)
	handler := handler.NewUsersHandler(service)

	router.POST("/register", handler.CreateUser)
	router.GET("/users", handler.FindAllUser)
	router.PUT("/user/:userId", handler.UpdateUser)
	router.GET("/user/:userId", handler.FindUserById)
	router.DELETE("/user/:userId", handler.DeleteUser)

	return *router
}
