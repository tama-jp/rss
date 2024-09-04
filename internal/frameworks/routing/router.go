package routing

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"github.com/tama-jp/rss/internal/adapter/controllers"
	"github.com/tama-jp/rss/internal/domain/response"
	"github.com/tama-jp/rss/internal/frameworks/config"
	db "github.com/tama-jp/rss/internal/frameworks/database"
	"github.com/tama-jp/rss/internal/frameworks/logger"
	"github.com/tama-jp/rss/internal/frameworks/middleware"
	"github.com/tama-jp/rss/internal/utils/message"
	"net/http"
	"strconv"
)

var Set = wire.NewSet(
	NewRouting,
)

type Routing struct {
	r      *gin.Engine
	Config *config.Config
	DB     *db.DataBase
	Gin    *gin.Engine
	Logger *logger.LogBase

	loginCtrl      *controllers.AccessTokenController
	jwtInteractor  *middleware.JwtInteractor
	corsInteractor *middleware.CorsInteractor

	userCtrl       *controllers.UserController
	systemUserCtrl *controllers.SystemUserController
}

func NewRouting(
	config *config.Config,
	db *db.DataBase,
	r *gin.Engine,
	logger *logger.LogBase,

	loginCtrl *controllers.AccessTokenController,
	jwtInteractor *middleware.JwtInteractor,
	corsInteractor *middleware.CorsInteractor,

	userCtrl *controllers.UserController,
	systemUserCtrl *controllers.SystemUserController,

) *Routing {

	return &Routing{
		r:      r,
		Config: config,
		DB:     db,
		Gin:    gin.Default(),
		Logger: logger,

		loginCtrl:      loginCtrl,
		jwtInteractor:  jwtInteractor,
		corsInteractor: corsInteractor,

		userCtrl:       userCtrl,
		systemUserCtrl: systemUserCtrl,
	}
}

func (r *Routing) Run() {

	port := r.Config.Rooting.Port

	fmt.Println("routing port:" + strconv.Itoa(port))
	r.Gin.Run(":" + strconv.Itoa(port))
}

func (r *Routing) Setup() {

	r.Gin.Use(r.corsInteractor.CorsMiddleware())

	r.Gin.GET("/ping", func(cc *gin.Context) {
		cc.JSON(200, gin.H{
			"message": "pong",
		})
	})

	v1 := r.Gin.Group("api/v1")

	auth := v1.Group("auth")

	// アクセストークン取得
	//	`GET /auth/access_token`
	auth.GET("/access_token", r.loginCtrl.GetAccessToken)

	auth.Use(r.jwtInteractor.JwtMiddleware())

	// アクセストークン再取得
	// `PUT /auth/access_token`
	auth.PUT("/access_token", r.loginCtrl.PutAccessToken)

	// アクセストークン削除
	// `DELETE /auth/access_token`
	auth.DELETE("/access_token", r.loginCtrl.DeleteAccessToken)

	user := v1.Group("user")

	// ユーザー情報取得
	// `GET /user/info`
	user.GET("/info/:user_id", r.userCtrl.Search)

	// パスワード変更
	// `PUT /user/password_change`
	user.PUT("/password_change", r.userCtrl.PasswordChange)

	system := v1.Group("system")

	// ユーザ情報登録
	// `POST /system/user`

	system.POST("/user", r.systemUserCtrl.PostUserInfo)

	systemUser := system.Group("user")

	// ユーザ情報取得
	// `GET /system/user`

	systemUser.GET("/:user_id", r.systemUserCtrl.SearchUserInfo)

	// ユーザ情報変更
	// `PUT /system/user`

	systemUser.PUT("/:user_id", r.systemUserCtrl.PutUserInfo)

	// ユーザ情報削除
	// `DELETE /system/user`
	systemUser.DELETE("/:user_id", r.systemUserCtrl.DeleteUserInfo)

	// ユーザ一覧取得
	// `GET /system/user/list`

	systemUser.GET("/list", r.systemUserCtrl.GetUserList)

	systemUserRole := system.Group("user_role")

	// ユーザー権限リスト
	// `GET /system/user_role/list`
	systemUserRole.GET("/list", r.systemUserCtrl.GetUserRoleList)

	// 静的ファイルを /static 配下で提供
	r.Gin.Static("/static", "./static")

	// SvelteKitなどのビルドファイルを提供
	r.Gin.Static("/_app", "./static/_app")

	// favicon.ico を提供
	r.Gin.StaticFile("/favicon.ico", "./static/favicon.ico")

	// ルートパス "/" で index.html を返す
	r.Gin.GET("/", func(c *gin.Context) {
		c.File("./static/index.html")
	})

	r.Gin.NoRoute(func(c *gin.Context) {
		messageStr := message.GetMsg(message.ERR006)
		res := response.Response{C: c}
		res.ErrorPayload(http.StatusBadRequest, message.ERR006, errors.New(messageStr))
		c.Abort()
	})

	r.Gin.NoMethod(func(c *gin.Context) {
		messageStr := message.GetMsg(message.ERR006)
		res := response.Response{C: c}
		res.ErrorPayload(http.StatusMethodNotAllowed, message.ERR006, errors.New(messageStr))
		c.Abort()
	})

}
