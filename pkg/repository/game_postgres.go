package repository

import (
	"Run_Hse_Run/pkg/model"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type GamePostgres struct {
	db *sqlx.DB
}

func (g *GamePostgres) GetRoomByCodePattern(code string) ([]model.Room, error) {
	var rooms []model.Room

	query := fmt.Sprintf(`SELECT * FROM %s rm WHERE rm.code LIKE $1`, roomsTable)
	err := g.db.Select(&rooms, query, code+"%")

	return rooms, err
}

func NewGamePostgres(db *sqlx.DB) *GamePostgres {
	return &GamePostgres{db: db}
}
