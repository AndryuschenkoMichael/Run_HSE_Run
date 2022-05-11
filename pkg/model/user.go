package model

type User struct {
	Id       int    `json:"id" db:"id"`
	Nickname string `json:"nickname" db:"nickname"`
	Email    string `json:"email" db:"email"`
	Image    int    `json:"image" db:"image"`
}
