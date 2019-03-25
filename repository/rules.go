package repository

import (
	"fmt"

	"bitbucket.org/magazine-ondemand/adminplace-api/models"
	"bitbucket.org/magazine-ondemand/adminplace-api/settings"
)

// GetAllRules Consulta todas as regras
func GetAllRules() ([]*models.Rules, error) {
	conn := settings.NewConn().ConnectDB().DB

	rows, err := conn.Query(`SELECT id_rules, type_rules, description FROM rules`)
	if err != nil {
		return nil, err
	}
	result := make([]*models.Rules, 0)
	for rows.Next() {
		i := new(models.Rules)
		err := rows.Scan(&i.IDRules, &i.TypeRules, &i.Description)
		fmt.Println(err)
		if err != nil {
			return nil, err
		}
		result = append(result, i)
	}

	return result, nil
}

// CreateRules Cria uma nova regra
func CreateRules(i *models.Rules) (int64, error) {
	conn := settings.NewConn().ConnectDB().DB

	res, err := conn.Exec(`insert rules set type_rules=?, description=?`, i.TypeRules, i.Description)
	if err != nil {
		return 0, err
	}

	id, _ := res.LastInsertId()
	return id, nil
}

//UpdateRules atualiza uma reagra
func UpdateRules(i *models.Rules) (int64, error) {

	conn := settings.NewConn().ConnectDB().DB
	res, err := conn.Exec(`update rules set type_rules=?, description=? where id_rules= ? `, i.TypeRules, i.Description, i.IDRules)
	if err != nil {
		return 0, err
	}

	id, _ := res.RowsAffected()
	return id, nil
}

// GetAllCrons Consulta todos os agendamentos de uma mensagem
func GetAllCrons() ([]*models.Cron, error) {
	conn := settings.NewConn().ConnectDB().DB

	rows, err := conn.Query(`SELECT id_cron, id_rules, minutes_cron, hours_cron, interval_cron, dayMonth_cron, month_cron, dayWeek_cron FROM cron`)
	if err != nil {
		return nil, err
	}
	result := make([]*models.Cron, 0)
	for rows.Next() {
		i := new(models.Cron)
		err := rows.Scan(&i.IDCron, &i.IDRules, &i.MinutesCron, &i.HourCron, &i.IntervalCron, &i.DayMonthCron, &i.MonthCron, &i.DayWeekCron)
		fmt.Println(err)
		if err != nil {
			return nil, err
		}
		result = append(result, i)
	}

	return result, nil
}

// CreateCrons Cria uma novo agendamento para uma mensagem
func CreateCrons(i *models.Cron) (int64, error) {
	conn := settings.NewConn().ConnectDB().DB

	res, err := conn.Exec(`insert cron set id_rules=?, minutes_cron=?, hours_cron=?, interval_cron=?, dayMonth_cron=?, month_cron=?, dayWeek_cron=?`, i.IDRules, i.MinutesCron, i.HourCron, i.IntervalCron, i.DayMonthCron, i.MonthCron, i.DayWeekCron)
	if err != nil {
		return 0, err
	}

	id, _ := res.LastInsertId()
	return id, nil
}

//UpdateCrons atualiza um agendamento para uma mensagem
func UpdateCrons(i *models.Cron) (int64, error) {
	conn := settings.NewConn().ConnectDB().DB
	res, err := conn.Exec(`update cron set id_rules=?, minutes_cron=?, hours_cron=?, interval_cron=?, dayMonth_cron=?, month_cron=?, dayWeek_cron where id_cron=?`, i.IDRules, i.MinutesCron, i.HourCron, i.IntervalCron, i.DayMonthCron, i.MonthCron, i.DayWeekCron, i.IDCron)
	if err != nil {
		return 0, err
	}

	id, _ := res.RowsAffected()
	return id, nil
}
