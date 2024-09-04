package middleware

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tama-jp/rss/internal/domain/response"
	"github.com/tama-jp/rss/internal/usecases/port"
	"github.com/tama-jp/rss/internal/utils/message"
	"net/http"
	"strconv"
	"strings"
)

type JwtInteractor struct {
	JwtRepo      port.JwtRepository
	UserRepo     port.UserPort
	userAuthRepo port.UserAuthPort
	loggerRepo   port.LoggerPort
}

func NewJwtInteractor(jwtRepo port.JwtRepository, userRepo port.UserPort, userAuthRepo port.UserAuthPort, loggerRepo port.LoggerPort) *JwtInteractor {
	return &JwtInteractor{UserRepo: userRepo, JwtRepo: jwtRepo, userAuthRepo: userAuthRepo, loggerRepo: loggerRepo}
}

func (interactor *JwtInteractor) PrintInfo(num string, group string, message string) {
	interactor.loggerRepo.PrintInfo(num, group, message)
}

func (interactor *JwtInteractor) PrintError(num string, group string, message string) {
	interactor.loggerRepo.PrintError(num, group, message)
}

func (interactor *JwtInteractor) PrintDebug(num string, group string, message string) {
	interactor.loggerRepo.PrintDebug(num, group, message)
}

func (interactor *JwtInteractor) JwtMiddleware() gin.HandlerFunc {
	interactor.loggerRepo.PrintInfo("Start", "JwtInteractor:JwtMiddleware", "")
	fmt.Println("Start", "JwtInteractor:JwtMiddleware")

	return func(c *gin.Context) {

		token, err := interactor.JwtRepo.GetToken(c)
		if err != nil {
			res := response.Response{C: c}
			messageStr := message.GetMsg(message.ERR005)

			interactor.loggerRepo.PrintInfo("End", "JwtInteractor:JwtMiddleware", messageStr)
			fmt.Println("End", "JwtInteractor:JwtMiddleware", messageStr)

			res.ErrorPayload(http.StatusUnauthorized, message.ERR005, errors.New(messageStr))
			c.Abort()
			return
		}

		auth, err := interactor.JwtRepo.ParseProc(token)

		if err != nil {
			res := response.Response{C: c}
			messageStr := message.GetMsg(message.ERR005)
			interactor.loggerRepo.PrintInfo("End", "JwtInteractor:JwtMiddleware", messageStr)
			fmt.Println("End", "JwtInteractor:JwtMiddleware", messageStr)
			res.ErrorPayload(http.StatusUnauthorized, message.ERR005, errors.New(messageStr))
			c.Abort()
			return
		}

		userAuth, err := interactor.userAuthRepo.FindAccessToken(token)

		if err != nil {
			res := response.Response{C: c}
			messageStr := message.GetMsg(message.ERR005)
			interactor.loggerRepo.PrintInfo("End", "JwtInteractor:JwtMiddleware", messageStr)
			fmt.Println("End", "JwtInteractor:JwtMiddleware", messageStr)
			res.ErrorPayload(http.StatusUnauthorized, message.ERR005, errors.New(messageStr))
			c.Abort()
			return
		}

		user, err := interactor.UserRepo.FindUserId(userAuth.UserID)
		if err != nil {
			return
		}

		if strings.Compare(user.UserName, auth.UserName) != 0 {
			res := response.Response{C: c}
			messageStr := message.GetMsg(message.ERR005)
			interactor.loggerRepo.PrintInfo("End", "JwtInteractor:JwtMiddleware", messageStr)
			fmt.Println("End", "JwtInteractor:JwtMiddleware", messageStr)
			res.ErrorPayload(http.StatusUnauthorized, message.ERR005, errors.New(messageStr))
			c.Abort()
		} else {

			interactor.loggerRepo.PrintInfo("End", "JwtInteractor:JwtMiddleware", "auth:"+"UserName:"+auth.UserName+"Iat:"+strconv.FormatInt(auth.Iat, 10)+"Exp:"+strconv.FormatInt(auth.Exp, 10))
			fmt.Println("End", "JwtInteractor:JwtMiddleware", "auth", "UserName", auth.UserName, "Iat", auth.Iat, "Exp:", auth.Exp)

			// ハンドラを内の処理に移行
			c.Next()
		}

	}
}
