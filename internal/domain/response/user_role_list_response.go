package response

type UserRoleListResponse struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	RoleName string `json:"role_name"`
}
