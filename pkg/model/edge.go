package model

type Edge struct {
	Id          int     `json:"id" db:"id"`
	StartRoomId int     `json:"start_room_id" db:"start_room_id"`
	EndRoomId   int     `json:"end_room_id" db:"end_room_id"`
	Cost        float64 `json:"cost" db:"cost"`
	CampusId    int     `json:"campus_id" db:"campus_id"`
}
