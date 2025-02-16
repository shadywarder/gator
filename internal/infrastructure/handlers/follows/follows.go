package follows

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/shadywarder/gator/internal/domain"
	"github.com/shadywarder/gator/internal/infrastructure/database"
)

// HandlerFollow subsribes user to a provided feed.
func HandlerFollow(s *domain.State, cmd *domain.Command, user *database.User) error {
	if len(cmd.Args) != 1 {
		return ErrFollowInvalidArgs
	}

	url := cmd.Args[0]

	feed, err := s.DB.GetFeedByURL(context.Background(), url)
	if err != nil {
		return ErrURLFeedRetrieve
	}

	_, err = s.DB.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return ErrFollowCreation
	}

	renderFeedFollows(user, &feed)

	return nil
}

// HandlerFollowing prints user's subscriptions.
func HandlerFollowing(s *domain.State, cmd *domain.Command, user *database.User) error {
	if len(cmd.Args) != 0 {
		return ErrFollowingInvalidArgs
	}

	follows, err := s.DB.GetFeedFollowsForUser(context.Background(), user.Name)
	if err != nil {
		return ErrRetrieveFeedFollows
	}

	fmt.Printf("%s follows the ensuing feeds:\n", user.Name)

	for _, follow := range follows {
		fmt.Printf("* %s\n", follow.FeedName)
	}

	return nil
}

// HandlerUnfollow unsubscribe user from a provided feed.
func HandlerUnfollow(s *domain.State, cmd *domain.Command, user *database.User) error {
	if len(cmd.Args) != 1 {
		return ErrUnfollowInvalidArgs
	}

	url := cmd.Args[0]

	if err := s.DB.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		Url:    url,
	}); err != nil {
		return ErrDeleteFeedFollow
	}

	return nil
}
