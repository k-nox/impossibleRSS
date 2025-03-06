package main

import (
	"embed"
	"impossiblerss/app"
	"impossiblerss/storage"
	"log"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/build
var assets embed.FS

//go:embed sqlite/migrations/*
var dbMigrations embed.FS

func main() {
	db, err := storage.New(":memory:", false)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Migrate(dbMigrations)
	if err != nil {
		log.Fatal(err)
	}

	app := app.New(db)

	// Create application with options
	err = wails.Run(&options.App{
		Title:  "impossiblerss",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.Startup,
		Bind: []interface{}{
			app.FeedList,
		},
	})
	if err != nil {
		println("Error:", err.Error())
	}
}
