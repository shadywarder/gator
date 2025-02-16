package main

import (
	"database/sql"
	"log/slog"
	"os"

	_ "github.com/lib/pq"

	"github.com/shadywarder/gator/internal/application"
	"github.com/shadywarder/gator/internal/config"
)

func main() {
	cfg, err := config.MustLoad()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	if len(os.Args) < 2 {
		slog.Error("not enough args!")
		os.Exit(1)
	}

	app := application.New(cfg, db)

	cmdName, cmdArgs := os.Args[1], os.Args[2:]

	if err := app.Run(cmdName, cmdArgs); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
