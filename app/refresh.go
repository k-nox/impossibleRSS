package app

import (
	"fmt"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func (fl *FeedList) setupRefreshers() {
	go func() {
		for {
			select {
			case err := <-fl.refreshErrs:
				runtime.EventsEmit(fl.ctx, string(RefreshError), err)
			case <-fl.ctx.Done():
				return
			}
		}
	}()
	for _, feed := range fl.feeds {
		// refresh once, then after every refreshRate tick
		fl.refreshOnce(feed)
		go fl.setupRefresher(feed)
	}
}

func (fl *FeedList) setupRefresher(feed *Feed) {
	ticker := time.Tick(fl.refreshRate)
	for range ticker {
		select {
		case <-fl.ctx.Done():
			return
		default:
			fl.refreshOnce(feed)
		}
	}
}

func (fl *FeedList) refreshOnce(feed *Feed) {
	url := feed.url()
	fetched, err := fl.parser.ParseURLWithContext(url, fl.ctx)
	if err != nil {
		fl.refreshErrs <- fmt.Errorf("error fetching feed %s: %w", url, err)
		return
	}

	for _, item := range fetched.Items {
		_, itemExists := feed.item(item.GUID)
		if !itemExists {
			newItem := fl.transformItem(item, url)
			err := feed.addItem(fl.ctx, newItem, fl.db)
			if err != nil {
				fl.refreshErrs <- fmt.Errorf("error adding new item with guid %s to feed %s: %w", newItem.GUID, url, err)
			}
			runtime.EventsEmit(fl.ctx, string(NewItem), newItem)
		}
	}
}
