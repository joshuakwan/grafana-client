package api

import (
	"github.com/joshuakwan/grafana-client/models"
	"github.com/go-resty/resty"
	"encoding/json"
)

func (client *Client) CreateGlobalUser(user *models.User) (*models.UserSuccessfulCreateMessage, error) {
	resp, err := resty.R().SetBasicAuth(client.AdminUser, client.AdminPassword).
		SetHeader("Content-Type","application/json").
		SetBody(user).Post(client.GrafanaURL + "api/admin/users")

	if err != nil {
		return nil, err
	}

	var message models.UserSuccessfulCreateMessage
	err = json.Unmarshal(resp.Body(), &message)
	if err != nil {
		return nil, err
	}

	return &message, nil
}
