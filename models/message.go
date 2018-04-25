package models

type Message struct {
	Message string `json:"message"`
}

type OrganizationSuccessfulPostMessage struct {
	OrgID   int    `json:"orgId"`
	Message string `json:"message"`
}

type DashboardSuccessfulDeleteMessage struct {
	Title string `json:"title"`
}

type DashboardSuccessfulPostMessage struct {
	ID      int    `json:"id"`
	UID     string `json:"uid"`
	URL     string `json:"url"`
	Status  string `json:"status"`
	Version int    `json:"version"`
	Slug    string `json:"slug"`
}

type DatasourceSuccessfulCreateMessage struct {
	Message string `json:"message"`
	ID      int    `json:"id"`
	Name    string `json:"name"`
}

type UserSuccessfulCreateMessage struct {
	ID      int    `json:"id"`
	Message string `json:"message"`
}

type APIKeySuccessfulCreateMessage struct {
	Name string `json:"name"`
	Key  string `json:"key"`
}
