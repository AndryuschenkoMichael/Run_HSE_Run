package repository

import (
	"Run_Hse_Run/pkg/model"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type FriendPostgres struct {
	db *sqlx.DB
}

func (f *FriendPostgres) AddFriend(userIdFrom, userIdTo int) error {
	var count int
	tx, err := f.db.Begin()
	if err != nil {
		return err
	}

	querySelect := fmt.Sprintf("SELECT Count(friends.id) FROM %s WHERE user_id1=$1 AND user_id2=$2", friendsTable)
	row := tx.QueryRow(querySelect, userIdFrom, userIdTo)
	err = row.Scan(&count)
	if err != nil || count > 0 {
		tx.Rollback()
		return nil
	}

	queryCreate := fmt.Sprintf("INSERT INTO %s (user_id1, user_id2) values ($1, $2)", friendsTable)
	_, err = f.db.Exec(queryCreate, userIdFrom, userIdTo)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (f *FriendPostgres) DeleteFriend(userIdFrom, userIdTo int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE user_id1=$1 AND user_id2=$2", friendsTable)
	_, err := f.db.Exec(query, userIdFrom, userIdTo)
	return err
}

func (f *FriendPostgres) GetFriends(userId int) ([]model.User, error) {
	var users []model.User

	query := fmt.Sprintf(`SELECT us.id, us.nickname, us.email, us.image FROM %s us 
								INNER JOIN %s fr on us.id = fr.user_id2 
								WHERE fr.user_id1 = $1`, usersTable, friendsTable)
	err := f.db.Select(&users, query, userId)

	return users, err
}

func NewFriendPostgres(db *sqlx.DB) *FriendPostgres {
	return &FriendPostgres{db: db}
}
