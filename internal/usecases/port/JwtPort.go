package port

import (
	"github.com/gin-gonic/gin"
	entity "github.com/tama-jp/rss/internal/domain/entities"
)

type JwtRepository interface {
	GetToken(c *gin.Context) (string, error)
	GetTargetUser(c *gin.Context) (*entity.Auth, error)
	GenerateTokenProc(userID string) (*entity.Auth, string, error)
	ParseProc(signedString string) (*entity.Auth, error)
}
