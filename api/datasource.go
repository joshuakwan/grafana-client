package api

import (
	"github.com/joshuakwan/grafana-client/models"
	"github.com/go-resty/resty"
	"encoding/json"
	"strconv"
)

func (c *Client) GetAllDatasources() ([]*models.Datasource, error) {
	resp, err := resty.R().SetHeader(AuthHeader, c.BearerToken).
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

func (c *Client) GetDatasourceByID(datasourceID int) (*models.Datasource, error) {
	resp, err := resty.R().SetHeader(AuthHeader, c.BearerToken).
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

func (c *Client) GetDatasourceByName(datasourceName string) (*models.Datasource, error) {
	resp, err := resty.R().SetHeader(AuthHeader, c.BearerToken).
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

func (c *Client) GetDatasourceIDByName(datasourceName string) (int, error) {
	resp, err := resty.R().SetHeader(AuthHeader, c.BearerToken).
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

func (c *Client) CreateDatasource(datasource *models.Datasource) (*models.DatasourceMessage, error) {
	resp, err := resty.R().SetHeader(AuthHeader, c.BearerToken).
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

func (c *Client) UpdateDatasource(datasourceID int, datasource *models.Datasource) (*models.DatasourceMessage, error) {
	resp, err := resty.R().SetHeader(AuthHeader, c.BearerToken).
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

func (c *Client) DeleteDatasourceByID(datasourceID int) (*models.Message, error) {
	resp, err := resty.R().SetHeader(AuthHeader, c.BearerToken).
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

func (c *Client) DeleteDatasourceByName(datasourceName string) (*models.Message, error) {
	resp, err := resty.R().SetHeader(AuthHeader, c.BearerToken).
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

// TODO
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
