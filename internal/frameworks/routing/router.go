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
	//systemAttendanceHolidayCtrl *controllers.SystemAttendanceHolidayController
	//attendanceCtrl              *controllers.AttendanceController
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
	//systemAttendanceHolidayCtrl *controllers.SystemAttendanceHolidayController,
	//attendanceCtrl *controllers.AttendanceController,

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
		//systemAttendanceHolidayCtrl: systemAttendanceHolidayCtrl,
		//attendanceCtrl:              attendanceCtrl,
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

	//attendance := v1.Group("attendance")

	////勤怠区分取得
	//// `GET /attendance/classification
	//attendance.GET("/classification", r.attendanceCtrl.GetClassification)
	//
	////勤怠詳細取得
	//// `GET /attendance/day/{user_id}/{year}/{month}/{day}`
	//attendance.GET("/day/:user_id/:year/:month/:day", r.attendanceCtrl.GetDay)
	//
	//// 勤怠詳細登録
	//// `POST /attendance/day/{user_id}/{year}/{month}/{day}`
	//attendance.POST("/day/:user_id/:year/:month/:day", r.attendanceCtrl.PostDay)
	//
	//// 勤怠詳細更新
	//// `PUT /attendance/day/{user_id}/{year}/{month}/{day}`
	//attendance.PUT("/day/:user_id/:year/:month/:day", r.attendanceCtrl.PutDay)
	//
	//// 勤怠詳細削除
	//// `DELETE /attendance/day/{user_id}/{year}/{month}/{day}`
	//
	//attendance.DELETE("/day/:user_id/:year/:month/:day", r.attendanceCtrl.DeleteDay)
	//
	//// 勤怠情報勤怠情報初期化
	//// `POST /attendance/day/initialize/{user_id}/{year}/{month}/{day}`
	//attendance.POST("/day/initialize/:user_id/:year/:month/:day", r.attendanceCtrl.PostInitializeDay)
	//
	//// 勤怠情報一覧
	//// `GET /attendance/header/month/list/{user_id}`
	//
	//attendance.GET("/header/month/list/:userid", r.attendanceCtrl.GetHeaderMonthList)
	//
	//// 1ヶ月勤怠情報勤怠情報取得
	//// `GET /attendance/month/{user_id}/{year}/{month}`
	//attendance.GET("/month/:user_id/:year/:month", r.attendanceCtrl.GetMonth)
	//
	//// 1ヶ月勤怠情報勤怠情報初期化
	//// `POST /attendance/month/initialize/{user_id}/{year}/{month}`
	//attendance.POST("/month/initialize/:user_id/:year/:month", r.attendanceCtrl.PostInitializeMonth)
	//
	//// 1ヶ月勤怠情報PDF出力
	//// `GET /attendance/month_pdf/{user_id}/{year}/{month}`
	//
	//attendance.GET("/month_pdf/:user_id/:year/:month", r.attendanceCtrl.GetMonthPdf)

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

	//systemAttendance := system.Group("attendance")

	//// 休日(年)一覧取得
	//// `GET /system/attendance/holiday/year/list`
	//systemAttendance.GET("/holiday/year/list", r.systemAttendanceHolidayCtrl.GetHolidayAllYearList)
	//
	//// 休日(公休/指定休日)一覧取得
	//// `GET /system/attendance/holiday/list`
	//
	//systemAttendance.GET("/holiday/list", r.systemAttendanceHolidayCtrl.GetHolidayList)
	//
	//// 休日登録
	//// `POST /system/attendance/holiday`
	//systemAttendance.POST("/holiday", r.systemAttendanceHolidayCtrl.PostHoliday)
	//
	//systemAttendanceHoliday := systemAttendance.Group("holiday")
	//
	////// 休日取得
	//// `GET /system/attendance/holiday`
	//
	//systemAttendanceHoliday.GET("/:holiday_id", r.systemAttendanceHolidayCtrl.GetHoliday)
	//
	//// 休日更新
	//// `PUT /system/attendance/holiday`
	//
	//systemAttendanceHoliday.PUT("/:holiday_id", r.systemAttendanceHolidayCtrl.PutHoliday)
	//
	//// 休日削除
	//// `DELETE /system/attendance/holiday`
	//
	//systemAttendanceHoliday.DELETE("/:holiday_id", r.systemAttendanceHolidayCtrl.DeleteHoliday)

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
