package debiturForm

import "fp_pinjaman_online/model/debiturFormDto"

type DebiturRepository interface {
	CreateDetailDebitur(req debiturFormDto.Debitur) error
	UpdatePhotoPaths(userId int, fotoKTP, fotoSelfie string) error
	GetDataByRole(role, status string, limit, offset int) ([]debiturFormDto.DetailDebitur, int, error)
}

type DebiturUseCase interface {
	CreateDetailDebitur(req debiturFormDto.Debitur) error
	UpdatePhotoPaths(userId int, fotoKTP, fotoSelfie string) error
	GetDataByRole(role, status string, page, size int) ([]debiturFormDto.DetailDebitur, int, error)
}