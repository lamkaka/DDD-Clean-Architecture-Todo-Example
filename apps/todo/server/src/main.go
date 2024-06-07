//go:generate go generate ./infras/postgres/client/ent
//go:generate go run github.com/google/wire/cmd/wire@v0.5.0

package main

import "context"

func main() {
	ctx := context.Background()

	server, err := initializeServer(ctx)
	if err != nil {
		panic(err)
	}
	server.Start(ctx)
}
