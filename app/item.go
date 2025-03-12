package app

import (
	"context"
	"fmt"
	"time"

	"github.com/k-nox/impossiblerss/storage"
	"github.com/k-nox/impossiblerss/storage/generated"
)

type Item struct {
	GUID          string     `json:"guid"`
	Title         string     `json:"title"`
	Authors       []string   `json:"authors"`
	Content       string     `json:"content"`
	Description   string     `json:"description"`
	PublishedDate *time.Time `json:"publishedDate" ts_type:"Date" ts_transform:"new Date(__VALUE__)"`
	FeedURL       string     `json:"feedURL"`
}

func (i Item) save(ctx context.Context, db storage.DB) error {
	err := db.CreateItem(ctx, generated.CreateItemParams{
		Guid:          i.GUID,
		Title:         storage.NewNullString(&i.Title),
		Content:       storage.NewNullString(&i.Content),
		Description:   storage.NewNullString(&i.Description),
		PublishedDate: storage.NewNullTime(i.PublishedDate),
		FeedUrl:       i.FeedURL,
	})
	if err != nil {
		return fmt.Errorf("error saving item with guid %s for feed %s: %w", i.GUID, i.FeedURL, err)
	}
	return nil
}
