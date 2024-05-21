package debtCollector

import (
	"fp_pinjaman_online/model/dto/debtCollectorDto"
	"fp_pinjaman_online/model/entity/debtCollectorEntity"
)

type DebtCollectorUseCase interface {
	CreateLogTugas(newLogPayload debtCollectorDto.NewLogTugasPayload) error
	GetLogTugasById(logTugasId string) (debtCollectorEntity.LogTugas, error)
	EditLogTugasById(logTugasId string, payload debtCollectorDto.UpdateLogTugasPayload) error
	LogTugasAuthorizationCheck(logTugasId string) (debtCollectorEntity.LogTugas, error)
}

type DebtCollectorRepository interface {
	SelectTugasById(tugasId string) (debtCollectorEntity.Tugas, error)
	InsertLogTugas(newLogPayload debtCollectorDto.NewLogTugasPayload) error
	UpdateLogTugasById(storedLog debtCollectorEntity.LogTugas, updateLogPayload debtCollectorDto.UpdateLogTugasPayload) error
	SelectLogTugasById(logTugasId string) (debtCollectorEntity.LogTugas, error)
	SoftDeleteLogTugasById(logTugasId string) error
}
