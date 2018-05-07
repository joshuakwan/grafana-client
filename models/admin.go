package models

// Password encapsulates a password object
type Password struct {
	Password string `json:"password"`
}

// APIKey encapsulates an API key object
type APIKey struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}
