package feeds

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/shadywarder/gator/internal/domain"
	"github.com/shadywarder/gator/internal/infrastructure/database"
	"github.com/shadywarder/gator/internal/infrastructure/netutil"
)

// scrapeFeeds scrape posts from certain feeds.
func scrapeFeeds(s *domain.State) error {
	feed, err := s.DB.GetNextFeedToFetch(context.Background())
	if err != nil {
		return err
	}

	if err := s.DB.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{
		LastFetchedAt: time.Now().UTC(),
		UpdatedAt:     time.Now().UTC(),
		ID:            feed.ID,
	}); err != nil {
		return err
	}

	fetched, err := netutil.FetchFeed(context.Background(), feed.Url)
	if err != nil {
		return err
	}

	for _, item := range fetched.Channel.Item {
		published, err := parsePublishedAt(item.PubDate)
		if err != nil {
			return err
		}

		_, err = s.DB.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   time.Now().UTC(),
			UpdatedAt:   time.Now().UTC(),
			Title:       item.Title,
			Url:         item.Link,
			Description: item.Description,
			PublishedAt: published,
			FeedID:      feed.ID,
		})

		if err != nil {
			var pqErr *pq.Error
			if errors.As(err, &pqErr) && pqErr.Code == "23505" {
				continue
			}

			return err
		}
	}

	return nil
}

// parsePublishedAt parses post's publication date.
func parsePublishedAt(pubDate string) (time.Time, error) {
	formats := []string{
		time.RFC1123,
		time.RFC1123Z,
		time.RFC822,
		time.RFC822Z,
		time.RFC3339,
	}

	for _, layout := range formats {
		t, err := time.Parse(layout, pubDate)
		if err == nil {
			return t.UTC(), nil
		}
	}

	return time.Time{}, ErrRSSWrongTimeFormat
}

// renderFeed renders feed.
func renderFeed(feed *database.Feed, username string) {
	fmt.Printf("* feed id: %v\n", feed.ID)
	fmt.Printf("* created_at: %v\n", feed.CreatedAt)
	fmt.Printf("* updated_at: %v\n", feed.UpdatedAt)
	fmt.Printf("* name: %v\n", feed.Name)
	fmt.Printf("* url: %v\n", feed.Url)
	fmt.Printf("* username: %v\n", username)
	fmt.Println("==============================")
}

// renderPost renders post.
func renderPost(index int, post *database.Post) {
	fmt.Printf("* item %v\n", index)
	fmt.Printf("* title: %v\n", post.Title)
	fmt.Printf("* url: %v\n", post.Url)
	fmt.Printf("* description: %v", post.Description)
	fmt.Printf("* pub. date: %v\n", post.PublishedAt)
	fmt.Println("==============================")
}
