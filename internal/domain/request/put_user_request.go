package request

import (
	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/tama-jp/rss/internal/utils/message"
)

type UserPutRequest struct {
	UserName       *string `json:"user_name"`
	LastName       *string `json:"last_name"`
	FirstName      *string `json:"first_name"`
	EmployeeNumber *string `json:"employee_number"`
	Password       *string `json:"password"`
	RoleBitCode    uint64  `json:"role_bit_code,uint"`
}

func (req UserPutRequest) Validate(c *gin.Context) error {
	return validation.ValidateStruct(&req,
		validation.Field(&req.UserName, validation.Required.Error(message.GetMsg(message.ERR001)),
			validation.RuneLength(1, 256).Error(message.GetMsg(message.ERR002, message.USER_ID, 256))),
		validation.Field(&req.LastName, validation.Required.Error(message.GetMsg(message.ERR001)),
			validation.RuneLength(1, 256).Error(message.GetMsg(message.ERR002, message.USER_ID, 256))),
		validation.Field(&req.FirstName, validation.Required.Error(message.GetMsg(message.ERR001)),
			validation.RuneLength(1, 256).Error(message.GetMsg(message.ERR002, message.USER_ID, 256))),
		validation.Field(&req.EmployeeNumber, validation.Required.Error(message.GetMsg(message.ERR001)),
			validation.RuneLength(1, 256).Error(message.GetMsg(message.ERR002, message.USER_ID, 256))),
		validation.Field(&req.Password, validation.Required.Error(message.GetMsg(message.ERR001)),
			validation.RuneLength(8, 256).Error(message.GetMsg(message.ERR003, message.PASSWORD))),
		validation.Field(&req.RoleBitCode, validation.Required.Error(message.GetMsg(message.ERR007))),
	)
}
