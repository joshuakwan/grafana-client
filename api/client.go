package api

var (
	AuthHeader = "Authorization"
)

type Client struct {
	BearerToken string
	AdminUser string
	AdminPassword string
	GrafanaURL string
}