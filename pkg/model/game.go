package model

type Game struct {
	UserIdFirst, RoomIdFirst   int
	UserIdSecond, RoomIdSecond int
}

type GameUsers struct {
	Id           int `json:"id" db:"id"`
	UserIdFirst  int `json:"user_id_first" db:"user_id_first"`
	UserIdSecond int `json:"user_id_second" db:"user_id_second"`
}
