package models

//programmer...
type Programmer struct {
	ID       uint   `json:"id"`
	FullName string `json:"full_name" db:"full_name"`
}
