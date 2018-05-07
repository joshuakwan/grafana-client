package models

// Message is used to return a grafana message
type Message struct {
	Message string `json:"message"`
}

// OrganizationSuccessfulPostMessage is used to return a message when an organization is created
type OrganizationSuccessfulPostMessage struct {
	OrgID   int    `json:"orgId"`
	Message string `json:"message"`
}
// DashboardSuccessfulDeleteMessage is used to return a message when a dashboard is deleted
type DashboardSuccessfulDeleteMessage struct {
	Title string `json:"title"`
}
// DashboardSuccessfulPostMessage is used to return a message when a dashboard is created/updated
type DashboardSuccessfulPostMessage struct {
	ID      int    `json:"id"`
	UID     string `json:"uid"`
	URL     string `json:"url"`
	Status  string `json:"status"`
	Version int    `json:"version"`
	Slug    string `json:"slug"`
}
// DatasourceMessage is used to return a message when a datasource is manipulated
type DatasourceMessage struct {
	Message string `json:"message"`
	ID      int    `json:"id"`
	Name    string `json:"name"`
}
// UserSuccessfulCreateMessage is used to return a message when a user is created
type UserSuccessfulCreateMessage struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
}
// APIKeySuccessfulCreateMessage is used to return a message when an API key is created
type APIKeySuccessfulCreateMessage struct {
	Name string `json:"name"`
	Key  string `json:"key"`
}
