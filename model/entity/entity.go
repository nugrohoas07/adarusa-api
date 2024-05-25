package entity

import "time"

type (
	Users struct {
		ID         int        `json:"id"`
		Email      string     `json:"email"`
		Password   string     `json:"password"`
		RoleID     int        `json:"role_id"`
		Status     string     `json:"status"`
		CreatedAt  time.Time  `json:"created_at"`
		UpdatedAt  time.Time  `json:"updated_at"`
		VerifiedAt *time.Time `json:"verified_at"`
		DeletedAt  *time.Time `json:"deleted_at"`
	}

	Rekening struct {
		ID            int    `json:"id"`
		UserID        int    `json:"user_id"`
		AccountNumber string `json:"account_number"`
		BankName      string `json:"bank_name"`
	}

	KontakDarurat struct {
		ID          int    `json:"id"`
		UserID      int    `json:"user_id"`
		Name        string `json:"name"`
		PhoneNumber string `json:"phone_number"`
	}

	UserJobDetail struct {
		ID            int     `json:"id"`
		UserID        int     `json:"user_id"`
		JobName       string  `json:"job_name"`
		Salary        float64 `json:"salary"`
		OfficeName    string  `json:"office_name"`
		OfficeContact string  `json:"office_contact"`
		Address       string  `json:"address"`
	}

	UserDetail struct {
		ID          int    `json:"id"`
		UserID      int    `json:"user_id"`
		LimitID     int    `json:"limit_id"`
		NIK         string `json:"nik"`
		FullName    string `json:"fullname"`
		PhoneNumber string `json:"phone_number"`
		Address     string `json:"address"`
		City        string `json:"city"`
		FotoKTP     string `json:"foto_ktp"`
		FotoSelfie  string `json:"foto_selfie"`
	}

	UserCompleteInfo struct {
		UserID              int        `json:"user_id"`
		Email               string     `json:"email"`
		Status              string     `json:"status"`
		AccountNumber       string     `json:"account_number"`
		BankName            string     `json:"bank_name"`
		EmergencyContact    string     `json:"emergency_contact_name"`
		EmergencyPhone      string     `json:"emergency_contact_phone"`
		JobName             string     `json:"job_name"`
		Gaji                float64    `json:"gaji"`
		OfficeName          string     `json:"office_name"`
		NIK                 string     `json:"nik"`
		FullName            string     `json:"fullname"`
		PersonalPhoneNumber string     `json:"personal_phone_number"`
		PersonalAddress     string     `json:"personal_address"`
		City                string     `json:"city"`
		FotoKTP             string     `json:"foto_ktp"`
		FotoSelfie          string     `json:"foto_selfie"`
		CreatedAt           time.Time  `json:"created_at"`
		UpdatedAt           time.Time  `json:"updated_at"`
		VerifiedAt          *time.Time `json:"verified_at"`
		DeletedAt           *time.Time `json:"deleted_at"`
	}
)
