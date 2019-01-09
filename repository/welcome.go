package repository

import (
	"bitbucket.org/dt_souza/adminplace-api/models"
	"bitbucket.org/dt_souza/adminplace-api/settings"
)

// GetWelcomeByID Consulta texto de boas vindas por id
func GetWelcomeByID(id int) (*models.Welcome, error) {
	conn := settings.NewConn().ConnectDB().DB

	row := conn.QueryRow(`select id, text, active from welcome where id = %d`, id)
	i := new(models.Welcome)
	err := row.Scan(&i.ID, &i.Text, &i.Active)
	if err != nil {
		return nil, err
	}

	return i, nil
}

// GetAllWelcome Consulta todos os textos de boas vindas
func GetAllWelcome() ([]*models.Welcome, error) {
	conn := settings.NewConn().ConnectDB().DB

	rows, err := conn.Query(`select id, text, active from welcome`)
	if err != nil {
		return nil, err
	}
	result := make([]*models.Welcome, 0)
	for rows.Next() {
		i := new(models.Welcome)
		err := rows.Scan(&i.ID, &i.Text, &i.Active)
		if err != nil {
			return nil, err
		}
		result = append(result, i)
	}

	return result, nil
}

// CreateWelcome Cadastra um novo texto de boas vindas
func CreateWelcome(i models.Welcome) (int64, error) {
	conn := settings.NewConn().ConnectDB().DB

	res, err := conn.Exec(`insert welcome set text=?, active=?`, i.Text, i.Active)
	if err != nil {
		return 0, err
	}

	id, _ := res.LastInsertId()
	return id, nil
}

//UpdateWelcome atualiza um texto de boas vindas
func UpdateWelcome(i models.Welcome) (int64, error) {
	conn := settings.NewConn().ConnectDB().DB

	res, err := conn.Exec(`update welcome set text=?, active=? where id=?`, i.Text, i.Active, i.ID)
	if err != nil {
		return 0, err
	}

	id, _ := res.RowsAffected()
	return id, nil
}

//DeleteWelcome exclui um tezto de boas vindas
func DeleteWelcome(i int) (int64, error) {
	conn := settings.NewConn().ConnectDB().DB

	res, err := conn.Exec(`delete from welcome where id=?`, i)
	if err != nil {
		return 0, err
	}

	id, _ := res.RowsAffected()
	return id, nil
}

// GetWelcomeByActive Consulta texto de boas vindas ativo
func GetWelcomeByActive(active bool) (*models.Welcome, error) {
	conn := settings.NewConn().ConnectDB().DB

	row := conn.QueryRow(`select id, text, active from welcome where active = %d`, active)
	i := new(models.Welcome)
	err := row.Scan(&i.ID, &i.Text, &i.Active)
	if err != nil {
		return nil, err
	}

	return i, nil
}
