package request

import (
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/tama-jp/rss/internal/utils/message"
)

type ReqUserPasswordChange struct {
	NewPassword string `json:"new_password"`
	OldPassword string `json:"old_password"`
}

func (req ReqUserPasswordChange) Validate(c *gin.Context) error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.NewPassword, validation.Required.Error(message.GetMsg(message.ERR001)),
			validation.RuneLength(8, 256).Error(message.GetMsg(message.ERR003, message.PASSWORD))),
		validation.Field(&req.OldPassword, validation.Required.Error(message.GetMsg(message.ERR001)),
			validation.RuneLength(8, 256).Error(message.GetMsg(message.ERR003, message.PASSWORD))),
	)
}
