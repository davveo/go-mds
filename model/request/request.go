package request

type MessageRequest struct {
	MesageBody   string `json:"messageBody" binding:"required"`
	MessageQueue string `json:"messageQueue" binding:"required"`
	MessageId    string `json:"messageId" binding:"required"`
	Extra        string `json:"extra"`
}
