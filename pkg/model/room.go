package model

type Room struct {
	Id       int    `json:"id" db:"id"`
	Code     string `json:"code" db:"code"`
	CampusId int    `json:"campus_id" db:"campus_id"`
}
