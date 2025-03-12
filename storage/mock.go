package storage

import (
	"context"
	"embed"
	"errors"
	"slices"

	"github.com/k-nox/impossiblerss/storage/generated"

	"modernc.org/sqlite"
	sqlite3 "modernc.org/sqlite/lib"
)

type mockDB struct {
	mockQuerier
}

func (m *mockDB) Migrate(migrationFiles embed.FS) error {
	return nil
}

type mockQuerier struct {
	feeds []generated.Feed
	items []generated.Item
}

var _ generated.Querier = (*mockQuerier)(nil)

// CreateFeed implements generated.Querier.
func (m *mockQuerier) CreateFeed(ctx context.Context, arg generated.CreateFeedParams) error {
	m.feeds = append(m.feeds, generated.Feed{
		Url:         arg.Url,
		Title:       arg.Title,
		Description: arg.Description,
	})
	return nil
}

// CreateItem implements generated.Querier.
func (m *mockQuerier) CreateItem(ctx context.Context, arg generated.CreateItemParams) error {
	isValidFeed := slices.ContainsFunc(m.feeds, func(f generated.Feed) bool {
		return f.Url == arg.FeedUrl
	})
	if !isValidFeed {
		return errors.New(sqlite.ErrorCodeString[sqlite3.SQLITE_CONSTRAINT])
	}
	m.items = append(m.items, generated.Item(arg))
	return nil
}

// GetFeeds implements generated.Querier.
func (m *mockQuerier) GetFeeds(ctx context.Context) ([]generated.Feed, error) {
	return m.feeds, nil
}

// GetItemsForFeed implements generated.Querier.
func (m *mockQuerier) GetItemsForFeed(ctx context.Context, feedUrl string) ([]generated.GetItemsForFeedRow, error) {
	out := []generated.GetItemsForFeedRow{}
	for _, item := range m.items {
		if item.FeedUrl == feedUrl {
			out = append(out, generated.GetItemsForFeedRow{
				Guid:          item.Guid,
				Title:         item.Title,
				Authors:       item.Authors,
				Content:       item.Content,
				Description:   item.Description,
				PublishedDate: item.PublishedDate,
			})
		}
	}
	return out, nil
}
