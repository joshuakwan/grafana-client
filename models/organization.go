package models

// OrganizationAddress holds the address info of an organization
type OrganizationAddress struct {
	Address1 string `json:"address1"`
	Address2 string `json:"address2"`
	City     string `json:"city"`
	ZipCode  string `json:"zipCode"`
	State    string `json:"state"`
	Country  string `json:"country"`
}

// GrafanaOrganization stands for an organization in grafana
type GrafanaOrganization struct {
	ID      int                 `json:"id"`
	Name    string              `json:"name"`
	Address OrganizationAddress `json:"address"`
}
