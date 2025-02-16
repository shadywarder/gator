package domain

import (
	"github.com/shadywarder/gator/internal/config"
	"github.com/shadywarder/gator/internal/infrastructure/database"
)

// State allows to conveniently perform operation over database.
type State struct {
	Cfg *config.Config
	DB  *database.Queries
}

// NewState instantiates a new State entity.
func NewState(cfg *config.Config, db *database.Queries) *State {
	return &State{
		Cfg: cfg,
		DB:  db,
	}
}
