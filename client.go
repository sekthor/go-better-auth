package gobetterauth

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"

	"github.com/sekthor/go-better-auth/models"
)

var _ Client = &betterAuthClient{}

type Client interface {
	getClient() *http.Client
	SignUpEmail(SignUpEmailRequest) (SignInEmailResponse, error)
	SignInEmail(SignInEmailRequest) (SignInEmailResponse, error)
	GetSession() (GetSessionResponse, error)
}

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
	User models.User `json:"user"`
}

type GetSessionResponse struct {
	Session models.Session `json:"session"`
	User    models.User    `json:"user"`
}

type betterAuthClient struct {
	ClientConfig
	token  string
	client *http.Client
}

type ClientConfig struct {
	BaseURL string
}

func NewBetterAuthClient(conf ClientConfig) (*betterAuthClient, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}

	return &betterAuthClient{
		ClientConfig: conf,
		client: &http.Client{
			Jar: jar,
		},
	}, nil
}

func (c *betterAuthClient) getClient() *http.Client {
	return c.client
}

func (c *betterAuthClient) SignUpEmail(req SignUpEmailRequest) (SignInEmailResponse, error) {
	url := fmt.Sprintf("%s/sign-up/email", c.BaseURL)
	return DoApiRequest[SignInEmailResponse](c, http.MethodPost, url, req)
}

func (c *betterAuthClient) SignInEmail(req SignInEmailRequest) (SignInEmailResponse, error) {
	url := fmt.Sprintf("%s/sign-in/email", c.BaseURL)
	return DoApiRequest[SignInEmailResponse](c, http.MethodPost, url, req)
}

func (c *betterAuthClient) GetSession() (GetSessionResponse, error) {
	url := fmt.Sprintf("%s/get-session", c.BaseURL)
	return DoApiRequest[GetSessionResponse](c, http.MethodGet, url, nil)
}
