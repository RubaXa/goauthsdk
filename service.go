package oauth

import (
	"net/http"
	"strings"
)

const (
	BaseURL = "https://oauth.mail.ru"
)

type Service struct {
	ClientID     EnvVar
	ClientSecret EnvVar
	RedirectURI  EnvVar
	Scopes       []string
}

type AuthResponse struct {
	Code        string
	State       string
	UserAgent   string
	RedirectURI string
}

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
	TokenType    string `json:"token_type"`
}

func (s Service) GetAuthURL() *OURL {
	return NewAuthURL(s)
}

func (s Service) ParseResponse(r *http.Request) AuthResponse {
	shceme := "http"
	if r.TLS != nil {
		shceme += "s"
	}

	q := r.URL.Query()
	u := shceme + "://" + r.Host + r.URL.String()

	return AuthResponse{
		Code:        q.Get("code"),
		State:       q.Get("state"),
		UserAgent:   r.UserAgent(),
		RedirectURI: u[0:strings.Index(u, "?")],
	}
}

func (s Service) Exchange(auth AuthResponse) (Token, error) {
	task := pullExchangeTask()
	go task.exec(s, auth)
	result := <-task.promise

	return result.token, result.err
}