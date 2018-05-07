package models

// Folder is a grafana folder
type Folder struct {
	ID        int    `json:"id"`
	UID       string `json:"uid"`
	Title     string `json:"title"`
	URL       string `json:"url"`
	HasACL    bool   `json:"hasAcl"`
	CanSave   bool   `json:"canSave"`
	CanEdit   bool   `json:"canEdit"`
	CanAdmin  bool   `json:"canAdmin"`
	CreatedBy string `json:"createdBy"`
	Created   string `json:"created"`
	UpdatedBy string `json:"updatedBy"`
	Updated   string `json:"updated"`
	Version   int    `json:"version"`
}
