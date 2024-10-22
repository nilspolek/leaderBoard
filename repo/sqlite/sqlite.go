package sqlite

import (
	"database/sql"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"github.com/nilspolek/leaderBoard/repo"
)

type SQLite struct {
	db *sql.DB
}

func New(path string) (repo.Repo, error) {
	var (
		out SQLite
		err error
	)
	out.db, err = sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}
	_, err = out.db.Exec("PRAGMA foreign_keys = ON")
	if err != nil {
		return nil, err
	}
	_, err = out.db.Exec(`
			CREATE TABLE IF NOT EXISTS users (
				id TEXT PRIMARY KEY,
				name TEXT NOT NULL,
				password TEXT NOT NULL
			);
			CREATE TABLE IF NOT EXISTS teams (
				id TEXT PRIMARY KEY,
				name TEXT NOT NULL
			);
			CREATE TABLE IF NOT EXISTS games (
				id TEXT PRIMARY KEY,
				team_a_id TEXT NOT NULL,
				team_a_score INTEGER NOT NULL,
				team_b_id TEXT NOT NULL,
				team_b_score INTEGER NOT NULL,
				name TEXT NOT NULL,
				FOREIGN KEY(team_a_id) REFERENCES teams(id),
				FOREIGN KEY(team_b_id) REFERENCES teams(id)
			);
		`)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (s SQLite) AddUser(user repo.User) error {
	_, err := s.db.Exec("INSERT INTO users (id, name, password) VALUES (?, ?, ?)", user.Uuid, user.Name, user.Password)
	return err
}

func (s SQLite) IsUserValid(user repo.User) (bool, error) {
	var (
		out bool
	)
	err := s.db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE name = ? AND password = ?)", user.Name, user.Password).Scan(&out)
	if err != nil {
		return false, err
	}
	return out, err
}
func (s SQLite) GetUser(id uuid.UUID) (repo.User, error) {
	var (
		out repo.User
	)
	err := s.db.QueryRow("SELECT name, password FROM users WHERE id = ?", id).Scan(&out.Name, &out.Password)
	out.Uuid = id
	return out, err
}

func (s SQLite) UpdateScore(gameId uuid.UUID, score repo.Score) error {
	_, err := s.db.Exec("UPDATE game SET team_b_score = ?, team_a_score = ? WHERE id = ?", score.TeamBScore, score.TeamAScore, gameId)
	return err
}

func (s SQLite) AddGame(game repo.Game) error {
	_, err := s.db.Exec("INSERT INTO games (id, team_a_id, team_a_score, team_b_id, team_b_score, name) VALUES (?, ?, ?, ?, ?, ?)", game.Uuid, game.TeamA.Uuid, game.TeamAScore, game.TeamB.Uuid, game.TeamBScore, game.Name)
	return err
}

func (s SQLite) AddTeam(team repo.Team) error {
	_, err := s.db.Exec("INSERT INTO teams (id, name) VALUES (?, ?)", team.Uuid, team.Name)
	return err
}

func (s SQLite) GetScore(gameId uuid.UUID) (repo.Score, error) {
	var (
		out repo.Score
	)
	err := s.db.QueryRow("SELECT team_a_score, team_b_score FROM games WHERE id = ?", gameId).Scan(&out.TeamAScore, &out.TeamBScore)
	return out, err
}

func (s SQLite) GetTeam(id uuid.UUID) (repo.Team, error) {
	var (
		out repo.Team
	)
	err := s.db.QueryRow("SELECT name FROM teams WHERE id = ?", id).Scan(&out.Name)
	out.Uuid = id
	return out, err
}

func (s SQLite) Close() error {
	return s.db.Close()
}
