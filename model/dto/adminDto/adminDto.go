package adminDto

type (
	RequestUpdateStatusUser struct {
		ID     int    `json:"id"`
		Status string `json:"status"`
	}

	RequestVerifyLoan struct {
		LoanID          int    `json:"loan_id"`
		UserID          int    `json:"user_id"`
		AdminID         int    `json:"admin_id"`
		StatusPengajuan string `json:"status_pengajuan"`
	}
	AdminResponse struct {
		ID     int    `json:"id"`
		Email  string `json:"email"`
		Status string `json:"status"`
	}

	LoanResponse struct {
		LoanID          int     `json:"loan_id"`
		UserID          int     `json:"user_id"`
		JumlahPinjaman  float64 `json:"jumlah_pinjaman"`
		Tenor           int     `json:"tenor"`
		BungaPerBulan   float64 `json:"bunga_per_bulan"`
		Description     string  `json:"description"`
		StatusPengajuan string  `json:"status_pengajuan"`
	}
)
