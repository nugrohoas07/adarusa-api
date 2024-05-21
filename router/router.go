package router

import (
	"database/sql"
	// "fp_pinjaman_online/src/checkHealth/checkHealthDelivery"
	// "fp_pinjaman_online/src/checkHealth/checkHealthRepository"
	// checkHealthUsecase "fp_pinjaman_online/src/checkHealth/checkHealthUseCase"

	"fp_pinjaman_online/src/debitur/debiturDelivery.go"
	"fp_pinjaman_online/src/debitur/debiturRepository"
	"fp_pinjaman_online/src/debitur/debiturUsecase"

	"github.com/gin-gonic/gin"
)

func InitRoute(v1Group *gin.RouterGroup, db *sql.DB) {
	// checkHealthRepo := checkHealthRepository.NewCheckHealthRepository(db)
	// checkHealthUC := checkHealthUsecase.NewCheckHealthUsecase(checkHealthRepo)
	// checkHealthDelivery.NewCheckHealthDelivery(v1Group, checkHealthUC)

	//debitur
	debiturRepository := debiturRepository.NewDebiturRepository(db)
	debiturUC := debiturUsecase.NewDebiturUsecase(debiturRepository)
	debiturDelivery.NewDebiturDelivery(v1Group, debiturUC)
}
