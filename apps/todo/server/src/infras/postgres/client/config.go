package postgres_client

type Config struct {
	Host      string
	Port      string
	User      string
	Password  string
	DBName    string
	SSLEnable bool
}
