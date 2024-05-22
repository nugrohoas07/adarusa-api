package debtCollector

import (
	"fp_pinjaman_online/model/dto/debtCollectorDto"
	"fp_pinjaman_online/model/dto/json"
	"fp_pinjaman_online/model/entity/debtCollectorEntity"
)

type DebtCollectorUseCase interface {
	CreateLogTugas(newLogPayload debtCollectorDto.NewLogTugasPayload) error
	GetLogTugasById(logTugasId string) (debtCollectorEntity.LogTugas, error)
	EditLogTugasById(logTugasId string, payload debtCollectorDto.UpdateLogTugasPayload) error
	LogTugasAuthorizationCheck(logTugasId string) (debtCollectorEntity.LogTugas, error)
	DeleteLogTugasById(logTugasId string) error
	GetAllLogTugas(tugasId string, page, size int) ([]debtCollectorEntity.LogTugas, json.Paging, error)
	GetAllLateDebtorByCity(dcId string, page, size int) ([]debtCollectorEntity.LateDebtor, json.Paging, error)
}

type DebtCollectorRepository interface {
	SelectTugasById(tugasId string) (debtCollectorEntity.Tugas, error)
	InsertLogTugas(newLogPayload debtCollectorDto.NewLogTugasPayload) error
	UpdateLogTugasById(storedLog debtCollectorEntity.LogTugas, updateLogPayload debtCollectorDto.UpdateLogTugasPayload) error
	SelectLogTugasById(logTugasId string) (debtCollectorEntity.LogTugas, error)
	SoftDeleteLogTugasById(logTugasId string) error
	SelectAllLogByTugasId(tugasId string, page, size int) ([]debtCollectorEntity.LogTugas, json.Paging, error)
	SelectAllLateDebitur(dcCity string, page, size int) ([]debtCollectorEntity.LateDebtor, json.Paging, error)
	SelectDebtCollectorById(id string) (debtCollectorEntity.DebtCollector, error) // TODO : it should be in users repository
}
