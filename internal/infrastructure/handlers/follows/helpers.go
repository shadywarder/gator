package follows

import (
	"fmt"

	"github.com/shadywarder/gator/internal/infrastructure/database"
)

func renderFeedFollows(user *database.User, feed *database.Feed) {
	fmt.Printf("* feed name: %v\n", feed.Name)
	fmt.Printf("* current user: %v\n", user.Name)
}
