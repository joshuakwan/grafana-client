package api

import (
	"encoding/json"
	"github.com/go-resty/resty"
	"github.com/joshuakwan/grafana-client/models"
	"strconv"
)

// GetAllDatasources returns all datasources
func (c *Client) GetAllDatasources() ([]*models.Datasource, error) {
	resp, err := resty.R().SetHeader(authHeader, c.BearerToken).
		Get(c.GrafanaURL + "api/datasources")
	if err != nil {
		return nil, err
	}

	var datasources []*models.Datasource
	if err = json.Unmarshal(resp.Body(), &datasources); err != nil {
		return nil, err
	}

	return datasources, nil
}

// GetDatasourceByID returns a datasource by its ID
func (c *Client) GetDatasourceByID(datasourceID int) (*models.Datasource, error) {
	resp, err := resty.R().SetHeader(authHeader, c.BearerToken).
		Get(c.GrafanaURL + "api/datasources/" + strconv.Itoa(datasourceID))
	if err != nil {
		return nil, err
	}

	var datasource *models.Datasource
	if err = json.Unmarshal(resp.Body(), &datasource); err != nil {
		return nil, err
	}

	return datasource, nil
}

// GetDatasourceByName returns a datasource by its name
func (c *Client) GetDatasourceByName(datasourceName string) (*models.Datasource, error) {
	resp, err := resty.R().SetHeader(authHeader, c.BearerToken).
		Get(c.GrafanaURL + "api/datasources/name/" + datasourceName)
	if err != nil {
		return nil, err
	}

	var datasource *models.Datasource
	if err = json.Unmarshal(resp.Body(), &datasource); err != nil {
		return nil, err
	}

	return datasource, nil
}

// GetDatasourceIDByName returns the ID of a datasource by its name
func (c *Client) GetDatasourceIDByName(datasourceName string) (int, error) {
	resp, err := resty.R().SetHeader(authHeader, c.BearerToken).
		Get(c.GrafanaURL + "api/datasources/id/" + datasourceName)
	if err != nil {
		return -1, err
	}

	var datasource *models.Datasource
	if err = json.Unmarshal(resp.Body(), &datasource); err != nil {
		return -1, err
	}

	return datasource.ID, nil
}

// CreateDatasource creates a datasource
func (c *Client) CreateDatasource(datasource *models.Datasource) (*models.DatasourceMessage, error) {
	resp, err := resty.R().SetHeader(authHeader, c.BearerToken).
		SetBody(datasource).Post(c.GrafanaURL + "api/datasources")
	if err != nil {
		return nil, err
	}

	var message models.DatasourceMessage
	if err = json.Unmarshal(resp.Body(), &message); err != nil {
		return nil, err
	}

	return &message, nil
}

// UpdateDatasource updates a datasource
func (c *Client) UpdateDatasource(datasourceID int, datasource *models.Datasource) (*models.DatasourceMessage, error) {
	resp, err := resty.R().SetHeader(authHeader, c.BearerToken).
		SetBody(datasource).Put(c.GrafanaURL + "api/datasources/" + strconv.Itoa(datasourceID))
	if err != nil {
		return nil, err
	}

	var message models.DatasourceMessage
	if err = json.Unmarshal(resp.Body(), &message); err != nil {
		return nil, err
	}

	return &message, nil
}

// DeleteDatasourceByID deletes a datasource by its ID
func (c *Client) DeleteDatasourceByID(datasourceID int) (*models.Message, error) {
	resp, err := resty.R().SetHeader(authHeader, c.BearerToken).
		Delete(c.GrafanaURL + "api/datasources/" + strconv.Itoa(datasourceID))

	if err != nil {
		return nil, err
	}

	var message models.Message
	if err = json.Unmarshal(resp.Body(), &message); err != nil {
		return nil, err
	}

	return &message, nil
}

// DeleteDatasourceByName deletes a datasource by its name
func (c *Client) DeleteDatasourceByName(datasourceName string) (*models.Message, error) {
	resp, err := resty.R().SetHeader(authHeader, c.BearerToken).
		Delete(c.GrafanaURL + "api/datasources/name/" + datasourceName)

	if err != nil {
		return nil, err
	}

	var message models.Message
	if err = json.Unmarshal(resp.Body(), &message); err != nil {
		return nil, err
	}

	return &message, nil
}

// AdminCreateDatasource creates a datasource with the admin basic authentication
// use with caution, and it might be removed in the future
func (c *Client) AdminCreateDatasource(
	orgID int,
	datasource *models.Datasource) (*models.DatasourceMessage, error) {
	err := c.AdminSwitchOrganization(orgID)
	if err != nil {
		return nil, err
	}

	resp, err := resty.R().SetBasicAuth(c.AdminUser, c.AdminPassword).
		SetBody(datasource).Post(c.GrafanaURL + "api/datasources")

	if err != nil {
		return nil, err
	}

	var message models.DatasourceMessage
	err = json.Unmarshal(resp.Body(), &message)
	if err != nil {
		return nil, err
	}

	return &message, nil
}
