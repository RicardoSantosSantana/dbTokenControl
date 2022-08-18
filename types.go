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

type Items struct {
	Id             string  `json:"id"`
	Title          string  `json:"title"`
	Subtitle       string  `json:"subtitle"`
	Price          float32 `json:"price"`
	Base_price     float32 `json:"base_price"`
	Original_price float32 `json:"original_price"`
	Permalink      string  `json:"permalink"`
	Thumbnail      string  `json:"thumbnail"`
	Pictures       string  `json:"pictures"`
	Description    string  `json:"description"`
}

type Pictures struct {
	Id         string `json:"id"`
	Url        string `json:"url"`
	Secure_url string `json:"secure_url"`
	Size       string `json:"size"`
	Max_size   string `json:"max_size"`
	Quality    string `json:"quality"`
}
type Description struct {
	Text         string `json:"text"`
	Plain_text   string `json:"plain_text"`
	Last_updated string `json:"last_updated"`
	Date_created string `json:"date_created"`
}
