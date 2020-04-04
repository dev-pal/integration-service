package cfg

type App struct {
	AppID            string   `json:"appId"`
	AppSecret        string   `json:"appSecret"`
	ClientSecret     string   `json:"clientSecret"`
	AuthorizeURL     string   `json:"authorizeUrl"`
	AccessTokenUrl   string   `json:"accessTokenUrl"`
	AuthorizedScopes []string `json:"authorizedScopes"`
}
