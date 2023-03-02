package service

import (
	"context"
	"database/sql"
	"fanchann/Go-APIWithArchitecture2/internals/models/domain"
	"fanchann/Go-APIWithArchitecture2/internals/models/web"
	"fanchann/Go-APIWithArchitecture2/internals/repository"
	"fanchann/Go-APIWithArchitecture2/utils"
	"time"
)

type UsersServiceImpl struct {
	DB   *sql.DB
	repo repository.IUsersRepository
}

func NewUsersServiceImpl(db *sql.DB, repo repository.IUsersRepository) IUsersService {
	return &UsersServiceImpl{
		DB:   db,
		repo: repo,
	}
}

func (service *UsersServiceImpl) CreateUser(ctx context.Context, form web.UsersRegister) web.WebResponse {
	tx, err := service.DB.Begin()
	utils.ErrorWithPanic(err)

	defer utils.TransactionCommitAndRollBack(tx)

	checkEmailIsValid, err := utils.EmailValidation(form.Email)
	if err != nil {

		return web.WebResponse{
			Status:  400,
			Message: "Email not valid",
			Data:    nil,
		}

	} else {

		domain := domain.Users{
			Name:       form.Name,
			Email:      checkEmailIsValid,
			Created_At: time.Now().String(),
		}

		responseToUser, err := service.repo.CreateUser(ctx, tx, &domain)

		if err != nil {
			return web.WebResponse{
				Status:  400,
				Message: "Email was registered",
				Data:    nil,
			}
		}

		return web.WebResponse{
			Status:  200,
			Message: "Success registered",
			Data:    responseToUser,
		}

	}
}

func (service *UsersServiceImpl) UpdateUser(ctx context.Context, form web.UsersEdit) web.WebResponse {
	tx, err := service.DB.Begin()
	utils.ErrorWithPanic(err)

	defer utils.TransactionCommitAndRollBack(tx)

	var originalStructure domain.Users

	originalStructure.Id = form.Id

	foundUser, err := service.repo.FindUserById(ctx, tx, &originalStructure)
	if err != nil {
		utils.ErrorWithPanic(err)
		return web.WebResponse{
			Status:  400,
			Message: "Data user not found",
			Data:    nil,
		}
	} else {
		updatedUser := service.repo.UpdateUser(ctx, tx, &foundUser)

		form.Id = updatedUser.Id
		form.Name = updatedUser.Name
		form.Email = updatedUser.Email
		form.Update_At = time.Now()

		// _ := strconv.Itoa(form.Id)

		return web.WebResponse{
			Status:  200,
			Message: "Success edit data",
			Data:    updatedUser,
		}
	}
}

func (service *UsersServiceImpl) DeleteUser(ctx context.Context, userId int) web.WebResponse {
	tx, err := service.DB.Begin()
	utils.ErrorWithPanic(err)
	defer utils.TransactionCommitAndRollBack(tx)

	var dataStructure domain.Users
	dataStructure.Id = userId

	foundUser, err := service.repo.FindUserById(ctx, tx, &dataStructure)
	if err != nil {
		utils.ErrorWithPanic(err)

		return web.WebResponse{
			Status:  400,
			Message: "Data user not found",
			Data:    nil,
		}

	} else {

		service.repo.DeleteUser(ctx, tx, &foundUser)
		return web.WebResponse{
			Status:  200,
			Message: "Succes delete",
			Data:    nil,
		}

	}
}

func (service *UsersServiceImpl) FindUserById(ctx context.Context, userId int) web.WebResponse {
	tx, err := service.DB.Begin()
	utils.ErrorWithPanic(err)

	var dataStructure domain.Users
	dataStructure.Id = userId

	foundUser, err := service.repo.FindUserById(ctx, tx, &dataStructure)
	cvToDomainResponse := domain.ResponseFromDomain{
		Id:         foundUser.Id,
		Name:       foundUser.Name,
		Created_At: foundUser.Created_At,
	}

	if err != nil {
		return web.WebResponse{
			Status:  400,
			Message: "Data user not found!",
			Data:    nil,
		}
	} else {
		return web.WebResponse{
			Status:  200,
			Message: "Data user found",
			Data:    cvToDomainResponse,
		}
	}
}

func (service *UsersServiceImpl) FindAllUsers(ctx context.Context) web.WebResponse {
	tx, err := service.DB.Begin()
	utils.ErrorWithPanic(err)

	users := service.repo.FindAllUsers(ctx, tx)

	return web.WebResponse{
		Status:  200,
		Message: "success get data",
		Data:    users,
	}
}
