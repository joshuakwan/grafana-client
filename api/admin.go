package api

import (
	"github.com/go-resty/resty"
	"errors"
	"github.com/joshuakwan/grafana-client/models"
	"encoding/json"
	"strconv"
)

// Settings
//   GET /api/admin/settings (basic auth)
func (c *Client) AdminGetSettings() (string, error) {
	resp, err := resty.R().SetBasicAuth(c.AdminUser, c.AdminPassword).
		Get(c.GrafanaURL + "api/admin/settings")
	if resp.StatusCode() != 200 {
		return "", errors.New("request failed")
	}
	return string(resp.Body()), err
}

// Grafana Stats
//   GET /api/admin/stats (basic auth)
func (c *Client) AdminGetStats() (string, error) {
	resp, err := resty.R().SetBasicAuth(c.AdminUser, c.AdminPassword).
		Get(c.GrafanaURL + "api/admin/stats")
	if resp.StatusCode() != 200 {
		return "", errors.New("request failed")
	}
	return string(resp.Body()), err
}

// Global Users
//   POST /api/admin/users (basic auth)
func (client *Client) AdminCreateGlobalUser(user *models.User) (*models.UserSuccessfulCreateMessage, error) {
	resp, err := resty.R().SetBasicAuth(client.AdminUser, client.AdminPassword).
		SetHeader("Content-Type", "application/json").
		SetBody(user).Post(client.GrafanaURL + "api/admin/users")
	if err != nil {
		return nil, err
	}

	var message models.UserSuccessfulCreateMessage
	if err = json.Unmarshal(resp.Body(), &message); err != nil {
		return nil, err
	}

	return &message, nil
}

// Password for User
//   PUT /api/admin/users/:id/password (basic auth)
func (client *Client) AdminChangePassword(userID int, password string) (*models.Message, error) {
	resp, err := resty.R().SetBasicAuth(client.AdminUser, client.AdminPassword).
		SetHeader("Content-Type", "application/json").
		SetBody(models.Password{Password: password}).
		Put(client.GrafanaURL + "api/admin/users/" + strconv.Itoa(userID) + "/password")
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

// Delete global User
//   DELETE /api/admin/users/:id (basic auth)
func (client *Client) AdminDeleteGlobalUser(userID int) (*models.Message, error) {
	resp, err := resty.R().SetBasicAuth(client.AdminUser, client.AdminPassword).
		SetHeader("Content-Type", "application/json").
		Delete(client.GrafanaURL + "api/admin/users/" + strconv.Itoa(userID))
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
