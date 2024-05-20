package debtCollectorUseCase

import (
	"fp_pinjaman_online/model/dto/debtCollectorDto"
	"fp_pinjaman_online/src/debtCollector"
)

type debtCollectorUseCase struct {
	debtCollRepo debtCollector.DebtCollectorRepository
}

func NewDebtCollectorUseCase(debtCollRepo debtCollector.DebtCollectorRepository) debtCollector.DebtCollectorUseCase {
	return &debtCollectorUseCase{debtCollRepo}
}

func (usecase *debtCollectorUseCase) CreateLogTugas(newLogPayload debtCollectorDto.NewLogTugasPayload) error {
	_, err := usecase.debtCollRepo.SelectTugasById(newLogPayload.TugasId)
	if err != nil {
		return err
	}

	err = usecase.debtCollRepo.InsertLogTugas(newLogPayload)
	if err != nil {
		return err
	}
	return nil
}
