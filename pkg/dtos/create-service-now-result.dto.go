package dtos

type CreateServiceNowResultDTO struct {
	Result struct {
		IssueType        string `json:"sys_class_name"`
		Description      string `json:"description"`
		ShortDescription string `json:"short_description"`
		Assignee         string `json:"assigned_to"`
		CreatedAt        string `json:"sys_created_on"`
		ProjectKey       string `json:"sys_id"`
		ProjectName      string `json:"parent"`
		TicketId         string `json:"number"`
	} `json:"result"`
}
