package debtCollectorUseCase

import (
	"fp_pinjaman_online/model/dto/debtCollectorDto"
	"fp_pinjaman_online/model/entity/debtCollectorEntity"
	"fp_pinjaman_online/src/debtCollector"
)

type debtCollectorUseCase struct {
	debtCollRepo debtCollector.DebtCollectorRepository
}

func NewDebtCollectorUseCase(debtCollRepo debtCollector.DebtCollectorRepository) debtCollector.DebtCollectorUseCase {
	return &debtCollectorUseCase{debtCollRepo}
}

// TODO
// Checking log tugas authorization by their user id
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

func (usecase *debtCollectorUseCase) GetLogTugasById(logTugasId string) (debtCollectorEntity.LogTugas, error) {
	log, err := usecase.debtCollRepo.SelectLogTugasById(logTugasId)
	if err != nil {
		return debtCollectorEntity.LogTugas{}, err
	}
	return log, nil
}

func (usecase *debtCollectorUseCase) EditLogTugasById(logTugasId string, payload debtCollectorDto.UpdateLogTugasPayload) error {
	storedLog, err := usecase.debtCollRepo.SelectLogTugasById(logTugasId)
	if err != nil {
		return err
	}

	err = usecase.debtCollRepo.UpdateLogTugasById(storedLog, payload)
	if err != nil {
		return err
	}
	return nil
}
