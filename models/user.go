package models

//User é a estrutura de usuário do aplicativo
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Group    string `json:"groupAccess"`
	Active   int    `json:"active"`
}

//Authenticate estrutura
type Authenticate struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
