package models

//Rules é a estrutura de uma regra de envio automático cadastrada
type Rules struct {
	IDRules     int64  `json:"idRules"`
	TypeRules   string `json:"typeRules"`
	Description string `json:"description"`
}

//Cron representa os agendamentos existentes para o envio de uma mensagem
type Cron struct {
	IDCron       int64 `json:"idCron"`
	IDRules      int64 `json:"idRules"`
	MinutesCron  int64 `json:"minutesCron"`
	HourCron     int64 `json:"hourCron"`
	IntervalCron int64 `json:"intervalCron"`
	DayMonthCron int64 `json:"dayMonthCron"`
	MonthCron    int64 `json:"monthCron"`
	DayWeekCron  int64 `json:"dayWeekCron"`
}
