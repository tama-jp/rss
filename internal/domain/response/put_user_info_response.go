package response

type PutUserInfoResponse struct {
	UserId          uint                        `json:"user_id"`
	EmployeeNumber  string                      `json:"employee_number"`
	UserName        string                      `json:"user_name"`
	LastName        string                      `json:"last_name"`
	FirstName       string                      `json:"first_name"`
	RoleBitCode     uint64                      `json:"role_bit_code"`
	RoleDescription UserRoleDescriptionResponse `json:"role_description"`
}
