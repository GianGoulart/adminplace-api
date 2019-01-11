package repository

import (
	"bitbucket.org/dt_souza/adminplace-api/models"
	"bitbucket.org/dt_souza/adminplace-api/settings"
)

// GetMessageByID Consulta uma mensagem por id
func GetMessageByID(id int) (*models.Message, error) {
	conn := settings.NewConn().ConnectDB().DB

	row := conn.QueryRow(`SELECT id, id_batch, id_workplace, send_time, receive_time, read_time FROM message where id=?`, id)
	i := new(models.Message)
	err := row.Scan(&i.ID, &i.IDBatch, &i.IDWorkplace, &i.SendTime, &i.ReceiveTime, &i.ReadTime)
	if err != nil {
		return nil, err
	}

	return i, nil
}

// GetMessageByBatch Consulta todas as mensagens de um lote
func GetMessageByBatch(idBatch int) ([]*models.Message, error) {
	conn := settings.NewConn().ConnectDB().DB

	rows, err := conn.Query(`SELECT id, id_batch, id_workplace, send_time, receive_time, read_time FROM message where id_batch=?`, idBatch)
	if err != nil {
		return nil, err
	}

	result := make([]*models.Message, 0)
	for rows.Next() {
		i := new(models.Message)
		err := rows.Scan(&i.ID, &i.IDBatch, &i.IDWorkplace, &i.SendTime, &i.ReceiveTime, &i.ReadTime)
		if err != nil {
			return nil, err
		}
		result = append(result, i)
	}

	return result, nil
}

// CreateMessage Cadastra uma nova mensagem
func CreateMessage(i models.Message) (int64, error) {
	conn := settings.NewConn().ConnectDB().DB

	res, err := conn.Exec(`insert message set id_batch=?, id_workplace=?, send_time=NOW()`, i.IDBatch, i.IDWorkplace)
	if err != nil {
		return 0, err
	}

	id, _ := res.LastInsertId()
	return id, nil
}

//UpdateReceivedMessage atualiza as mensagens de um colabroador com a data de recebimento
func UpdateReceivedMessage(idWorkplace string, idIntegration int) (int64, error) {
	conn := settings.NewConn().ConnectDB().DB

	res, err := conn.Exec(`update message set receive_time=NOW() where id_workplace=? and id_integration=?`, idWorkplace, idIntegration)
	if err != nil {
		return 0, err
	}

	id, _ := res.RowsAffected()
	return id, nil
}

//UpdateReadedMessage atualiza as mensagens de um colabroador com a data de leitura
func UpdateReadedMessage(idWorkplace string, idIntegration int) (int64, error) {
	conn := settings.NewConn().ConnectDB().DB

	res, err := conn.Exec(`update message set read_time=NOW() where id_workplace=? and id_integration=?`, idWorkplace, idIntegration)
	if err != nil {
		return 0, err
	}

	id, _ := res.RowsAffected()
	return id, nil
}
