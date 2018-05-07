package api

var (
	AuthHeader = "Authorization"
)

// Client encapsulates a Grafana API client
type Client struct {
	BearerToken   string
	AdminUser     string
	AdminPassword string
	GrafanaURL    string
}
