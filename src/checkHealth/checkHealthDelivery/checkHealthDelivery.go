package checkHealthDelivery

import (
	"fp_pinjaman_online/model/dto/checkHealthDto"
	"fp_pinjaman_online/model/dto/json"
	"fp_pinjaman_online/pkg/validation"
	"fp_pinjaman_online/src/checkHealth"

	"github.com/gin-gonic/gin"
)

type checkHealthDelivery struct {
	checkHealthUC checkHealth.CheckHealthUseCase
}

func NewCheckHealthDelivery(v1Group *gin.RouterGroup, checkHealthUC checkHealth.CheckHealthUseCase) {
	handler := checkHealthDelivery{
		checkHealthUC: checkHealthUC,
	}
	checkHealthGroup := v1Group.Group("/checkHealth")
	{
		checkHealthGroup.GET("/version", handler.GetVersion)
		checkHealthGroup.POST("/version/save", handler.SaveVersion)
	}
}

func (c *checkHealthDelivery) GetVersion(ctx *gin.Context) {
	version, err := c.checkHealthUC.GetVersion()
	if err != nil {
		json.NewResponseError(ctx, err.Error(), "01", "01")
		return
	}
	json.NewResponseSuccess(ctx, version, "success", "01", "01")
}

func (c *checkHealthDelivery) SaveVersion(ctx *gin.Context) {
	var req checkHealthDto.VersionRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		validationError := validation.GetValidationError(err)
		if len(validationError) > 0 {
			json.NewResponseBadRequestValidator(ctx, validationError, "bad request", "01", "02")
			return
		}
	}

	json.NewResponseSuccess(ctx, req.Version, "success", "01", "01")
}
