package userUseCase

import (
	"errors"
	"fp_pinjaman_online/model/userDto"
	"fp_pinjaman_online/pkg/middleware"
	"fp_pinjaman_online/pkg/validation"
	"fp_pinjaman_online/src/users"
)

type userUC struct {
	userRepo users.UserRepository
}

func NewUserUseCase(userRepo users.UserRepository) users.UserUseCase {
	return &userUC{userRepo}
}

func (useCase *userUC) CreateUser(req userDto.CreateRequest, roleId int) error {
	exists, err := useCase.userRepo.UserExists(req.Email)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("user already exist")
	}

	req.Password, err = validation.HashedPassword(req.Password)
	if err != nil {
		return err
	}

    return useCase.userRepo.CreateUser(req, roleId)
}

func (useCase *userUC) Login(req userDto.LoginRequest) (string, error) {
	hashedPassword, err := useCase.userRepo.Login(req)
	if err != nil {
		return "", nil
	}
	if hashedPassword == "" {
		return "", errors.New("invalid email or password")
	}

	// Get user by their email
	user, err := useCase.userRepo.GetUserByEmail(req.Email)
	if err != nil {
		return "", err
	}

	token, err := middleware.GenerateTokenJwt(user.Id, user.Email, user.Roles, 1)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (useCase *userUC) GetUserByEmail(email string) (userDto.User, error) {
	return useCase.userRepo.GetUserByEmail(email)
}