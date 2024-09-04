package response

type DeleteAccessTokenResponse struct {
	UserName    string `json:"user_name"`
	AccessToken string `json:"access_token"`
}
