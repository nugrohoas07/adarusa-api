package debiturFormDto


type Debitur struct {
    DetailUser      DetailDebitur   `json:"detail_user"`
    UserJobs        UserJobs        `json:"user_job_detail"`
    EmergencyContact EmergencyContact `json:"emergency"`
}

type DetailDebitur struct {
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

type UserJobs struct {
    UserID      int    `json:"user_id,omitempty"`
    JobName      string `json:"job_name" binding:"required"`
    Salary       float64 `json:"salary" binding:"required"`
    OfficeName   string `json:"office_name" binding:"required"`
    OfficeContact string `json:"office_contact" binding:"required"`
    OfficeAddress string `json:"office_address" binding:"required"`
}

type EmergencyContact struct {
    UserID      int    `json:"user_id,omitempty"`
    Name        string `json:"name" binding:"required"`
    PhoneNumber string `json:"phone_number" binding:"required"`
}
