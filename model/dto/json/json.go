package json

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type (
	// JSONResponse - struct for json response success
	jsonResponse struct {
		Code    string      `json:"responseCode"`
		Message string      `json:"message,omitempty"`
		Data    interface{} `json:"data,omitempty"`
	}

	jsonResponseWithPaging struct {
		Code    string      `json:"responseCode"`
		Message string      `json:"message,omitempty"`
		Data    interface{} `json:"data,omitempty"`
		Paging  *Paging     `json:"paging,omitempty"`
	}

	// JSONResponse - struct for json response error
	jsonErrorResponse struct {
		Code    string `json:"responseCode"`
		Message string `json:"message,omitempty"`
		Error   string `json:"error,omitempty"`
	}

	Paging struct {
		Page      int `json:"page,omitempty"`
		TotalData int `json:"totalData,omitempty"`
	}

	ValidationField struct {
		FieldName string `json:"field"`
		Message   string `json:"message"`
	}

	jsonBadRequestResponse struct {
		Code             string            `json:"responseCode"`
		Message          string            `json:"message"`
		ErrorDescription []ValidationField `json:"error_description,omitempty"`
	}
)

func NewResponseSuccess(c *gin.Context, result interface{}, message, serviceCode, responseCode string) {
	c.JSON(http.StatusOK, jsonResponse{
		Code:    "200" + serviceCode + responseCode,
		Message: message,
		Data:    result,
	})
}

func NewResponseSuccessWithPaging(c *gin.Context, result interface{}, paging Paging, message, serviceCode, responseCode string) {
	c.JSON(http.StatusOK, jsonResponseWithPaging{
		Code:    "200" + serviceCode + responseCode,
		Message: message,
		Data:    result,
		Paging:  &paging,
	})
}

func NewResponseBadRequestValidator(c *gin.Context, validationField []ValidationField, message, serviceCode, errorCode string) {
	c.JSON(http.StatusBadRequest, jsonBadRequestResponse{
		Code:             "400" + serviceCode + errorCode,
		Message:          message,
		ErrorDescription: validationField,
	})
}

func NewResponseBadRequest(c *gin.Context, message, serviceCode, errorCode string) {
	c.JSON(http.StatusNotFound, jsonResponse{
		Code:    "400" + serviceCode + errorCode,
		Message: message,
	})
}

func NewResponseError(c *gin.Context, err, serviceCode, errorCode string) {
	log.Error().Msg(err)
	c.JSON(http.StatusInternalServerError, jsonErrorResponse{
		Code:    "500" + serviceCode + errorCode,
		Message: "internal server error",
		Error:   err,
	})
}

func NewResponseNotFound(c *gin.Context, message, serviceCode, errorCode string) {
	c.JSON(http.StatusNotFound, jsonResponse{
		Code:    "404" + serviceCode + errorCode,
		Message: message,
	})
}

func NewResponseForbidden(c *gin.Context, message, serviceCode, errorCode string) {
	c.JSON(http.StatusForbidden, jsonResponse{
		Code:    "403" + serviceCode + errorCode,
		Message: message,
	})
}

func NewResponseUnauthorized(c *gin.Context, message, serviceCode, errorCode string) {
	c.JSON(http.StatusUnauthorized, jsonResponse{
		Code:    "401" + serviceCode + errorCode,
		Message: message,
	})
}

func NewAbortForbidden(c *gin.Context, message, serviceCode, errorCode string) {
	c.AbortWithStatusJSON(http.StatusForbidden, jsonResponse{
		Code:    "403" + serviceCode + errorCode,
		Message: message,
	})
}

func NewAbortUnauthorized(c *gin.Context, message, serviceCode, errorCode string) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, jsonResponse{
		Code:    "401" + serviceCode + errorCode,
		Message: message,
	})
}
