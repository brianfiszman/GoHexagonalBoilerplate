package dtos

type CreateTicketDTO struct {
	CallerID         string `json:"caller_id"`
	Description      string `json:"description"`
	ShortDescription string `json:"short_description"`
}
