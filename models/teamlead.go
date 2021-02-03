package models

//teamlead...
type Teamlead struct {
	ID       uint   `json:"id"`
	FullName string `json:"full_name" db:"full_name"`
}
