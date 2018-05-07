package models

// AlertNotificationSettings encapsulates an alert notification setting object
type AlertNotificationSettings struct {
	HTTPMethod string `json:"httpMethod"`
	URL        string `json:"url"`
}

// AlertNotification encapsulates an alert notification object
type AlertNotification struct {
	ID        int                       `json:"id"`
	Name      string                    `json:"name"`
	Type      string                    `json:"type"`
	IsDefault bool                      `json:"isDefault"`
	Settings  AlertNotificationSettings `json:"settings"`
	Created   string                    `json:"created"`
	Updated   string                    `json:"updated"`
}
