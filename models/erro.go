package models

// Error monta o retorno padrão
type Error struct {
	DeveloperMessage string `json:"developerMessage"`
	UserMessage      string `json:"userMessage"`
	ErrorCode        int    `json:"errorCode"`
}

// NewError monta o retorno padrão
func NewError(dm string, um string, ec int) *Error {
	r := &Error{
		DeveloperMessage: dm,
		UserMessage:      um,
		ErrorCode:        ec}
	return r
}
