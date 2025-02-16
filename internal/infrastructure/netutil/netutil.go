package netutil

import (
	"context"
	"encoding/xml"
	"html"
	"io"
	"net/http"

	"github.com/shadywarder/gator/internal/domain"
)

// FetchFeed fetches feeds using provided URL and format fetched data.
func FetchFeed(ctx context.Context, feedURL string) (*domain.RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, feedURL, http.NoBody)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var data domain.RSSFeed

	if err := xml.Unmarshal(body, &data); err != nil {
		return nil, err
	}

	data.Channel.Title = html.UnescapeString(data.Channel.Title)
	data.Channel.Description = html.UnescapeString(data.Channel.Description)

	for i, item := range data.Channel.Item {
		item.Title = html.UnescapeString(item.Title)
		item.Description = html.UnescapeString(item.Description)
		data.Channel.Item[i] = item
	}

	return &data, nil
}
