package models

type DatasourceSuccessfulCreateMessage struct {
	Message string `json:"message"`
	ID      int    `json:"id"`
	Name    string `json:"name"`
}

type JsonData struct {
	KeepCookies []string `json:"keepCookies" yaml:"keepCookies"`
}

type Datasource struct {
	ID          int      `json:"id"`
	OrgID       int      `json:"orgId"`
	Name        string   `json:"name"`
	Type        string   `json:"type"`
	TypeLogoURL string   `json:"typeLogoUrl"`
	Access      string   `json:"access"`
	URL         string   `json:"url"`
	Database    string   `json:"database"`
	BasicAuth   bool     `json:"basicAuth"`
	IsDefault   bool     `json:"isDefault"`
	JsonData    JsonData `json:"jsonData"`
	ReadOnly    bool     `json:"readOnly"`
}
