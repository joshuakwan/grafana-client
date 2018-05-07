package api

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/go-resty/resty"
	"github.com/joshuakwan/grafana-client/models"
)

// AdminGetSettings gets Grafana settings
//   GET /api/admin/settings (basic auth)
func (c *Client) AdminGetSettings() (string, error) {
	resp, err := resty.R().SetBasicAuth(c.AdminUser, c.AdminPassword).
		Get(c.GrafanaURL + "api/admin/settings")
	if resp.StatusCode() != 200 {
		return "", errors.New("request failed")
	}
	return string(resp.Body()), err
}

// AdminGetStats gets Grafana stats
//   GET /api/admin/stats (basic auth)
func (c *Client) AdminGetStats() (string, error) {
	resp, err := resty.R().SetBasicAuth(c.AdminUser, c.AdminPassword).
		Get(c.GrafanaURL + "api/admin/stats")
	if resp.StatusCode() != 200 {
		return "", errors.New("request failed")
	}
	return string(resp.Body()), err
}

// AdminCreateGlobalUser creates a new global user
//   POST /api/admin/users (basic auth)
func (c *Client) AdminCreateGlobalUser(user *models.User) (*models.UserSuccessfulCreateMessage, error) {
	resp, err := resty.R().SetBasicAuth(c.AdminUser, c.AdminPassword).
		SetHeader("Content-Type", "application/json").
		SetBody(user).Post(c.GrafanaURL + "api/admin/users")
	if err != nil {
		return nil, err
	}

	var message models.UserSuccessfulCreateMessage
	if err = json.Unmarshal(resp.Body(), &message); err != nil {
		return nil, err
	}

	return &message, nil
}

// AdminChangePassword changes password for a user
//   PUT /api/admin/users/:id/password (basic auth)
func (c *Client) AdminChangePassword(userID int, password string) (*models.Message, error) {
	resp, err := resty.R().SetBasicAuth(c.AdminUser, c.AdminPassword).
		SetHeader("Content-Type", "application/json").
		SetBody(models.Password{Password: password}).
		Put(c.GrafanaURL + "api/admin/users/" + strconv.Itoa(userID) + "/password")
	if err != nil {
		return nil, err
	}

	var message models.Message
	if err = json.Unmarshal(resp.Body(), &message); err != nil {
		return nil, err
	}

	return &message, nil
}

// TODO Permissions
//   PUT /api/admin/users/:id/permissions (basic auth)

// AdminDeleteGlobalUser deletes a global user
//   DELETE /api/admin/users/:id (basic auth)
func (c *Client) AdminDeleteGlobalUser(userID int) (*models.Message, error) {
	resp, err := resty.R().SetBasicAuth(c.AdminUser, c.AdminPassword).
		SetHeader("Content-Type", "application/json").
		Delete(c.GrafanaURL + "api/admin/users/" + strconv.Itoa(userID))
	if err != nil {
		return nil, err
	}

	var message models.Message
	if err = json.Unmarshal(resp.Body(), &message); err != nil {
		return nil, err
	}

	return &message, nil
}

// TODO Pause all alerts
//   POST /api/admin/pause-all-alerts (basic auth)
