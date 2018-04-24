package api

import (
	"github.com/joshuakwan/grafana-client/models"
	"github.com/go-resty/resty"
	"encoding/json"
)

func (client *Client) GetAllDatasources() ([]*models.Datasource, error) {
	resp, err := resty.R().SetHeader(AuthHeader, client.BearerToken).
		Get(client.GrafanaURL + "api/datasources")

	if err != nil {
		return nil, err
	}

	var datasources []*models.Datasource
	err = json.Unmarshal(resp.Body(), &datasources)
	if err != nil {
		return nil, err
	}

	return datasources, nil
}

func (client *Client) CreateDatasource(datasource *models.Datasource) (*models.DatasourceSuccessfulCreateMessage, error) {
	resp, err := resty.R().SetHeader(AuthHeader, client.BearerToken).
		SetBody(datasource).Post(client.GrafanaURL + "api/datasources")

	if err != nil {
		return nil, err
	}

	var message models.DatasourceSuccessfulCreateMessage
	err = json.Unmarshal(resp.Body(), &message)
	if err != nil {
		return nil, err
	}

	return &message, nil
}

func (client *Client) DeleteDatasource(datasourceName string) (*models.Message, error) {
	resp, err:=resty.R().SetHeader(AuthHeader, client.BearerToken).
		Delete(client.GrafanaURL + "api/datasources/name/" + datasourceName)

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