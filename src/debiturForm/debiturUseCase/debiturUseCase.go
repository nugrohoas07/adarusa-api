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

func (dbt *DebiturUseCase) GetDataByRole(role, status string, page, size int) ([]debiturFormDto.DetailDebitur, int, error) {
    offset := (page - 1) * size
    debitur, totalData, err := dbt.debiturRepository.GetDataByRole(role, status, size, offset)
    if err != nil {
        return nil, 0, err
    }
    return debitur, totalData, nil
}
