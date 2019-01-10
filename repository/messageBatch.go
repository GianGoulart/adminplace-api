package repository

import (
	"bitbucket.org/dt_souza/adminplace-api/models"
	"bitbucket.org/dt_souza/adminplace-api/settings"
)

// GetMessageBatchByID Consulta um lote de mensagens por id
func GetMessageBatchByID(id int) (*models.MessageBatch, error) {
	conn := settings.NewConn().ConnectDB().DB

	row := conn.QueryRow(`SELECT id, text, id_user_send, send_time, id_integration FROM message_batch where id=?`, id)
	i := new(models.MessageBatch)
	err := row.Scan(&i.ID, &i.Text, &i.IDUserSend, &i.SendTime, &i.IDIntegration)
	if err != nil {
		return nil, err
	}

	return i, nil
}

// CreateMessageBatch Cadastra um novo lote de mensagens
func CreateMessageBatch(i models.MessageBatch) (int64, error) {
	conn := settings.NewConn().ConnectDB().DB

	res, err := conn.Exec(`insert message_batch set text=?, id_user_send=?, send_time=NOW(), id_integration=?`, i.Text, i.IDUserSend, i.IDIntegration)
	if err != nil {
		return 0, err
	}

	id, _ := res.LastInsertId()
	return id, nil
}
