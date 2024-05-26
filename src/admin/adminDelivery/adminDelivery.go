package adminDelivery

import (
	"fp_pinjaman_online/model/dto/adminDto"
	"fp_pinjaman_online/model/dto/json"
	adminInterface "fp_pinjaman_online/src/admin"
	"strconv"

	"github.com/gin-gonic/gin"
)

type adminDelivery struct {
	adminUc adminInterface.AdminUsecase
}

func NewAdminDelivery(v1Group *gin.RouterGroup, adminUc adminInterface.AdminUsecase) {
	handler := adminDelivery{
		adminUc: adminUc,
	}
	userGroup := v1Group.Group("/users")
	userGroup.PATCH("/:id/verify", handler.VerifyAndUpdateUser)
}

func (a *adminDelivery) VerifyAndUpdateUser(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		json.NewResponseBadRequest(ctx, err.Error(), "01", "02")
		return
	}

	var req adminDto.RequestUpdateStatusUser
	if err := ctx.ShouldBindJSON(&req); err != nil {
		json.NewResponseBadRequest(ctx, err.Error(), "01", "02")
		return
	}

	req.ID = id

	res, err := a.adminUc.VerifyAndUpdateUser(req)
	if err != nil {
		json.NewResponseError(ctx, err.Error(), "01", "01")
		return
	}

	json.NewResponseSuccess(ctx, res, "User status updated successfully", "01", "02")
}
