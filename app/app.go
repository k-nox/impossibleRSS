package app

import (
	"context"
	"impossiblerss/storage"
)

// App struct
type App struct {
	ctx      context.Context
	FeedList *FeedList
}

// NewApp creates a new App application struct
func New(db storage.DB) *App {
	return &App{
		FeedList: newFeedList(db),
	}
}

// Startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
	a.FeedList.ctx = ctx
}
