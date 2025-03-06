package app

import (
	"context"
	"impossiblerss/config"
	"impossiblerss/sqlite"
	"impossiblerss/storage"
)

// App struct
type App struct {
	ctx      context.Context
	FeedList *FeedList
}

// NewApp creates a new App application struct
func New(cfg *config.Config) (*App, error) {
	db, err := storage.New(&cfg.Database)
	if err != nil {
		return nil, err
	}
	err = db.Migrate(sqlite.Migrations)
	if err != nil {
		return nil, err
	}
	return &App{
		FeedList: newFeedList(db),
	}, nil
}

// Startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
	a.FeedList.ctx = ctx
}
