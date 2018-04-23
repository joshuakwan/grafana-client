package models

type AlertNotificationSettings struct {
	HttpMethod string `json:"httpMethod"`
	URL        string `json:"url"`
}

type AlertNotification struct {
	ID        int                       `json:"id"`
	Name      string                    `json:"name"`
	Type      string                    `json:"type"`
	IsDefault bool                      `json:"isDefault"`
	Settings  AlertNotificationSettings `json:"settings"`
	Created   string                    `json:"created"`
	Updated   string                    `json:"updated"`
}
