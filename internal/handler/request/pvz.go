package request

type CreatePVZInput struct {
	City string `json:"city" binding:"required"`
}
