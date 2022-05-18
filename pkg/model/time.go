package model

type Time struct {
	Id     int `json:"id" db:"id"`
	GameId int `json:"game_id" db:"game_id"`
	UserId int `json:"user_id" db:"user_id"`
	Time   int `json:"time" db:"time"`
}
