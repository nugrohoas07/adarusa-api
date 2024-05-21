package debtCollectorDelivery

import (
	"fp_pinjaman_online/model/dto/debtCollectorDto"
	"fp_pinjaman_online/model/dto/json"
	"fp_pinjaman_online/pkg/validation"
	"fp_pinjaman_online/src/debtCollector"
	"strings"

	"github.com/gin-gonic/gin"
)

type debtCollectorDelivery struct {
	debtCollUC debtCollector.DebtCollectorUseCase
}

func NewDebtCollectorDelivery(v1Group *gin.RouterGroup, debtCollUC debtCollector.DebtCollectorUseCase) {
	handler := debtCollectorDelivery{
		debtCollUC: debtCollUC,
	}
	dcGroup := v1Group.Group("/debt-collector")
	{
		dcGroup.POST("/tugas")                                 // claim tugas ?
		dcGroup.POST("/log-tugas/create", handler.AddLogTugas) // membuat log tugas baru
		dcGroup.GET("/log-tugas")                              // get all log
		dcGroup.GET("/log-tugas/:id", handler.GetLogTugas)     // get log detail
		dcGroup.PUT("/log-tugas/:id")                          // edit log
		dcGroup.DELETE("/log-tugas/:id")                       // hapus log
	}
}

func (d *debtCollectorDelivery) AddLogTugas(ctx *gin.Context) {
	var payload debtCollectorDto.NewLogTugasPayload
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		validationError := validation.GetValidationError(err)
		if len(validationError) > 0 {
			json.NewResponseBadRequestValidator(ctx, validationError, "bad request", "01", "02")
			return
		}
		json.NewResponseBadRequest(ctx, "invalid payload", "01", "02")
		return
	}

	err := d.debtCollUC.CreateLogTugas(payload)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			json.NewResponseNotFound(ctx, err.Error(), "01", "01")
			return
		}
		json.NewResponseError(ctx, err.Error(), "01", "02")
		return
	}

	json.NewResponseSuccess(ctx, nil, "success", "01", "01")
}

func (d *debtCollectorDelivery) GetLogTugas(ctx *gin.Context) {
	var param debtCollectorDto.Param
	if err := ctx.ShouldBindUri(&param); err != nil {
		validationError := validation.GetValidationError(err)
		if len(validationError) > 0 {
			json.NewResponseBadRequestValidator(ctx, validationError, "bad request", "01", "02")
			return
		}
	}

	log, err := d.debtCollUC.GetLogTugasById(param.ID)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			json.NewResponseNotFound(ctx, err.Error(), "01", "01")
			return
		}
		json.NewResponseError(ctx, err.Error(), "01", "02")
		return
	}

	json.NewResponseSuccess(ctx, log, "success", "01", "01")
}
