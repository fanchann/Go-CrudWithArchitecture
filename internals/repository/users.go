package repository

import (
	"context"
	"database/sql"
	"fanchann/Go-APIWithArchitecture2/internals/models/domain"
)

type IUsersRepository interface {
	CreateUser(ctx context.Context, tx *sql.Tx, form *domain.Users) (domain.Users, error)
	UpdateUser(ctx context.Context, tx *sql.Tx, form *domain.Users) domain.Users
	DeleteUser(ctx context.Context, tx *sql.Tx, form *domain.Users)
	FindAllUsers(ctx context.Context, tx *sql.Tx) []domain.ResponseFromDomain
	FindUserById(ctx context.Context, tx *sql.Tx, form *domain.Users) (domain.Users, error)
}
