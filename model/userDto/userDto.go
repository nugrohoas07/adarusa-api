package userDto

import "fp_pinjaman_online/model/dto/json"

type (
	CreateRequest struct {
		Name string
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=8,password"`
		Roles string
	}

	LoginRequest struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	Response struct {
		ResponseCode string      `json:"responseCode"`
		Data         []User      `json:"data"`
		Paging       json.Paging `json:"paging"`
	}

	User struct {
		Id       string `json:"id"`
		Fullname string `json:"fullname"`
		Email    string `json:"email"`
		Password string `json:"password"`
		Roles string `json:"role"`
	}

	Update struct {
		Fullname string `json:"name" binding:"omitempty"`
		Password string `json:"password" binding:"omitempty,min=8,password"`
		Email    string
	}
)