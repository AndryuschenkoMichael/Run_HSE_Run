package repository

import (
	"Run_Hse_Run/pkg/model"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type UsersPostgres struct {
	db *sqlx.DB
}

func (u *UsersPostgres) GetUsersByNicknamePattern(nickname string) ([]model.User, error) {
	var users []model.User

	query := fmt.Sprintf(`SELECT * FROM %s us WHERE us.nickname LIKE $1`, usersTable)
	err := u.db.Select(&users, query, nickname+"%")

	return users, err
}

func (u *UsersPostgres) GetUserById(userId int) (model.User, error) {
	var user model.User
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", usersTable)
	err := u.db.Get(&user, query, userId)

	return user, err
}

func NewUsersPostgres(db *sqlx.DB) *UsersPostgres {
	return &UsersPostgres{db: db}
}
