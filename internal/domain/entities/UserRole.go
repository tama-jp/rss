package entity

type UserRole struct {
	Model
	Name     string `gorm:"comment:和名"`
	RoleName string `gorm:"comment:権限名"`
	BitCode  uint64 `gorm:"comment:ビットコード"`
}

const (
	UserRoleNoAuthority uint64 = 1 << iota
	UserRoleDefault     uint64 = 1 << iota
	UserRoleSuperUser   uint64 = 1 << iota
)
