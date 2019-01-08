package models

//MessageSend é a estrutura para envio de mensgens básicas
type MessageSend struct {
	MessagingType string `json:"messaging_type"`
	Recipient     struct {
		ID string `json:"id"`
	} `json:"recipient"`
	MessageData struct {
		Text string `json:"text"`
	} `json:"message"`
}

//MessageResponse é a estrutura de resposta das mensagens básicas com sucesso
type MessageResponse struct {
	RecipientID string `json:"recipient_id"`
	MessageID   string `json:"message_id"`
}

//MensagensGenericasReq é a estrutra de requisição de envio de mensagens por colaborador
type MensagensGenericasReq struct {
	Employees []struct {
		Email string `json:"email"`
	} `json:"employees"`
	Message string `json:"message"`
}

//MensagensGenericasRes é a estrutra de resposta do envio de mensagens por colaborador
type MensagensGenericasRes struct {
	IDLote string   `json:"idLote"`
	Count  int      `json:"count"`
	Send   []Send   `json:"send"`
	Errors []Errors `json:"errors"`
}

//Send ...
type Send struct {
	EmployeeID  string `json:"employee_id"`
	RecipientID string `json:"recipient_id"`
	MessageID   string `json:"message_id"`
}

// Errors ...
type Errors struct {
	EmployeeID string `json:"employee_id"`
	Message    string `json:"message"`
}
