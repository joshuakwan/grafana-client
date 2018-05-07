package api

import (
	"encoding/json"
	"github.com/go-resty/resty"
	"github.com/joshuakwan/grafana-client/models"
)

/*
200 – Created
400 – Errors (invalid json, missing or invalid fields, etc)
401 – Unauthorized
403 – Access denied
412 – Precondition failed
*/
func (c *Client) postDashboard(dashboard *models.GrafanaDashboard) (*models.DashboardSuccessfulPostMessage, error) {
	resp, err := resty.R().SetHeader(authHeader, c.BearerToken).
		SetBody(dashboard).Post(c.GrafanaURL + "api/dashboards/db/")
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

// CreateDashboard creates a new dashboard
func (c *Client) CreateDashboard(dashboard *models.GrafanaDashboard) (*models.DashboardSuccessfulPostMessage, error) {
	return c.postDashboard(dashboard)
}

// CreateDashboardFromJSON creates a new dashboard from a JSON input
func (c *Client) CreateDashboardFromJSON(jsonData []byte) (*models.DashboardSuccessfulPostMessage, error) {
	var dashboard models.GrafanaDashboard

	if err := json.Unmarshal(jsonData, &dashboard); err != nil {
		return nil, err
	}

	if dashboard.Meta == nil {
		dashboard.Meta = &models.Meta{CanSave: true, Slug: dashboard.Dashboard.Title}
	}

	return c.postDashboard(&dashboard)
}

// UpdateDashboard updates a dashboard
func (c *Client) UpdateDashboard(uid string, dashboard *models.GrafanaDashboard) (*models.DashboardSuccessfulPostMessage, error) {
	dashboard.Dashboard.UID = uid
	targetDashboard, err := c.GetDashboardByUID(uid)
	if err != nil {
		return nil, err
	}

	dashboard.Dashboard.Version = targetDashboard.Dashboard.Version

	return c.postDashboard(dashboard)
}

// UpdateDashboardFromJSON updates a dashboard from a JSON input
func (c *Client) UpdateDashboardFromJSON(uid string, jsonData []byte) (*models.DashboardSuccessfulPostMessage, error) {
	var dashboard models.GrafanaDashboard

	err := json.Unmarshal(jsonData, &dashboard)
	if err != nil {
		return nil, err
	}

	return c.UpdateDashboard(uid, &dashboard)
}

// GetDashboardByUID returns a dashboard by its UID
func (c *Client) GetDashboardByUID(uid string) (*models.GrafanaDashboard, error) {
	resp, err := resty.R().SetHeader(authHeader, c.BearerToken).
		Get(c.GrafanaURL + "api/dashboards/uid/" + uid)
	if err != nil {
		return nil, err
	}

	var dashboard *models.GrafanaDashboard
	if err = json.Unmarshal(resp.Body(), &dashboard); err != nil {
		return nil, err
	}

	return dashboard, nil
}

// GetDashboardUID returns the UID of a dashboard
func (c *Client) GetDashboardUID(dashboardTitle string, dashboardFolderTitle string) (string, error) {
	var results []*models.SearchResult
	var uid string

	resp, err := resty.R().SetHeader(authHeader, c.BearerToken).
		SetQueryParam("query", dashboardTitle).
		Get(c.GrafanaURL + "api/search/")
	if err != nil {
		return "", err
	}

	if err = json.Unmarshal(resp.Body(), &results); err != nil {
		return "", err
	}

	for _, result := range results {
		if result.FolderTitle == dashboardFolderTitle {
			uid = result.UID
		}
	}

	return uid, nil
}

// GetDashboardTags returns the tags of a dashboard
func (c *Client) GetDashboardTags() ([]*models.TagResult, error) {
	resp, err := resty.R().SetHeader(authHeader, c.BearerToken).
		Get(c.GrafanaURL + "api/dashboards/tags")
	if err != nil {
		return nil, err
	}

	var results []*models.TagResult
	if err = json.Unmarshal(resp.Body(), &results); err != nil {
		return nil, err
	}

	return results, nil
}

// DeleteDashboardByUID deletes a dashboard by its UID
func (c *Client) DeleteDashboardByUID(uid string) (*models.DashboardSuccessfulDeleteMessage, error) {
	resp, err := resty.R().SetHeader(authHeader, c.BearerToken).
		Delete(c.GrafanaURL + "api/dashboards/uid/" + uid)
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

////////////// TODO

func (c *Client) adminPostDashboard(orgID int, dashboard *models.GrafanaDashboard) (*models.DashboardSuccessfulPostMessage, error) {
	err := c.AdminSwitchOrganization(orgID)
	if err != nil {
		return nil, err
	}

	resp, err := resty.R().SetBasicAuth(c.AdminUser, c.AdminPassword).
		SetBody(dashboard).Post(c.GrafanaURL + "api/dashboards/db/")
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

// AdminCreateDashboard creates a dashboard with the admin basic authentication
// use with caution, and it might be removed in the future
func (c *Client) AdminCreateDashboard(orgID int, dashboard *models.GrafanaDashboard) (*models.DashboardSuccessfulPostMessage, error) {
	return c.adminPostDashboard(orgID, dashboard)
}

// AdminCreateDashboardFromJSON creates a dashboard with the admin basic authentication from a JSON input
// use with caution, and it might be removed in the future
func (c *Client) AdminCreateDashboardFromJSON(orgID int, jsonData []byte) (*models.DashboardSuccessfulPostMessage, error) {
	var dashboard models.GrafanaDashboard

	err := json.Unmarshal(jsonData, &dashboard)
	if err != nil {
		return nil, err
	}

	if dashboard.Meta == nil {
		dashboard.Meta = &models.Meta{CanSave: true, Slug: dashboard.Dashboard.Title}
	}

	return c.adminPostDashboard(orgID, &dashboard)
}

// AdminUpdateDashboard updates a dashboard with the admin basic authentication
// use with caution, and it might be removed in the future
func (c *Client) AdminUpdateDashboard(orgID int, uid string, dashboard *models.GrafanaDashboard) (*models.DashboardSuccessfulPostMessage, error) {
	dashboard.Dashboard.UID = uid
	targetDashboard, err := c.GetDashboardByUID(uid)
	if err != nil {
		return nil, err
	}

	dashboard.Dashboard.Version = targetDashboard.Dashboard.Version

	return c.adminPostDashboard(orgID, dashboard)
}

// AdminUpdateDashboardFromJSON updates a dashboard with the admin basic authentication from a JSON input
// use with caution, and it might be removed in the future
func (c *Client) AdminUpdateDashboardFromJSON(orgID int, uid string, jsonData []byte) (*models.DashboardSuccessfulPostMessage, error) {
	var dashboard models.GrafanaDashboard

	err := json.Unmarshal(jsonData, &dashboard)
	if err != nil {
		return nil, err
	}

	return c.AdminUpdateDashboard(orgID, uid, &dashboard)
}

// AdminDeleteDashboardByUID deletes a dashboard with the admin basic authentication
// use with caution, and it might be removed in the future
func (c *Client) AdminDeleteDashboardByUID(orgID int, uid string) (*models.DashboardSuccessfulDeleteMessage, error) {
	err := c.AdminSwitchOrganization(orgID)
	if err != nil {
		return nil, err
	}

	resp, err := resty.R().SetBasicAuth(c.AdminUser, c.AdminPassword).
		Delete(c.GrafanaURL + "api/dashboards/uid/" + uid)

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
