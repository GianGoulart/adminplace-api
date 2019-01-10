package models

import "time"

//MessageBatch é a estrutura de lotes de mensagens
type MessageBatch struct {
	ID            int       `json:"id"`
	Text          string    `json:"text"`
	IDUserSend    int       `json:"idUserSend"`
	SendTime      time.Time `json:"sendTime"`
	IDIntegration int       `json:"idIntegration"`
}
