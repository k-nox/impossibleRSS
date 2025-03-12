package app

import (
	"context"

	"github.com/k-nox/impossiblerss/config"
	"github.com/k-nox/impossiblerss/sqlite"
	"github.com/k-nox/impossiblerss/storage"
)

// App struct
type App struct {
	ctx      context.Context
	FeedList *FeedList
}

// NewApp creates a new App application struct
func New(cfg *config.Config) (*App, error) {
	db, err := storage.New(cfg.Database)
	if err != nil {
		return nil, err
	}
	err = db.Migrate(sqlite.Migrations)
	if err != nil {
		return nil, err
	}
	feedList, err := newFeedList(db, cfg.Feeds)
	if err != nil {
		return nil, err
	}
	return &App{
		FeedList: feedList,
	}, nil
}

// Startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
	a.FeedList.ctx = ctx
}

func (a *App) OnDOMReady(ctx context.Context) {
	a.FeedList.setupRefreshers()
}
