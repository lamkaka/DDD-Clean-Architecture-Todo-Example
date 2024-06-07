package env

import (
	"context"
	postgres_client "todo-server/infras/postgres/client"
)

func NewPostgresConfig(ctx context.Context, parser Parser) (postgres_client.Config, error) {
	var e postgresConfigEnv
	err := parser.Parse(ctx, &e)
	if err != nil {
		return postgres_client.Config{}, err
	}

	return postgres_client.Config(e), nil
}

type postgresConfigEnv struct {
	Host      string `env:"POSTGRES_HOST"`
	Port      string `env:"POSTGRES_PORT" envDefault:"5432"`
	User      string `env:"POSTGRES_USER"`
	Password  string `env:"POSTGRES_PASSWORD"`
	DBName    string `env:"POSTGRES_DB_NAME"`
	SSLEnable bool   `env:"POSTGRES_SSL_ENABLE" envDefault:"false"`
}
