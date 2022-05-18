package repository

import (
	"Run_Hse_Run/pkg/model"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type GamePostgres struct {
	db *sqlx.DB
}

func (g *GamePostgres) GetGame(gameId int) (model.GameUsers, error) {
	var game model.GameUsers
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", gamesTable)
	err := g.db.Get(&game, query, gameId)

	return game, err
}

func (g *GamePostgres) GetTime(gameId, userId int) (model.Time, error) {
	var time model.Time
	query := fmt.Sprintf("SELECT * FROM %s WHERE game_id=$1 AND user_id=$2", timesTable)
	err := g.db.Get(&time, query, gameId, userId)

	return time, err
}

func (g *GamePostgres) AddGame(userIdFirst, userIdSecond int) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (user_id_first, user_id_second) values ($1, $2) RETURNING id", gamesTable)
	row := g.db.QueryRow(query, userIdFirst, userIdSecond)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (g *GamePostgres) AddTime(gameId, userId, time int) error {
	query := fmt.Sprintf("INSERT INTO %s (game_id, user_id, time) values ($1, $2, $3)", timesTable)
	_, err := g.db.Exec(query, gameId, userId, time)

	return err
}

func (g *GamePostgres) UpdateTime(gameId, userId, time int) error {
	query := fmt.Sprintf("UPDATE %s SET time = $1 WHERE game_id = $2 AND user_id = $3", timesTable)
	_, err := g.db.Exec(query, time, gameId, userId)

	return err
}

func (g *GamePostgres) GetRoomById(roomId int) (model.Room, error) {
	var room model.Room
	query := fmt.Sprintf(`SELECT * FROM %s rm WHERE rm.id = $1`, roomsTable)
	err := g.db.Get(&room, query, roomId)

	return room, err
}

func (g *GamePostgres) AddCall(userIdFirst, userIdSecond, roomIdFirst int) (model.Game, error) {
	var count int
	tx, err := g.db.Begin()
	if err != nil {
		return model.Game{}, err
	}

	querySelect := fmt.Sprintf("SELECT Count(id) FROM %s WHERE user_id_first=$1 AND user_id_second=$2", callsTable)
	row := tx.QueryRow(querySelect, userIdSecond, userIdFirst)
	err = row.Scan(&count)
	if err != nil {
		tx.Rollback()
		return model.Game{}, err
	}

	if count > 0 {
		var call model.Call
		query1 := fmt.Sprintf("DELETE FROM %s WHERE user_id_first=$1 AND user_id_second=$2", callsTable)
		query2 := fmt.Sprintf("SELECT * FROM %s WHERE user_id_first=$1 AND user_id_second=$2 LIMIT 1", callsTable)
		row1 := tx.QueryRow(query2, userIdSecond, userIdFirst)
		err = row1.Scan(&call.Id, &call.UserIdFirst, &call.RoomIdFirst, &call.UserIdSecond)
		if err != nil {
			tx.Rollback()
			return model.Game{}, err
		}

		_, err = tx.Exec(query1, userIdSecond, userIdFirst)

		if err != nil {
			tx.Rollback()
			return model.Game{}, err
		}

		return model.Game{
			UserIdFirst:  userIdFirst,
			RoomIdFirst:  roomIdFirst,
			UserIdSecond: userIdSecond,
			RoomIdSecond: call.RoomIdFirst,
		}, tx.Commit()
	} else {
		query1 := fmt.Sprintf("INSERT INTO %s (user_id_first, room_id_first, user_id_second) values ($1, $2, $3)",
			callsTable)

		_, err := tx.Exec(query1, userIdFirst, roomIdFirst, userIdSecond)

		if err != nil {
			tx.Rollback()
			return model.Game{}, err
		}

		return model.Game{
			UserIdFirst:  -1,
			RoomIdFirst:  -1,
			UserIdSecond: -1,
			RoomIdSecond: -1,
		}, tx.Commit()
	}
}

func (g *GamePostgres) DeleteCall(userIdFirst, userIdSecond int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE user_id_first=$1 AND user_id_second=$2", callsTable)
	_, err := g.db.Exec(query, userIdFirst, userIdSecond)
	return err
}

func (g *GamePostgres) GetListOfEdges(startRoomId int) ([]model.Edge, error) {
	var edges []model.Edge

	query := fmt.Sprintf(`SELECT * FROM %s ed WHERE ed.start_room_id = $1`, edgesTable)
	err := g.db.Select(&edges, query, startRoomId)

	return edges, err
}

func (g *GamePostgres) GetEdge(startRoomId, endRoomId int) (model.Edge, error) {
	var edges []model.Edge

	query := fmt.Sprintf(`SELECT * FROM %s ed WHERE ed.start_room_id = $1 AND ed.end_room_id = $2`, edgesTable)
	err := g.db.Select(&edges, query, startRoomId, endRoomId)

	if len(edges) == 0 {
		return model.Edge{}, errors.New(fmt.Sprintf("not found edge with start_room_id = %d, end_room_id = %d",
			startRoomId, endRoomId))
	}

	return edges[0], err
}

func (g *GamePostgres) GetRoomByCodePattern(code string, campusId int) ([]model.Room, error) {
	var rooms []model.Room

	query := fmt.Sprintf(`SELECT * FROM %s rm WHERE rm.campus_id = $1 AND rm.code LIKE $2`, roomsTable)
	err := g.db.Select(&rooms, query, campusId, code+"%")

	return rooms, err
}

func NewGamePostgres(db *sqlx.DB) *GamePostgres {
	return &GamePostgres{db: db}
}
