package schemas

type Ticket struct {
	Id string `json:"id"`
	TicketId string `json:"ticketId"`
	Caller string `json:"caller"`
	Description string `json:"description"`
	ShortDescription string `json:"shortDescription"`
}