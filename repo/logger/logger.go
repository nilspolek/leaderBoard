package logger

import (
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/nilspolek/leaderBoard/repo"
)

type Logger struct {
	next repo.Repo
}

func New(repo repo.Repo) repo.Repo {
	return Logger{
		next: repo,
	}
}

func (l Logger) GetTeam(id uuid.UUID) (repo.Team, error) {
	defer func(start time.Time) {
		log.Printf("GetTeam(%+v) took %+v", id, time.Since(start))
	}(time.Now())
	return l.next.GetTeam(id)
}

func (l Logger) GetScore(id uuid.UUID) (repo.Score, error) {
	defer func(start time.Time) {
		log.Printf("GetScore(%+v) took %+v", id, time.Since(start))
	}(time.Now())
	return l.next.GetScore(id)
}

func (l Logger) UpdateScore(id uuid.UUID, score repo.Score) error {
	defer func(start time.Time) {
		log.Printf("UpdateScore(%+v, %+v) took %+v", id, score, time.Since(start))
	}(time.Now())
	return l.next.UpdateScore(id, score)
}

func (l Logger) AddGame(game repo.Game) error {
	defer func(start time.Time) {
		log.Printf("AddGame(%+v) took %+v", game, time.Since(start))
	}(time.Now())
	return l.next.AddGame(game)
}

func (l Logger) AddUser(user repo.User) error {
	defer func(start time.Time) {
		log.Printf("AddUser(%+v) took %+v", user, time.Since(start))
	}(time.Now())
	return l.next.AddUser(user)
}

func (l Logger) GetUser(id uuid.UUID) (repo.User, error) {
	defer func(start time.Time) {
		log.Printf("GetUser(%+v) took %+v", id, time.Since(start))
	}(time.Now())
	return l.next.GetUser(id)
}

func (l Logger) AddTeam(team repo.Team) error {
	defer func(start time.Time) {
		log.Printf("AddTeam(%+v) took %+v", team, time.Since(start))
	}(time.Now())
	return l.next.AddTeam(team)
}

func (l Logger) IsUserValid(u repo.User) (bool, error) {
	defer func(start time.Time) {
		log.Printf("IsUserValid() took %+v", time.Since(start))
	}(time.Now())
	return l.next.IsUserValid(u)
}
