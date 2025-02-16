package users

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/shadywarder/gator/internal/domain"
	"github.com/shadywarder/gator/internal/infrastructure/database"
)

// HandlerLogin sets the provided user as a current one.
func HandlerLogin(s *domain.State, cmd *domain.Command) error {
	if len(cmd.Args) != 1 {
		return ErrInvalidUserName
	}

	name := cmd.Args[0]

	_, err := s.DB.GetUserByName(context.Background(), name)
	if err != nil {
		return ErrUserExistance
	}

	if err := s.Cfg.SetUser(name); err != nil {
		return err
	}

	log.Printf("user %s has been set!", s.Cfg.CurrentUserName)

	return nil
}

// HandlerRegisterUser register a new user and sets him as a current one.
func HandlerRegisterUser(s *domain.State, cmd *domain.Command) error {
	if len(cmd.Args) != 1 {
		return ErrInvalidUserName
	}

	name := cmd.Args[0]

	user, err := s.DB.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
	})
	if err != nil {
		return ErrUserCreation
	}

	if err := s.Cfg.SetUser(user.Name); err != nil {
		return err
	}

	fmt.Println("user created successfully!")
	render(&user)

	return nil
}

// HandlerUsers prints all information about users.
func HandlerUsers(s *domain.State, _ *domain.Command) error {
	users, err := s.DB.SelectUsers(context.Background())
	if err != nil {
		return ErrRetrieveUser
	}

	for _, user := range users {
		line := fmt.Sprintf("* %v", user)
		if user == s.Cfg.CurrentUserName {
			line += " (current)"
		}

		line += "\n"
		fmt.Print(line)
	}

	return nil
}
