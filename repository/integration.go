package repository

import (
	"bitbucket.org/dt_souza/adminplace-api/models"
	"bitbucket.org/dt_souza/adminplace-api/settings"
)

// GetIntegrationByID Consulta integração por id
func GetIntegrationByID(id int) (*models.Integration, error) {
	conn := settings.NewConn().ConnectDB().DB

	row := conn.QueryRow(`select id, nome, descricao, token, secret, verify from integracao where id = %d`, id)
	i := new(models.Integration)
	err := row.Scan(&i.ID, &i.Name, &i.Description, &i.Token, &i.Secret, &i.Verify)
	if err != nil {
		return nil, err
	}

	return i, nil
}

// GetAllIntegration Consulta todas as integrações
func GetAllIntegration() ([]*models.Integration, error) {
	conn := settings.NewConn().ConnectDB().DB

	rows, err := conn.Query(`select id, nome, descricao, token, secret, verify from integracao`)
	if err != nil {
		return nil, err
	}
	result := make([]*models.Integration, 0)
	for rows.Next() {
		i := new(models.Integration)
		err := rows.Scan(&i.ID, &i.Name, &i.Description, &i.Token, &i.Secret, &i.Verify)
		if err != nil {
			return nil, err
		}
		result = append(result, i)
	}

	return result, nil
}

// CreateIntegration Cria uma nova integração
func CreateIntegration(i models.Integration) (int64, error) {
	conn := settings.NewConn().ConnectDB().DB

	res, err := conn.Exec(`insert integracao set nome = ?, descricao = ?, token = ?, secret = ?, verify = ?`, i.Name, i.Description, i.Token, i.Secret, i.Verify)
	if err != nil {
		return 0, err
	}

	id, _ := res.LastInsertId()
	return id, nil
}

//UpdateIntegration atualiza uma integração
func UpdateIntegration(i models.Integration) (int64, error) {
	conn := settings.NewConn().ConnectDB().DB

	res, err := conn.Exec(`update integracao set nome = ?, descricao = ?, token = ?, secret = ?, verify = ? where id = ?`, i.Name, i.Description, i.Token, i.Secret, i.Verify, i.ID)
	if err != nil {
		return 0, err
	}

	id, _ := res.RowsAffected()
	return id, nil
}

//DeleteIntegration exclui uma integração
func DeleteIntegration(i int) (int64, error) {
	conn := settings.NewConn().ConnectDB().DB

	res, err := conn.Exec(`delete from integracao where id = ?`, i)
	if err != nil {
		return 0, err
	}

	id, _ := res.RowsAffected()
	return id, nil
}