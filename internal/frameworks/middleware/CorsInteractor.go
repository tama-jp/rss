package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/tama-jp/rss/internal/utils/message"
	"time"
)

type CorsInteractor struct {
}

func NewCorsInteractor() *CorsInteractor {
	return &CorsInteractor{}
}

func (interactor *CorsInteractor) CorsMiddleware() gin.HandlerFunc {
	return cors.New(cors.Config{
		// アクセスを許可したいアクセス元
		//AllowOrigins: []string{"*"},
		AllowOrigins: []string{
			"https://kodawari.club",
			"https://acs-pro.com",
			"http://localhost:3000",
			"http://localhost:5173",
			"http://localhost:9999",
			"http://localhost:8888",
			"http://localhost2:8888",

			"https://web.staging.acs-pro.work",
			"https://frontend.staging.acs-pro.work",
		},
		// アクセスを許可したいHTTPメソッド(以下の例だとPUTやDELETEはアクセスできません)
		AllowMethods: []string{
			"PUT",
			"PATCH",
			"DELETE",
			"POST",
			"GET",
			"OPTIONS",
		},
		// 許可したいHTTPリクエストヘッダ
		AllowHeaders: []string{
			"Access-Control-Allow-Credentials",
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"Authorization",
			"X-CSRF-Token",
			"Origin",
			message.X_USER_NAME,
			message.X_PASSWORD,
		},
		ExposeHeaders: []string{"Content-Length"},
		// cookieなどの情報を必要とするかどうか
		AllowCredentials: true,
		//AllowAllOrigins:  true,
		//AllowOriginFunc: func(origin string) bool {
		//	return origin == "*"
		//},
		// preflightリクエストの結果をキャッシュする時間
		MaxAge: 24 * time.Hour,
	})
}
