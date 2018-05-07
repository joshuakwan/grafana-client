package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-resty/resty"
	"github.com/joshuakwan/grafana-client/models"
	"strconv"
)

/*
GetFolders returns all folders that the authenticated user has permission to view.
GET /api/folders
*/
func (c *Client) GetFolders() ([]*models.Folder, error) {
	resp, err := resty.R().SetHeader(authHeader, c.BearerToken).
		Get(c.GrafanaURL + "api/folders")
	if err != nil {
		return nil, err
	}

	var results []*models.Folder
	if err = json.Unmarshal(resp.Body(), &results); err != nil {
		return nil, err
	}

	return results, nil
}

/*
GetFolderByUID will return the folder given the folder uid.
GET /api/folders/:uid

Status Codes:

200 – Found
401 – Unauthorized
403 – Access Denied
404 – Folder not found
*/
func (c *Client) GetFolderByUID(uid string) (*models.Folder, error) {
	resp, err := resty.R().SetHeader(authHeader, c.BearerToken).
		Get(c.GrafanaURL + "api/folders/" + uid)
	if err != nil {
		return nil, err
	}

	switch resp.StatusCode() {
	case 200:
		var folder *models.Folder
		if err = json.Unmarshal(resp.Body(), &folder); err != nil {
			return nil, err
		}
		return folder, nil
	case 401:
		return nil, errors.New("Unauthorized")
	case 403:
		return nil, errors.New("Access Denied")
	case 404:
		return nil, errors.New("Folder not found")
	}

	return nil, errors.New("In the middle of nowhere")
}

/*
CreateFolder creates a new folder
POST /api/folders

Status Codes:

200 – Created
400 – Errors (invalid json, missing or invalid fields, etc)
401 – Unauthorized
403 – Access Denied
*/
func (c *Client) CreateFolder(title string) (*models.Folder, error) {
	resp, err := resty.R().SetHeader(authHeader, c.BearerToken).
		SetBody(fmt.Sprintf(`{"title":"%s"}`, title)).
		SetHeader("Content-Type", "application/json").
		Post(c.GrafanaURL + "api/folders")
	if err != nil {
		return nil, err
	}

	switch resp.StatusCode() {
	case 200:
		var folder *models.Folder
		if err = json.Unmarshal(resp.Body(), &folder); err != nil {
			return nil, err
		}
		return folder, nil
	case 401:
		return nil, errors.New("Unauthorized")
	case 403:
		return nil, errors.New("Access Denied")
	case 400:
		return nil, errors.New("Errors (invalid json, missing or invalid fields, etc)")
	}

	return nil, errors.New("In the middle of nowhere")
}

/* TODO
Update folder
PUT /api/folders/:uid

Updates an existing folder identified by uid.

Example Request:

PUT /api/folders/nErXDvCkzz HTTP/1.1
Accept: application/json
Content-Type: application/json
Authorization: Bearer eyJrIjoiT0tTcG1pUlY2RnVKZTFVaDFsNFZXdE9ZWmNrMkZYbk

{
"title":"Department DEF",
"version": 1
}
JSON Body schema:

uid – Provide another unique identifier than stored to change the unique identifier.
title – The title of the folder.
version – Provide the current version to be able to update the folder. Not needed if overwrite=true.
overwrite – Set to true if you want to overwrite existing folder with newer version.
Example Response:

HTTP/1.1 200
Content-Type: application/json

{
"id":1,
"uid": "nErXDvCkzz",
"title": "Departmenet DEF",
"url": "/dashboards/f/nErXDvCkzz/department-def",
"hasAcl": false,
"canSave": true,
"canEdit": true,
"canAdmin": true,
"createdBy": "admin",
"created": "2018-01-31T17:43:12+01:00",
"updatedBy": "admin",
"updated": "2018-01-31T17:43:12+01:00",
"version": 1
}
Status Codes:

200 – Updated
400 – Errors (invalid json, missing or invalid fields, etc)
401 – Unauthorized
403 – Access Denied
404 – Folder not found
412 – Precondition failed
The 412 status code is used for explaing that you cannot update the folder and why. There can be different reasons for this:

The folder has been changed by someone else, status=version-mismatch
The response body will have the following properties:

HTTP/1.1 412 Precondition Failed
Content-Type: application/json; charset=UTF-8
Content-Length: 97

{
"message": "The folder has been changed by someone else",
"status": "version-mismatch"
}
*/

/*
DeleteFolderByUID an existing folder identified by uid together with all dashboards stored in the folder, if any. This operation cannot be reverted.

DELETE /api/folders/:uid

Status Codes:

200 – Deleted
401 – Unauthorized
403 – Access Denied
404 – Folder not found
*/
func (c *Client) DeleteFolderByUID(uid string) (*models.Message, error) {
	resp, err := resty.R().SetHeader(authHeader, c.BearerToken).
		Delete(c.GrafanaURL + "api/folders/" + uid)
	if err != nil {
		return nil, err
	}

	switch resp.StatusCode() {
	case 200:
		var message *models.Message
		if err = json.Unmarshal(resp.Body(), &message); err != nil {
			return nil, err
		}
		return message, nil
	case 401:
		return nil, errors.New("Unauthorized")
	case 403:
		return nil, errors.New("Access Denied")
	case 404:
		return nil, errors.New("Folder not found")
	}

	return nil, errors.New("In the middle of nowhere")
}

/*
GetFolderByID will return the folder identified by id.
GET /api/folders/:id

Status Codes:

200 – Found
401 – Unauthorized
403 – Access Denied
404 – Folder not found
*/
func (c *Client) GetFolderByID(id int) (*models.Folder, error) {
	resp, err := resty.R().SetHeader(authHeader, c.BearerToken).
		Get(c.GrafanaURL + "api/folders/" + strconv.Itoa(id))
	if err != nil {
		return nil, err
	}

	switch resp.StatusCode() {
	case 200:
		var folder *models.Folder
		if err = json.Unmarshal(resp.Body(), &folder); err != nil {
			return nil, err
		}
		return folder, nil
	case 401:
		return nil, errors.New("Unauthorized")
	case 403:
		return nil, errors.New("Access Denied")
	case 404:
		return nil, errors.New("Folder not found")
	}

	return nil, errors.New("In the middle of nowhere")
}
