package models

type OrganizationAddress struct {
	Address1 string `json:"address1"`
	Address2 string `json:"address2"`
	City     string `json:"city"`
	ZipCode  string `json:"zipCode"`
	State    string `json:"state"`
	Country  string `json:"country"`
}

type GrafanaOrganization struct {
	ID      int                 `json:"id"`
	Name    string              `json:"name"`
	Address OrganizationAddress `json:"address"`
}
