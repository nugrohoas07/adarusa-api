package users

import "fp_pinjaman_online/model/userDto"

type UserRepository interface {
	CreateUser(req userDto.CreateRequest) error
	Login(req userDto.LoginRequest) (string, error)
	GetUserById(id string) (userDto.User, error)
	GetUserCount(email, password string) (int, error)
	UpdateUser(id string, req userDto.Update) error
	UserExists(email string) (bool, error)
	UserExistById(id string) (bool, error)
	GetUserByEmail(email string) (userDto.User, error)
}

type UserUseCase interface {
	CreateUser(req userDto.CreateRequest) error
	Login(req userDto.LoginRequest) (string, error)
	UpdateUser(id string, req userDto.Update) error
}