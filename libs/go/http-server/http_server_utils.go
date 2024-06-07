package http_server

import (
	"fmt"
	"strings"
)

func getAddress(config Config) string {
	if strings.HasPrefix(config.ServerPort, ":") {
		return config.ServerPort
	}
	return fmt.Sprintf(":%s", config.ServerPort)
}
