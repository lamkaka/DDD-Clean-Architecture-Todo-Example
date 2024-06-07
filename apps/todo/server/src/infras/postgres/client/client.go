package postgres_client

import (
	"context"
	"fmt"

	"github.com/lib/pq"

	"todo-server/infras/postgres/client/ent"
)

type Client = *ent.Client

func NewClient(ctx context.Context, config Config) (Client, error) {
	err := createDatabase(ctx, config)
	if err != nil {
		return nil, err
	}

	client, err := ent.Open("postgres", getDataSourceName(config))
	if err != nil {
		return nil, err
	}

	// Run the auto migration tool.
	err = client.Schema.Create(ctx)
	if err != nil {
		return nil, err
	}

	return client, nil
}

func createDatabase(ctx context.Context, config Config) error {
	bootstrapClient, err := newBootstrapClient(ctx, config)
	if err != nil {
		return err
	}

	_, err = bootstrapClient.ExecContext(ctx, fmt.Sprintf("CREATE DATABASE %s", config.DBName))
	if isDuplicateDatabaseError(err) {
		return nil
	}
	return err
}

// Create ent client for creating the database
func newBootstrapClient(ctx context.Context, config Config) (Client, error) {
	bootstrapConfig := config
	bootstrapConfig.DBName = "postgres"

	return ent.Open("postgres", getDataSourceName(bootstrapConfig))
}

func getDataSourceName(config Config) string {
	sslMode := "require"
	if !config.SSLEnable {
		sslMode = "disable"
	}
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", config.Host, config.Port, config.User, config.Password, config.DBName, sslMode)
}

func isDuplicateDatabaseError(err error) bool {
	sqlErr, isSQLErr := err.(*pq.Error)
	// postgresql reference: https://www.postgresql.org/docs/8.2/errcodes-appendix.html
	// 42P04 == (DUPLICATE DATABASE) duplicate_database
	return isSQLErr && sqlErr.Code == "42P04"
}
