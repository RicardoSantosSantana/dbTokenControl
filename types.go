package dbTokenControl

type TokenResponse struct {
	Access_token  string `json:"access_token"`
	Token_type    string `json:"token_type"`
	Expires_in    int    `json:"expires_in"`
	Scope         string `json:"scope"`
	User_id       int    `json:"user_id"`
	Refresh_token string `json:"refresh_token"`
}
type StringConnection struct {
	DbName     string
	DbHost     string
	DbUser     string
	DbPassword string
	DbPort     string
}
