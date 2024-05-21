package userDelivery

import (
	"fp_pinjaman_online/model/dto/json"
	"fp_pinjaman_online/model/userDto"
	"fp_pinjaman_online/pkg/validation"
	"fp_pinjaman_online/src/users"

	"github.com/gin-gonic/gin"
)

type userDelivery struct {
	userUC users.UserUseCase
}

func NewUserDelivery(v1Group *gin.RouterGroup, userUc users.UserUseCase) {
	handler := userDelivery{
		userUC: userUc,
	}
	userGroup := v1Group.Group("/users")
	userGroup.POST("/login", handler.login)
	userGroup.POST("/create", handler.createUser)
}

func (c *userDelivery) createUser(ctx *gin.Context) {
	var req userDto.CreateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		validationError := validation.GetValidationError(err)
		if len(validationError) > 0 {
			json.NewResponseBadRequestValidator(ctx, validationError, "bad request", "01", "02")
			return
		}
	}

	err := c.userUC.CreateUser(req)
	if err != nil {
		json.NewResponseError(ctx, err.Error(), "01", "01")
		return
	}
	json.NewResponseSuccess(ctx, nil, "success", "01", "01")
}

func (c *userDelivery) login(ctx *gin.Context) {
	var req userDto.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		validationErro := validation.GetValidationError(err)
		if len(validationErro) > 0 {
			json.NewResponseBadRequestValidator(ctx, validationErro, "bad request", "01", "01")
			return
		}
	}

	token, err := c.userUC.Login(req)
	if err != nil {
		json.NewResponseError(ctx, err.Error(), "01", "01")
		return
	}
	json.NewResponseSuccess(ctx, token, "success", "01", "01")
}