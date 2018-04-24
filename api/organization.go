package api

import (
	"github.com/joshuakwan/grafana-client/models"
	"github.com/go-resty/resty"
	"encoding/json"
)

func (client *Client) GetOrganization(orgName string) (*models.GrafanaOrganization, error) {
	resp, err := resty.R().SetBasicAuth(client.AdminUser, client.AdminPassword).
		Get(client.GrafanaURL + "api/orgs/name/" + orgName)

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

func (client *Client) CreateOrganization(organization *models.GrafanaOrganization) (*models.OrganizationSuccessfulPostMessage, error) {
	resp, err := resty.R().SetBasicAuth(client.AdminUser, client.AdminPassword).
		SetBody(organization).Post(client.GrafanaURL + "api/orgs/")

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
