package gobetterauth

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type ExtendedUser struct {
	User
	Role       string
	Banned     bool
	BanReason  string
	BanExpires time.Time
}

type ListUsersResponse struct {
	Users []ExtendedUser `json:"users"`
}

func (c *Client) GetUser(id string) (ExtendedUser, error) {
	url := fmt.Sprintf("%s/admin/get-user?id=%s", c.config.BaseURL, id)
	res, err := invokeApiRequest[ExtendedUser](c.client, http.MethodGet, url, nil, nil)
	return res, err
}

func (c *Client) ListUsers(filters url.Values) ([]ExtendedUser, error) {
	url := fmt.Sprintf("%s/admin/list-users", c.config.BaseURL)
	res, err := invokeApiRequest[ListUsersResponse](c.client, http.MethodGet, url, nil, filters)
	return res.Users, err
}
