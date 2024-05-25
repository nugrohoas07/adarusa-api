package dcFormDto

import "fp_pinjaman_online/model/dto/json"

type (
	DetailDC struct {
		UserID      int    `json:"user_id,omitempty"`
		LimitID     int    `json:"limit_id" binding:"required"`
		Nik         string `json:"nik" binding:"required"`
		Fullname    string `json:"fullname" binding:"required"`
		PhoneNumber string `json:"phone_number" binding:"required"`
		Address     string `json:"address" binding:"required"`
		City        string `json:"city" binding:"required"`
		FotoKtp     string `json:"foto_ktp"`
		FotoSelfie  string `json:"foto_selfie"`
	}

	Response struct {
		ResponseCode int         `json:"responseCode"`
		Data         []DetailDC  `json:"data"`
		Paging       json.Paging `json:"paging"`
	}
)