package models

// JSONData is the json data of a grafana datasource
type JSONData struct {
	KeepCookies []string `json:"keepCookies" yaml:"keepCookies"`
}

// Datasource holds a datasource object
type Datasource struct {
	ID          int       `json:"id,omitempty"`
	OrgID       int       `json:"orgId,omitempty"`
	Name        string    `json:"name,omitempty"`
	Type        string    `json:"type,omitempty"`
	TypeLogoURL string    `json:"typeLogoUrl,omitempty"`
	Access      string    `json:"access,omitempty"`
	URL         string    `json:"url,omitempty"`
	Database    string    `json:"database,omitempty"`
	BasicAuth   bool      `json:"basicAuth,omitempty"`
	IsDefault   bool      `json:"isDefault,omitempty"`
	JSONData    *JSONData `json:"jsonData,omitempty"`
	ReadOnly    bool      `json:"readOnly,omitempty"`
}
