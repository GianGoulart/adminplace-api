package models

import "time"

//Employee é a estrutura de colaboradores da empresa
type Employee struct {
	ID               int       `json:"id"`
	FirstName        string    `json:"firstName"`
	LastName         string    `json:"lastName"`
	Name             string    `json:"name"`
	Email            string    `json:"email"`
	JobTitle         string    `json:"jobTitle"`
	Department       string    `json:"department"`
	EmployeeNumber   int       `json:"employeeNumber"`
	IDWorkplace      string    `json:"idWorkplace"`
	AccountClaimTime time.Time `json:"accountClaimTime"`
	Welcome          int       `json:"welcome"`
}
