package employees

import (
	"net/http"
	"strconv"

	"bitbucket.org/magazine-ondemand/adminplace-api/controllers/utils"
	"bitbucket.org/magazine-ondemand/adminplace-api/models"
	"bitbucket.org/magazine-ondemand/adminplace-api/repository"

	"github.com/gorilla/mux"
)

// GetEmployeeByID rota: /employee/{id}
func GetEmployeeByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	employee, err := repository.GetEmployeeByID(id)
	utils.ResponseRequest(w, employee, err)
}

// GetEmployeeByAny rota: /employee/search
func GetEmployeeByAny(w http.ResponseWriter, r *http.Request) {
	utils.ValidationRequest(w, r)
	obj := utils.DecoderRequest(r, &models.Employee{})
	i := obj.(*models.Employee)
	employee, err := repository.GetEmployeeByAny(i)
	utils.ResponseRequest(w, employee, err)

}

// GetAllEmployee rota: /employee
func GetAllEmployee(w http.ResponseWriter, r *http.Request) {
	employee, err := repository.GetAllEmployee()
	utils.ResponseRequest(w, employee, err)
}

// CreateEmployee rota: /employee
func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	utils.ValidationRequest(w, r)
	obj := utils.DecoderRequest(r, &models.Employee{})
	i := obj.(*models.Employee)
	employee, err := repository.CreateEmployee(i)
	utils.ResponseRequest(w, employee, err)
}

// UpdateEmployee rota: /employee
func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	utils.ValidationRequest(w, r)
	obj := utils.DecoderRequest(r, &models.Employee{})
	i := obj.(*models.Employee)

	employee, err := repository.UpdateEmployee(i)
	utils.ResponseRequest(w, employee, err)
}

// DeleteEmployee rota: /employee/:id
func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	employee, err := repository.DeleteEmployee(id)
	utils.ResponseRequest(w, employee, err)
}

// GetEmployeeByWelcome rota: /employee/welcome
func GetEmployeeByWelcome(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	vbool, _ := strconv.Atoi(vars["bool"])
	employee, err := repository.GetEmployeeByWelcome(vbool)
	utils.ResponseRequest(w, employee, err)
}
