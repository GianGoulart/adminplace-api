package repository

import (
	"bitbucket.org/dt_souza/adminplace-api/models"
	"bitbucket.org/dt_souza/adminplace-api/settings"
)

// GetEmployeeByID Consulta colaborador por id
func GetEmployeeByID(id int) (*models.Employee, error) {
	conn := settings.NewConn().ConnectDB().DB

	row := conn.QueryRow(`SELECT id, first_name, last_name, name, email, job_title, department, employee_number, id_workplace, account_claim_time, welcome FROM employee where id = %d`, id)
	i := new(models.Employee)
	err := row.Scan(&i.ID, &i.FirstName, &i.LastName, &i.Name, &i.Email, &i.JobTitle, &i.Department, &i.EmployeeNumber, &i.IDWorkplace, &i.AccountClaimTime, &i.Welcome)
	if err != nil {
		return nil, err
	}

	return i, nil
}

// GetAllEmployee Consulta todas as integrações
func GetAllEmployee() ([]*models.Employee, error) {
	conn := settings.NewConn().ConnectDB().DB

	rows, err := conn.Query(`SELECT id, first_name, last_name, name, email, job_title, department, employee_number, id_workplace, account_claim_time, welcome FROM employee`)
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

// CreateEmployee Cria uma nova integração
func CreateEmployee(i models.Employee) (int64, error) {
	conn := settings.NewConn().ConnectDB().DB

	res, err := conn.Exec(`insert employess set first_name=?, last_name=?, name=?, email=?, job_title=?, department=?, employee_number=?, id_workplace=?, account_claim_time=?, welcome=?`, i.FirstName, i.LastName, i.Name, i.Email, i.JobTitle, i.Department, i.EmployeeNumber, i.IDWorkplace, i.AccountClaimTime, i.Welcome)
	if err != nil {
		return 0, err
	}

	id, _ := res.LastInsertId()
	return id, nil
}

//UpdateEmployee atualiza uma integração
func UpdateEmployee(i models.Employee) (int64, error) {
	conn := settings.NewConn().ConnectDB().DB

	res, err := conn.Exec(`update employess set first_name=?, last_name=?, name=?, email=?, job_title=?, department=?, employee_number=?, id_workplace=?, account_claim_time=?, welcome=? where id=?`, i.FirstName, i.LastName, i.Name, i.Email, i.JobTitle, i.Department, i.EmployeeNumber, i.IDWorkplace, i.AccountClaimTime, i.Welcome, i.ID)
	if err != nil {
		return 0, err
	}

	id, _ := res.RowsAffected()
	return id, nil
}

//DeleteEmployee exclui uma integração
func DeleteEmployee(i int) (int64, error) {
	conn := settings.NewConn().ConnectDB().DB

	res, err := conn.Exec(`delete from employess where id=?`, i)
	if err != nil {
		return 0, err
	}

	id, _ := res.RowsAffected()
	return id, nil
}
