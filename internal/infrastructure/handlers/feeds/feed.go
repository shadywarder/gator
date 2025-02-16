package feeds

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/shadywarder/gator/internal/domain"
	"github.com/shadywarder/gator/internal/infrastructure/database"
)

// HandlerAggregate fetches feeds every timeBetweenReqs time interval.
func HandlerAggregate(s *domain.State, cmd *domain.Command) error {
	if len(cmd.Args) != 1 {
		return ErrAggIntervalAbsence
	}

	timeBetweenReqs := cmd.Args[0]

	duration, err := time.ParseDuration(timeBetweenReqs)
	if err != nil {
		return ErrAggInvalidTimeFormat
	}

	log.Printf("collecting feeds every %v", duration)
	ticker := time.NewTicker(duration)

	for ; ; <-ticker.C {
		if err := scrapeFeeds(s); err != nil {
			return err
		}
	}
}

// HandlerAddFeed creates a new feed.
func HandlerAddFeed(s *domain.State, cmd *domain.Command, user *database.User) error {
	if len(cmd.Args) != 2 {
		return ErrAddFeedInvalidArgs
	}

	name, url := cmd.Args[0], cmd.Args[1]

	feed, err := s.DB.CreateFeed(context.Background(),
		database.CreateFeedParams{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			Name:      name,
			Url:       url,
			UserID:    user.ID,
		},
	)
	if err != nil {
		return ErrFeedCreation
	}

	_, err = s.DB.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return ErrFeedFollow
	}

	fmt.Printf("%+v\n", user)
	fmt.Printf("%+v\n", feed)

	return nil
}

// HandlerFeeds all information about feeds.
func HandlerFeeds(s *domain.State, _ *domain.Command, _ *database.User) error {
	feeds, err := s.DB.GetFeeds(context.Background())
	if err != nil {
		return ErrRetrieveFeeds
	}

	if len(feeds) == 0 {
		return ErrEmptyFeeds
	}

	log.Printf("found %v feeds:\n", len(feeds))

	for i := range len(feeds) {
		user, err := s.DB.GetUsedByID(context.Background(), feeds[i].UserID)
		if err != nil {
			return ErrUnknownUser
		}

		renderFeed(&feeds[i], user.Name)
	}

	return nil
}

// HandlerBrowse prints max(2, limit) recent feeds.
func HandlerBrowse(s *domain.State, cmd *domain.Command) error {
	if len(cmd.Args) > 1 {
		return ErrInvalidBrowseArgs
	}

	limit := 2

	if len(cmd.Args) != 0 {
		value, err := strconv.Atoi(cmd.Args[0])
		if err != nil {
			return ErrInvalidBrowseArgs
		}

		limit = value
	}

	posts, err := s.DB.GetPosts(context.Background())
	if err != nil {
		return err
	}

	for i := 0; i < min(limit, len(posts)); i++ {
		post := posts[i]
		renderPost(i+1, &post)
	}

	return nil
}
