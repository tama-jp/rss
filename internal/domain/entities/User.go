package entity

import "time"

type User struct {
	Model
	UserName       string     `gorm:"index:idx_name,unique; comment:ユーザ名"`
	LastName       string     `gorm:"comment:性"`
	FirstName      string     `gorm:"comment:名"`
	EmployeeNumber string     `gorm:"index:idx_user_employee_number;default:'no_number';comment:社員番号"`
	Password       string     `gorm:"comment:パスワード"`
	LastLoginAt    *time.Time `gorm:"comment:最終ログイン日時"`
	RoleBitCode    uint64     `gorm:"comment:ユーザ権限"`
}
