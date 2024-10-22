package main

import (
	"github.com/nilspolek/goLog"
	"github.com/nilspolek/leaderBoard/repo/logger"
	"github.com/nilspolek/leaderBoard/repo/sqlite"
	"github.com/nilspolek/leaderBoard/transport"
)

func main() {
	repo, err := sqlite.New("./db.sqlite")
	if err != nil {
		goLog.Error("Error: %v", err)
		return
	}
	goLog.Info("Connected to database")
	repo = logger.New(repo)
	goLog.Info("Logger initialized")
	server := transport.New(repo, transport.ServerConfig{Addr: ":8080"})
	server.SetupRoutes()
	goLog.Info("Server started on port %s", server.Config.Addr)
	server.ServeHTTP()
}
