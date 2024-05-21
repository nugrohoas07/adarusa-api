package debtCollectorDelivery

import (
	"fp_pinjaman_online/model/dto/debtCollectorDto"
	"fp_pinjaman_online/model/dto/json"
	"fp_pinjaman_online/pkg/validation"
	"fp_pinjaman_online/src/debtCollector"
	"strconv"
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
		dcGroup.GET("/late-debitur")  // get all debitur nunggak
		dcGroup.POST("/tugas/create") // claim tugas ?
		dcGroup.GET("/tugas")         // get all tugas atau user yang pernah di tagih
		// endpoint minta bayaran ???
		dcGroup.GET("tugas/:id/log-tugas", handler.GetAllLogTugas) // get all log
		dcGroup.POST("/log-tugas/create", handler.AddLogTugas)     // membuat log tugas baru
		dcGroup.GET("/log-tugas/:id", handler.GetLogTugas)         // get log detail
		dcGroup.PUT("/log-tugas/:id", handler.EditLogTugas)        // edit log
		dcGroup.DELETE("/log-tugas/:id", handler.DeleteLogTugas)   // hapus log
	}
}

// TODO
// add upload file for job proof ?
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

func (d *debtCollectorDelivery) EditLogTugas(ctx *gin.Context) {
	var param debtCollectorDto.Param
	if err := ctx.ShouldBindUri(&param); err != nil {
		validationError := validation.GetValidationError(err)
		if len(validationError) > 0 {
			json.NewResponseBadRequestValidator(ctx, validationError, "bad request", "01", "02")
			return
		}
	}

	var payload debtCollectorDto.UpdateLogTugasPayload
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		validationError := validation.GetValidationError(err)
		if len(validationError) > 0 {
			json.NewResponseBadRequestValidator(ctx, validationError, "bad request", "01", "02")
			return
		}
		json.NewResponseBadRequest(ctx, "invalid payload", "01", "01")
		return
	}

	err := d.debtCollUC.EditLogTugasById(param.ID, payload)
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

func (d *debtCollectorDelivery) DeleteLogTugas(ctx *gin.Context) {
	var param debtCollectorDto.Param
	if err := ctx.ShouldBindUri(&param); err != nil {
		validationError := validation.GetValidationError(err)
		if len(validationError) > 0 {
			json.NewResponseBadRequestValidator(ctx, validationError, "bad request", "01", "02")
			return
		}
	}

	err := d.debtCollUC.DeleteLogTugasById(param.ID)
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

func (d *debtCollectorDelivery) GetAllLogTugas(ctx *gin.Context) {
	var param debtCollectorDto.Param
	if err := ctx.ShouldBindUri(&param); err != nil {
		validationError := validation.GetValidationError(err)
		if len(validationError) > 0 {
			json.NewResponseBadRequestValidator(ctx, validationError, "bad request", "01", "02")
			return
		}
	}

	var queryParams debtCollectorDto.Query
	if err := ctx.ShouldBindQuery(&queryParams); err != nil {
		validationError := validation.GetValidationError(err)
		if len(validationError) > 0 {
			json.NewResponseBadRequestValidator(ctx, validationError, "bad request", "01", "02")
			return
		}
	}
	page, _ := strconv.Atoi(queryParams.Page)
	size, _ := strconv.Atoi(queryParams.Size)

	logsList, paging, err := d.debtCollUC.GetAllLogTugas(param.ID, page, size)
	if err != nil {
		json.NewResponseError(ctx, err.Error(), "01", "01")
		return
	}

	if len(logsList) == 0 {
		json.NewResponseSuccess(ctx, nil, "data not found", "01", "01")
		return
	}

	json.NewResponseSuccessWithPaging(ctx, logsList, paging, "", "01", "02")
}
