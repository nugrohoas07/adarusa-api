package users

import "fp_pinjaman_online/model/userDto"

type UserRepository interface {
	CreateUser(req userDto.CreateRequest, roleId int) error
	Login(req userDto.LoginRequest) (string, error)
	UserExists(email string) (bool, error)
	GetUserByEmail(email string) (userDto.User, error)
}

type UserUseCase interface {
	CreateUser(req userDto.CreateRequest, roleId int) error
	Login(req userDto.LoginRequest) (string, error)
	GetUserByEmail(email string) (userDto.User, error)
}