package debtCollector

import (
	"fp_pinjaman_online/model/dto/debtCollectorDto"
	"fp_pinjaman_online/model/entity/debtCollectorEntity"
)

type DebtCollectorUseCase interface {
	CreateLogTugas(newLogPayload debtCollectorDto.NewLogTugasPayload) error
}

type DebtCollectorRepository interface {
	SelectTugasById(tugasId string) (debtCollectorEntity.Tugas, error)
	InsertLogTugas(newLogPayload debtCollectorDto.NewLogTugasPayload) error
}
