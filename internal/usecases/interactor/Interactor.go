package interactor

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	entity "github.com/tama-jp/rss/internal/domain/entities"
	"github.com/tama-jp/rss/internal/domain/response"
	"github.com/tama-jp/rss/internal/usecases/port"
	"github.com/tama-jp/rss/internal/utils/message"
)

var Set = wire.NewSet(
	NewAccessTokenInteractor,
	NewUserInteractor,
	NewSystemUserInteractor,
)

func getTargetUserInfo(c *gin.Context,
	jwtRepo port.JwtRepository,
	userRepo port.UserPort,
	userAuthRepo port.UserAuthPort,
	loggerRepo port.LoggerPort,
) (*entity.User, *response.UserRoleDescriptionResponse, error) {
	loggerRepo.PrintInfo("Start", "interactor:getTargetUserInfo", "")
	fmt.Println("Start", "interactor:getTargetUserInfo")

	token, err := jwtRepo.GetToken(c)

	if err != nil {
		messageStr := message.GetMsg(message.ERR005)
		return nil, nil, errors.New(messageStr)
	}

	userAuth, err := userAuthRepo.FindAccessToken(token)

	if err != nil {
		messageStr := message.GetMsg(message.ERR005)
		return nil, nil, errors.New(messageStr)
	}

	if userAuth == nil {
		messageStr := message.GetMsg(message.ERR005)
		return nil, nil, errors.New(messageStr)
	}

	auth, err := jwtRepo.GetTargetUser(c)

	fmt.Println("2", "interactor:getTargetUserInfo")

	if err != nil {
		messageStr := message.GetMsg(message.ERR005)
		return nil, nil, errors.New(messageStr)
	}
	fmt.Println("*** userName ****", auth.UserName)

	// Users の取得
	targetUser, err := userRepo.FindUserName(auth.UserName)
	if err != nil {
		return nil, nil, err
	}

	if targetUser.ID == 0 {
		return nil, nil, err
	}

	userRoleDescription := getResUserRoleDescription(targetUser.RoleBitCode, loggerRepo)

	loggerRepo.PrintInfo("End", "interactor:getTargetUserInfo", "")
	fmt.Println("End", "interactor:getTargetUserInfo")

	return targetUser, userRoleDescription, nil
}

func getUser(
	userId uint,
	userRepo port.UserPort,
	loggerRepo port.LoggerPort,
) (*entity.User, *response.UserRoleDescriptionResponse, error) {
	loggerRepo.PrintInfo("Start", "interactor:getUser", "")
	fmt.Println("Start", "interactor:getUser")

	user, err := userRepo.FindUserId(uint(userId))
	if err != nil {
		return nil, nil, err
	}

	// Users の取得
	targetUser, err := userRepo.FindUserName(user.UserName)
	if err != nil {
		return nil, nil, err
	}

	if targetUser.ID == 0 {
		return nil, nil, err
	}

	userRoleDescription := getResUserRoleDescription(targetUser.RoleBitCode, loggerRepo)

	loggerRepo.PrintInfo("End", "interactor:getUser", "")
	fmt.Println("End", "interactor:getUser")

	return targetUser, userRoleDescription, nil
}

func getResUserRoleDescription(
	userRoleNum uint64,
	loggerRepo port.LoggerPort,
) *response.UserRoleDescriptionResponse {
	loggerRepo.PrintInfo("Start", "interactor:getgetResUserRoleDescriptionUser", "")
	fmt.Println("Start", "interactor:getResUserRoleDescription")

	var userRoleNoAuthority bool
	var userRoleDefault bool
	var userRoleSuperUser bool
	var userRoleSEO bool

	if userRoleNum&entity.UserRoleNoAuthority > 0 {
		userRoleNoAuthority = true
	} else {
		userRoleNoAuthority = false
	}

	if userRoleNum&entity.UserRoleDefault > 0 {
		userRoleDefault = true
	} else {
		userRoleDefault = false
	}

	if userRoleNum&entity.UserRoleSuperUser > 0 {
		userRoleSuperUser = true
	} else {
		userRoleSuperUser = false
	}

	loggerRepo.PrintInfo("End", "interactor:getgetResUserRoleDescriptionUser", "")
	fmt.Println("End", "interactor:getResUserRoleDescription")

	return &response.UserRoleDescriptionResponse{
		NoAuthority: userRoleNoAuthority,
		Default:     userRoleDefault,
		SuperUser:   userRoleSuperUser,
		Ceo:         userRoleSEO,
	}
}
