package app

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/k-nox/impossiblerss/config"
	"github.com/k-nox/impossiblerss/storage"
	"github.com/k-nox/impossiblerss/storage/generated"

	"github.com/microcosm-cc/bluemonday"
	"github.com/mmcdole/gofeed"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type FeedList struct {
	ctx         context.Context
	db          storage.DB
	parser      *gofeed.Parser
	sanitizer   *bluemonday.Policy
	refreshRate time.Duration
	feeds       []*Feed
	refreshErrs chan error
}

func newFeedList(db storage.DB, cfg *config.Feeds) (*FeedList, error) {
	fl := &FeedList{
		db:          db,
		parser:      gofeed.NewParser(),
		sanitizer:   bluemonday.UGCPolicy(),
		refreshRate: time.Second * time.Duration(cfg.RefreshRate),
		refreshErrs: make(chan error),
	}

	feeds, err := fl.getFeedsFromDB()
	if err != nil {
		return nil, err
	}
	fl.feeds = feeds
	return fl, nil
}

func (fl *FeedList) AddFeed(feedURL string) (*Feed, error) {
	fetched, err := fl.parser.ParseURLWithContext(feedURL, fl.ctx)
	if err != nil {
		return nil, fmt.Errorf("error adding new feed %s: %w", feedURL, err)
	}
	feed := &Feed{
		Title:       fetched.Title,
		Description: fetched.Description,
		URL:         feedURL,
		Items:       map[string]Item{},
		mutex:       &sync.RWMutex{},
	}

	err = fl.db.CreateFeed(fl.ctx, generated.CreateFeedParams{
		Url:         feed.URL,
		Title:       storage.NewNullString(&feed.Title),
		Description: storage.NewNullString(&feed.Description),
	})
	if err != nil {
		return nil, fmt.Errorf("error saving new feed %s: %w", feedURL, err)
	}

	for _, fetchedItem := range fetched.Items {
		item := fl.transformItem(fetchedItem, feedURL)
		err := feed.addItem(fl.ctx, item, fl.db)
		if err != nil {
			runtime.EventsEmit(fl.ctx, string(RefreshError), err)
		}
	}

	go fl.setupRefresher(feed)

	return feed, nil
}

func (fl *FeedList) Feeds() []*Feed {
	return fl.feeds
}

func (fl *FeedList) getFeedsFromDB() ([]*Feed, error) {
	feedRows, err := fl.db.GetFeeds(fl.ctx)
	if err != nil {
		return nil, fmt.Errorf("error querying for all feeds: %w", err)
	}

	feeds := make([]*Feed, len(feedRows))
	for idx, row := range feedRows {
		items, err := fl.getItemsFromDB(row.Url)
		if err != nil {
			return nil, err
		}
		feeds[idx] = &Feed{
			Title:       row.Title.String,
			Description: row.Description.String,
			URL:         row.Url,
			Items:       items,
			mutex:       &sync.RWMutex{},
		}
	}
	return feeds, nil
}

func (fl *FeedList) getItemsFromDB(feedURL string) (map[string]Item, error) {
	itemRows, err := fl.db.GetItemsForFeed(fl.ctx, feedURL)
	if err != nil {
		return nil, fmt.Errorf("error querying for items for feed %s: %w", feedURL, err)
	}

	items := make(map[string]Item, len(itemRows))
	for _, row := range itemRows {
		items[row.Guid] = Item{
			GUID:  row.Guid,
			Title: row.Title.String,
			// Authors: []string,
			Content:       row.Content.String,
			Description:   row.Description.String,
			PublishedDate: &row.PublishedDate.Time,
			FeedURL:       feedURL,
		}
	}

	return items, nil
}

func (fl *FeedList) transformItem(raw *gofeed.Item, feedURL string) Item {
	content := raw.Content
	if content == "" {
		content = raw.Custom["content"]
	}
	content = fl.sanitizer.Sanitize(content)
	authors := make([]string, len(raw.Authors))
	for idx, author := range raw.Authors {
		authors[idx] = author.Name
	}

	return Item{
		GUID:          raw.GUID,
		Title:         raw.Title,
		Authors:       authors,
		Content:       content,
		Description:   raw.Description,
		PublishedDate: raw.PublishedParsed,
		FeedURL:       feedURL,
	}
}
