package app

import (
	"context"
	"sync"

	"github.com/k-nox/impossiblerss/storage"
)

type Feed struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	URL         string  `json:"link"`
	Items       []*Item `json:"items"`
	mutex       *sync.RWMutex
	itemGuids   map[string]int
}

func (f *Feed) url() string {
	f.mutex.RLock()
	out := f.URL
	f.mutex.RUnlock()
	return out
}

func (f *Feed) item(guid string) (*Item, bool) {
	var out *Item
	f.mutex.RLock()
	idx, exists := f.itemGuids[guid]
	if exists {
		out = f.Items[idx]
	}
	f.mutex.RUnlock()
	return out, exists
}

func (f *Feed) addItem(ctx context.Context, item *Item, db storage.DB) error {
	err := item.save(ctx, db)
	if err != nil {
		return err
	}
	f.mutex.Lock()
	f.Items = append(f.Items, item)
	f.itemGuids[item.GUID] = len(f.Items) - 1
	f.mutex.Unlock()
	return nil
}
