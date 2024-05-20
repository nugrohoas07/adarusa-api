package debtCollectorEntity

type (
	Tugas struct {
		ID          string `json:"id"`
		UserId      string `json:"userId"`
		CollectorId string `json:"collectorId"`
		Status      string `json:"status"`
	}

	LogTugas struct {
		ID          string `json:"id"`
		TugasId     string `json:"tugasId"`
		Description string `json:"description"`
		CreatedAt   string `json:"createdAt,omitempty"`
		UpdatedAt   string `json:"updatedAt,omitempty"`
	}
)
