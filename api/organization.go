package api

import (
	"encoding/json"
	"errors"
	"github.com/go-resty/resty"
	"github.com/joshuakwan/grafana-client/models"
	"strconv"
)

/*
GetCurrentOrganization gets current Organisation
GET /api/org/
*/
func (c *Client) GetCurrentOrganization() (*models.GrafanaOrganization, error) {
	resp, err := resty.R().SetHeader(authHeader, c.BearerToken).
		Get(c.GrafanaURL + "api/org/")
	if err != nil {
		return nil, err
	}

	var org *models.GrafanaOrganization
	if err = json.Unmarshal(resp.Body(), &org); err != nil {
		return nil, err
	}
	return org, nil
}

/*
GetOrganizationByID gets Organisation by Id
GET /api/orgs/:orgId
*/
func (c *Client) GetOrganizationByID(orgID int) (*models.GrafanaOrganization, error) {
	resp, err := resty.R().SetBasicAuth(c.AdminUser, c.AdminPassword).
		Get(c.GrafanaURL + "api/orgs/" + strconv.Itoa(orgID))
	if err != nil {
		return nil, err
	}

	var org *models.GrafanaOrganization
	if err = json.Unmarshal(resp.Body(), &org); err != nil {
		return nil, err
	}
	return org, nil
}

/*
GetOrganizationByName gets Organisation by Name
GET /api/orgs/name/:orgName
*/
func (c *Client) GetOrganizationByName(orgName string) (*models.GrafanaOrganization, error) {
	resp, err := resty.R().SetBasicAuth(c.AdminUser, c.AdminPassword).
		Get(c.GrafanaURL + "api/orgs/name/" + orgName)
	if err != nil {
		return nil, err
	}

	var org *models.GrafanaOrganization
	if err = json.Unmarshal(resp.Body(), &org); err != nil {
		return nil, err
	}
	return org, nil
}

/*
AdminCreateOrganization creates Organisation
POST /api/orgs
*/
func (c *Client) AdminCreateOrganization(organization *models.GrafanaOrganization) (*models.OrganizationSuccessfulPostMessage, error) {
	resp, err := resty.R().SetBasicAuth(c.AdminUser, c.AdminPassword).
		SetBody(organization).Post(c.GrafanaURL + "api/orgs/")
	if err != nil {
		return nil, err
	}

	var message models.OrganizationSuccessfulPostMessage
	if err = json.Unmarshal(resp.Body(), &message); err != nil {
		return nil, err
	}
	return &message, nil
}

/*
Update current Organisation
PUT /api/org

Example Request:

PUT /api/org HTTP/1.1
Accept: application/json
Content-Type: application/json
Authorization: Bearer eyJrIjoiT0tTcG1pUlY2RnVKZTFVaDFsNFZXdE9ZWmNrMkZYbk

{
  "name":"Main Org."
}
Example Response:

HTTP/1.1 200
Content-Type: application/json

{"message":"Organization updated"}
Get all users within the actual organisation
GET /api/org/users

Example Request:

GET /api/org/users HTTP/1.1
Accept: application/json
Content-Type: application/json
Authorization: Bearer eyJrIjoiT0tTcG1pUlY2RnVKZTFVaDFsNFZXdE9ZWmNrMkZYbk
Example Response:

HTTP/1.1 200
Content-Type: application/json

[
  {
    "orgId":1,
    "userId":1,
    "email":"admin@mygraf.com",
    "login":"admin",
    "role":"Admin"
  }
]
Add a new user to the actual organisation
POST /api/org/users

Adds a global user to the actual organisation.

Example Request:

POST /api/org/users HTTP/1.1
Accept: application/json
Content-Type: application/json
Authorization: Bearer eyJrIjoiT0tTcG1pUlY2RnVKZTFVaDFsNFZXdE9ZWmNrMkZYbk

{
  "role": "Admin",
  "loginOrEmail": "admin"
}
Example Response:

HTTP/1.1 200
Content-Type: application/json

{"message":"User added to organization"}
Updates the given user
PATCH /api/org/users/:userId

Example Request:

PATCH /api/org/users/1 HTTP/1.1
Accept: application/json
Content-Type: application/json
Authorization: Bearer eyJrIjoiT0tTcG1pUlY2RnVKZTFVaDFsNFZXdE9ZWmNrMkZYbk

{
  "role": "Viewer",
}
Example Response:

HTTP/1.1 200
Content-Type: application/json

{"message":"Organization user updated"}
Delete user in actual organisation
DELETE /api/org/users/:userId

Example Request:

DELETE /api/org/users/1 HTTP/1.1
Accept: application/json
Content-Type: application/json
Authorization: Bearer eyJrIjoiT0tTcG1pUlY2RnVKZTFVaDFsNFZXdE9ZWmNrMkZYbk
Example Response:

HTTP/1.1 200
Content-Type: application/json

{"message":"User removed from organization"}
Organisations
Search all Organisations
GET /api/orgs

Example Request:

GET /api/orgs HTTP/1.1
Accept: application/json
Content-Type: application/json
Authorization: Bearer eyJrIjoiT0tTcG1pUlY2RnVKZTFVaDFsNFZXdE9ZWmNrMkZYbk
Note: The api will only work when you pass the admin name and password to the request http url, like http://admin:admin@localhost:3000/api/orgs

Example Response:

HTTP/1.1 200
Content-Type: application/json

[
  {
    "id":1,
    "name":"Main Org."
  }
]
Update Organisation
PUT /api/orgs/:orgId

Update Organisation, fields Address 1, Address 2, City are not implemented yet.

Example Request:

PUT /api/orgs/1 HTTP/1.1
Accept: application/json
Content-Type: application/json
Authorization: Bearer eyJrIjoiT0tTcG1pUlY2RnVKZTFVaDFsNFZXdE9ZWmNrMkZYbk

{
  "name":"Main Org 2."
}
Example Response:

HTTP/1.1 200
Content-Type: application/json

{"message":"Organization updated"}
Get Users in Organisation
GET /api/orgs/:orgId/users

Example Request:

GET /api/orgs/1/users HTTP/1.1
Accept: application/json
Content-Type: application/json
Authorization: Bearer eyJrIjoiT0tTcG1pUlY2RnVKZTFVaDFsNFZXdE9ZWmNrMkZYbk
Note: The api will only work when you pass the admin name and password to the request http url, like http://admin:admin@localhost:3000/api/orgs/1/users

Example Response:

HTTP/1.1 200
Content-Type: application/json
[
  {
    "orgId":1,
    "userId":1,
    "email":"admin@mygraf.com",
    "login":"admin",
    "role":"Admin"
  }
]
Add User in Organisation
POST /api/orgs/:orgId/users

Example Request:

POST /api/orgs/1/users HTTP/1.1
Accept: application/json
Content-Type: application/json
Authorization: Bearer eyJrIjoiT0tTcG1pUlY2RnVKZTFVaDFsNFZXdE9ZWmNrMkZYbk

{
  "loginOrEmail":"user",
  "role":"Viewer"
}
Example Response:

HTTP/1.1 200
Content-Type: application/json

{"message":"User added to organization"}
Update Users in Organisation
PATCH /api/orgs/:orgId/users/:userId

Example Request:

PATCH /api/orgs/1/users/2 HTTP/1.1
Accept: application/json
Content-Type: application/json
Authorization: Bearer eyJrIjoiT0tTcG1pUlY2RnVKZTFVaDFsNFZXdE9ZWmNrMkZYbk

{
  "role":"Admin"
}
Example Response:

HTTP/1.1 200
Content-Type: application/json

{"message":"Organization user updated"}
Delete User in Organisation
DELETE /api/orgs/:orgId/users/:userId

Example Request:

DELETE /api/orgs/1/users/2 HTTP/1.1
Accept: application/json
Content-Type: application/json
Authorization: Bearer eyJrIjoiT0tTcG1pUlY2RnVKZTFVaDFsNFZXdE9ZWmNrMkZYbk
Example Response:

HTTP/1.1 200
Content-Type: application/json

{"message":"User removed from organization"}
*/

// DeleteOrganization deletes an organization
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

// AdminSwitchOrganization switch the admin to another organization
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

// CreateOrganizationAdminKey creates an admin API key of an organization
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

// AdminAddOrganizationUser adds a user to an organization
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
