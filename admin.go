package gobetterauth

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/sekthor/go-better-auth/models"
)

type ExtendedUser struct {
	models.User
	Role       string
	Banned     bool
	BanReason  string
	BanExpires time.Time
}

type ListUserQuery struct {
	SearchValue    string
	SearchField    string
	SearchOperator string
	Limit          int
	Offset         int
	SortBy         string
	SortDirection  string
	FilterField    string
	FilterValue    string
	FilterOperator string
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
