package service

import (
	"context"
	"fanchann/Go-APIWithArchitecture2/internals/models/web"
)

type IUsersService interface {
	CreateUser(ctx context.Context, form web.UsersRegister) web.WebResponse
	UpdateUser(ctx context.Context, form web.UsersEdit) web.WebResponse
	DeleteUser(ctx context.Context, userId int) web.WebResponse
	FindUserById(ctx context.Context, userId int) web.WebResponse
	FindAllUsers(ctx context.Context) web.WebResponse
}
