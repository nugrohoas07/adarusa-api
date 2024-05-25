package adminDto

type (
	RequestUpdateStatusUser struct {
		ID     int    `json:"id"`
		Status string `json:"status"`
	}
	AdminResponse struct {
		ID         int    `json:"id"`
		Email      string `json:"email"`
		Status     string `json:"status"`
		VerifiedAt string `json:"verified_at,omitempty"`
	}
)
