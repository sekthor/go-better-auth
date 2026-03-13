package gobetterauth

import (
	"net/http"
	"net/http/cookiejar"
)

type Client struct {
	config ClientConfig
	token  string
	client *http.Client
}

type ClientConfig struct {
	BaseURL string
}

func NewBetterAuthClient(conf ClientConfig) (*Client, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}

	client := &Client{
		config: conf,
		client: &http.Client{
			Jar: jar,
		},
	}

	return client, nil
}
