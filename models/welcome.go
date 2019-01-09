package models

//Welcome é a estrutura do texto de boas vindas
type Welcome struct {
	ID     int    `json:"id"`
	Text   string `json:"text"`
	Active bool   `json:"active"`
}
