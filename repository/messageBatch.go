package repository

import (
	"fmt"
	"strconv"

	"bitbucket.org/magazine-ondemand/adminplace-api/models"
	"bitbucket.org/magazine-ondemand/adminplace-api/settings"
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

// GetMessageBatchByAny Consulta usuário por qualquer dados
func GetMessageBatchByAny(i *models.MessageBatch) ([]*models.MessageBatch, error) {
	conn := settings.NewConn().ConnectDB().DB

	sql := fmt.Sprintf(`SELECT * FROM message_batch where id_user_send = ` + strconv.Itoa(i.IDUserSend))
	if i.ID > 0 {
		sql = sql + ` AND id =` + strconv.Itoa(i.ID)
	}
	if i.IDIntegration > 0 {
		sql = sql + ` AND id_integration = ` + strconv.Itoa(i.IDIntegration)
	}

	fmt.Println(sql)
	rows, err := conn.Query(sql)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	result := make([]*models.MessageBatch, 0)
	for rows.Next() {
		i := new(models.MessageBatch)
		err := rows.Scan(&i.ID, &i.Text, &i.IDUserSend, &i.SendTime, &i.IDIntegration)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		result = append(result, i)
	}
	return result, nil

}
