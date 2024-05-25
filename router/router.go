package router

import (
	"database/sql"

	"fp_pinjaman_online/src/admin/adminDelivery"
	"fp_pinjaman_online/src/admin/adminUsecase"
	adminRepository "fp_pinjaman_online/src/admin/repository"

	// "fp_pinjaman_online/src/checkHealth/checkHealthDelivery"
	// "fp_pinjaman_online/src/checkHealth/checkHealthRepository"
	// checkHealthUsecase "fp_pinjaman_online/src/checkHealth/checkHealthUseCase"

	"fp_pinjaman_online/src/debitur/debiturDelivery.go"
	"fp_pinjaman_online/src/debitur/debiturRepository"
	"fp_pinjaman_online/src/debitur/debiturUsecase"

	"fp_pinjaman_online/src/checkHealth/checkHealthDelivery"
	"fp_pinjaman_online/src/checkHealth/checkHealthRepository"
	checkHealthUsecase "fp_pinjaman_online/src/checkHealth/checkHealthUseCase"
	"fp_pinjaman_online/src/debtCollector/debtCollectorDelivery"
	"fp_pinjaman_online/src/debtCollector/debtCollectorRepository"
	"fp_pinjaman_online/src/debtCollector/debtCollectorUseCase"
	"fp_pinjaman_online/src/users/userDelivery"
	"fp_pinjaman_online/src/users/userRepository"
	"fp_pinjaman_online/src/users/userUseCase"

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
	checkHealthRepo := checkHealthRepository.NewCheckHealthRepository(db)
	checkHealthUC := checkHealthUsecase.NewCheckHealthUsecase(checkHealthRepo)
	checkHealthDelivery.NewCheckHealthDelivery(v1Group, checkHealthUC)

	adminRepo := adminRepository.NewAdminRepository(db)
	adminUC := adminUsecase.NewAdminUsecase(adminRepo)
	adminDelivery.NewAdminDelivery(v1Group, adminUC)

	debtCollectorRepo := debtCollectorRepository.NewDebtCollectorRepository(db)
	debtCollectorUC := debtCollectorUseCase.NewDebtCollectorUseCase(debtCollectorRepo)
	debtCollectorDelivery.NewDebtCollectorDelivery(v1Group, debtCollectorUC)

	userRepository := userRepository.NewUserRepository(db)
	userUC := userUseCase.NewUserUseCase(userRepository)
	userDelivery.NewUserDelivery(v1Group, userUC)

}
