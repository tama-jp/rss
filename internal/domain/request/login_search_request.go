package request

import (
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/tama-jp/rss/internal/utils/message"
)

type LoginSearchRequest struct {
	UserName string `json:"x-user-name"`
	Password string `json:"x-password"`
}

func (req LoginSearchRequest) Validate(c *gin.Context) error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.UserName, validation.Required.Error(message.GetMsg(message.ERR001)),
			validation.RuneLength(1, 256).Error(message.GetMsg(message.ERR002, message.USER_ID, 256))),
		validation.Field(&req.Password, validation.Required.Error(message.GetMsg(message.ERR001)),
			validation.RuneLength(8, 256).Error(message.GetMsg(message.ERR003, message.PASSWORD))),
	)
}
