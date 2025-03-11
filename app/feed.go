package app

import (
	"context"
	"impossiblerss/storage"
	"sync"
)

type Feed struct {
	Title       string          `json:"title"`
	Description string          `json:"description"`
	URL         string          `json:"link"`
	Items       map[string]Item `json:"items"`
	mutex       *sync.RWMutex
}

func (f *Feed) url() string {
	f.mutex.RLock()
	out := f.URL
	f.mutex.RUnlock()
	return out
}

func (f *Feed) item(guid string) (Item, bool) {
	f.mutex.RLock()
	item, exists := f.Items[guid]
	f.mutex.RUnlock()
	return item, exists
}

func (f *Feed) addItem(ctx context.Context, item Item, db storage.DB) error {
	err := item.save(ctx, db)
	if err != nil {
		return err
	}
	f.mutex.Lock()
	f.Items[item.GUID] = item
	f.mutex.Unlock()
	return nil
}
