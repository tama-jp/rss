package request

import (
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type UserSearchRequest struct {
}

func (req UserSearchRequest) Validate(c *gin.Context) error {
	return validation.ValidateStruct(&req, nil)
}
