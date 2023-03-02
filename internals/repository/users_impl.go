package repository

import (
	"context"
	"database/sql"
	"errors"
	"fanchann/Go-APIWithArchitecture2/internals/models/domain"
	"fanchann/Go-APIWithArchitecture2/utils"
)

type UsersRepositoryImpl struct{}

func NewUsersRepositoryImpl() IUsersRepository {
	return &UsersRepositoryImpl{}
}

func (repo *UsersRepositoryImpl) CreateUser(ctx context.Context, tx *sql.Tx, form *domain.Users) (domain.Users, error) {
	createQuery := `insert into users (name,email) values (?,?)`
	result, err := tx.ExecContext(ctx, createQuery, form.Name, form.Email)
	if err != nil {
		return domain.Users{}, err
	}

	id, err := result.LastInsertId()
	utils.ErrorWithPanic(err)

	form.Id = int(id)

	return *form, nil
}

func (repo *UsersRepositoryImpl) UpdateUser(ctx context.Context, tx *sql.Tx, form *domain.Users) domain.Users {
	updateQuery := `update users set `
	var newUpdateQuery []interface{}

	if form.Name != "" {
		updateQuery += `name = ?, `
		newUpdateQuery = append(newUpdateQuery, form.Name)
	}

	if form.Email != "" {
		updateQuery += `email = ?, `
		newUpdateQuery = append(newUpdateQuery, form.Email)
	}

	updateQuery = updateQuery[:len(updateQuery)-2]
	updateQuery += `where id = ?`
	newUpdateQuery = append(newUpdateQuery, form.Id)

	result, err := tx.ExecContext(ctx, updateQuery, newUpdateQuery...)
	utils.ErrorWithPanic(err)

	idAfterUpdate, err := result.RowsAffected()
	utils.ErrorWithPanic(err)

	form.Id = int(idAfterUpdate)

	return *form
}

func (repo *UsersRepositoryImpl) DeleteUser(ctx context.Context, tx *sql.Tx, form *domain.Users) {
	deleteQuery := `delete from users where id = ?`
	_, err := tx.ExecContext(ctx, deleteQuery, form.Id)
	utils.ErrorWithPanic(err)
}

func (repo *UsersRepositoryImpl) FindAllUsers(ctx context.Context, tx *sql.Tx) []domain.ResponseFromDomain {
	var usersData []domain.ResponseFromDomain
	getAllQuery := `select id,name,register_at from users`

	rows, err := tx.QueryContext(ctx, getAllQuery)
	utils.ErrorWithPanic(err)
	defer rows.Close()

	for rows.Next() {
		var userData domain.ResponseFromDomain

		err := rows.Scan(&userData.Id, &userData.Name, &userData.Created_At)
		utils.ErrorWithPanic(err)

		usersData = append(usersData, userData)
	}

	return usersData
}

func (repo *UsersRepositoryImpl) FindUserById(ctx context.Context, tx *sql.Tx, form *domain.Users) (domain.Users, error) {
	var userData domain.Users

	findUserWithIdQuery := `select  id,name,register_at from users where id = ?`
	rows, err := tx.QueryContext(ctx, findUserWithIdQuery, form.Id)
	utils.ErrorWithPanic(err)
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&userData.Id, &userData.Name, &userData.Created_At)
		utils.ErrorWithPanic(err)
		return userData, nil
	} else {
		return userData, errors.New("user not found!")
	}
}
