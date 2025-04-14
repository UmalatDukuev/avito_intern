package request

type AddProductInput struct {
	PvzID string `json:"pvz_id" binding:"required"`
	Type  string `json:"type" binding:"required"`
	//ReceptionID string `json:"reception_id"`
}
