package repository

import (
	"fmt"
	"strconv"

	"bitbucket.org/magazine-ondemand/adminplace-api/models"
	"bitbucket.org/magazine-ondemand/adminplace-api/settings"
)

// GetIntegrationByID Consulta integração por id
func GetIntegrationByID(id int) (*models.Integration, error) {
	conn := settings.NewConn().ConnectDB().DB

	row := conn.QueryRow(`select id, name, description, token, secret, verify, active from integration where id=?`, id)
	i := new(models.Integration)
	err := row.Scan(&i.ID, &i.Name, &i.Description, &i.Token, &i.Secret, &i.Verify, &i.Active)
	if err != nil {
		return nil, err
	}

	return i, nil
}

// GetIntegrationByAny Consulta integração por id
func GetIntegrationByAny(b *models.Integration) ([]*models.Integration, error) {
	conn := settings.NewConn().ConnectDB().DB

	sql := fmt.Sprintf(`SELECT id, name, description, token, secret, verify, active from integration where 1 = 1 `)
	if b.Name != "" {
		sql = sql + `AND name like '%` + b.Name + `%' `
	}

	if b.ID != 0 {
		sql = sql + `AND id = ` + strconv.Itoa(b.ID)
	}
	if b.Active != 0 {
		sql = sql + `AND active = ` + strconv.Itoa(b.Active)
	}

	fmt.Println(sql)

	rows, err := conn.Query(sql)
	if err != nil {
		return nil, err
	}

	result := make([]*models.Integration, 0)
	for rows.Next() {
		i := new(models.Integration)
		err := rows.Scan(&i.ID, &i.Name, &i.Description, &i.Token, &i.Secret, &i.Verify, &i.Active)
		if err != nil {
			return nil, err
		}
		result = append(result, i)
	}
	return result, nil
}

// GetIntegrationByVerify Consulta integração por verify_token
func GetIntegrationByVerify(verify string) (*models.Integration, error) {
	conn := settings.NewConn().ConnectDB().DB

	row := conn.QueryRow(`select id, name, description, token, secret, verify, active from integration where verify=?`, verify)
	i := new(models.Integration)
	err := row.Scan(&i.ID, &i.Name, &i.Description, &i.Token, &i.Secret, &i.Verify, &i.Active)
	if err != nil {
		return nil, err
	}

	return i, nil
}

// GetAllIntegration Consulta todas as integrações
func GetAllIntegration() ([]*models.Integration, error) {
	conn := settings.NewConn().ConnectDB().DB

	rows, err := conn.Query(`select id, name, description, token, secret, verify, active from integration`)
	if err != nil {
		return nil, err
	}
	result := make([]*models.Integration, 0)
	for rows.Next() {
		i := new(models.Integration)
		err := rows.Scan(&i.ID, &i.Name, &i.Description, &i.Token, &i.Secret, &i.Verify, &i.Active)
		if err != nil {
			return nil, err
		}
		result = append(result, i)
	}

	return result, nil
}

// CreateIntegration Cria uma nova integração
func CreateIntegration(i *models.Integration) (int64, error) {
	fmt.Println(i.Name)
	conn := settings.NewConn().ConnectDB().DB
	res, err := conn.Exec(`insert integration set name=?, description=?, token=?, secret=?, verify=?, active=?`, i.Name, i.Description, i.Token, i.Secret, i.Verify, i.Active)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	id, _ := res.LastInsertId()
	return id, nil
}

//UpdateIntegration atualiza uma integração
func UpdateIntegration(i *models.Integration) (int64, error) {
	conn := settings.NewConn().ConnectDB().DB

	res, err := conn.Exec(`update integration set name=?, description=?, secret=?, verify=?, active=? where id=?`, i.Name, i.Description, i.Secret, i.Verify, i.Active, i.ID)
	if err != nil {
		return 0, err
	}

	id, _ := res.RowsAffected()
	return id, nil
}

//DeleteIntegration exclui uma integração
func DeleteIntegration(i int) (int64, error) {
	conn := settings.NewConn().ConnectDB().DB

	res, err := conn.Exec(`delete from integration where id=?`, i)
	if err != nil {
		return 0, err
	}

	id, _ := res.RowsAffected()
	return id, nil
}
