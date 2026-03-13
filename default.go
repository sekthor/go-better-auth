package gobetterauth

import (
	"fmt"
	"net/http"
)

type SignInEmailRequest struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	CallbackUrl string `json:"callbackURL"`
	RememberMe  bool   `json:"rememberMe"`
}

type SignUpEmailRequest struct {
	SignInEmailRequest
	Name  string `json:"name"`
	Image string `json:"image"`
}

type SignInEmailResponse struct {
	User User `json:"user"`
}

type GetSessionResponse struct {
	Session Session `json:"session"`
	User    User    `json:"user"`
}

func (c *Client) SignUpEmail(req SignUpEmailRequest) (SignInEmailResponse, error) {
	url := fmt.Sprintf("%s/sign-up/email", c.config.BaseURL)
	return invokeApiRequest[SignInEmailResponse](c.client, http.MethodPost, url, req, nil)
}

func (c *Client) SignInEmail(req SignInEmailRequest) (SignInEmailResponse, error) {
	url := fmt.Sprintf("%s/sign-in/email", c.config.BaseURL)
	return invokeApiRequest[SignInEmailResponse](c.client, http.MethodPost, url, req, nil)
}

func (c *Client) GetSession() (GetSessionResponse, error) {
	url := fmt.Sprintf("%s/get-session", c.config.BaseURL)
	return invokeApiRequest[GetSessionResponse](c.client, http.MethodGet, url, nil, nil)
}
