package debtCollectorRepository

import (
	"fp_pinjaman_online/model/dto/debtCollectorDto"
	"fp_pinjaman_online/model/dto/json"
	"fp_pinjaman_online/model/entity/debtCollectorEntity"

	"github.com/stretchr/testify/mock"
)

type DebtCollectorRepositoryMock struct {
	Mock mock.Mock
}

func (dm *DebtCollectorRepositoryMock) SelectTugasById(tugasId string) (debtCollectorEntity.Tugas, error) {
	args := dm.Mock.Called(tugasId)
	if args.Get(1) != nil {
		return debtCollectorEntity.Tugas{}, args.Error(1)
	}
	return args.Get(0).(debtCollectorEntity.Tugas), nil
}

func (dm *DebtCollectorRepositoryMock) InsertLogTugas(newLogPayload debtCollectorDto.NewLogTugasPayload) error {
	args := dm.Mock.Called(newLogPayload)
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return nil
}

func (dm *DebtCollectorRepositoryMock) UpdateLogTugasById(storedLog debtCollectorEntity.LogTugas, updateLogPayload debtCollectorDto.UpdateLogTugasPayload) error {
	args := dm.Mock.Called(storedLog, updateLogPayload)
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return nil
}

func (dm *DebtCollectorRepositoryMock) SelectLogTugasById(logTugasId string) (debtCollectorEntity.LogTugas, error) {
	args := dm.Mock.Called(logTugasId)
	if args.Get(1) != nil {
		return debtCollectorEntity.LogTugas{}, args.Error(1)
	}
	return args.Get(0).(debtCollectorEntity.LogTugas), nil
}

func (dm *DebtCollectorRepositoryMock) SoftDeleteLogTugasById(logTugasId string) error {
	args := dm.Mock.Called(logTugasId)
	if args.Get(0) != nil {
		return args.Error(0)
	}
	return nil
}

func (dm *DebtCollectorRepositoryMock) SelectAllLogByTugasId(tugasId string, page, size int) ([]debtCollectorEntity.LogTugas, json.Paging, error) {
	args := dm.Mock.Called(tugasId, page, size)
	if args.Get(2) != nil {
		return []debtCollectorEntity.LogTugas{}, json.Paging{}, args.Error(2)
	}
	return args.Get(0).([]debtCollectorEntity.LogTugas), args.Get(1).(json.Paging), nil
}

func (dm *DebtCollectorRepositoryMock) SelectAllLateDebitur(dcCity string, page, size int) ([]debtCollectorEntity.LateDebtor, json.Paging, error) {
	args := dm.Mock.Called(dcCity, page, size)
	if args.Get(2) != nil {
		return nil, json.Paging{}, args.Error(2)
	}
	return args.Get(0).([]debtCollectorEntity.LateDebtor), args.Get(1).(json.Paging), nil
}

func (dm *DebtCollectorRepositoryMock) SelectDebtCollectorById(id string) (debtCollectorEntity.DebtCollector, error) {
	args := dm.Mock.Called(id)
	if args.Get(1) != nil {
		return debtCollectorEntity.DebtCollector{}, args.Error(1)
	}
	return args.Get(0).(debtCollectorEntity.DebtCollector), nil
}
