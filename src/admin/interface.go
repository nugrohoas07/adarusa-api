package adminInterface

import (
	"fp_pinjaman_online/model/dto/adminDto"
	"fp_pinjaman_online/model/entity"
)

type AdminRepository interface {
	RetrieveUserStatusById(id int) (*entity.UserCompleteInfo, error)
	UpdateUserStatus(id int, status string) error
}

type AdminUsecase interface {
	VerifyAndUpdateUser(req adminDto.RequestUpdateStatusUser) (adminDto.AdminResponse, error)
}
