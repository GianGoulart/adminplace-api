package repository

import (
	"fmt"
	"strconv"

	"bitbucket.org/magazine-ondemand/adminplace-api/models"
	"bitbucket.org/magazine-ondemand/adminplace-api/settings"
)

// Authenticate autentica usuário
func Authenticate(user models.Authenticate) (*models.User, error) {
	conn := settings.NewConn().ConnectDB().DB

	row := conn.QueryRow(`select id, name, email, password, groupAccess, active from user where email=? and password=? and active = ?`, user.Email, user.Password, true)

	usuario := new(models.User)
	err := row.Scan(&usuario.ID, &usuario.Name, &usuario.Email, &usuario.Password, &usuario.Group, &usuario.Active)

	if err != nil {
		return nil, err
	}

	return usuario, nil
}

// GetUserByID Consulta usuário por id
func GetUserByID(id int) (*models.User, error) {
	conn := settings.NewConn().ConnectDB().DB

	row := conn.QueryRow(`select id, name, email, password, groupAccess, active from user where id=?`, id)
	i := new(models.User)
	err := row.Scan(&i.ID, &i.Name, &i.Email, &i.Password, &i.Group, &i.Active)
	if err != nil {
		return nil, err
	}

	return i, nil
}

// GetAllUser Consulta todos usuários
func GetAllUser() ([]*models.User, error) {
	conn := settings.NewConn().ConnectDB().DB

	rows, err := conn.Query(`select id, name, email, groupAccess, password , active from user`)
	if err != nil {
		return nil, err
	}
	result := make([]*models.User, 0)
	for rows.Next() {
		i := new(models.User)
		err := rows.Scan(&i.ID, &i.Name, &i.Email, &i.Group, &i.Password, &i.Active)
		if err != nil {
			return nil, err
		}
		result = append(result, i)
	}
	return result, nil
}

// GetUserByAny Consulta usuário por qualquer dados
func GetUserByAny(e *models.User) ([]*models.User, error) {
	conn := settings.NewConn().ConnectDB().DB

	sql := fmt.Sprintf(`SELECT id, name, email, groupAccess, password, active FROM user where 1 = 1 `)
	if e.Name != "" {
		sql = sql + `AND name like '%` + e.Name + `%' `
	}
	if e.Email != "" {
		sql = sql + `AND email like '%` + e.Email + `%' `
	}
	if e.Group != "" {
		sql = sql + `AND groupAccess = '` + e.Group + `' `
	}
	if e.Active != 0 {
		sql = sql + `AND active = ` + strconv.Itoa(e.Active)
	}
	if e.ID != 0 {
		sql = sql + `AND id = ` + strconv.Itoa(e.ID)
	}
	rows, err := conn.Query(sql)
	if err != nil {
		return nil, err
	}

	result := make([]*models.User, 0)
	for rows.Next() {
		i := new(models.User)
		err := rows.Scan(&i.ID, &i.Name, &i.Email, &i.Group, &i.Password, &i.Active)
		if err != nil {
			return nil, err
		}
		result = append(result, i)
	}
	return result, nil

}

// CreateUser Cria um novo usuário
func CreateUser(i *models.User) (int64, error) {
	fmt.Println(i)
	conn := settings.NewConn().ConnectDB().DB
	res, err := conn.Exec(`insert user set name=?, email=?, password=? ,active=?, groupAccess=?`, i.Name, i.Email, i.Password, i.Active, i.Group)
	if err != nil {
		return 0, err
	}
	id, _ := res.LastInsertId()
	fmt.Println(err)

	return id, nil
}

//UpdateUser atualiza um usuário
func UpdateUser(i *models.User) (int64, error) {
	conn := settings.NewConn().ConnectDB().DB

	res, err := conn.Exec(`update user set name=?, email=?, groupAccess=?,password=? ,active=? where id=?`, i.Name, i.Email, i.Group, i.Password, i.Active, i.ID)
	if err != nil {
		return 0, err
	}
	fmt.Println(res)
	id, _ := res.RowsAffected()
	return id, nil
}

//DeleteUser exclui um usuário
func DeleteUser(i int) (int64, error) {
	conn := settings.NewConn().ConnectDB().DB

	res, err := conn.Exec(`delete from user where id=?`, i)
	if err != nil {
		return 0, err
	}

	id, _ := res.RowsAffected()
	return id, nil
}
