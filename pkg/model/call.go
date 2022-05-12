package model

type Call struct {
	Id           int `json:"id" db:"id"`
	UserIdFirst  int `json:"user_id_first" db:"user_id_first"`
	RoomIdFirst  int `json:"room_id_first" db:"room_id_first"`
	UserIdSecond int `json:"user_id_second" db:"user_id_second"`
}
