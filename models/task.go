package models

import "database/sql"

//task...
type Tasks struct {
	ID     uint64       `json:"id"`
	Title  string       `json:"title"`
	TID    uint64       `json:"t_id"`
	PID    uint64       `json:"p_id"`
	Endt   sql.NullBool `json:"end_task"`
	Check  sql.NullBool `json:"check_task" `
	Finish sql.NullBool `json:"finish_task" `
}
