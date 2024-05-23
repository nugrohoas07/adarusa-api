package debiturForm

import "fp_pinjaman_online/model/debiturFormDto"

type DebiturRepository interface {
	CreateDetailDebitur(req debiturFormDto.Debitur) error
	UpdatePhotoPaths(userId int, fotoKTP, fotoSelfie string) error
}

type DebiturUseCase interface {
	CreateDetailDebitur(req debiturFormDto.Debitur) error
	UpdatePhotoPaths(userId int, fotoKTP, fotoSelfie string) error
}