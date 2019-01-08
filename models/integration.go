package models

//Integration é a estrutura para as integrações utilizadas
type Integration struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Token       string `json:"token"`
	Secret      string `json:"secret"`
	Verify      string `json:"verify"`
}
