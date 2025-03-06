package app

import (
	"context"
	"database/sql"
	"fmt"
	"impossiblerss/storage"
	"impossiblerss/storage/generated"
	"time"

	"github.com/microcosm-cc/bluemonday"
	"github.com/mmcdole/gofeed"
)

type Feed struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"link"`
}

type Item struct {
	GUID          string    `json:"guid"`
	Title         string    `json:"title"`
	Authors       []string  `json:"authors"`
	Content       string    `json:"content"`
	Description   string    `json:"description"`
	PublishedDate time.Time `json:"publishedDate" ts_type:"Date" ts_transform:"new Date(__VALUE__)"`
	FeedURL       string    `json:"feedURL"`
}

type FeedList struct {
	ctx       context.Context
	db        *storage.DB
	parser    *gofeed.Parser
	sanitizer *bluemonday.Policy
}

func newFeedList(db *storage.DB) *FeedList {
	return &FeedList{
		db:        db,
		parser:    gofeed.NewParser(),
		sanitizer: bluemonday.UGCPolicy(),
	}
}

func (fl *FeedList) GetFeeds() ([]Feed, error) {
	feedRows, err := fl.db.GetFeeds(fl.ctx)
	if err != nil {
		return nil, fmt.Errorf("error querying for all feeds: %w", err)
	}

	feeds := make([]Feed, len(feedRows))
	for idx, row := range feedRows {
		feeds[idx] = Feed{
			Title:       row.Title.String,
			Description: row.Description.String,
			URL:         row.Url,
		}
	}
	return feeds, nil
}

func (fl *FeedList) GetItems(feedURL string) ([]Item, error) {
	itemRows, err := fl.db.GetItemsForFeed(fl.ctx, feedURL)
	if err != nil {
		return nil, fmt.Errorf("error querying for items for feed %s: %w", feedURL, err)
	}

	items := make([]Item, len(itemRows))
	for idx, row := range itemRows {
		items[idx] = Item{
			GUID:  row.Guid,
			Title: row.Title.String,
			// Authors: []string,
			Content:       row.Content.String,
			Description:   row.Description.String,
			PublishedDate: row.PublishedDate.Time,
			FeedURL:       feedURL,
		}
	}

	return items, nil
}

func (fl *FeedList) RefreshFeed(feedURL string) error {
	feed, err := fl.parser.ParseURLWithContext(feedURL, fl.ctx)
	if err != nil {
		return fmt.Errorf("error parsing feed: %w", err)
	}

	fl.saveItems(feed.Items, feedURL)

	return nil
}

func (fl *FeedList) AddFeed(feedURL string) error {
	feed, err := fl.parser.ParseURLWithContext(feedURL, fl.ctx)
	if err != nil {
		return fmt.Errorf("error parsing feed: %w", err)
	}

	err = fl.db.CreateFeed(fl.ctx, generated.CreateFeedParams{
		Url:         feedURL,
		Title:       nullStringFromPtr(&feed.Title),
		Description: nullStringFromPtr(&feed.Description),
	})
	if err != nil {
		return fmt.Errorf("error saving feed: %w", err)
	}

	fl.saveItems(feed.Items, feedURL)

	return nil
}

func (fl *FeedList) saveItems(items []*gofeed.Item, feedURL string) {
	for _, item := range items {
		err := fl.db.CreateItem(fl.ctx, generated.CreateItemParams{
			Guid:  item.GUID,
			Title: nullStringFromPtr(&item.Title),
			// Authors: sql.NullString,
			Description:   nullStringFromPtr(&item.Description),
			Content:       nullStringFromPtr(&item.Content),
			PublishedDate: nullTimeFromPtr(item.PublishedParsed),
			FeedUrl:       feedURL,
		})
		if err != nil {
			// TODO: handle
			panic(err)
		}
	}
}

func nullStringFromPtr(s *string) sql.NullString {
	if s == nil || *s == "" {
		return sql.NullString{}
	}

	return sql.NullString{
		String: *s,
		Valid:  true,
	}
}

func nullTimeFromPtr(t *time.Time) sql.NullTime {
	if t == nil || t.IsZero() {
		return sql.NullTime{}
	}

	return sql.NullTime{
		Time:  *t,
		Valid: true,
	}
}
