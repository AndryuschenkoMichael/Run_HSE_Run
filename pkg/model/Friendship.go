package model

type Friendship struct {
	Id         int `json:"id" db:"id"`
	UserIdFrom int `json:"user_id_from" db:"user_id1"`
	UserIdTo   int `json:"user_id_to" db:"user_id2"`
}
