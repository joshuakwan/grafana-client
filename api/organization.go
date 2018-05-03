package api

import (
	"github.com/joshuakwan/grafana-client/models"
	"github.com/go-resty/resty"
	"encoding/json"
	"errors"
	"strconv"
)

func (c *Client) GetOrganization(orgName string) (*models.GrafanaOrganization, error) {
	resp, err := resty.R().SetBasicAuth(c.AdminUser, c.AdminPassword).
		Get(c.GrafanaURL + "api/orgs/name/" + orgName)
	if err != nil {
		return nil, err
	}

	var org *models.GrafanaOrganization
	err = json.Unmarshal(resp.Body(), &org)
	if err != nil {
		return nil, err
	}

	return org, nil
}

func (c *Client) CreateOrganization(organization *models.GrafanaOrganization) (*models.OrganizationSuccessfulPostMessage, error) {
	resp, err := resty.R().SetBasicAuth(c.AdminUser, c.AdminPassword).
		SetBody(organization).Post(c.GrafanaURL + "api/orgs/")
	if err != nil {
		return nil, err
	}

	var message models.OrganizationSuccessfulPostMessage
	err = json.Unmarshal(resp.Body(), &message)
	if err != nil {
		return nil, err
	}

	return &message, nil
}

func (c *Client) DeleteOrganization(orgID int) (*models.Message, error) {
	resp, err := resty.R().SetBasicAuth(c.AdminUser, c.AdminPassword).
		Delete(c.GrafanaURL + "api/orgs/" + strconv.Itoa(orgID))
	if err != nil {
		return nil, err
	}

	var message models.Message
	err = json.Unmarshal(resp.Body(), &message)
	if err != nil {
		return nil, err
	}

	return &message, nil
}

func (c *Client) AdminSwitchOrganization(orgID int) error {
	resp, err := resty.R().SetBasicAuth(c.AdminUser, c.AdminPassword).
		Post(c.GrafanaURL + "api/user/using/" + strconv.Itoa(orgID))
	if err != nil {
		return err
	}

	// {"message":"Active organization changed"}
	var messageOrg models.Message
	err = json.Unmarshal(resp.Body(), &messageOrg)
	if err != nil {
		return err
	}
	if messageOrg.Message != "Active organization changed" {
		return errors.New("Fail to switch active organization for the admin")
	}

	return nil
}

// curl -X POST http://admin:admin@localhost:3000/api/user/using/<id of new org>
// curl -X POST
//      -H "Content-Type: application/json"
//      -d '{"name":"apikeycurl", "role": "Admin"}'
//      http://admin:admin@localhost:3000/api/auth/keys
func (c *Client) CreateOrganizationAdminKey(orgID int) (*models.APIKeySuccessfulCreateMessage, error) {
	err := c.AdminSwitchOrganization(orgID)
	if err != nil {
		return nil, err
	}

	resp, err := resty.R().SetBasicAuth(c.AdminUser, c.AdminPassword).
		SetBody(`{"name":"adminIntegrationKey", "role":"Admin"}`).
		SetHeader("Content-Type", "application/json").
		Post(c.GrafanaURL + "api/auth/keys")
	if err != nil {
		return nil, err
	}

	var messageKey models.APIKeySuccessfulCreateMessage
	err = json.Unmarshal(resp.Body(), &messageKey)
	if err != nil {
		return nil, err
	}

	return &messageKey, nil
}

// POST /api/orgs/:orgId/users
// {
//  "loginOrEmail":"user",
//  "role":"Viewer"
// }
// Use admin basic auth here to make things simple
func (c *Client) AdminAddOrganizationUser(orgID int, userLogin string, role string) (*models.Message, error) {
	resp, err := resty.R().SetBasicAuth(c.AdminUser, c.AdminPassword).
		SetBody(map[string]interface{}{"loginOrEmail": userLogin, "role": role}).
		SetHeader("Content-Type", "application/json").
		Post(c.GrafanaURL + "api/orgs/" + strconv.Itoa(orgID) + "/users")
	if err != nil {
		return nil, err
	}

	var message models.Message
	err = json.Unmarshal(resp.Body(), &message)
	if err != nil {
		return nil, err
	}

	return &message, nil
}
