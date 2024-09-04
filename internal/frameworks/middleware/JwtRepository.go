package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	entity "github.com/tama-jp/rss/internal/domain/entities"
	"github.com/tama-jp/rss/internal/usecases/port"
	"github.com/tama-jp/rss/internal/utils/message"
	"strings"
	"time"
)

//type JwtRepository interface {
//	GetToken(c *gin.Context) (string, error)
//	GetTargetUser(c *gin.Context) (*entity.Auth, error)
//	GenerateTokenProc(userID string) (*entity.Auth, string, error)
//	ParseProc(signedString string) (*entity.Auth, error)
//}

type jwtRepository struct {
	auth *entity.Auth
}

func NewJwtRepository() port.JwtRepository {
	auth := &entity.Auth{}
	return &jwtRepository{auth: auth}
}

const (
	secret = "2FMd5FNSqS/nW2wWJy5S3ppjSHhUnLt8HuwBkTD6HqfPfBBDlykwLA=="

	// userNameKey はユーザーの ID を表す。
	userNameKey = "user_name"

	// iat と exp は登録済みクレーム名。それぞれの意味は https://tools.ietf.org/html/rfc7519#section-4.1 を参照。{
	iatKey = "iat"
	expKey = "exp"
	// }

	// lifetime は jwt の発行から失効までの期間を表す。
	lifetime = message.API_LIFE_TIME_ACCESS_TOKEN * time.Second
)

func (repository *jwtRepository) GetToken(c *gin.Context) (string, error) {
	strAuthorization := c.Request.Header["Authorization"]

	slice := strings.Split(strAuthorization[0], " ")

	var token = slice[1]
	return token, nil
}

func (repository *jwtRepository) GetTargetUser(c *gin.Context) (*entity.Auth, error) {
	token, err := repository.GetToken(c)
	if err != nil {
		fmt.Println(err)
		messageStr := message.GetMsg(message.ERR004)
		return nil, fmt.Errorf(" %s", messageStr)

	}

	auth, err := repository.ParseProc(token)

	if err != nil {
		messageStr := message.GetMsg(message.ERR005)
		return nil, fmt.Errorf(" %s", messageStr)
	}
	fmt.Println("*** userName ****", auth.UserName)

	return auth, nil
}

func (repository *jwtRepository) ParseProc(signedString string) (*entity.Auth, error) {
	token, err := jwt.Parse(signedString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return "", fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, fmt.Errorf("%s is expired [%s]", signedString, err)
			} else {
				return nil, fmt.Errorf("%s is invalid [%s]", signedString, err)
			}
		} else {
			return nil, fmt.Errorf("%s is invalid [ %s]", signedString, err)
		}
	}

	if token == nil {
		return nil, fmt.Errorf("not found token in %s", signedString)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("not found claims in %s", signedString)
	}
	userName, ok := claims[userNameKey].(string)
	if !ok {
		return nil, fmt.Errorf("not found %s in %s", userNameKey, signedString)
	}
	iat, ok := claims[iatKey].(float64)
	if !ok {
		return nil, fmt.Errorf("not found %s in %s", iatKey, signedString)
	}

	exp, ok := claims[expKey].(float64)
	if !ok {
		return nil, fmt.Errorf("not found %s in %s", exp, signedString)
	}

	return &entity.Auth{
		UserName: userName,
		Iat:      int64(iat),
		Exp:      int64(exp),
	}, nil
}
func (repository *jwtRepository) GenerateTokenProc(userName string) (*entity.Auth, string, error) {
	//fmt.Println("*** Generate_token_proc ***")
	now := time.Now()

	auth := &entity.Auth{
		UserName: userName,
		Iat:      now.Unix(),
		Exp:      now.Add(lifetime).Unix(),
	}

	tokenData := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		userNameKey: auth.UserName,
		iatKey:      auth.Iat,
		expKey:      auth.Exp,
	})

	token, err := tokenData.SignedString([]byte(secret))

	return auth, token, err
}
