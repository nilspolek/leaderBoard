package transport

import (
	"github.com/google/uuid"
	"github.com/nilspolek/leaderBoard/repo"
)

type UserDTO struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (dto UserDTO) ToUser() repo.User {
	return repo.User{
		Name:     dto.Name,
		Password: dto.Password,
		Uuid:     uuid.New(),
	}
}

type GameDTO struct {
	Name  string `json:"name"`
	TeamA string `json:"teamA"`
	TeamB string `json:"teamB"`
}

func (dto GameDTO) ToGame() repo.Game {
	return repo.Game{
		Uuid:       uuid.New(),
		Name:       dto.Name,
		TeamA:      repo.Team{Name: dto.TeamA, Uuid: uuid.New()},
		TeamAScore: 0,
		TeamB:      repo.Team{Name: dto.TeamB, Uuid: uuid.New()},
		TeamBScore: 0,
	}
}

type TeamDTO struct {
	Name string `json:"name"`
}

func (dto TeamDTO) ToTeam() repo.Team {
	return repo.Team{
		Name: dto.Name,
		Uuid: uuid.New(),
	}
}
