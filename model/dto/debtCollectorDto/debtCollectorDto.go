package debtCollectorDto

type (
	NewLogTugasPayload struct {
		TugasId     string `json:"tugasId" binding:"required,number"`
		Description string `json:"description" binding:"required"`
	}

	UpdateLogTugasPayload struct {
		Description string `json:"description" binding:"omitempty"`
	}

	NewTugasPayload struct {
		UserId string `json:"userId" binding:"required,number"`
	}

	Param struct {
		ID string `uri:"id" binding:"required,number"`
	}

	Query struct {
		Page string `form:"page" binding:"omitempty,number"`
		Size string `form:"size" binding:"omitempty,number"`
	}
)
