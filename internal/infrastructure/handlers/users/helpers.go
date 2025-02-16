package users

import (
	"fmt"

	"github.com/shadywarder/gator/internal/infrastructure/database"
)

func render(user *database.User) {
	fmt.Printf("id: %s\n", user.ID)
	fmt.Printf("name: %s\n", user.Name)
}
