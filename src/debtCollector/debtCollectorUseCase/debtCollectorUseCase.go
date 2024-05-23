package debtCollectorUseCase

import (
	"fp_pinjaman_online/model/dto/debtCollectorDto"
	"fp_pinjaman_online/model/dto/json"
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
func (usecase *debtCollectorUseCase) LogTugasAuthorizationCheck(logTugasId string) (debtCollectorEntity.LogTugas, error) {
	// check if log tugas with input id exist
	log, err := usecase.debtCollRepo.SelectLogTugasById(logTugasId)
	if err != nil {
		return debtCollectorEntity.LogTugas{}, err
	}
	// authorization here

	return log, nil
}

func (usecase *debtCollectorUseCase) GetAllLateDebtorByCity(dcId string, page, size int) ([]debtCollectorEntity.LateDebtor, json.Paging, error) {
	loggedDc, err := usecase.debtCollRepo.SelectDebtCollectorById(dcId)
	if err != nil {
		return nil, json.Paging{}, err
	}

	lateDebtorsList, paging, err := usecase.debtCollRepo.SelectAllLateDebitur(loggedDc.City, page, size)
	if err != nil {
		return nil, json.Paging{}, err
	}
	return lateDebtorsList, paging, nil
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

func (usecase *debtCollectorUseCase) GetLogTugasById(logTugasId string) (debtCollectorEntity.LogTugas, error) {
	log, err := usecase.LogTugasAuthorizationCheck(logTugasId)
	if err != nil {
		return debtCollectorEntity.LogTugas{}, err
	}
	return log, nil
}

func (usecase *debtCollectorUseCase) EditLogTugasById(logTugasId string, payload debtCollectorDto.UpdateLogTugasPayload) error {
	storedLog, err := usecase.LogTugasAuthorizationCheck(logTugasId)
	if err != nil {
		return err
	}

	err = usecase.debtCollRepo.UpdateLogTugasById(storedLog, payload)
	if err != nil {
		return err
	}
	return nil
}

func (usecase *debtCollectorUseCase) DeleteLogTugasById(logTugasId string) error {
	_, err := usecase.LogTugasAuthorizationCheck(logTugasId)
	if err != nil {
		return err
	}

	err = usecase.debtCollRepo.SoftDeleteLogTugasById(logTugasId)
	if err != nil {
		return err
	}

	return nil
}

func (usecase *debtCollectorUseCase) GetAllLogTugas(tugasId string, page, size int) ([]debtCollectorEntity.LogTugas, json.Paging, error) {
	logsList, paging, err := usecase.debtCollRepo.SelectAllLogByTugasId(tugasId, page, size)
	if err != nil {
		return nil, json.Paging{}, err
	}
	return logsList, paging, nil
}

func (usecase *debtCollectorUseCase) ClaimTugas(dcId string, payload debtCollectorDto.NewTugasPayload) error {
	// mengambil kota dc
	loggedDc, err := usecase.debtCollRepo.SelectDebtCollectorById(dcId)
	if err != nil {
		return err
	}

	_, err = usecase.debtCollRepo.SelectLateDebiturById(payload.UserId, loggedDc.City)
	if err != nil {
		return err
	}

	err = usecase.debtCollRepo.CreateClaimTugas(dcId, payload.UserId)
	if err != nil {
		return err
	}

	return nil
}
