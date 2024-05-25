package users

import (
	"fp_pinjaman_online/model/dcFormDto"
	"fp_pinjaman_online/model/debiturFormDto"
	"fp_pinjaman_online/model/userDto"
)

type UserRepository interface {
	CreateUser(req userDto.CreateRequest, roleId int) error
	Login(req userDto.LoginRequest) (string, error)
	UserExists(email string) (bool, error)
	GetUserByEmail(email string) (userDto.User, error)
	CreateDetailDebitur(req debiturFormDto.Debitur) error
	CreateDetailDc(req dcFormDto.DetailDC) error
	UpdatePhotoPaths(userId int, fotoKTP, fotoSelfie string) error
	GetDataByRole(role, status string, limit, offset int) ([]debiturFormDto.DetailDebitur, int, error)
	GetFullname(userId int) (string, error)
}

type UserUseCase interface {
	CreateUser(req userDto.CreateRequest, roleId int) error
	Login(req userDto.LoginRequest) (string, error)
	GetUserByEmail(email string) (userDto.User, error)
	CreateDetailDebitur(req debiturFormDto.Debitur) error
	CreateDetailDc(req dcFormDto.DetailDC) error
	UpdatePhotoPaths(userId int, fotoKTP, fotoSelfie string) error
	GetDataByRole(role, status string, page, size int) ([]debiturFormDto.DetailDebitur, int, error)
	GetFullname(userId int) (string, error)
}