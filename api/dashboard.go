package api

import (
	"github.com/joshuakwan/grafana-client/models"
	"github.com/go-resty/resty"
	"encoding/json"
)

func (client *Client) GetDashboardByUID(uid string) (*models.GrafanaDashboard, error) {
	resp, err := resty.R().SetHeader(AuthHeader, client.BearerToken).
		Get(client.GrafanaURL + "api/dashboards/uid/" + uid)
	if err != nil {
		return nil, err
	}

	var dashboard *models.GrafanaDashboard
	err = json.Unmarshal(resp.Body(), &dashboard)
	if err != nil {
		return nil, err
	}

	return dashboard, nil
}

func (client *Client) GetDashboardUID(dashboardTitle string, dashboardFolderTitle string) (string, error) {
	var results []*models.SearchResult
	var uid string

	resp, err := resty.R().SetHeader(AuthHeader, client.BearerToken).
		SetQueryParam("query", dashboardTitle).
		Get(client.GrafanaURL + "api/search/")
	if err != nil {
		return "", err
	}

	err = json.Unmarshal(resp.Body(), &results)
	if err != nil {
		return "", err
	}

	for _, result := range (results) {
		if result.FolderTitle == dashboardFolderTitle {
			uid = result.UID
		}
	}

	return uid, nil
}

func (client *Client) postDashboard(dashboard *models.GrafanaDashboard) (*models.DashboardSuccessfulPostMessage, error) {
	resp, err := resty.R().SetHeader(AuthHeader, client.BearerToken).
		SetBody(dashboard).Post(client.GrafanaURL + "api/dashboards/db/")
	if err != nil {
		return nil, err
	}

	var message models.DashboardSuccessfulPostMessage
	err = json.Unmarshal(resp.Body(), &message)
	if err != nil {
		return nil, err
	}

	return &message, nil
}

func (client *Client) CreateDashboardFromJSON(jsonData []byte) (*models.DashboardSuccessfulPostMessage, error) {
	var dashboard models.GrafanaDashboard

	err := json.Unmarshal(jsonData, &dashboard)
	if err != nil {
		return nil, err
	}

	if dashboard.Meta == nil {
		dashboard.Meta = &models.Meta{CanSave: true, Slug: dashboard.Dashboard.Title}
	}

	return client.postDashboard(&dashboard)
}

func (client *Client) UpdateDashboardFromJSON(uid string, jsonData []byte) (*models.DashboardSuccessfulPostMessage, error) {
	var dashboard models.GrafanaDashboard

	err := json.Unmarshal(jsonData, &dashboard)
	if err != nil {
		return nil, err
	}

	return client.UpdateDashboard(uid, &dashboard)
}

func (client *Client) CreateDashboard(dashboard *models.GrafanaDashboard) (*models.DashboardSuccessfulPostMessage, error) {
	return client.postDashboard(dashboard)
}

func (client *Client) UpdateDashboard(uid string, dashboard *models.GrafanaDashboard) (*models.DashboardSuccessfulPostMessage, error) {
	dashboard.Dashboard.UID = uid
	targetDashboard, err := client.GetDashboardByUID(uid)
	if err != nil {
		return nil, err
	}

	dashboard.Dashboard.Version = targetDashboard.Dashboard.Version

	return client.postDashboard(dashboard)
}

func (client *Client) DeleteDashboardByUID(uid string) (*models.DashboardSuccessfulDeleteMessage, error) {
	resp, err := resty.R().SetHeader(AuthHeader, client.BearerToken).
		Delete(client.GrafanaURL + "api/dashboards/uid/" + uid)

	if err != nil {
		return nil, err
	}

	var message models.DashboardSuccessfulDeleteMessage
	err = json.Unmarshal(resp.Body(), &message)
	if err != nil {
		return nil, err
	}

	return &message, nil
}

func (client *Client) adminPostDashboard(orgID int, dashboard *models.GrafanaDashboard) (*models.DashboardSuccessfulPostMessage, error) {
	err := client.AdminSwitchOrganization(orgID)
	if err != nil {
		return nil, err
	}

	resp, err := resty.R().SetBasicAuth(client.AdminUser, client.AdminPassword).
		SetBody(dashboard).Post(client.GrafanaURL + "api/dashboards/db/")
	if err != nil {
		return nil, err
	}

	var message models.DashboardSuccessfulPostMessage
	err = json.Unmarshal(resp.Body(), &message)
	if err != nil {
		return nil, err
	}

	return &message, nil
}

func (client *Client) AdminCreateDashboard(orgID int, dashboard *models.GrafanaDashboard) (*models.DashboardSuccessfulPostMessage, error) {
	return client.adminPostDashboard(orgID, dashboard)
}

func (client *Client) AdminCreateDashboardFromJSON(orgID int, jsonData []byte) (*models.DashboardSuccessfulPostMessage, error) {
	var dashboard models.GrafanaDashboard

	err := json.Unmarshal(jsonData, &dashboard)
	if err != nil {
		return nil, err
	}

	if dashboard.Meta == nil {
		dashboard.Meta = &models.Meta{CanSave: true, Slug: dashboard.Dashboard.Title}
	}

	return client.adminPostDashboard(orgID, &dashboard)
}

func (client *Client) AdminUpdateDashboard(orgID int, uid string, dashboard *models.GrafanaDashboard) (*models.DashboardSuccessfulPostMessage, error) {
	dashboard.Dashboard.UID = uid
	targetDashboard, err := client.GetDashboardByUID(uid)
	if err != nil {
		return nil, err
	}

	dashboard.Dashboard.Version = targetDashboard.Dashboard.Version

	return client.adminPostDashboard(orgID, dashboard)
}

func (client *Client) AdminUpdateDashboardFromJSON(orgID int, uid string, jsonData []byte) (*models.DashboardSuccessfulPostMessage, error) {
	var dashboard models.GrafanaDashboard

	err := json.Unmarshal(jsonData, &dashboard)
	if err != nil {
		return nil, err
	}

	return client.AdminUpdateDashboard(orgID, uid, &dashboard)
}

func (client *Client) AdminDeleteDashboardByUID(orgID int, uid string) (*models.DashboardSuccessfulDeleteMessage, error) {
	err := client.AdminSwitchOrganization(orgID)
	if err != nil {
		return nil, err
	}

	resp, err := resty.R().SetBasicAuth(client.AdminUser, client.AdminPassword).
		Delete(client.GrafanaURL + "api/dashboards/uid/" + uid)

	if err != nil {
		return nil, err
	}

	var message models.DashboardSuccessfulDeleteMessage
	err = json.Unmarshal(resp.Body(), &message)
	if err != nil {
		return nil, err
	}

	return &message, nil
}
