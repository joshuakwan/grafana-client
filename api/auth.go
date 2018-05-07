package api

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty"
	"github.com/joshuakwan/grafana-client/models"
)

// AuthGetAPIKeys returns a list of API Keys
//   GET /api/auth/keys
func (c *Client) AuthGetAPIKeys() ([]*models.APIKey, error) {
	resp, err := resty.R().SetHeader(authHeader, c.BearerToken).
		Get(c.GrafanaURL + "api/auth/keys")
	if err != nil {
		return nil, err
	}

	var keys []*models.APIKey
	if err = json.Unmarshal(resp.Body(), &keys); err != nil {
		return nil, err
	}

	return keys, nil
}

// AuthCreateAPIKey creates an API Key
//   POST /api/auth/keys
func (c *Client) AuthCreateAPIKey(keyName string, role string) (*models.APIKeySuccessfulCreateMessage, error) {
	resp, err := resty.R().SetHeader(authHeader, c.BearerToken).
		SetBody(fmt.Sprintf(`{"name":"%s", "role":"%s"}`, keyName, role)).
		SetHeader("Content-Type", "application/json").
		Post(c.GrafanaURL + "api/auth/keys")
	if err != nil {
		return nil, err
	}

	var message models.APIKeySuccessfulCreateMessage
	if err = json.Unmarshal(resp.Body(), &message); err != nil {
		return nil, err
	}

	return &message, nil
}

// AuthDeleteAPIKey deletes an API Key
//   DELETE /api/auth/keys/:id
func (c *Client) AuthDeleteAPIKey(keyID int) (*models.Message, error) {
	resp, err := resty.R().SetHeader(authHeader, c.BearerToken).
		Delete(c.GrafanaURL + fmt.Sprintf("api/auth/keys/%d", keyID))
	if err != nil {
		return nil, err
	}

	var message models.Message
	if err = json.Unmarshal(resp.Body(), &message); err != nil {
		return nil, err
	}

	return &message, nil
}
