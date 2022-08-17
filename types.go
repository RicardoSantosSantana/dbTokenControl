package dbTokenControl

type Token struct {
	Access_token  string `json:"access_token"`
	Token_type    string `json:"token_type"`
	Expires_in    int    `json:"expires_in"`
	Scope         string `json:"scope"`
	User_id       int    `json:"user_id"`
	Refresh_token string `json:"refresh_token"`
}
type STRConn struct {
	DbName     string `json:"dbname"`
	DbHost     string `json:"dbhost"`
	DbUser     string `json:"dbuser"`
	DbPassword string `json:"dbpassword"`
	DbPort     string `json:"dbport"`
}
