//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	//_ "github.com/tama-jp/rss/cmd/app"
	"github.com/tama-jp/rss/internal/adapter/controllers"
	"github.com/tama-jp/rss/internal/adapter/gateways/repository"
	"github.com/tama-jp/rss/internal/frameworks/config"
	"github.com/tama-jp/rss/internal/frameworks/database"
	"github.com/tama-jp/rss/internal/frameworks/logger"
	"github.com/tama-jp/rss/internal/frameworks/middleware"
	"github.com/tama-jp/rss/internal/frameworks/routing"
	"github.com/tama-jp/rss/internal/usecases/interactor"
	"github.com/tama-jp/rss/pkg/app"
)

var superSet = wire.NewSet(
	controllers.Set,
	interactor.Set,
	repository.Set,
	routing.Set,
	middleware.Set,
	app.NewApp,
)

func InitializeApp(conf *config.Config, dbConn *db.DataBase, r *gin.Engine, logger *logger.LogBase) app.App {
	wire.Build(superSet)
	return app.App{}
}
