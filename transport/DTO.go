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
