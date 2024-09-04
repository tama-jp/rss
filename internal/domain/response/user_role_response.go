package response

type UserRoleDescriptionResponse struct {
	NoAuthority bool `json:"no_authority"`
	Default     bool `json:"default"`
	SuperUser   bool `json:"super_user"`
	Ceo         bool `json:"ceo"`
}
