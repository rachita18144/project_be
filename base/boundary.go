package base

const (
	// Server Configurations
	SERVER_TYPE_LOCALHOST   Config = 1
	SERVER_TYPE_DEVELOPMENT Config = 2
	SERVER_TYPE_PRODUCTION  Config = 3

	// Database Constants
	DB_PORTAL = "portal"
	COL_USERS = "users"
	COL_OFORM = "oform"

	// Server Details
	SERVER_VERSION = "1.0.0"
)

var (
	RUNNING_SERVER_TYPE Config
	BASE_URL            string
	PORT                string
	MONGO_BASE_URL      string
)

type Config int

func SetupServer(c Config) (string, error) {
	return setupServer(c)
}
