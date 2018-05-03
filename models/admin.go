package models

type Password struct {
	Password string `json:"password"`
}

type APIKey struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}
