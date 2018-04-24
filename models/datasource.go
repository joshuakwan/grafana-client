package models

type JsonData struct {
	KeepCookies []string `json:"keepCookies" yaml:"keepCookies"`
}

type Datasource struct {
	ID          int      `json:"id,omitempty"`
	OrgID       int      `json:"orgId,omitempty"`
	Name        string   `json:"name,omitempty"`
	Type        string   `json:"type,omitempty"`
	TypeLogoURL string   `json:"typeLogoUrl,omitempty"`
	Access      string   `json:"access,omitempty"`
	URL         string   `json:"url,omitempty"`
	Database    string   `json:"database,omitempty"`
	BasicAuth   bool     `json:"basicAuth,omitempty"`
	IsDefault   bool     `json:"isDefault,omitempty"`
	JsonData    JsonData `json:"jsonData,omitempty"`
	ReadOnly    bool     `json:"readOnly,omitempty"`
}
