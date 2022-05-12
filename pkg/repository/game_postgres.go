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

func (g *GamePostgres) GetRoomById(roomId int) (model.Room, error) {
	var room model.Room
	query := fmt.Sprintf(`SELECT * FROM %s rm WHERE rm.id = $1`, roomsTable)
	err := g.db.Get(&room, query, roomId)

	return room, err
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
