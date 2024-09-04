package app

import (
	"github.com/tama-jp/rss/internal/frameworks/config"
	db "github.com/tama-jp/rss/internal/frameworks/database"
	"github.com/tama-jp/rss/internal/frameworks/routing"
)

type App struct {
	conf *config.Config
	db   *db.DataBase
	r    *routing.Routing
}

func NewApp(conf *config.Config, db *db.DataBase, r *routing.Routing) App {
	return App{
		conf: conf,
		db:   db,
		r:    r,
	}
}

func (app *App) Start() {
	app.r.Setup()
	app.r.Run()
}
