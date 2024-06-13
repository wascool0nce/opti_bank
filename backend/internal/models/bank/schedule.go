package bank

import "encoding/json"

// Время работы отделения
type Schedule struct {
	Days  string `json:"days"`
	Hours string `json:"hours"`
}

// Перегрузка обнаружения переменной в БД
func (s *Schedule) Scan(value interface{}) error {
	if value == nil {
		*s = Schedule{}
		return nil
	}
	var str string
	err := json.Unmarshal(value.([]byte), &str)
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(str), s)
}
