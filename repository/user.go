package repository

import (
	"bitbucket.org/dt_souza/adminplace-api/models"
	"bitbucket.org/dt_souza/adminplace-api/settings"
)

// GetUserByID Consulta usuário por id
func GetUserByID(id int) (*models.User, error) {
	conn := settings.NewConn().ConnectDB().DB

	row := conn.QueryRow(`select id, nome, email, senha from user where id = %d`, id)
	i := new(models.User)
	err := row.Scan(&i.ID, &i.Name, &i.Email, &i.Password)
	if err != nil {
		return nil, err
	}

	return i, nil
}

// GetAllUser Consulta todos usuários
func GetAllUser() ([]*models.User, error) {
	conn := settings.NewConn().ConnectDB().DB

	rows, err := conn.Query(`select id, nome, email, senha from user`)
	if err != nil {
		return nil, err
	}
	result := make([]*models.User, 0)
	for rows.Next() {
		i := new(models.User)
		err := rows.Scan(i.ID, &i.Name, &i.Email, &i.Password)
		if err != nil {
			return nil, err
		}
		result = append(result, i)
	}

	return result, nil
}

// CreateUser Cria um novo usuário
func CreateUser(i models.User) (int64, error) {
	conn := settings.NewConn().ConnectDB().DB

	res, err := conn.Exec(`insert user set nome = ?, email = ?, senha = ?`, i.Name, i.Email, i.Password)
	if err != nil {
		return 0, err
	}

	id, _ := res.LastInsertId()
	return id, nil
}

//UpdateUser atualiza um usuário
func UpdateUser(i models.User) (int64, error) {
	conn := settings.NewConn().ConnectDB().DB

	res, err := conn.Exec(`update user set nome = ?, email = ?, senha = ? where id = ?`, i.Name, i.Email, i.Password, i.ID)
	if err != nil {
		return 0, err
	}

	id, _ := res.RowsAffected()
	return id, nil
}

//DeleteUser exclui um usuário
func DeleteUser(i int) (int64, error) {
	conn := settings.NewConn().ConnectDB().DB

	res, err := conn.Exec(`delete from user where id = ?`, i)
	if err != nil {
		return 0, err
	}

	id, _ := res.RowsAffected()
	return id, nil
}
