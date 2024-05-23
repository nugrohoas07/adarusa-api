package debiturUseCase

import (
	"fp_pinjaman_online/model/debiturFormDto"
	"fp_pinjaman_online/src/debiturForm"
)

type DebiturUseCase struct {
	debiturRepository debiturForm.DebiturRepository
}

func NewDebiturUseCase(repo debiturForm.DebiturRepository) *DebiturUseCase {
	return &DebiturUseCase{debiturRepository: repo}
}

func (dbt *DebiturUseCase) CreateDetailDebitur(debitur debiturFormDto.Debitur) error {
	return dbt.debiturRepository.CreateDetailDebitur(debitur)
}

func (dbt *DebiturUseCase) UpdatePhotoPaths(userId int, fotoKTP, fotoSelfie string) error {
	return dbt.debiturRepository.UpdatePhotoPaths(userId, fotoKTP, fotoSelfie)
}