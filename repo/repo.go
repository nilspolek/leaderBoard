package repo

import "github.com/google/uuid"

type Repo interface {
	AddGame(game Game) error
	UpdateScore(gameId uuid.UUID, score Score) error
	GetScore(gameId uuid.UUID) (Score, error)
	AddTeam(team Team) error
	GetTeam(id uuid.UUID) (Team, error)
	AddUser(user User) error
	GetUser(id uuid.UUID) (User, error)
	IsUserValid(user User) (bool, error)
}

type Score struct {
	TeamAScore int
	TeamA      Team
	TeamBScore int
	TeamB      Team
}

type Team struct {
	Name string
	Uuid uuid.UUID
}

type Game struct {
	Uuid       uuid.UUID
	Name       string
	TeamA      Team
	TeamAScore int
	TeamB      Team
	TeamBScore int
}

type User struct {
	Name     string
	Password string
	Uuid     uuid.UUID
}
