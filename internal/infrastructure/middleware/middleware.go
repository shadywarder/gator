package middleware

import (
	"context"

	"github.com/shadywarder/gator/internal/domain"
	"github.com/shadywarder/gator/internal/infrastructure/database"
)

// Login that implements login logic.
func Login(handler func(s *domain.State,
	cmd *domain.Command, user *database.User) error) func(*domain.State, *domain.Command) error {
	return func(s *domain.State, cmd *domain.Command) error {
		user, err := s.DB.GetUserByName(context.Background(), s.Cfg.CurrentUserName)
		if err != nil {
			return ErrReceiveUser
		}

		return handler(s, cmd, &user)
	}
}
