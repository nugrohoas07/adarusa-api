package debtCollectorDto

type (
	NewLogTugasPayload struct {
		TugasId     string `json:"tugasId" binding:"required,number"`
		Description string `json:"description" binding:"required"`
	}

	UpdateLogTugasPayload struct {
		TugasId     string `json:"tugasId" binding:"omitempty,number"`
		Description string `json:"description" binding:"omitempty"`
	}
)
