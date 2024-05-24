package router

import (
	"database/sql"
	"fp_pinjaman_online/src/checkHealth/checkHealthDelivery"
	"fp_pinjaman_online/src/checkHealth/checkHealthRepository"
	checkHealthUsecase "fp_pinjaman_online/src/checkHealth/checkHealthUseCase"
	"fp_pinjaman_online/src/users/userDelivery"
	"fp_pinjaman_online/src/users/userRepository"
	"fp_pinjaman_online/src/users/userUseCase"

	"github.com/gin-gonic/gin"
)

func InitRoute(v1Group *gin.RouterGroup, db *sql.DB) {
	checkHealthRepo := checkHealthRepository.NewCheckHealthRepository(db)
	checkHealthUC := checkHealthUsecase.NewCheckHealthUsecase(checkHealthRepo)
	checkHealthDelivery.NewCheckHealthDelivery(v1Group, checkHealthUC)

	userRepository := userRepository.NewUserRepository(db)
	userUC := userUseCase.NewUserUseCase(userRepository)
	userDelivery.NewUserDelivery(v1Group, userUC)
}
