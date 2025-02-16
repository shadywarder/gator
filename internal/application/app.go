package application

import (
	"database/sql"

	"github.com/shadywarder/gator/internal/config"
	"github.com/shadywarder/gator/internal/domain"
	"github.com/shadywarder/gator/internal/infrastructure/database"
	"github.com/shadywarder/gator/internal/infrastructure/handlers/feeds"
	"github.com/shadywarder/gator/internal/infrastructure/handlers/follows"
	"github.com/shadywarder/gator/internal/infrastructure/handlers/users"
	"github.com/shadywarder/gator/internal/infrastructure/handlers/util"
	"github.com/shadywarder/gator/internal/infrastructure/middleware"
)

// Application represents the main aggregator entity.
type Application struct {
	commands *domain.Commands
	state    *domain.State
}

// New instantiates a new Application entity.
func New(cfg *config.Config, db *sql.DB) *Application {
	dbQueries := database.New(db)

	state := domain.NewState(cfg, dbQueries)

	commands := domain.NewCommands()

	commands.Register("login", users.HandlerLogin)
	commands.Register("register", users.HandlerRegisterUser)
	commands.Register("reset", util.HandlerResetTable)
	commands.Register("users", users.HandlerUsers)
	commands.Register("agg", feeds.HandlerAggregate)
	commands.Register("addfeed", middleware.Login(feeds.HandlerAddFeed))
	commands.Register("feeds", middleware.Login(feeds.HandlerFeeds))
	commands.Register("follow", middleware.Login(follows.HandlerFollow))
	commands.Register("following", middleware.Login(follows.HandlerFollowing))
	commands.Register("unfollow", middleware.Login(follows.HandlerUnfollow))
	commands.Register("browse", feeds.HandlerBrowse)

	return &Application{
		commands: commands,
		state:    state,
	}
}

// Run executes command with respect to provided command name and its arguments.
func (a *Application) Run(cmdName string, cmdArgs []string) error {
	return a.commands.Call(a.state, domain.NewCommand(cmdName, cmdArgs))
}
