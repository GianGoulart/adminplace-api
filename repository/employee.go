package repository

import (
	"fmt"
	"strconv"

	"bitbucket.org/magazine-ondemand/adminplace-api/models"
	"bitbucket.org/magazine-ondemand/adminplace-api/settings"
)

// GetEmployeeByID Consulta colaborador por id
func GetEmployeeByID(id int) (*models.Employee, error) {
	conn := settings.NewConn().ConnectDB().DB

	row := conn.QueryRow(`SELECT id, first_name, last_name, name, email, job_title, department, employee_number, id_workplace, account_claim_time, welcome FROM employee where id=?`, id)
	i := new(models.Employee)
	err := row.Scan(&i.ID, &i.FirstName, &i.LastName, &i.Name, &i.Email, &i.JobTitle, &i.Department, &i.EmployeeNumber, &i.IDWorkplace, &i.AccountClaimTime, &i.Welcome)
	if err != nil {
		return nil, err
	}

	return i, nil
}

// GetEmployeeByAny Consulta colaborador por qualquer dados
func GetEmployeeByAny(e *models.Employee) ([]*models.Employee, error) {
	conn := settings.NewConn().ConnectDB().DB

	sql := fmt.Sprintf(`SELECT id, first_name, last_name, name, email, job_title, department, employee_number, id_workplace, welcome FROM employee where 1 = 1 `)
	if e.Name != "" {
		sql = sql + `AND name like '%` + e.Name + `%' `
	}
	if e.Email != "" {
		sql = sql + `AND email like '%` + e.Email + `%' `
	}
	if e.JobTitle != "" {
		sql = sql + `AND job_title like '%` + e.JobTitle + `%' `
	}
	if e.EmployeeNumber != 0 {
		sql = sql + `AND employee_number = ` + strconv.Itoa(e.EmployeeNumber)
	}

	rows, err := conn.Query(sql)
	if err != nil {
		return nil, err
	}

	result := make([]*models.Employee, 0)
	for rows.Next() {
		i := new(models.Employee)
		err := rows.Scan(&i.ID, &i.FirstName, &i.LastName, &i.Name, &i.Email, &i.JobTitle, &i.Department, &i.EmployeeNumber, &i.IDWorkplace, &i.Welcome)
		if err != nil {
			return nil, err
		}
		result = append(result, i)
	}

	return result, nil

}

// GetAllEmployee Consulta todas as integrações
func GetAllEmployee() ([]*models.Employee, error) {
	conn := settings.NewConn().ConnectDB().DB

	rows, err := conn.Query(`SELECT id, first_name, last_name, name, email, job_title, department, employee_number, id_workplace, welcome FROM employee`)
	if err != nil {
		return nil, err
	}
	result := make([]*models.Employee, 0)
	for rows.Next() {
		i := new(models.Employee)
		err := rows.Scan(&i.ID, &i.FirstName, &i.LastName, &i.Name, &i.Email, &i.JobTitle, &i.Department, &i.EmployeeNumber, &i.IDWorkplace, &i.Welcome)
		fmt.Println(err)
		if err != nil {
			return nil, err
		}
		result = append(result, i)
	}

	return result, nil
}

// CreateEmployee Cria uma nova integração
func CreateEmployee(i *models.Employee) (int64, error) {
	conn := settings.NewConn().ConnectDB().DB

	res, err := conn.Exec(`insert employee set first_name=?, last_name=?, name=?, email=?, job_title=?, department=?, employee_number=?, id_workplace=?, account_claim_time=?, welcome=?`, i.FirstName, i.LastName, i.Name, i.Email, i.JobTitle, i.Department, i.EmployeeNumber, i.IDWorkplace, i.AccountClaimTime, i.Welcome)
	if err != nil {
		return 0, err
	}

	id, _ := res.LastInsertId()
	return id, nil
}

//UpdateEmployee atualiza uma integração
func UpdateEmployee(i *models.Employee) (int64, error) {
	conn := settings.NewConn().ConnectDB().DB
	fmt.Println(i)
	res, err := conn.Exec(`update employee set first_name=?, last_name=?, name=?, email=?, job_title=?, department=? where id=?`, i.FirstName, i.LastName, i.Name, i.Email, i.JobTitle, i.Department, i.ID)
	if err != nil {
		return 0, err
	}

	id, _ := res.RowsAffected()
	return id, nil
}

//DeleteEmployee exclui uma integração
func DeleteEmployee(i int) (int64, error) {
	conn := settings.NewConn().ConnectDB().DB

	res, err := conn.Exec(`delete from employee where id=?`, i)
	if err != nil {
		return 0, err
	}

	id, _ := res.RowsAffected()
	return id, nil
}

// GetEmployeeByWelcome Consulta colaborador pela flag de bem vindo
func GetEmployeeByWelcome(welcome int) ([]*models.Employee, error) {
	conn := settings.NewConn().ConnectDB().DB

	rows, err := conn.Query(`SELECT id, first_name, last_name, name, email, job_title, department, employee_number, id_workplace, account_claim_time, welcome FROM employee where welcome=?`, welcome)
	if err != nil {
		return nil, err
	}
	result := make([]*models.Employee, 0)
	for rows.Next() {
		i := new(models.Employee)
		err := rows.Scan(&i.ID, &i.FirstName, &i.LastName, &i.Name, &i.Email, &i.JobTitle, &i.Department, &i.EmployeeNumber, &i.IDWorkplace, &i.AccountClaimTime, &i.Welcome)
		if err != nil {
			return nil, err
		}
		result = append(result, i)
	}

	return result, nil
}
