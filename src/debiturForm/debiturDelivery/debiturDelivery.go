package debiturDelivery

import (
	"fp_pinjaman_online/config/cloudinary"
	"fp_pinjaman_online/model/debiturFormDto"
	"fp_pinjaman_online/model/dto/json"
	"fp_pinjaman_online/pkg/middleware"
	"fp_pinjaman_online/pkg/validation"
	"fp_pinjaman_online/src/debiturForm"
	"strconv"

	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/gin-gonic/gin"
)

type debiturDelivery struct {
	debtUseCase debiturForm.DebiturUseCase
}

func NewDebiturDelivery(v1Group *gin.RouterGroup, useCase debiturForm.DebiturUseCase) {
	handler := debiturDelivery{
		debtUseCase: useCase,
	}
	debtGroup := v1Group.Group("/debitur")
	debtGroup.Use(middleware.JWTAuth())
	{
		debtGroup.POST("/form", handler.createDetailDebitur)
		debtGroup.POST("/:id/form/upload", handler.uploadFiles)
		debtGroup.GET("/users/:roles", handler.getDataByRole)
	}
}

func (c *debiturDelivery) createDetailDebitur(ctx *gin.Context) {
	userIdStr, exists := ctx.Get("userId")
	if !exists {
		json.NewAbortUnauthorized(ctx, "unauthorized", "01", "01")
		return
	}
	userId, err := strconv.Atoi(userIdStr.(string))
	if err != nil {
		json.NewResponseError(ctx, "invalid userId format", "01", "01")
		return
	}

	var debt debiturFormDto.Debitur
	if err := ctx.ShouldBindJSON(&debt); err != nil {
		validationError := validation.GetValidationError(err)
		if len(validationError) > 0 {
			json.NewResponseBadRequest(ctx, "bad request", "01", "01")
			return
		}
	}

	// Set userID dari token JWT ke dalam struct debt
	debt.DetailUser.UserID = userId
    debt.UserJobs.UserID = userId
    debt.EmergencyContact.UserID = userId

	// Pengecekan jika userID dalam body berbeda dengan userID dari token
	if debt.DetailUser.UserID != userId {
		json.NewAbortForbidden(ctx, "forbidden: cannot modify another user's data", "01", "01")
		return
	}

	err = c.debtUseCase.CreateDetailDebitur(debt)
	if err != nil {
		json.NewResponseError(ctx, err.Error(), "01", "01")
		return
	}

	json.NewResponseSuccess(ctx, nil, "success", "01", "01")
}

func (c *debiturDelivery) uploadFiles(ctx *gin.Context) {
	// get userId and roles from context
	roles, _ := ctx.Get("roleName")
	userIdStr, exists := ctx.Get("userId")
	if !exists {
		json.NewResponseUnauthorized(ctx, "unauthorized", "01", "01")
		return
	}

	userId, err := strconv.Atoi(userIdStr.(string))
	if err != nil {
		json.NewResponseError(ctx, "invalid userID", "01", "01")
		return
	}

	debtIdStrPost := ctx.PostForm("user_id")
	debtId, err := strconv.Atoi(debtIdStrPost)
	if err != nil || debtId != userId {
		json.NewResponseBadRequest(ctx, "invalid or mismatched user ID", "01", "01")
		return
	}

	fileKTP, err := ctx.FormFile("foto_ktp")
	if err != nil {
		json.NewResponseBadRequest(ctx, "no foto_ktp file is received", "01", "01")
		return
	}

	fileSelfie, err := ctx.FormFile("foto_selfie")
	if err != nil {
		json.NewResponseBadRequest(ctx, "no foto_selfie is received", "01", "01")
		return
	}

	// upload to cloudinary
	ktpFile, err := fileKTP.Open()
	if err != nil {
		json.NewResponseError(ctx, err.Error(), "01", "01")
		return
	}
	defer ktpFile.Close()

	ktpUploadResp, err := cloudinary.Cloudinary.Upload.Upload(ctx, ktpFile, uploader.UploadParams{
		Folder: "uploads/" + roles.(string) + "/" + strconv.Itoa(debtId) + "/ktp",
	})
	if err != nil {
		json.NewResponseError(ctx, err.Error(), "01", "01")
		return
	}

	// upload foto selfie to cloudinary
	selfieFile, err := fileSelfie.Open()
	if err != nil {
		json.NewResponseError(ctx, err.Error(), "01", "01")
		return
	}
	defer selfieFile.Close()

	selfieFileResp, err := cloudinary.Cloudinary.Upload.Upload(ctx, selfieFile, uploader.UploadParams{
		Folder: "uploads/" + roles.(string) + "/" + strconv.Itoa(debtId) + "/selfie",
	})
	if err != nil {
		json.NewResponseError(ctx, err.Error(), "01", "01")
		return
	}

	err = c.debtUseCase.UpdatePhotoPaths(debtId, ktpUploadResp.SecureURL, selfieFileResp.SecureURL)
	if err != nil {
		json.NewResponseError(ctx, err.Error(), "01", "01")
		return
	}

	json.NewResponseSuccess(ctx, nil, "success", "01", "02")
}

func (c *debiturDelivery) getDataByRole(ctx *gin.Context) {
    role := ctx.Param("roles")
    pageStr := ctx.DefaultQuery("page", "1")
    sizeStr := ctx.DefaultQuery("size", "5")
    status := ctx.DefaultQuery("status", "")

    page, err := strconv.Atoi(pageStr)
    if err != nil {
        json.NewResponseBadRequest(ctx, "bad request: invalid page parameter", "01", "01")
        return
    }
    size, err := strconv.Atoi(sizeStr)
    if err != nil {
        json.NewResponseBadRequest(ctx, "bad request: invalid size parameter", "01", "01")
        return
    }

    debitur, totalData, err := c.debtUseCase.GetDataByRole(role, status, page, size)
    if err != nil {
        json.NewResponseError(ctx, err.Error(), "01", "01")  // Provide detailed error message
        return
    }
    if len(debitur) == 0 {
        json.NewResponseSuccess(ctx, "", "success", "01", "02")
        return
    }

    response := debiturFormDto.Response{
        ResponseCode: 200,
        Data:         debitur,
        Paging:       json.Paging{Page: page, TotalData: totalData},
    }

    json.NewResponseSuccess(ctx, response, "success", "01", "01")
}
