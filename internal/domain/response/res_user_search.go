package response

type UserSearchResponse struct {
	UserId          uint                        `json:"user_id"`
	UserName        string                      `json:"user_name"`
	LastName        string                      `json:"last_name"`
	FirstName       string                      `json:"first_name"`
	EmployeeNumber  string                      `json:"employee_number"`
	RoleBitCode     uint64                      `json:"role_bit_code"`
	RoleDescription UserRoleDescriptionResponse `json:"role_description"`
}
