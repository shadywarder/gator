package util

import (
	"context"
	"log"

	"github.com/shadywarder/gator/internal/domain"
)

// HandlerResetTable deletes all users.
func HandlerResetTable(s *domain.State, _ *domain.Command) error {
	if err := s.DB.ReseteTable(context.Background()); err != nil {
		return ErrTableReset
	}

	log.Println("database reset successfully!")

	return nil
}
