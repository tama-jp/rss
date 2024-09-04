package response

type GetAccessTokenResponse struct {
	UserID             uint                        `json:"user_id"`
	EmployeeNumber     string                      `json:"employee_number"`
	UserName           string                      `json:"user_name"`
	LastName           string                      `json:"last_name"`
	FirstName          string                      `json:"first_name"`
	AccessToken        string                      `json:"access_token"`
	AccessTokenExpires int64                       `json:"access_token_expires"`
	RoleBitCode        uint64                      `json:"role_bit_code"`
	RoleDescription    UserRoleDescriptionResponse `json:"role_description"`
}
