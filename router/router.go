package router

import (
	"database/sql"
	"fp_pinjaman_online/src/checkHealth/checkHealthDelivery"
	"fp_pinjaman_online/src/checkHealth/checkHealthRepository"
	checkHealthUsecase "fp_pinjaman_online/src/checkHealth/checkHealthUseCase"
	"fp_pinjaman_online/src/debtCollector/debtCollectorDelivery"
	"fp_pinjaman_online/src/debtCollector/debtCollectorRepository"
	"fp_pinjaman_online/src/debtCollector/debtCollectorUseCase"

	"github.com/gin-gonic/gin"
)

func InitRoute(v1Group *gin.RouterGroup, db *sql.DB) {
	checkHealthRepo := checkHealthRepository.NewCheckHealthRepository(db)
	checkHealthUC := checkHealthUsecase.NewCheckHealthUsecase(checkHealthRepo)
	checkHealthDelivery.NewCheckHealthDelivery(v1Group, checkHealthUC)

	debtCollectorRepo := debtCollectorRepository.NewDebtCollectorRepository(db)
	debtCollectorUC := debtCollectorUseCase.NewDebtCollectorUseCase(debtCollectorRepo)
	debtCollectorDelivery.NewDebtCollectorDelivery(v1Group, debtCollectorUC)
}
